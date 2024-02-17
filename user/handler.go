package main

import (
	"context"
	user "github.com/FrazierLei/MyCC98/common/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Signup implements the UserServiceImpl interface.
func (s *UserServiceImpl) Signup(ctx context.Context, req *user.SignupRequest) (resp *user.SignupResponse, err error) {
	// TODO: Your code here...
	return
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	// TODO: Your code here...
	return
}

// FindOrCreate implements the UserServiceImpl interface.
func (s *UserServiceImpl) FindOrCreate(ctx context.Context, req *user.FindOrCreateRequest) (resp *user.FindOrCreateResponse, err error) {
	// TODO: Your code here...
	return
}

// Profile implements the UserServiceImpl interface.
func (s *UserServiceImpl) Profile(ctx context.Context, req *user.ProfileRequest) (resp *user.ProfileResponse, err error) {
	// TODO: Your code here...
	return
}

// UpdateNonSensitiveInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateNonSensitiveInfo(ctx context.Context, req *user.UpdateNonSensitiveInfoRequest) (resp *user.UpdateNonSensitiveInfoResponse, err error) {
	// TODO: Your code here...
	return
}
