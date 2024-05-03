package lsp

type Request struct {
	// Each request has its own params object
	Jsonrpc string `json:"jsonrpc"` // Should always be "2.0"
	Id      int    `json:"id"`
	Method  string `json:"method"`
}

type Response struct {
	// Each response has its own result object
	Jsonrpc string `json:"jsonrpc"`      // Should always be "2.0"
	Id      *int   `json:"id,omitempty"` // Should be int or string on request response, or null on notification response
}

type ResponseError struct {
	Code    int    `json:"code"` // A number indicating the error type that occurred
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"` // A primitive or structured value that contains additional information about the error. Can be omitted
}

type Notification struct {
	// Each notification has its own params object
	Jsonrpc string `json:"jsonrpc"` // Should always be "2.0"
	Method  string `json:"method"`
}
