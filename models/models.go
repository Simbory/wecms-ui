package models

type LoginModel struct {
	Email 		string "field:email"
	Password 	string "field:pwd"
	RememberMe 	string "field:remember"
}

