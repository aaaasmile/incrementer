package main

import (
	"app/check"
	"log"
	"strings"
	"time"
)

func main() {

	rootJs := `C:\Projects\Denzel\Denzel_2017_Upgrade\golang\Project_Denzel_WS\Project_Denzel_WS\OBWebservice\static\js\vue\main.js`
	dateTagXml := "2022-10-31"

	parsedDate, err := time.Parse("2006-01-02", dateTagXml)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	ck := check.Checker{
		RootJs:  strings.ReplaceAll(rootJs, "\\", "/"),
		TagDate: parsedDate,
	}
	if err := ck.Process(); err != nil {
		log.Fatal("Error: ", err)
	}

	log.Println("That's all folks!")

}
