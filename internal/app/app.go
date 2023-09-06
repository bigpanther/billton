package app

import (
	"log"
	"os"
	"sync"
	"time"

	"log/slog"

	"github.com/bigpanther/warrant/internal/firebase"
	"github.com/bigpanther/warrant/internal/middleware"
	"github.com/bigpanther/warrant/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/pop/v6"
	slogformatter "github.com/samber/slog-formatter"
	sloggin "github.com/samber/slog-gin"
)

var logger = slog.New(
	slogformatter.NewFormatterHandler(
		slogformatter.TimezoneConverter(time.UTC),
		slogformatter.TimeFormatter(time.DateTime, nil),
	)(
		slog.NewTextHandler(os.Stdout, nil),
	),
)

// App represents an instance of the application
func App() (*gin.Engine, error) {
	db := InitDb()
	firebase, err := InitFirebase()
	if err != nil {
		logger.Error("failed it initialize connection to firebase:", err)
		return nil, err
	}
	engine := gin.Default()

	engine.Use(sloggin.New(logger))
	engine.Use(middleware.Database(db))
	engine.MaxMultipartMemory = 8 << 20
	setupRouter(engine, firebase)
	return engine, nil
}

// db is a connection to the database to be used
// throughout the application.
var db *pop.Connection
var once sync.Once

// InitDb initializes a db connection
func InitDb() *pop.Connection {
	var err error
	env := envy.Get("GO_ENV", "dev")
	once.Do(func() {
		db, err = pop.Connect(env)
		if err != nil {
			log.Fatal(env, err)
		}
	})
	pop.Debug = env == "dev"
	return db
}

const currentUserKey = "current_user"

func loggedInUser(c *gin.Context) *models.User {
	return c.Value(currentUserKey).(*models.User)
}

// InitFirebase initializes a firebase instance
func InitFirebase() (f firebase.Firebase, err error) {
	env := envy.Get("GO_ENV", "dev")
	isProd := env == "production"
	if isProd {
		return firebase.New()
	}
	return firebase.NewFake()
}
