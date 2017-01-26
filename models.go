package charlotte

type PageModel struct {
	Name string
}

type DashboardEntryModel struct {
	Name    string
	Content string
}

type AuthenticationEntryModel struct {
}

type ConnectorEntryModel struct {
}

type ThingModel struct {
	Id          string `json:"id"`
	Created     int64  `json:"created"`
	LastUpdated int64  `json:"lastUpdated"`
	Schema      string `json:"schema"`
	State       string `json:"state"`
}

type MessageEntryModel struct {
}
