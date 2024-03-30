package grpc

import (
	"context"
	"example.com/m/internal/genproto/user_pb/pb"
	"example.com/m/internal/user/app"
	"time"
)

type UserServer struct {
	userUseCase app.UserUseCase
	pb.UnimplementedUserServiceServer
}

func NewUserGrpcServer(userUseCase app.UserUseCase) pb.UserServiceServer {
	return &UserServer{userUseCase: userUseCase}
}

func (u *UserServer) SignUpAdmin(ctx context.Context, req *pb.NewUser) (*pb.Error, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := u.userUseCase.SignUpAdmin(req)

	if err != nil {
		return &pb.Error{
			Message: err.Error(),
		}, nil
	}

	return nil, nil
}

func (u *UserServer) SignUpUser(ctx context.Context, req *pb.NewUser) (*pb.Error, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := u.userUseCase.SignUpUser(req)

	if err != nil {
		return &pb.Error{
			Message: err.Error(),
		}, nil
	}

	return nil, nil
}

func (u *UserServer) SignInUser(ctx context.Context, req *pb.UserCredentials) (*pb.SignInResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := u.userUseCase.SignInUser(req.Email, req.Password)

	if err != nil {
		return &pb.SignInResponse{
			Success: false,
		}, nil
	}

	return &pb.SignInResponse{Success: true}, nil
}

func (u *UserServer) GetUsers(ctx context.Context, req *pb.UserRequest) (*pb.UsersResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	users, err := u.userUseCase.GetUsers(req)

	if err != nil {
		return nil, err
	}

	var pbUsers []*pb.UserResponse
	for _, user := range users.GetUsers() {

		pbUsers = append(pbUsers, user)
	}

	response := &pb.UsersResponse{
		Users: pbUsers,
	}
	return response, nil
}

func (u *UserServer) GetUser(ctx context.Context, req *pb.UserId) (*pb.UserResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user, err := u.userUseCase.GetUser(req)
	if err != nil {
		return &pb.UserResponse{}, nil
	}

	return user, nil
}

func (u *UserServer) UpdateUser(ctx context.Context, req *pb.UserUpdate) (*pb.Error, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := u.userUseCase.Update(req)

	if err != nil {
		return &pb.Error{Message: "failed to update"}, nil
	}

	return nil, nil
}

func (u *UserServer) DeleteUser(ctx context.Context, req *pb.UserId) (*pb.Error, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := u.userUseCase.Delete(req)

	if err != nil {
		return &pb.Error{Message: "failed to delete"}, nil
	}

	return nil, nil
}
