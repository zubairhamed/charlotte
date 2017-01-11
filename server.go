package charlotte

import (
	"github.com/NYTimes/gziphandler"
	"github.com/alecthomas/template"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func NewServer() Server {
	return &CharlotteServer{
		connectors: []Connector{},
	}
}

type CharlotteServer struct {
	connectors []Connector
	tpl        *template.Template
}

func (s *CharlotteServer) RegisterConnector(conn Connector) error {
	s.connectors = append(s.connectors, conn)

	return nil
}

func (s *CharlotteServer) Start() (err error) {
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

	// Services
	r.HandleFunc("/service/dashboards", s.handleServiceListDashboards)
	r.HandleFunc("/service/dashboard/{id}", s.handleServiceLoadDashboard)
	r.HandleFunc("/service/dashboard/{id}", s.handleServiceDeleteDashboard).Methods("DELETE")
	r.HandleFunc("/service/dashboard/{id}", s.handleServiceSaveDashboard).Methods("PUT")

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
