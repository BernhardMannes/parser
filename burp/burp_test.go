package burp

import (
	"log"
	"os"
	"testing"
)

func getIssuesFromFile(fileName string) *BurpRun {
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	issues, err := ParseBurp(data)
	if err != nil {
		log.Fatal(err)
	}

	return issues
}

func TestOneRequest(t *testing.T) {
	issues := getIssuesFromFile("./testfiles/1_request.xml")
	//issues = getIssuesFromFile("./burp/testfiles/3_requests.xml")
	println(len(issues.Items))
}

/*
func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b`+name+`\b`)
	msg, err := Hello("Gladys")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

*/
