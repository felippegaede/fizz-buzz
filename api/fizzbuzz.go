package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/felippegaede/fizz-buzz/config"
	"github.com/gin-gonic/gin"
)

type fizzBuzzRequest struct {
	Number int `uri:"number" binding:"required,min=1"`
}

func (server *Server) fizzBuzz(ctx *gin.Context) {
	var req fizzBuzzRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	n := req.Number
	var response []string

	for i := 0; i < n; i++ {
		if i%3 == 0 && i%5 == 0 {
			response = append(response, "FizzBuzz")
		} else if i%3 == 0 {
			response = append(response, "Fizz")
		} else if i%5 == 0 {
			response = append(response, "Buzz")
		} else {
			response = append(response, fmt.Sprint(i))

		}
	}

	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	ctx.JSON(http.StatusOK, gin.H{"app": config.AppName, "response": response})
}
