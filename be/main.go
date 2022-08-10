package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/itsapep/golang-sample-xss/model"
	"github.com/microcosm-cc/bluemonday"
)

// go github.com/gin-gonic/gin
// go get github.com/microcosm-cc/bluemonday

func NewUser(user model.User) *model.User {
	p := bluemonday.UGCPolicy()
	user.Username = p.Sanitize(user.Username)
	user.FirstName = p.Sanitize(user.FirstName)
	user.LastName = p.Sanitize(user.LastName)
	return &user
}

func main() {
	users := make([]model.User, 0)
	apiHost := os.Getenv("API_HOST")
	apiPort := os.Getenv("API_PORT")
	listenAddress := fmt.Sprintf("%s:%s", apiHost, apiPort)

	routerEngine := gin.Default()
	routerGroup := routerEngine.Group("/api")
	routerGroup.POST("/user", func(ctx *gin.Context) {
		var newUser model.User
		err := ctx.ShouldBindJSON(&newUser)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		// users = append(users, newUser)
		// sanitised
		users = append(users, *NewUser(newUser))
		ctx.JSON(http.StatusOK, gin.H{
			"message": "SUCCESS",
			"data":    newUser,
		})
	})

	routerGroup.GET("/user", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "SUCCESS",
			"data":    users,
		})
	})

	err := routerEngine.Run(listenAddress)
	if err != nil {
		panic(err)
	}
}
