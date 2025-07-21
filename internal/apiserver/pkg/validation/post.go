package validation

import (
	"context"

	v1 "github.com/TobyIcetea/fastgo/pkg/api/apiserver/v1"
)

func (v *Validator) ValidateCreatePostRequest(ctx context.Context, rq *v1.CreatePostRequest) error {
	return nil
}

func (v *Validator) ValidateUpdatePostRequest(ctx context.Context, rq *v1.UpdatePostRequest) error {
	return nil
}

func (v *Validator) ValidateDeletePostRequest(ctx context.Context, rq *v1.DeletePostRequest) error {
	return nil
}

func (v *Validator) ValidateGetPostRequest(ctx context.Context, rq *v1.GetPostRequest) error {
	return nil
}

func (v *Validator) ValidateListPostRequest(ctx context.Context, rq *v1.ListPostRequest) error {
	return nil
}
