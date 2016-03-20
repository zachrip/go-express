package express

import (
	"net/http"
	"time"
	"log"
	"regexp"
)

type Listener struct {
	handler func(w http.ResponseWriter, r *http.Request)
	path string
}

type express_server struct {
	server *http.Server
	listener *express_listener
}

func (s express_server) Listen() express_server {
	s.server.ListenAndServe()

	return s;
}

func (s express_server) Get(path string, listener func (w http.ResponseWriter, r *http.Request)) express_server {
	log.Printf("Added route GET %s", path);
	
	regexpPath := "^" + path + "$"
	
	s.listener.getRoutes = append(s.listener.getRoutes, *&Listener{handler: listener, path: regexpPath});
	
	return s;
}

func (s express_server) Post(path string, listener func (w http.ResponseWriter, r *http.Request)) express_server {
	log.Printf("Added route POST %s", path);
	
	regexpPath := "^" + path + "$"
	
	s.listener.getRoutes = append(s.listener.postRoutes, *&Listener{handler: listener, path: regexpPath});
	
	return s;
}

type express_listener struct {
	getRoutes []Listener
	postRoutes []Listener
}

func(ex express_listener) ServeHTTP (w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s", r.Method, r.URL.Path);
		
	switch(r.Method) {
		case "GET":
			for _, e := range ex.getRoutes {
				match, err := regexp.MatchString(e.path, r.URL.Path)
				
				if(err != nil) {
					log.Fatal(err)
				}
				
				if(match) {
					e.handler(w, r);
					
					break;
				}
			}
		case "POST":
			for _, e := range ex.postRoutes {
				match, err := regexp.MatchString(e.path, r.URL.Path)
				
				if(err != nil) {
					log.Fatal(err)
				}
				
				if(match) {
					e.handler(w, r);
					
					break;
				}
			}
	}
}

func Server(hostname string) *express_server {
	l := &express_listener{make([]Listener, 0), make([]Listener, 0)}

	s := &http.Server{
		Addr:           hostname,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		Handler: l,
	}		
	
	n := &express_server{server: s, listener: l}
	
	return n;
}