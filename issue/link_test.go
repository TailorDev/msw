package issue_test

import (
	"testing"

	"github.com/TailorDev/msw/issue"
)

func TestGetBufferText(t *testing.T) {
	l := issue.Link{Name: "Hello, World", URL: "https://example.org"}

	expected := "Hello, World: https://example.org."
	if text := l.GetBufferText(); text != expected {
		t.Fatalf("Expected '%s', got: '%s'", expected, text)
	}
}

func TestGetBufferTextReplacesPlaceholders(t *testing.T) {
	l := issue.Link{
		Name:   "Hello, World",
		URL:    "https://example.org",
		Buffer: "%name% % %name%",
	}

	expected := "Hello, World % Hello, World"
	if text := l.GetBufferText(); text != expected {
		t.Fatalf("Expected '%s', got: '%s'", expected, text)
	}

	l.Buffer = "%url% == URL"
	expected = "https://example.org == URL"
	if text := l.GetBufferText(); text != expected {
		t.Fatalf("Expected '%s', got: '%s'", expected, text)
	}
}
