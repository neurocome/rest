package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/neurocome/rest/entity"
	"github.com/neurocome/rest/services"
	"github.com/neurocome/rest/validators"
	"gopkg.in/go-playground/validator.v9"
)

type EmployeController interface {
	FindAll() []entity.Employe
	Save(ctx *gin.Context) error
	// ShowAll(ctx *gin.Context)
}

type controller struct {
	service services.EmployeService
}

var validate *validator.Validate

func New(service services.EmployeService) EmployeController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateTitle)
	return &controller{
		service: service,
	}
}
func (c *controller) FindAll() []entity.Employe {
	return c.service.FindAll()
}
func (c *controller) Save(ctx *gin.Context) error {
	var employe entity.Employe
	err := ctx.ShouldBindJSON(&employe)
	fmt.Print(err)
	if err != nil {
		return err
	}
	err = validate.Struct(employe)
	if err != nil {
		return err
	}
	c.service.Save(employe)
	return nil
}

// func (req *controller) ShowAll(ctx *gin.Context)  {
// 	employes := req.service.FindAll()
// 	data := gin.H{
// 		"title" : "employes",
// 		"data"	: employes,
// 	}
// 	ctx.HTML(http.StatusOK, "index.html", data)
// }
