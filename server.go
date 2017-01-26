package charlotte

import (
	"encoding/json"
	"github.com/NYTimes/gziphandler"
	"github.com/alecthomas/template"
	"github.com/gorilla/mux"
	"github.com/nats-io/go-nats"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func NewServer() Server {
	return &CharlotteServer{
		connectors: []Connector{},
	}
}

type CharlotteServer struct {
	connectors []Connector
	tpl        *template.Template
	nc         *nats.Conn
}

func (s *CharlotteServer) RegisterConnector(conn Connector) error {
	s.connectors = append(s.connectors, conn)

	return nil
}

func (s *CharlotteServer) Start() (err error) {

	nc, err := CreateNatsClient(nats.DefaultURL)
	s.nc = nc

	for _, v := range s.connectors {
		go v.Start()
	}

	s.startDashboard()

	return nil
}

func (s *CharlotteServer) createTemplateFuncMap() template.FuncMap {
	var templateMap = template.FuncMap{}
	templateMap["eq"] = func(a, b interface{}) bool {
		return a == b
	}
	return templateMap
}

func (s *CharlotteServer) startDashboard() {
	templateMap := s.createTemplateFuncMap()

	var tpl = template.New("").Funcs(templateMap).Delims("#{", "}#")
	s.tpl = tpl

	err := filepath.Walk("./webapp/tpl", func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".html") {
			_, err = tpl.ParseFiles(path)
			if err != nil {
				log.Println(err)
			}
		}
		return err
	})

	if err != nil {
		panic(err.Error())
	}

	r := mux.NewRouter()

	// r.HandleFunc("/api/relations/{id}/{levels}", s.handleGetThingsRelations)
	r.HandleFunc("/", s.handleViewDashboard)
	r.HandleFunc("/?{dashboard}", s.handleViewDashboard)
	r.HandleFunc("/things", s.handleViewThings)
	r.HandleFunc("/things/add", s.handleViewAddThings)
	r.HandleFunc("/auth", s.handleViewAuthentication)
	r.HandleFunc("/connectors", s.handleViewConnectors)
	r.HandleFunc("/messages", s.handleViewMessages)
	r.HandleFunc("/settings", s.handleViewSettings)

	// Services
	r.HandleFunc("/service/dashboards", s.handleServiceListDashboards)
	r.HandleFunc("/service/dashboard/{id}", s.handleServiceLoadDashboard)
	r.HandleFunc("/service/dashboard/{id}", s.handleServiceDeleteDashboard).Methods("DELETE")
	r.HandleFunc("/service/dashboard/{id}", s.handleServiceSaveDashboard).Methods("PUT")

	// Things Service
	/*
		/services/things 		GET 	list
		/services/things		POST	create new
		/services/things/:id	GET		get by id
		/services/things/:id	PUT		update by id
		/services/things/:id	DELETE	delete by id

	*/
	r.HandleFunc("/service/things", s.handleServiceThingsList).Methods("GET")
	r.HandleFunc("/service/things/{id}", s.handleServiceThingsAdd).Methods("POST")
	r.HandleFunc("/service/things/{id}", s.handleServiceThingsGet).Methods("GET")
	r.HandleFunc("/service/things/{id}", s.handleServiceThingsUpdate).Methods("PUT")
	r.HandleFunc("/service/things/{id}", s.handleServiceThingsDelete).Methods("DELETE")

	// Connector Service
	r.HandleFunc("/service/connectors", s.handleServiceConnectorsList).Methods("GET")
	r.HandleFunc("/service/connectors/{id}", s.handleServiceConnectorGet).Methods("GET")
	r.HandleFunc("/service/connectors/{id}", s.handleServiceConnectorUpdate).Methods("PUT")

	// Auth Service
	r.HandleFunc("/service/auth", s.handleServiceAuthList).Methods("GET")
	r.HandleFunc("/service/auth", s.handleServiceAuthAdd).Methods("POST")
	r.HandleFunc("/service/auth/{id}", s.handleServiceAuthGet).Methods("GET")
	r.HandleFunc("/service/auth/{id}", s.handleServiceAuthUpdate).Methods("PUT")
	r.HandleFunc("/service/auth/search", s.handleServiceAuthSearch)

	// Messages Service
	// List
	// Get
	// Search

	// Settings Service
	// Get
	// Update

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./webapp/static/")))
	http.Handle("/", gziphandler.GzipHandler(r))

	log.Println("Running Dashboard on :8010. Ready to rock!")
	http.ListenAndServe(":8010", nil)
}

