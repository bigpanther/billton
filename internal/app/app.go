package app

import (
	"log"
	"os"
	"sync"
	"time"

	"log/slog"

	"github.com/bigpanther/billton/internal/firebase"
	"github.com/bigpanther/billton/internal/middleware"
	"github.com/bigpanther/billton/internal/models"
	"github.com/gin-gonic/gin"
	slogformatter "github.com/samber/slog-formatter"
	sloggin "github.com/samber/slog-gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
var db *gorm.DB
var once sync.Once

// InitDb initializes a db connection
func InitDb() *gorm.DB {
	var err error

	dsn := "host=localhost user=postgres password=postgres dbname=billton-dev port=5432 sslmode=disable TimeZone=UTC"
	once.Do(func() {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
		if err != nil {
			log.Fatal(env(), err)
		}
	})
	return db
}

const currentUserKey = "current_user"

func loggedInUser(c *gin.Context) *models.User {
	return c.Value(currentUserKey).(*models.User)
}
func env() string {
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "dev"
	}
	return env
}

// InitFirebase initializes a firebase instance
func InitFirebase() (f firebase.Firebase, err error) {
	isProd := env() == "production"
	if isProd {
		return firebase.New()
	}
	return firebase.NewFake()
}
