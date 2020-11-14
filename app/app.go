package app

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/utm-cssc/log/config"
	"github.com/utm-cssc/log/handler"
	"gorm.io/gorm"
	"net/http"
	"os"
)

type App struct {
	db      *gorm.DB
	router  *mux.Router
	origin  handlers.CORSOption
	methods handlers.CORSOption
	headers handlers.CORSOption
}

type routeHandler func(w http.ResponseWriter, r *http.Request)
type hdlr func(db *gorm.DB, w http.ResponseWriter, r *http.Request)

func (app *App) Init(_ *config.DBConfig) {
	app.router = mux.NewRouter()
	app.origin = handlers.AllowedOrigins([]string{os.Getenv("FRONTEND_URL")})
	app.methods = handlers.AllowedMethods([]string{http.MethodPost})
	app.headers = handlers.AllowedHeaders([]string{"Content-Type"})
	app.setRoutes()
}

func (app *App) setRoutes() {
	app.Post("/ask-jack", app.Handle(handler.AskJack))
}

func (app *App) Post(path string, f routeHandler) {
	app.router.HandleFunc(path, f).Methods(http.MethodPost)
}

// Run - Main run function to startup logger
func (app *App) Run(port string) {
	err := http.ListenAndServe(port, handlers.CORS(app.origin, app.methods, app.headers)(app.router))
	if err != nil {
		panic(err)
	}
}

func (app *App) Handle(h hdlr) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		h(app.db, w, r)
	}
}
