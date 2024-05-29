package handlers

func Login(username string, password string) string {
	//более сложной обработки логина/пароля задание не требует
	if username == "admin" && password == "password" {
		return "True"
	} else {
		return "False"
	}
}
