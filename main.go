package main

import (
	"app/check"
	"log"
	"time"
)

func main() {

	rootJs := `D:\Projects\go-lang\cup-service\cup-service\static\js\vue\main.js`
	dateTagXml := "2022-10-31"
	parsedDate, err := time.Parse("2006-01-02", dateTagXml)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	ck := check.Checker{
		RootJs:  rootJs,
		TagDate: parsedDate,
	}
	if err := ck.Process(); err != nil {
		log.Fatal("Error: ", err)
	}

	log.Println("That's all folks!")

}
