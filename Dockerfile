# # this gives a huge image which has all dependencies in it
# FROM golang:1.20.4-alpine3.18
# #  here we are specifying that working directory inside image is "/app"
# WORKDIR /app
# # here we are copying everything from the current directory into "/app" of image that will be generated
# COPY . .
# # This is the place where we enter 
# RUN go build -o main main.go

# #the port to expose- doesnt have any consequnce on build but is used for when people want to implement it
# EXPOSE 8080

# #command to run on start of file
# CMD ["app/main"]

# BUILD STAGE
FROM golang:1.20.4-alpine3.18 AS builder
#  here we are specifying that working directory inside image is "/app"
WORKDIR /app
# here we are copying everything from the current directory into "/app" of image that will be generated
COPY . .
# COPY ./go.mod .
# RUN go mod download
# This is the place where we enter 
RUN go build -o main main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz

# RUN STAGE
# this specifies the base file for image is from alpine:3.18
FROM alpine:3.18
WORKDIR /app 
#  we are copying files from builder to working directroy in image
COPY --from=builder /app/main .
COPY --from=builder /app/migrate ./migrate
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./migration

# #the port to expose- doesnt have any consequnce on build but is used for when people want to implement it
EXPOSE 8080

# #command to run on start of file
CMD ["/app/main"]
ENTRYPOINT [ "/app/start.sh" ]