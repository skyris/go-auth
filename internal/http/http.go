package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"

	"github.com/skyris/auth-server/internal/env"
	"github.com/skyris/auth-server/internal/usecase"
)

type Options struct{}

type Controller struct {
	ucUser  usecase.User
	router  *gin.Engine
	options Options
}

func New(ucUser usecase.User, options Options) *Controller {
	d := &Controller{
		ucUser:  ucUser,
		options: options,
	}
	d.initRouter()

	return d
}

func (d *Controller) SetOptions(options Options) {
	if d.options != options {
		d.options = options
	}
}

func (d *Controller) Run() error {
	addr := fmt.Sprintf("%s:%s", env.HOST, env.PORT)
	log.Println("Server is running on port", env.PORT)

	return d.router.Run(addr)
}
