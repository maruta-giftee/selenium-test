// main.go
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/maruta-giftee/selenium-test/parser"
	"github.com/maruta-giftee/selenium-test/selenium"
	"github.com/maruta-giftee/selenium-test/selenium/chrome"
)

type scenario struct {
	id  int
	url string
}

func worker(scCh <-chan scenario, drvCh chan selenium.Driver) {
	for sc := range scCh {
		// New SeleniumDriver
		driver := chrome.NewDriver()
		if err := driver.Start(); err != nil {
			log.Fatalf("Failed to start driver: %v", err)
		}

		// New Page from SeleniumDriver
		page, err := driver.NewPage()
		if err != nil {
			log.Fatalf("Failed to open page:%v", err)
			return
		}
		if err := page.Navigate(sc.url); err != nil {
			log.Fatalf("Failed to navigate:%v", err)
		}
		page.Screenshot(fmt.Sprintf("./_tools/amex/output/%4d.png", sc.id))
		log.Printf("OK: %d (%s)\n", sc.id, sc.url)
		drvCh <- driver
	}
}

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

	// Parse file by parser
	inputs := parser.ParseList(file)
	if err != nil {
		log.Fatalf("Failed in parse csv file :%v", err)
		return
	}

	var wg sync.WaitGroup
	scChan := make(chan scenario)
	defer close(scChan)

	drvChan := make(chan selenium.Driver)
	defer close(drvChan)

	go func() {
		for {
			driver := <-drvChan
			driver.Stop()
			wg.Done()
		}
	}()

	for i := 1; i <= 15; i++ {
		go worker(scChan, drvChan)
	}

	// run
	for i, input := range inputs {
		wg.Add(1)
		scChan <- scenario{i, input}
	}
	wg.Wait()
}
