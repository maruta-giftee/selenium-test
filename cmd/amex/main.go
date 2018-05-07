// main.go
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/maruta-giftee/selenium-test/parser"
	"github.com/maruta-giftee/selenium-test/selenium/chrome"
)

func main() {
	// Open file
	filePath := flag.String("f", "file", "file path")
	flag.Parse()
	log.Println("file path: " + *filePath)
	file, err := os.Open(*filePath)
	if err != nil {
		log.Fatalf("open file error %#v\n", err)
		return
	}
	defer file.Close()

	// New SeleniumDriver
	driver := chrome.NewDriver()
	if err := driver.Start(); err != nil {
		log.Fatalf("Failed to start driver:%v", err)
	}
	defer driver.Stop()

	// New Page from SeleniumDriver
	page, err := driver.NewPage()
	if err != nil {
		log.Fatalf("Failed to open page:%v", err)
		return
	}

	// Parse CSV file by parser
	inputs := parser.ParseList(file)
	if err != nil {
		log.Fatalf("Failed in parse csv file :%v", err)
		return
	}

	// run
	for i, input := range inputs {
		if err := page.Navigate(input); err != nil {
			log.Fatalf("Failed to navigate:%v", err)
		}
		page.Screenshot(fmt.Sprintf("./_tools/amex/output/%4d.png", i))
		log.Printf("OK: %d (%s)\n", i, input)
	}
}
