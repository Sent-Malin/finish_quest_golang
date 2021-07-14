package models

//структура для хранения данных валют с api
type Struct_currency struct {
	Ccy      string `json:"ccy"`
	Base_ccy string `json:"base_ccy"`
	Buy      string `json:"buy"`
	Sale     string `json:"sale"`
}

//структура для хранения данных введенных в запросе create
type Struct_contein_data_create struct {
	Currency1 string `json:"currency1"`
	Currency2 string `json:"currency2"`
}

//структура для хранения данных введенных в запросе convert
type Struct_contein_data_convert struct {
	CurrencyFrom string `json:"currencyFrom"`
	CurrencyTo   string `json:"currencyTo"`
	Value        int    `json:"value"`
}

type Config struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
}

func NewConfig() *Config {
	return &Config{
		Host:     "",
		Port:     0,
		User:     "",
		Password: "",
		Dbname:   "",
	}
}
