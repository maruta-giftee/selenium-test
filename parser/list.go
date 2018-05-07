package parser

import (
	"bufio"
	"os"
)

func ParseList(file *os.File) []string {
	// Scannerで読み込む
	// lines := []string{}
	lists := make([]string, 0, 1000) // ある程度多めに取っとく
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// appendで追加
		lists = append(lists, scanner.Text())
	}
	return lists
}
