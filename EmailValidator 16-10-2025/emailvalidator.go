package main

import (
	"fmt"
	"regexp"
)

func emailValidator(email string) bool {
	r, _ := regexp.Compile(`^[^.].+[^.]@(\w+\\.)+\w\w+"`)
	return r.MatchString(email)
}

func main() {
	fmt.Println(emailValidator("avalid@email.com"))
	fmt.Println(emailValidator(".notavalid@email.com"))
	fmt.Println(emailValidator("alsonotavalid@email.c"))
}
