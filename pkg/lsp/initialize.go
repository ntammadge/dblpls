package lsp

type InitializeRequest struct {
	Request
	Params InitializeParams `json:"params"`
}

type InitializeParams struct {
	ClientInfo   ClientInfo         `json:"clientInfo"`
	Capabilities ClientCapabilities `json:"capabilities"`
}

type ClientInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type ClientCapabilities struct {
}

type InitializeResponse struct {
	Response
	Result InitializeResult `json:"result"`
}

type InitializeResult struct {
	Capabilities ServerCapabilities `json:"capabilities"`
	ServerInfo   ServerInfo         `json:"serverInfo"`
}

type ServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type ServerCapabilities struct {
	TextDocumentSync int `json:"textDocumentSync"`
}

func NewInitializeResponse(id int) InitializeResponse {
	return InitializeResponse{
		Response: Response{
			Jsonrpc: "2.0",
			Id:      &id,
		},
		Result: InitializeResult{
			ServerInfo: ServerInfo{
				Name:    "dblpls",
				Version: "0.1.0", // TODO: automate/parameterize this
			},
			Capabilities: ServerCapabilities{
				TextDocumentSync: 2, // Incremental
			},
		},
	}
}
