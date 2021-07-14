package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

	models "github.com/Sent-Malin/finish_quest_golang"
	"github.com/Sent-Malin/finish_quest_golang/signal"
	"github.com/gin-gonic/gin"
)

//функция обработки запроса на создание записи
func Handling_query_create(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)

	if err != nil {
		log.Fatal(err)
	}

	var data_for_create = models.Struct_contein_data_create{}

	serr := json.Unmarshal(body, &data_for_create)

	if serr != nil {
		log.Fatal(serr)
	}
	//вызов функции вставки в базу данных и проверка результата
	if signal.Create_in_bd(data_for_create) == true {

		string_for_out := fmt.Sprintf(`%s %s`, data_for_create.Currency1, data_for_create.Currency2)

		c.JSON(200, gin.H{
			"message":     "Запись добавлена",
			"was created": string_for_out,
		})
	} else {
		c.JSON(200, gin.H{
			"message":     "Операция завершена",
			"was created": "в базе данных уже есть данная строка",
		})
	}
}
