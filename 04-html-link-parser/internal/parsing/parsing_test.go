package parsing

import (
	"bytes"
	"os"
	"reflect"
	"testing"
)

func TestTesting(t *testing.T) {
	num1 := 2
	num2 := 2
	want := 4
	result := num1 + num2
	if result != want {
		t.Fatalf("Sum of %v and %v = %v, want %v", num1, num2, result, want)
	}
}

func TestEx1(t *testing.T) {
	htmlData, _ := os.ReadFile("test_data/ex1.html")
	want := []Link{{Href: "/other-page", Text: "A link to another page"}}

	r := bytes.NewReader(htmlData)
	result, err := ParseLinks(r)
	if err != nil {
		t.Fatalf("Error parsing links: %v", err)
	}
	if !reflect.DeepEqual(result, want) {
		t.Fatalf("ParseLinks(ex1) = %v, want %v", result, want)
	}
}

func TestEx2(t *testing.T) {
	htmlData, _ := os.ReadFile("test_data/ex2.html")
	want := []Link{
		{
			Href: "https://www.twitter.com/joncalhoun",
			Text: "Check me out on twitter",
		},
		{
			Href: "https://github.com/gophercises",
			Text: "Gophercises is on Github!",
		},
	}

	r := bytes.NewReader(htmlData)
	result, err := ParseLinks(r)
	if err != nil {
		t.Fatalf("Error parsing links: %v", err)
	}
	if !reflect.DeepEqual(result, want) {
		t.Fatalf("ParseLinks(ex2) = %v, want %v", result, want)
	}
}

func TestEx3(t *testing.T) {
	htmlData, _ := os.ReadFile("test_data/ex3.html")
	want := []Link{
		{
			Href: "#",
			Text: "Login",
		},
		{
			Href: "/lost",
			Text: "Lost? Need help?",
		},
		{
			Href: "https://twitter.com/marcusolsson",
			Text: "@marcusolsson",
		},
	}

	r := bytes.NewReader(htmlData)
	result, err := ParseLinks(r)
	if err != nil {
		t.Fatalf("Error parsing links: %v", err)
	}
	if !reflect.DeepEqual(result, want) {
		t.Fatalf("ParseLinks(ex3) = %v, want %v", result, want)
	}
}

func TestEx4(t *testing.T) {
	htmlData, _ := os.ReadFile("test_data/ex4.html")
	want := []Link{
		{
			Href: "/dog-cat",
			Text: "dog cat",
		},
	}

	r := bytes.NewReader(htmlData)
	result, err := ParseLinks(r)
	if err != nil {
		t.Fatalf("Error parsing links: %v", err)
	}
	if !reflect.DeepEqual(result, want) {
		t.Fatalf("ParseLinks(ex4) = %v, want %v", result, want)
	}
}
