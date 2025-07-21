package handler

import (
	"github.com/TobyIcetea/fastgo/internal/apiserver/biz"
	"github.com/TobyIcetea/fastgo/internal/apiserver/pkg/validation"
)

// Handler 处理博客模块的请求
type Handler struct {
	biz biz.IBiz
	val *validation.Validator
}

// NewHandler 创建新的 Handler 实例
func NewHandler(biz biz.IBiz, val *validation.Validator) *Handler {
	return &Handler{
		biz: biz,
		val: val,
	}
}
