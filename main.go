package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

type Person struct {
    ID string `uri:"id" binding:"required,uuid"`
    Name string `uri:"name" binding:"required"`
}

func getPerson(context *gin.Context) {
    var person Person
    if err := context.Copy().ShouldBindUri(&person); err != nil {
        context.JSON(
            http.StatusBadRequest,
            gin.H {
                "message": err.Error(),
            },
        )
    }

    context.JSON(
        http.StatusOK,
        gin.H {
            "name": person.Name,
            "uuid": person.ID,
        },
    )
}

func main() {
    route := gin.Default()
    route.GET("/:name/:id", getPerson)

    route.Run()
}
