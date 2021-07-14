package signal

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	models "github.com/Sent-Malin/finish_quest_golang"
)

//функция обновления курса валют в бд
func Update_in_bd() {
	//адрес банка
	url := "https://api.privatbank.ua/p24api/pubinfo?json&exchange&coursid=5"

	var netClient = http.Client{
		Timeout: time.Second * 10,
	}
	//каждые 2 минуты мы получаем данные по указанному url и обновляем бд
	for true {
		res, err := netClient.Get(url)

		if err != nil {
			log.Fatal(err)
		}

		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)

		if err != nil {
			log.Fatal(err)
		}

		var currency_list = []models.Struct_currency{}

		serr := json.Unmarshal(body, &currency_list)

		if serr != nil {
			log.Fatal(serr)
		}

		conf := loadConfig()
		//строка подключения
		psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Port, conf.User, conf.Password, conf.Dbname)

		db, err := sql.Open("postgres", psqlconn)

		if err != nil {
			panic(err)
		}

		defer db.Close()
		//Проходимся по всем значениям валют, если находим совпадение в бд, обновляем
		for _, record := range currency_list {
			//переводим данные валюты в формат real для базы данных
			curs_buy, _ := strconv.ParseFloat(record.Buy, 24)

			curs_sale, _ := strconv.ParseFloat(record.Sale, 24)
			if curs_sale != 0 {
				curs_sale = 1 / curs_sale
			}
			//Получаем данные о текущей дате для базы данных
			string_time_now := time.Now().String()

			date_time_now := string_time_now[0:10]
			//строка для обновления данных конвертации с 1 на 2 валюту
			str_sql_update1 := fmt.Sprintf(`update "DataValue" SET курс=%f, "последнее время обновления"='%s' WHERE Валюта1='%s' AND Валюта2='%s'`, curs_buy, date_time_now, record.Ccy, record.Base_ccy)
			//строка для обновления данных конвертации со 2 на 1 валюту
			str_sql_update2 := fmt.Sprintf(`update "DataValue" SET курс=%f, "последнее время обновления"='%s' WHERE Валюта1='%s' AND Валюта2='%s'`, curs_sale, date_time_now, record.Base_ccy, record.Ccy)

			_, e := db.Exec(str_sql_update1)

			if e != nil {
				panic(e)
			}

			_, er := db.Exec(str_sql_update2)

			if er != nil {
				panic(er)
			}
		}

		time.Sleep(120 * time.Second)
	}
}
