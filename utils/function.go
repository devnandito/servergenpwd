package utils

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
)

type Menu struct {
	Url string
	Show string
	Create string
	Put string
	Delete string
	Detail string
	Change string
	Generate string
}

var (
	lowerCharSet = "abcdefghijklmnopqrstuvwxyz"
	upperCharSet = "ABCDEFGHIGKLMNOPQRSTUVWXYZ"
	specialCharSet = "!#$%&*"
	numberSet = "0123456789"
	allCharSet = lowerCharSet + upperCharSet + specialCharSet + numberSet
	minSpecialChar = 1
	minNum = 1
	minUpperCase = 1
	// i = 5.0
)

func GetMenu() []Menu {
	m := []Menu{
		{Url: "/", Show: "show", Create: "create", Put: "put", Delete: "delete", Detail: "detail", Generate: "generate"},
		{Url: "password", Show: "show", Create: "create", Put: "put", Delete: "delete", Detail: "detail", Generate: "generate"},
	}
	return m
}

func Ckeck(e error) {
	if e != nil {
		log.Panic(e)
	}
}

func GeneratePassword(passwordLength int) string {
	var password strings.Builder

	//Set special character
	for i := 0; i < minSpecialChar; i++ {
		random := rand.Intn(len(specialCharSet))
		password.WriteString(string(specialCharSet[random]))
	}

	//Set numeric
	for i := 0; i < minNum; i++ {
		random := rand.Intn(len(numberSet))
		password.WriteString(string(numberSet[random]))
	}

	//Set uppercase
	for i := 0; i < minUpperCase; i++ {
		random := rand.Intn(len(numberSet))
		password.WriteString(string(upperCharSet[random]))
	}

	remaingLength := passwordLength - minSpecialChar - minNum - minUpperCase
	for i := 0; i < remaingLength; i ++ {
		random := rand.Intn(len(allCharSet))
		password.WriteString(string(allCharSet[random]))
	}
	
	inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j],inRune[i]
	})

	return string(inRune)
}

func GetContentMail(gree, system, lenPwd, username, pwd string) ([]string, []byte) {
	var txt string
	var lines = []string {
		gree,
		fmt.Sprintf("Por este medio remito las credenciales de acceso a SIPAP %s", system),
		"Tener en cuenta las siguientes políricas de seguridad al cambiar la contraseña:",
		"- No debe ser una palabra que pueda estar en algún diccionario, ni relacionado al nombre de usuario o entidad.",
		"- No debe ser correlativo al anterior, ni similar en las últimas 10 contraseñas.",
		fmt.Sprintf("- Debe contener por los menos una mayúscula, una minúscula, un número y un carácter especial y de longitud mínima de %s.", lenPwd),
		"- Vigencia mínima de contraseña, de 1(un) días. Un solo cambio por día es posible.",
		"- Escriba su contraseña en un bloc de notas para luego copiar y pegar en el sistema.",
		"USUARIO",
		username,
		"CONTRASEÑA",
		pwd,
		"Favor dar acuse.",
		"Saludos cordiales.",
	}


	for _, line := range lines {
		txt = txt +"\n " + line
	}

	data := []byte(string(txt))

	return lines, data
}