package main

import (
	"bufio"
	"os"
	"pay-later-service/service"
	"pay-later-service/storage"
)

func main() {
	in := bufio.NewScanner(os.Stdin)

	app := service.Context{}
	app.DBClient = storage.GetClient()
	for in.Scan() {
		text := in.Text()
		app.Command(text)
	}
}
