package gapi

import (
	"context"
	"database/sql"

	db "github.com/angrypenguin1995/simple__bank/db/sqlc"
	"github.com/angrypenguin1995/simple__bank/pb"
	"github.com/angrypenguin1995/simple__bank/util"
	"github.com/angrypenguin1995/simple__bank/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {

	violations := ValidateLoginUserRequest(req)

	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	user, err := server.store.GetUser(ctx, req.GetUsername())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "User %s not found", req.GetUsername())
		}
		// ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return nil, status.Errorf(codes.Internal, "Failed to Search user %s", req.GetUsername())
	}
	err = util.CheckPassword(req.GetPassword(), user.HashedPassword)
	if err != nil {
		// ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return nil, status.Errorf(codes.Unauthenticated, "Given Username and Password dont match %s", err)
	}

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(
		user.Username,
		server.config.AccessTokenDuration,
	)

	if err != nil {
		// ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return nil, status.Errorf(codes.Internal, "error creating  Access token %s", err)
	}

	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(
		user.Username,
		server.config.RefreshTokenDuration,
	)

	if err != nil {
		// ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return nil, status.Errorf(codes.Internal, "error creating Refresh token %s", err)
	}
	metadata := server.extractMetadata(ctx)
	session, err := server.store.CreateSession(ctx, db.CreateSessionParams{
		ID:           refreshPayload.ID,
		Username:     user.Username,
		RefreshToken: refreshToken,
		UserAgent:    metadata.UserAgent, //ctx.Request.UserAgent(),
		ClientIp:     metadata.ClientIP,  //ctx.ClientIP(),
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
	})

	if err != nil {
		// ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return nil, status.Errorf(codes.Internal, "error creating DB session %s", err)
	}

	rsp := &pb.LoginUserResponse{
		User:                  convertUser(user),
		SessionId:             session.ID.String(),
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresat:  timestamppb.New(accessPayload.ExpiredAt),
		RefreshTokenExpiresat: timestamppb.New(refreshPayload.ExpiredAt),
	}
	return rsp, nil
}

func ValidateLoginUserRequest(req *pb.LoginUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, fieldViolation("username ", err))
	}

	if err := val.ValidatePassword(req.GetPassword()); err != nil {
		violations = append(violations, fieldViolation("password", err))
	}

	return violations
}
