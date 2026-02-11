package main

import(
	"file-manager/internal/config"
	"fmt"
)


func main(){
	cfg, err := config.NewConfig()
	if err != nil{
		fmt.Printf("error: %w", err)
		return
	}
	fmt.Println(cfg)
}