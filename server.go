package charlotte

func NewServer() Server {
	return &CharlotteServer{
		connectors: []Connector{},
	}
}

type CharlotteServer struct {
	connectors []Connector
}

func (s *CharlotteServer) RegisterConnector(conn Connector) error {
	s.connectors = append(s.connectors, conn)

	return nil
}

func (s *CharlotteServer) Start() (err error) {
	for _, v := range s.connectors {
		go v.Start()
	}
	return nil
}
