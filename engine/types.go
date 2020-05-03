package engine

// Request type
type Request struct {
	URL        string
	ParserFunc func([]byte) ParseResult
}

// ParseResult get from requst parser
type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

// NilParser return empty parseresult
func NilParser([]byte) ParseResult {
	return ParseResult{}
}
