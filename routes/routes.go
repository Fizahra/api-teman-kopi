package routes

import (
	"api_temankopi/config"
	"api_temankopi/handler"
	"api_temankopi/repository"
	"api_temankopi/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Route(r *gin.Engine) {
	var (
		db              *gorm.DB                   = config.SetupDBConnection()
		staffRepository repository.StaffRepository = repository.NewStaffRepository(db)
		staffService    service.StaffService       = service.NewStaffService(staffRepository)
		staffController handler.StaffHandler       = handler.NewStaffHandler(staffService)
	)
	Route := r.Group("staff")
	{
		Route.GET("/", staffController.ViewStaff)
		Route.POST("/", staffController.AddStaff)
		Route.PUT("/:id", staffController.EditStaff)
		Route.DELETE("/:id", staffController.DeleteStaff)
	}
}
