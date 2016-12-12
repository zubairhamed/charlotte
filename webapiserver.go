package charlotte

func NewWebApiServer() *WebApiServer {
	return &WebApiServer{}
}

type WebApiServer struct {

}

func (s *WebApiServer) Start() error {
	// TODO Startup WebServer @ Port 8080

	// For any calls, make to NATS using Request-Response
	return nil
}