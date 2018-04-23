// main.go
package main

import (
	"flag"
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
	inputs, err := parser.ParseCSV(file)
	if err != nil {
		log.Fatalf("Failed in parse csv file :%v", err)
		return
	}

	// run
	for _, input := range inputs {
		if err := page.Navigate(input.URL); err != nil {
			log.Fatalf("Failed to navigate:%v", err)
		}
		// <p> 要素を取得
		xpath := page.AllByXPath("//p")
		// 力技でアイテムの整合性を確認
		item, err := xpath.At(1).Text()
		if err != nil || item != input.Item {
			log.Fatalf("Mismatch Item: Value = %#v, Error = %v", item, err)
		}
		// 力技で店舗の整合性を確認
		name, err := xpath.At(3).Text()
		if err != nil || name != "店舗名:"+input.Name {
			log.Fatalf("Mismatch Name: Value = %#v, Error = %v", name, err)
		}
		// 力技で期限の整合性を確認
		limit, err := xpath.At(6).Text()
		if err != nil || limit != "引換期限"+input.Limit {
			log.Fatalf("Mismatch Limit: Value = %#v, Error = %v", limit, err)
		}
		page.Screenshot("./_tools/output/" + input.Name + ".png")
		log.Printf("OK: %s ( %s )\n", input.Name, input.URL)
	}
}
