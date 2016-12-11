package charlotte

type Connector interface {
	Start() error
	Stop() error
}

type ConnectorManager interface {
	Register(Connector) error
	Unregister(Connector) error
}

type Server interface {
	RegisterConnector(Connector) error
	Start() error
}