package models

type User struct {
	Id_user               uint `json:"id_user" gorm:"primaryKey;autoIncrement:true"`
	Nama, Email, Username string
	Password              string
}

//Tulis jawab code nya di sini

type buyers struct {
	id_buyer   int `json:"id_buyer" gorm:"primaryKey;autoIncrement=true"`
	name_buyer string
	email      string
	password   string
}

type seller struct {
	id_seller   int `json:"id_seller" gorm:"primaryKey;autoIncrement=true"`
	name_seller string
	email       string
	password    string
}

type product struct {
	id_product    int `json:"id_product" gorm:"primaryKey;autoIncrement=true"`
	name_product  string
	exp_date      string
	price_product int
}
