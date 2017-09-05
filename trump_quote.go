package trumpQuote

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
)

const (
	ENDPOINT           = "https://api.whatdoestrumpthink.com/api/"
	QUOTE_RANDOM       = "v1/quotes/random"
	QUOTE_PERSONALIZED = "v1/quotes/personalized?q="
	QUOTES             = "v1/quotes"
	VALIDATE_PATTERN   = "[\\w\\s]+"
)

type Nattributes []string

type Quote struct {
	Message          string                 `json:"message"`
	NLPAttributesRaw map[string]interface{} `json:"nlp_attributes"`
	NLPAttributes    struct {
		QuoteStructure []Nattributes
	}
}

type QuotePersonalized struct {
	Message          string                 `json:"message"`
	Nickname         string                 `json:"nickname"`
	NLPAttributesRaw map[string]interface{} `json:"nlp_attributes"`
	NLPAttributes    struct {
		Pronoun        string
		QuoteStructure []Nattributes
	}
}

type Quotes struct {
	Messages struct {
		Personalized    []interface{} `json:"personalized"`
		NonPersonalized []interface{} `json:"non_personalized"`
	} `json:"messages"`
}

func callAPI(url string, target interface{}) error {
	//two parts: call the API using the requested url, and
	// decode the Json to get a compatible structure
	//The specific calling methods will take care of a
	//more specific conversion
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(target)
	if err != nil {
		return err
	}
	return nil

}

func convertNattributes(nattrs *[]Nattributes, raw []interface{}) {

	//convert the Raw NLPAttributes into the correct type/structure
	for _, v := range raw {
		var nlpattr Nattributes
		for _, attr := range v.([]interface{}) {
			nlpattr = append(nlpattr, attr.(string))
		}
		*nattrs = append(*nattrs, nlpattr)
	}
}

func NewRandom() (*Quote, error) {
	url := fmt.Sprintf("%s%s", ENDPOINT, QUOTE_RANDOM)
	quote := &Quote{}
	err := callAPI(url, quote)
	if err != nil {
		return nil, err
	}
	attrs, _ := quote.NLPAttributesRaw["quote_structure"].([]interface{})
	convertNattributes(&quote.NLPAttributes.QuoteStructure, attrs)

	return quote, nil
}

func NewPersonalized(s string) (*QuotePersonalized, error) {
	//Todo validate s
	if _, err := regexp.MatchString(VALIDATE_PATTERN, s); err != nil {
		return nil, errors.New("Not valid Nickname provided")
	}
	url := fmt.Sprintf("%s%s%s", ENDPOINT, QUOTE_PERSONALIZED, s)
	quotePersonalized := &QuotePersonalized{}
	err := callAPI(url, quotePersonalized)
	if err != nil {
		return nil, err
	}

	quotePersonalized.NLPAttributes.Pronoun, _ = quotePersonalized.NLPAttributesRaw["pronoun"].(string)
	attrs, _ := quotePersonalized.NLPAttributesRaw["quote_structure"].([]interface{})
	convertNattributes(&quotePersonalized.NLPAttributes.QuoteStructure, attrs)
	return quotePersonalized, nil
}

func All() (*Quotes, error) {

	url := fmt.Sprintf("%s%s", ENDPOINT, QUOTES)
	quotes := &Quotes{}
	err := callAPI(url, quotes)
	if err != nil {
		return nil, err
	}
	return quotes, nil
}
