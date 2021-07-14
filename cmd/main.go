package main

import (
	"github.com/Sent-Malin/finish_quest_golang/internal/app/api"
	"github.com/Sent-Malin/finish_quest_golang/signal"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.POST("/create", api.Handling_query_create)

	router.POST("/convert", api.Handling_query_convert)

	//запуск цикла обновления
	go signal.Update_in_bd()

	router.Run()
}
