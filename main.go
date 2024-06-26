package main

import (
	"fmt"
	"os"

	"github.com/taiki-nd/actions_notification/app"
	"github.com/taiki-nd/actions_notification/app/services"
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

	err := app.InitApp(env)
	if err != nil {
		os.Exit(1)
	}

	err = services.SendMessage()
	if err != nil {
		os.Exit(1)
	}
}
