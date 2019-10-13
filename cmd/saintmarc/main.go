// main.go
package main

import (
	"flag"
	"log"
	"os"
	"strconv"

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
		log.Fatalf("Failed to start driver: %v", err)
	}
	defer driver.Stop()

	// New Page from SeleniumDriver
	page, err := driver.NewPage()
	if err != nil {
		log.Fatalf("Failed to open page: %v", err)
		return
	}

	// Parse CSV file by parser
	inputs, err := parser.ParseCSV(file)
	if err != nil {
		log.Fatalf("Failed in parse csv file: %v", err)
		return
	}

	// run
	for i, input := range inputs {
		if i%50 > 0 {
			continue
		}

		if err := page.Navigate(input.URL); err != nil {
			log.Fatalf("Failed to navigate: %v", err)
		}
		// <p> 要素を取得
		xpath := page.AllByClass("gjn-textare")
		// 力技で期限の整合性を確認
		textare, err := xpath.At(0).Text()
		if err != nil || textare != "アンケートにお答えいただきありがとうございます。" {
			log.Printf("Mismatch textare: Value = %#v, Error = %v", textare, err)
			page.Screenshot("./_tools/saintmarc/output/" + strconv.Itoa(i) + ".png")
			log.Fatalf("URL: %v", input.URL)
		}
		log.Printf("OK: %s ( %s )\n", input.Name, input.URL)
	}
}
