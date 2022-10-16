package main

import (
	"fmt"

	"github.com/spf13/viper"
)

// type Serve

func main() {
	v := viper.New()
	v.SetConfigFile("D:/project_gaos/go_books/golang_books_progams/exercise/config.yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	fmt.Println(v.Get("name"))
}