func (s *CharlotteServer) renderTemplate(w http.ResponseWriter, tmpl string, p interface{}) {
	w.Header().Set("Content-Type", "text/html")
	err := s.tpl.ExecuteTemplate(w, tmpl, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *CharlotteServer) handleViewDashboard(rw http.ResponseWriter, req *http.Request) {
	m := &PageModel{
		Name: "home",
	}
	s.renderTemplate(rw, "home", m)
}

func (s *CharlotteServer) handleViewThings(rw http.ResponseWriter, req *http.Request) {
	m := &PageModel{
		Name: "things",
	}
	s.renderTemplate(rw, "things", m)
}

func (s *CharlotteServer) handleViewAddThings(rw http.ResponseWriter, req *http.Request) {
	m := &PageModel{
		Name: "things-add",
	}
	s.renderTemplate(rw, "things-add", m)
}

func (s *CharlotteServer) handleViewAuthentication(rw http.ResponseWriter, req *http.Request) {
	m := &PageModel{
		Name: "auth",
	}
	s.renderTemplate(rw, "auth", m)
}

func (s *CharlotteServer) handleViewConnectors(rw http.ResponseWriter, req *http.Request) {
	m := &PageModel{
		Name: "connectors",
	}
	s.renderTemplate(rw, "connectors", m)
}

func (s *CharlotteServer) handleViewMessages(rw http.ResponseWriter, req *http.Request) {

	m := &PageModel{
		Name: "messages",
	}
	s.renderTemplate(rw, "messages", m)
}

func (s *CharlotteServer) handleViewSettings(rw http.ResponseWriter, req *http.Request) {

	m := &PageModel{
		Name: "settings",
	}
	s.renderTemplate(rw, "settings", m)
}

// Services
func (s *CharlotteServer) handleServiceListDashboards(rw http.ResponseWriter, req *http.Request) {
	log.Println("List Dashboards")
}

func (s *CharlotteServer) handleServiceSaveDashboard(rw http.ResponseWriter, req *http.Request) {
	log.Println("Save Dashboard")
}

func (s *CharlotteServer) handleServiceLoadDashboard(rw http.ResponseWriter, req *http.Request) {
	log.Println("Load Dashboard")
}

func (s *CharlotteServer) handleServiceDeleteDashboard(rw http.ResponseWriter, req *http.Request) {
	log.Println("Delete Dashboard")
}

func (s *CharlotteServer) handleServiceThingsList(rw http.ResponseWriter, req *http.Request) {

	things := []*ThingModel{}

	t := &ThingModel{}
	t.Id = "abc123"
	t.Created = time.Now().UnixNano() / 1000000
	t.LastUpdated = time.Now().UnixNano() / 1000000
	t.Schema = "{}"
	t.State = "{}"

	things = append(things, t)

	s.writeJsonModel(things, rw)
}

func (s *CharlotteServer) handleServiceThingsAdd(rw http.ResponseWriter, req *http.Request) {
	log.Println("Add Thing Instance")

	var msg map[string]interface{}
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&msg)
	if err != nil {
		err.Error()
	}

	// TODO: Call request-response to add thing
	/*
		// Requests
		var response string
		err := c.Request("help", "help me", &response, 10*time.Millisecond)
		if err != nil {
			fmt.Printf("Request failed: %v\n", err)
		}

		// Replying
		c.Subscribe("help", func(subj, reply string, msg string) {
			c.Publish(reply, "I can help!")
		})
	 */

	log.Println(msg["id"])
	log.Println(msg["Schema"])
	log.Println(msg["State"])
}

func (s *CharlotteServer) handleServiceThingsGet(rw http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]

	log.Println("Get Thing", id)

	t := &ThingModel{}
	t.Id = "abc123"
	t.Created = time.Now().UnixNano() / 1000000
	t.LastUpdated = time.Now().UnixNano() / 1000000
	t.Schema = "{}"
	t.State = "{}"

	// TODO: Send GET Request to ThingRegistry

	s.writeJsonModel(t, rw)
}

func (s *CharlotteServer) handleServiceThingsUpdate(rw http.ResponseWriter, req *http.Request) {
	log.Println("Update Thing Instance")
}

func (s *CharlotteServer) handleServiceThingsDelete(rw http.ResponseWriter, req *http.Request) {

}

func (s *CharlotteServer) handleServiceConnectorsList(rw http.ResponseWriter, req *http.Request) {

}

func (s *CharlotteServer) handleServiceConnectorGet(rw http.ResponseWriter, req *http.Request) {

}

func (s *CharlotteServer) handleServiceConnectorUpdate(rw http.ResponseWriter, req *http.Request) {

}

func (s *CharlotteServer) handleServiceAuthList(rw http.ResponseWriter, req *http.Request) {

}

func (s *CharlotteServer) handleServiceAuthAdd(rw http.ResponseWriter, req *http.Request) {

}

func (s *CharlotteServer) handleServiceAuthGet(rw http.ResponseWriter, req *http.Request) {

}

func (s *CharlotteServer) handleServiceAuthUpdate(rw http.ResponseWriter, req *http.Request) {

}

func (s *CharlotteServer) handleServiceAuthSearch(rw http.ResponseWriter, req *http.Request) {

}

func (s *CharlotteServer) writeJsonModel(m interface{}, rw http.ResponseWriter) {
	j, _ := json.Marshal(m)
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(j)

}
