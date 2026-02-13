package main

import(
	"file-manager/internal/config"
	"file-manager/internal/parser"
	"file-manager/internal/report"
	"log"
	
	"context"
)


func main(){

	cfg, err := config.NewConfig()
	if err != nil{
		log.Fatal(cfg, err)
		return
	}
	
	parsedFile, err := parser.ParseTSVFile("test_file.tsv")
	if err != nil{
		log.Fatal(err)
	}
}