package auth

import (
	context "context"
)

type RegisterRequest struct {
	Email    string
	Password string
}

type RegisterResponse struct {
	UserId string
	Token  string
}

type LoginRequest struct {
	Email    string
	Password string
}

type LoginResponse struct {
	UserId string
	Token  string
}

type GetUserProfileRequest struct {
	UserId string
}

type GetUserProfileResponse struct {
	UserId    string
	Email     string
	CreatedAt int64
}

// AuthServiceServer is the server API for AuthService service.
type AuthServiceServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	GetUserProfile(context.Context, *GetUserProfileRequest) (*GetUserProfileResponse, error)
}

// UnimplementedAuthServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct{}

func (UnimplementedAuthServiceServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, nil
}

func (UnimplementedAuthServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, nil
}

func (UnimplementedAuthServiceServer) GetUserProfile(context.Context, *GetUserProfileRequest) (*GetUserProfileResponse, error) {
	return nil, nil
}
