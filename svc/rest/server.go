package rest

import (
	"strings"
	"web-Scraping-test/svc/configs"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

type RestServer struct {
	Log *logrus.Logger
	c   *configs.Config
}

func NewServer(c *configs.Config) (*RestServer, error) {

	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)

	return &RestServer{
		Log: logger,
		c:   c,
	}, nil
}

func (s *RestServer) Run() error {
	r := gin.Default()

	r.Use(cors.New(cors.Options{
		AllowedOrigins:   strings.Split(s.c.CORSHosts, ","),
		AllowedMethods:   []string{"HEAD", "GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	}))

	pprof.Register(r)

	openAccessed := r.Group("/")
	{
		openAccessed.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{"ping": "pong"})
		})
	}

	apis := r.Group("/apis")
	{
		apis.POST("/visit_url", s.VisitUrls)
	}

	err := r.Run(s.c.HostPort)
	if err != nil {
		return errors.Errorf("serving on %s failed: %v", s.c.HostPort, err)
	}

	return nil
}
