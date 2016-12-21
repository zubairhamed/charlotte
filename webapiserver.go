package charlotte

func NewWebApiServer() *WebApiServer {
	return &WebApiServer{
		webinterfaces: []WebInterface{},
	}
}

type WebApiServer struct {
	webinterfaces []WebInterface
}

func (s *WebApiServer) RegisterWebInterface(i WebInterface) error {
	s.webinterfaces = append(s.webinterfaces, i)

	return nil
}

func (s *WebApiServer) Start() error {
	for _, v := range s.webinterfaces {
		go v.Start()
	}

	// For any calls, make to NATS using Request-Response
	return nil
}
