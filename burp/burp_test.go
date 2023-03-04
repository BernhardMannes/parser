package burp

import (
	"github.com/google/go-cmp/cmp"
	"os"
	"strings"
	"testing"
)

func TestWrongInput(t *testing.T) {
	run, err := Parse([]byte("test"))
	if err == nil || err.Error() != "EOF" {
		t.Fatalf("received not the expected EOF error message")
	}

	if run.Items != nil {
		t.Fatalf("no items expected")
	}
}

func TestOneRequest(t *testing.T) {
	burpRun, _ := ParseFromFile("./test/1_request.xml")
	if len(burpRun.Items) != 1 {
		t.Fatalf("Expected one item")
	}

	expected := &Run{
		Items: []Item{
			{
				Time:     "Thu Mar 02 20:06:44 CET 2023",
				Url:      "http://10.129.204.197/api/check-username.php?u=maria",
				Host:     "10.129.204.197",
				Port:     80,
				Protocol: "http",
				Method:   "GET",
				Path:     "/api/check-username.php?u=maria",
				Ext:      "php",
				Request: Request{
					Base64: true,
					Data:   "R0VUIC9hcGkvY2hlY2stdXNlcm5hbWUucGhwP3U9bWFyaWEgSFRUUC8xLjENCkhvc3Q6IDEwLjEyOS4yMDQuMTk3DQpVc2VyLUFnZW50OiBNb3ppbGxhLzUuMCAoV2luZG93cyBOVCAxMC4wOyBXaW42NDsgeDY0KSBBcHBsZVdlYktpdC81MzcuMzYgKEtIVE1MLCBsaWtlIEdlY2tvKSBDaHJvbWUvMTEwLjAuNTQ4MS43OCBTYWZhcmkvNTM3LjM2DQpBY2NlcHQ6ICovKg0KUmVmZXJlcjogaHR0cDovLzEwLjEyOS4yMDQuMTk3L3NpZ251cC5waHANCkFjY2VwdC1FbmNvZGluZzogZ3ppcCwgZGVmbGF0ZQ0KQWNjZXB0LUxhbmd1YWdlOiBlbi1VUyxlbjtxPTAuOQ0KQ29ubmVjdGlvbjogY2xvc2UNCg0K",
				},
				Status:         200,
				ResponseLength: 213,
				MimeType:       "JSON",
				Response: Response{
					Base64: true,
					Data:   "SFRUUC8xLjEgMjAwIE9LDQpEYXRlOiBUaHUsIDAyIE1hciAyMDIzIDE5OjA3OjI1IEdNVA0KU2VydmVyOiBBcGFjaGUvMi40LjU0IChXaW42NCkgUEhQLzguMS4xMw0KWC1Qb3dlcmVkLUJ5OiBQSFAvOC4xLjEzDQpDb250ZW50LUxlbmd0aDogMTgNCkNvbm5lY3Rpb246IGNsb3NlDQpDb250ZW50LVR5cGU6IGFwcGxpY2F0aW9uL2pzb24NCg0KeyJzdGF0dXMiOiJ0YWtlbiJ9",
				},
				Comment: "",
			},
		},
	}

	assertEquals(t, expected, burpRun)

	data, _ := os.ReadFile("./test/expectedDecodedRequest.txt")
	assertEquals(t, string(data), strings.ReplaceAll(burpRun.Items[0].GetRequest(), "\r\n", "\n"))

	data, _ = os.ReadFile("./test/expectedDecodedResponse.txt")
	assertEquals(t, string(data), strings.ReplaceAll(burpRun.Items[0].GetResponse(), "\r\n", "\n"))
}

func TestOneRequestNotEncoded(t *testing.T) {
	_, err := ParseFromFile("./test/1_request_not_encoded.xml")

	if err.Error() != "xml: unsupported version \"1.1\"; only version 1.0 is supported" {
		t.Fatalf("Expected only xml 1.0 is supported message")
	}
}

func assertEquals(t *testing.T, expected, actual interface{}) {
	if !cmp.Equal(expected, actual) {
		print(cmp.Diff(expected, actual))
		t.Fatalf("Difference detedted")
	}
}

func TestParseByWrongFileName(t *testing.T) {
	burpRun, err := ParseFromFile("wrongFileName.xml")
	if err == nil {
		t.Fatalf("No error received")
	}
	if err.Error() != "open wrongFileName.xml: no such file or directory" {
		t.Fatalf("Wrong error received: " + err.Error())
	}

	if burpRun != nil {
		t.Fatalf("No item expected")
	}
}

func TestParseByFileName(t *testing.T) {
	burpRun, err := ParseFromFile("./test/1_request.xml")
	if err != nil {
		t.Fatalf("No error expected, but received: " + err.Error())
	}

	if burpRun == nil || burpRun.Items == nil || len(burpRun.Items) != 1 {
		t.Fatalf("Not received the expected content")
	}
}
