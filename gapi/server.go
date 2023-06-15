package gapi

import (
	"fmt"

	db "github.com/angrypenguin1995/simple__bank/db/sqlc"
	"github.com/angrypenguin1995/simple__bank/pb"
	"github.com/angrypenguin1995/simple__bank/token"
	"github.com/angrypenguin1995/simple__bank/util"
)

// Server serves grpc requests
type Server struct {
	pb.UnimplementedSimpleBankServer
	store      db.Store
	tokenMaker token.Maker
	config     util.Config
}

// New Server creates a new GRPC server
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	// tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey) //uncomment this line and comment above line for using JWT instead of PASETO for token
	if err != nil {
		return nil, fmt.Errorf("Failed to create a tokenmaker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
