package charlotte

func NewWebFunctionsServer() *WebFunctionsServer {
	return &WebFunctionsServer{}
}

type WebFunctionsServer struct {
}

func (s *WebFunctionsServer) Start() error {
	return nil
}
