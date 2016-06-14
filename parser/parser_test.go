package parser_test

import (
	"testing"

	"github.com/TailorDev/msw/parser"
)

func TestParse(t *testing.T) {
	issue, err := parser.Parse("../fixtures/2016-10-13.yml")
	if err != nil {
		t.Fatalf("TestParse: %s", err)
	}

	if issue.Number != 123 {
		t.Fatalf("Expected 123, got %d", issue.Number)
	}

	if issue.Date.Format("2006-01-02") != "2016-10-13" {
		t.Fatalf("Expected 2016-10-13, got %s", issue.Date)
	}

	if issue.WelcomeText != "Hello, World!\n" {
		t.Fatalf("Expected 'Hello, World!\\n', got %s", issue.WelcomeText)
	}

	if len(issue.Categories) != 1 {
		t.Fatalf("Expected 3 categories, got: %d", len(issue.Categories))
	}

	if title := issue.Categories[0].Title; title != "Cat. 1" {
		t.Fatalf("Expected 'Cat. 1', got: %s", title)
	}

	if nbLinks := len(issue.Categories[0].Links); nbLinks != 1 {
		t.Fatalf("Expected no links, got: %d", nbLinks)
	}

	link := issue.Categories[0].Links[0]

	if link.Name != "Link #1" {
		t.Fatalf("Expected 'Link #1', got: %s", link.Name)
	}

	if link.URL != "http://example.org" {
		t.Fatalf("Expected 'http://example.org', got: %s", link.URL)
	}

	if link.Abstract != "This is the abstract of the first link.\n" {
		t.Fatalf("Expected 'This is the abstract of the first link.\\n', got: %s", link.Abstract)
	}
}

func TestParseUnexistentFile(t *testing.T) {
	_, err := parser.Parse("foobar.yml")
	if err == nil {
		t.Fatalf("TestParseUnexistentFile should return an error")
	}
}

func TestParseInvalidFilename(t *testing.T) {
	_, err := parser.Parse("parser_test.go")
	if err == nil {
		t.Fatalf("TestParseInvalidFilename should return an error")
	}

	if err.Error() != "Invalid filename (parser_test.go), should match 'YYYY-MM-DD.yml'" {
		t.Fatalf("Got: %s", err.Error())
	}
}