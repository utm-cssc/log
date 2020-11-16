package app

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/utm-cssc/log/app/model"
	"github.com/utm-cssc/log/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
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

func (app *App) Init(dbConfig *config.DBConfig) {
	dbFormat :=
		fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=America/New_York",
			dbConfig.Host,
			dbConfig.Port,
			dbConfig.User,
			dbConfig.Password,
			dbConfig.Name,
		)
	db, err := gorm.Open(postgres.Open(dbFormat), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	if err != nil {
		log.Fatal("Unable to connect to database\n", err)
	}
	app.db = db
	app.router = mux.NewRouter()
	app.origin = handlers.AllowedOrigins([]string{os.Getenv("FRONTEND_URL")})
	app.methods = handlers.AllowedMethods([]string{http.MethodPost})
	app.headers = handlers.AllowedHeaders([]string{"Content-Type"})
	app.setRoutes()
	log.Println("Connected to Database")
	if db.Migrator().CreateTable(model.NewAskJackLog()) != nil {
		log.Println("Base tables not created")
	}
}

func (app *App) setRoutes() {
	app.Post("/ask-jack", app.Handle(AddQuestionEntry))
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

// Handlers
func AddQuestionEntry(db *gorm.DB, _ http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		os.Exit(1)
	}
	form := url.Values{}
	for key, values := range r.PostForm {
		fmt.Println(key, values)
		form[key] = values
	}
	_, err := http.PostForm("https://formspree.io/xwkrdzyg", form)
	ip, port, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(ip)
	fmt.Println(port)
	p, err := strconv.Atoi(port)
	if err != nil {
		fmt.Println("Port is unprocessable")
		os.Exit(1)
	}
	askJackLog := model.AskJackLog{
		IP:       ip,
		Port:     p,
		Question: form["Question"][0],
		Email:    form["Email"][0],
	}
	err = db.Create(&askJackLog).Error
	if err != nil {
		fmt.Println("Unable to log request")
	}
}
