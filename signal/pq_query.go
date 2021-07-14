package signal

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	models "github.com/Sent-Malin/finish_quest_golang"
	_ "github.com/lib/pq"
	"gopkg.in/yaml.v2"
)

func loadConfig() *models.Config {
	con := models.NewConfig()
	yamFile, err := ioutil.ReadFile("./config/app.yaml")
	if err != nil {
		log.Fatalf("Error #%v", err)
	}
	err = yaml.Unmarshal(yamFile, con)
	if err != nil {
		log.Fatalf("Error #%v", err)
	}
	return con
}

//Функция добавления записи в базу, возвращает true в случае успеха
func Create_in_bd(data_create models.Struct_contein_data_create) bool {

	conf := loadConfig()

	//строка подключения
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Port, conf.User, conf.Password, conf.Dbname)

	db, err := sql.Open("postgres", psqlconn)

	if err != nil {
		panic(err)
	}

	defer db.Close()
	//Проверка базы данных на наличие полученной позиции
	str_sql_select := fmt.Sprintf(`SELECT Валюта1 FROM "DataValue" WHERE Валюта1='%s' AND Валюта2='%s'`, data_create.Currency1, data_create.Currency2)

	res, err := db.Query(str_sql_select)

	if err != nil {
		panic(err)
	}
	//Если результат запроса не содержит строк, добавляем запись
	if res.Next() == false {

		str_sql_insert := fmt.Sprintf(`insert into "DataValue" ("Валюта1", "Валюта2", "курс", "последнее время обновления") values('%s', '%s', 0, null)`, data_create.Currency1, data_create.Currency2)

		_, e := db.Exec(str_sql_insert)

		if e != nil {
			panic(e)
		}
		return true
	}
	return false
}

//функция, возвращающая курс указанной валюты из бд
func Get_currency_from_bd(data_for_find models.Struct_contein_data_convert) float64 {

	conf := loadConfig()
	//строка подключения
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Port, conf.User, conf.Password, conf.Dbname)

	db, err := sql.Open("postgres", psqlconn)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	str_sql_select := fmt.Sprintf(`SELECT курс FROM "DataValue" WHERE Валюта1='%s' AND Валюта2='%s'`, data_for_find.CurrencyFrom, data_for_find.CurrencyTo)

	res, err := db.Query(str_sql_select)

	if err != nil {
		panic(err)
	}

	var str_for_out_query string
	//если результат запроса есть
	if res.Next() == true {
		//считываем значение в строку
		res.Scan(&str_for_out_query)
		//переводим в формат float
		curs, _ := strconv.ParseFloat(str_for_out_query, 24)

		return curs
	}
	//при отсутствии записи в бд возвращаем -1
	return -1
}
