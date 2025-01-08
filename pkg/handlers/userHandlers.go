package handlers

import (
	"github.com/Venukishore-R/CODECRAFT_BW_02/internal/app/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Create(ctx *gin.Context) {
	var user *models.User

	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{
			"message": "invalid json body",
			"error":   err.Error(),
		})
		return
	}

	if err := h.Service.Create(user); err != nil {
		ctx.JSON(500, gin.H{
			"message": "failed to create user",
			"error":   err.Error(),
		})

		return
	}

	ctx.JSON(201, gin.H{
		"message": "user created successfully",
		"user":    user,
	})
}

func (h *Handler) FindAll(ctx *gin.Context) {
	users, err := h.Service.FindAll()
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "failed to find users",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "users retrieved successfully",
		"users":   users,
	})
}

func (h *Handler) FindByEmail(ctx *gin.Context) {
	email := ctx.Param("email")

	user, err := h.Service.FindByEmail(email)
	if err != nil {
		ctx.JSON(404, gin.H{
			"message": "user not found",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "user retrieved successfully",
		"user":    user,
	})
}

func (h *Handler) Update(ctx *gin.Context) {
	email := ctx.Param("email")

	var user *models.User

	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{
			"message": "invalid json body",
			"error":   err.Error(),
		})
		return
	}

	if err := h.Service.Update(email, user); err != nil {
		ctx.JSON(500, gin.H{
			"message": "failed to update user",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "user updated successfully",
		"user":    user,
	})
}

func (h *Handler) Delete(ctx *gin.Context) {
	email := ctx.Param("email")

	if err := h.Service.Delete(email); err != nil {
		ctx.JSON(500, gin.H{
			"message": "failed to delete user",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "user deleted successfully",
		"email":   email,
	})
}
