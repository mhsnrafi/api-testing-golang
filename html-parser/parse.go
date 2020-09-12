package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func html_parser(str string){
	// Request the HTML page.
	res, err := http.Get(str)
	if err != nil{
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	// Load the HTML document
	html, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	// show the HTML code as a string %s
	fmt.Printf("%s\n", html)
	
	pageContent := string(html)
	pageTitle := getTitle(pageContent)

	//Calling Page Title
	fmt.Printf("Page title: %s\n", pageTitle)

	//Calling Html Version
	htmlVersion := getHtmlVersion(pageContent)
	fmt.Println(htmlVersion)


}

func getTitle(pageContent string) []byte{
	// Find a title
	titleStartIndex := strings.Index(pageContent, "<title>")
	if titleStartIndex == -1 {
		fmt.Println("No title element found")
		os.Exit(0)
	}
	// The start index of the title is the index of the first
	// character, the < symbol. We don't want to include
	// <title> as part of the final value, so let's offset
	// the index by the number of characers in <title>
	titleStartIndex += 7

	// Find the index of the closing tag
	titleEndIndex := strings.Index(pageContent, "</title>")
	if titleEndIndex == -1 {
		fmt.Println("No closing tag for title found.")
		os.Exit(0)
	}

	// (Optional)
	// Copy the substring in to a separate variable so the
	// variables with the full document data can be garbage collected
	pageTitle := []byte(pageContent[titleStartIndex:titleEndIndex])

	// Print out the result
	return pageTitle

}

func getHtmlVersion(pageContent string) string{
	doctTypeMap := make(map[string]string)

	//Html Versions Declaration
	doctTypeMap["html5"] = "<!doctype html>"
	doctTypeMap["HTML4.01-Strict"] = "<!doctype html public \"-//w3c//dtd html 4.01//en\">"
	doctTypeMap["HTML4.01-Transitional"] = "<!doctype html public \"-//w3c//dtd html 4.01 transitional//en\">"
	doctTypeMap["HTML4.01-Frameset"] = "<!doctype html public \"-//w3c//dtd html  4.01 frameset//en\">"

	for key,value := range doctTypeMap{
		if strings.Contains(strings.ToLower(pageContent), value){
			return key
		}
	}
	return "No version found"
}
func main() {
	url := "https://www.w3.org/TR/html401/"
	html_parser(url)

}