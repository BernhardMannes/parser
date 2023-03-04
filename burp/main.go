package main

import (
	"burp/burp"
	"fmt"
	"log"
	"os"
)

func getIssuesFromFile(fileName string) *burp.BurpRun {
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	issues, err := burp.ParseBurp(data)
	if err != nil {
		log.Fatal(err)
	}

	return issues
}

func main() {
	fmt.Println("Burp Module to read saved Burp requests in XML format")

	issues := getIssuesFromFile("./burp/testfiles/1_request.xml")
	issues = getIssuesFromFile("./burp/testfiles/3_requests.xml")

	print(issues.Items[0].GetRequest())

	println(len(issues.Items))
}
