package service

import (
    "net/http"
    "os"

    "github.com/codegangsta/negroni"
    "github.com/gorilla/mux"
    "github.com/unrolled/render"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {

    formatter := render.New(render.Options{
		Directory:  "templates",
        Extensions: []string{".html"},
        IndentJSON: true,
    })

    n := negroni.Classic()
    mx := mux.NewRouter()

    initRoutes(mx, formatter)

    n.UseHandler(mx)
    return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
    webRoot := os.Getenv("WEBROOT")
    if len(webRoot) == 0 {
        if root, err := os.Getwd(); err != nil {
            panic("Could not retrive working directory")
        } else {
            webRoot = root
            //fmt.Println(root)
        }
    }

    // test 1
    mx.PathPrefix("/assets").Handler(http.StripPrefix("/assets", http.FileServer(http.Dir(webRoot+"/assets/"))))
    // test 2
    mx.HandleFunc("/api/test", apiTestHandler(formatter)).Methods("GET")
    // test 3
    mx.HandleFunc("/templates", homeHandler(formatter)).Methods("GET")
    //test 4
    mx.HandleFunc("/login", loginHandler(formatter)).Methods("GET")
    mx.HandleFunc("/login", tableHandler(formatter)).Methods("POST")
    //test 5
    mx.HandleFunc("/unknown", unknown()).Methods("GET")
    mx.PathPrefix("/").Handler(http.FileServer(http.Dir(webRoot + "/assets/")))

	
}