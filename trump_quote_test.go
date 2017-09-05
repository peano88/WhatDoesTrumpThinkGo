package trumpQuote

import (
	"fmt"
	"regexp"
	"testing"
)

func TestQuoteRandom(t *testing.T) {
	quote, err := NewRandom()
	if err != nil {
		message := fmt.Sprintf("No error should have been raised but: %s", err)
		t.Fatalf(message)
	}

	if len(quote.Message) == 0 {
		t.Error("No Random quote received")
	}

	if len(quote.NLPAttributes.QuoteStructure) == 0 {
		t.Error("No attributes received")
	}

	for _, attrs := range quote.NLPAttributes.QuoteStructure {
		if len(attrs) == 0 {
			t.Error(" EMpty attributes Transfered")
		}
	}
}
func TestQuotePesonalizedBadName(t *testing.T) {

	_, err := NewPersonalized("&$X")
	if err == nil {
		t.Fatalf("A bad Name Error should have been raised")
	}
}

func TestQuotePersonalized(t *testing.T) {

	quote, err := NewPersonalized("XYZ")
	if err != nil {
		message := fmt.Sprintf("No error should have been raised but: %s", err)
		t.Fatalf(message)
	}

	if len(quote.Message) == 0 {
		t.Error("No Random quote received")
	}
	// Check if personalized preference worked
	reg := regexp.MustCompile("XYZ")
	if !reg.MatchString(quote.Message) {
		t.Errorf("Not a proper personalized")
	}

	if len(quote.NLPAttributes.QuoteStructure) == 0 {
		t.Error("No attributes received")
	}

	for _, attrs := range quote.NLPAttributes.QuoteStructure {
		if len(attrs) == 0 {
			t.Error(" EMpty attributes Transfered")
		}
	}
}

func TestQuotesAll(t *testing.T) {
	quotes, err := All()
	if err != nil {
		message := fmt.Sprintf("No error should have been raised but: %s", err)
		t.Fatalf(message)
	}

	if &quotes.Messages == nil {
		t.Fatalf("No messages received")
	}

	if len(quotes.Messages.Personalized) == 0 {
		t.Error("No Personalized messages received")
	}

	if len(quotes.Messages.NonPersonalized) == 0 {
		t.Error("No Non-Personalized messages received")
	}
}
