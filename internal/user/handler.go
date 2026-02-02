package user

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		Service: service,
	}
}

// Request Structures
type RegisterRequest struct {
	FullName  string `json:"full_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	ContactNo string `json:"contact_no" binding:"required"`
}
type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// --- API Functions ---
func (h *Handler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := User{
		FullName:  req.FullName,
		Email:     req.Email,
		Password:  req.Password,
		ContactNo: &req.ContactNo,
	}

	if err := h.Service.Register(&user); err != nil {
		c.JSON(500, gin.H{"error": "Failed to register"})
		return
	}
	c.JSON(200, gin.H{
		"message": "User created successfully",
	})
}

func (h *Handler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	token, err := h.Service.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"token": token})
}
