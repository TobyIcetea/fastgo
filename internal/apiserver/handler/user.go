package handler

import (
	"log/slog"

	"github.com/TobyIcetea/fastgo/internal/pkg/core"
	v1 "github.com/TobyIcetea/fastgo/pkg/api/apiserver/v1"
	"github.com/gin-gonic/gin"
	"github.com/onexstack/onexstack/pkg/errorsx"
)

// Login 用户登录并返回 JWT Token.
func (h *Handler) Login(c *gin.Context) {
	slog.Info("Login function called")

	var rq v1.LoginRequest
	if err := c.ShouldBindJSON(&rq); err != nil {
		core.WriteResponse(c, nil, errorsx.ErrBind)
		return
	}

	// 小作业：请你自行补全校验代码

	resp, err := h.biz.UserV1().Login(c.Request.Context(), &rq)
	if err != nil {
		core.WriteResponse(c, nil, err)
		return
	}

	core.WriteResponse(c, resp, nil)
}

// RefreshToken 刷新 JWT Token.
func (h *Handler) RefreshToken(c *gin.Context) {
	slog.Info("Refresh token function called")

	var rq v1.RefreshTokenRequest
	if err := c.ShouldBindJSON(&rq); err != nil {
		core.WriteResponse(c, nil, errorsx.ErrBind)
		return
	}

	// 小作业：请你自行补全校验代码

	resp, err := h.biz.UserV1().RefreshToken(c.Request.Context(), &rq)
	if err != nil {
		core.WriteResponse(c, nil, err)
		return
	}

	core.WriteResponse(c, resp, nil)
}

// ChangeUserPassword 修改用户密码.
func (h *Handler) ChangePassword(c *gin.Context) {
	slog.Info("Change password function called")

	var rq v1.ChangePasswordRequest
	if err := c.ShouldBindJSON(&rq); err != nil {
		core.WriteResponse(c, nil, errorsx.ErrBind)
		return
	}

	// 小作业：请你自行补全校验代码

	resp, err := h.biz.UserV1().ChangePassword(c.Request.Context(), &rq)
	if err != nil {
		core.WriteResponse(c, nil, err)
		return
	}

	core.WriteResponse(c, resp, nil)
}

// CreateUser 创建新用户
func (h *Handler) CreateUser(c *gin.Context) {
	slog.Info("Create user function called")

	var rq v1.CreateUserRequest
	if err := c.ShouldBindJSON(&rq); err != nil {
		core.WriteResponse(c, nil, errorsx.ErrBind)
		return
	}

	if err := h.val.ValidateCreateUserRequest(c.Request.Context(), &rq); err != nil {
		core.WriteResponse(c, nil, err)
		return
	}

	resp, err := h.biz.UserV1().Create(c.Request.Context(), &rq)
	if err != nil {
		core.WriteResponse(c, nil, err)
		return
	}

	core.WriteResponse(c, resp, nil)
}

// UpdateUser 更新用户信息
func (h *Handler) UpdateUser(c *gin.Context) {
	slog.Info("Update user function called")

	var rq v1.UpdateUserRequest
	if err := c.ShouldBindJSON(&rq); err != nil {
		core.WriteResponse(c, nil, errorsx.ErrBind)
		return
	}

	if err := h.val.ValidateUpdateUserRequest(c.Request.Context(), &rq); err != nil {
		core.WriteResponse(c, nil, err)
		return
	}

	resp, err := h.biz.UserV1().Update(c.Request.Context(), &rq)
	if err != nil {
		core.WriteResponse(c, nil, err)
		return
	}

	core.WriteResponse(c, resp, nil)
}

// DeleteUser 删除用户
func (h *Handler) DeleteUser(c *gin.Context) {
	slog.Info("Delete user function called")

	var rq v1.DeleteUserRequest
	if err := c.ShouldBindJSON(&rq); err != nil {
		core.WriteResponse(c, nil, errorsx.ErrBind)
		return
	}

	// 补全校验代码
	if err := h.val.ValidateDeleteUserRequest(c.Request.Context(), &rq); err != nil {
		core.WriteResponse(c, nil, err)
		return
	}

	resp, err := h.biz.UserV1().Delete(c.Request.Context(), &rq)
	if err != nil {
		core.WriteResponse(c, nil, err)
		return
	}

	core.WriteResponse(c, resp, nil)
}

// GetUser 获取用户信息
func (h *Handler) GetUser(c *gin.Context) {
	slog.Info("Get user function called")

	var rq v1.GetUserRequest
	if err := c.ShouldBindJSON(&rq); err != nil {
		core.WriteResponse(c, nil, errorsx.ErrBind)
		return
	}

	// 补全校验代码
	if err := h.val.ValidateGetUserRequest(c.Request.Context(), &rq); err != nil {
		core.WriteResponse(c, nil, err)
		return
	}

	resp, err := h.biz.UserV1().Get(c.Request.Context(), &rq)
	if err != nil {
		core.WriteResponse(c, nil, err)
		return
	}

	core.WriteResponse(c, resp, nil)
}

// ListUser 列出用户信息
func (h *Handler) ListUser(c *gin.Context) {
	slog.Info("List user function called")

	var rq v1.ListUserRequest
	if err := c.ShouldBindJSON(&rq); err != nil {
		core.WriteResponse(c, nil, errorsx.ErrBind)
		return
	}

	// 补全校验代码
	if err := h.val.ValidateListUserRequest(c.Request.Context(), &rq); err != nil {
		core.WriteResponse(c, nil, err)
		return
	}

	resp, err := h.biz.UserV1().List(c.Request.Context(), &rq)
	if err != nil {
		core.WriteResponse(c, nil, err)
		return
	}

	core.WriteResponse(c, resp, nil)
}
