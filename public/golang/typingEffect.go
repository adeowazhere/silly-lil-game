package typingEffect

import (
	"fmt"
	"os"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func getFile(HTMLFileName string) (*goquery.Document, error) {
	HTMLFile, err := os.Open(HTMLFileName)
	if err != nil {
		fmt.Println("Error: ", err)
		return nil, err
	}
	defer HTMLFile.Close()

	doc, err := goquery.NewDocumentFromReader(HTMLFile)
	if err != nil {
		fmt.Println("Error: ", err)
		return nil, err
	}
	return doc, nil
}

func applyStylingEffect(HTMLFileName string) {
	Document, err := getFile(HTMLFileName)

	time.Sleep(2 * time.Second)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	typedDialogueLine := Document.Find(".text-typing")
	DialogueLineToBeTyped := typedDialogueLine.Next()

	if DialogueLineToBeTyped.Length() == 0 {
		return
	}

}
