package validation

import (
	"context"
	"errors"

	v1 "github.com/TobyIcetea/fastgo/pkg/api/apiserver/v1"
)

func (v *Validator) ValidateCreateUserRequest(ctx context.Context, rq *v1.CreateUserRequest) error {
	// Validate username
	if rq.Username == "" {
		return errors.New("Username cannot be empty")
	}
	if len(rq.Username) < 4 || len(rq.Username) > 32 {
		return errors.New("username must be between 4 and 32 characters")
	}

	// Validate password
	if rq.Password == "" {
		return errors.New("Password cannot be empty")
	}
	if len(rq.Password) < 8 || len(rq.Password) > 64 {
		return errors.New("Password must be between 8 and 64 characters")
	}

	// Validate nickname(if probided)
	if rq.Nickname != nil && *rq.Nickname != "" {
		if len(*rq.Nickname) > 32 {
			return errors.New("Nickname cannot exceed 32 characters")
		}
	}

	// Validate email
	if rq.Email == "" {
		return errors.New("Email cannot be empty")
	}

	// Validate phone number
	if rq.Phone == "" {
		return errors.New("Phone number cannot be empty")
	}

	return nil
}

func (v *Validator) ValidateUpdateUserRequest(ctx context.Context, rq *v1.UpdateUserRequest) error {
	return nil
}

func (v *Validator) ValidateDeleteUserRequest(ctx context.Context, rq *v1.DeleteUserRequest) error {
	return nil
}

func (v *Validator) ValidateGetUserRequest(ctx context.Context, rq *v1.GetUserRequest) error {
	return nil
}

func (v *Validator) ValidateListUserRequest(ctx context.Context, rq *v1.ListUserRequest) error {
	return nil
}
