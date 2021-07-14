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

//функция обработки запроса на конвертацию значений
func Handling_query_convert(c *gin.Context) {

	body, err := io.ReadAll(c.Request.Body)

	if err != nil {
		log.Fatal(err)
	}

	var data_for_convert = models.Struct_contein_data_convert{}

	serr := json.Unmarshal(body, &data_for_convert)

	if serr != nil {
		log.Fatal(serr)
	}
	//присваиваем переменной значение курса валюты из бд
	curs := signal.Get_currency_from_bd(data_for_convert)

	if curs > 0 {
		out_cur := float64(data_for_convert.Value) * curs
		string_for_out := fmt.Sprintf(`%d from %s in %s`, data_for_convert.Value, data_for_convert.CurrencyFrom, data_for_convert.CurrencyTo)
		c.JSON(200, gin.H{
			"convert": string_for_out,
			"result":  out_cur,
		})
	} else if curs == 0 {
		c.JSON(200, gin.H{
			"message": "данное значение найдено, но еще не обновилось",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "данное значение не найдено",
		})
	}
}
