package app

import (
	"github.com/gin-gonic/gin"
)

func UserRoutes(rg *gin.RouterGroup) {
	//rg.GET("/user/:userid", warrantyByUser)
	//rg.POST("/users", createUser)
}

// func createUser(c *gin.Context) {
// 	db := c.Value("db").(*pop.Connection)

// 	u := &models.User{}
// 	err := c.BindJSON(u)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": "error creating user",
// 		})
// 		log.Println(err)
// 		return
// 	}
// 	verrs, err := db.ValidateAndCreate(u, "ID")
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"message": "error saving user",
// 		})
// 		log.Println(err)

// 		return
// 	}
// 	if verrs.HasAny() {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": verrs,
// 		})
// 		log.Println(verrs)
// 		return
// 	}
// 	c.IndentedJSON(http.StatusOK, u)
// }
