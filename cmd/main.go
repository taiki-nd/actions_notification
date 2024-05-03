package main

import (
	"actionsnotification/app"
	"actionsnotification/app/services"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("too many arguments. Please input only one argument.")
		os.Exit(1)
	}

	env := ""
	// Argsの中身を一件ずつ出力します
	for _, v := range os.Args {
		env = v
	}

	if env == "local" {
		err := app.InitAppOnLocal()
		if err != nil {
			os.Exit(1)
		}
	} else {
		err := app.InitApp()
		if err != nil {
			os.Exit(1)
		}
	}
	err := services.SendMessage()
	if err != nil {
		os.Exit(1)
	}
}
