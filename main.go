package main

import (
	"app/check"
	"log"
	"time"
)

func main() {

	rootDir := `D:\Projects\go-lang\cup-service\cup-service\static\js`
	filterOnDir := []string{"vue", "local"}
	dateTagXml := "2022-10-31"
	parsedDate, err := time.Parse("2006-01-02", dateTagXml)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	ck := check.Checker{
		RootDir:     rootDir,
		FilterOnDir: filterOnDir,
		TagDate:     parsedDate,
	}
	if err := ck.Process(); err != nil {
		log.Fatal("Error: ", err)
	}

	log.Println("That's all folks!")

}
