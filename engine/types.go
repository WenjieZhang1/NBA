package engine

// Request type
type Request struct {
	URL        string
	ParserFunc func([]byte) ParseResult
}

// ParseResult get from request parser
type ParseResult struct {
	Requests []Request
	Items    []Item
}

// Item to store to elastic search
type Item struct {
	URL     string
	Id      string
	Type    string
	Payload interface{}
}

// NilParser return empty parseresult
func NilParser([]byte) ParseResult {
	return ParseResult{}
}
