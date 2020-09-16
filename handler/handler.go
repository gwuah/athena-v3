package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gwuah/api/lib/dispatch"
	"github.com/gwuah/api/middleware"
	"github.com/gwuah/api/repository"
	"github.com/gwuah/api/utils/jwt"
	"github.com/gwuah/api/utils/secure"
	"gorm.io/gorm"
)

type Handler struct {
	DB   *gorm.DB
	Repo *repository.Repository
	JWT  jwt.Service
	Sec  *secure.Service
}

func New(DB *gorm.DB, jwt jwt.Service, sec *secure.Service) *Handler {
	repo := repository.New(DB)
	return &Handler{DB, repo, jwt, sec}
}

func (h *Handler) Register(v1 *gin.RouterGroup) {

	dispatch := dispatch.New(h.DB)

	v1.GET("/customer/realtime/:id", dispatch.HandleConnection("customer"))
	v1.GET("/electron/realtime/:id", dispatch.HandleConnection("electron"))

	v1.POST("/signup", h.Signup)
	v1.POST("/login", h.Login)
	v1.POST("/otp/verify", h.VerifyOTP)
	v1.GET("/refresh/:token", h.Refresh)

	customers := v1.Group("/customers", middleware.JWT(h.JWT))
	customers.GET("/", h.ListCustomers)
	customers.GET("/:id", h.ViewCustomer)
	customers.POST("/", h.CreateCustomer)

}