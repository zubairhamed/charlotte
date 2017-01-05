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
	r.HandleFunc("/", s.handleDashboard)
	r.HandleFunc("/?{dashboard}", s.handleDashboard)

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./webapp/static/")))
	http.Handle("/", gziphandler.GzipHandler(r))

	log.Println("Running Dashboard on :8010. Ready to rock!")
	http.ListenAndServe(":8010", nil)

	/*
		/
		/?<dashboardname>
		/things/?<id>
		/things/add
		/connectors/?<id>
	*/
}

func (s *CharlotteServer) renderTemplate(w http.ResponseWriter, tmpl string, p interface{}) {
	w.Header().Set("Content-Type", "text/html")
	err := s.tpl.ExecuteTemplate(w, tmpl, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *CharlotteServer) handleDashboard(rw http.ResponseWriter, req *http.Request) {
	m := &PageModel{
		Name: "home",
	}

	s.renderTemplate(rw, "home", m)
}

type PageModel struct {
	Name string
}
