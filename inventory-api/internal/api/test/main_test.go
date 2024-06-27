package test

import (
	"log"
	"os"
	"testing"

	"github.com/bobbybof/inventory-api/config"
	"github.com/bobbybof/inventory-api/internal/api"
	"github.com/bobbybof/inventory-api/internal/repository"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store repository.Store) *api.Server {
	config, err := config.NewConfig("../../../.env")
	if err != nil {
		log.Fatal("cannot load env: ", err)
	}

	server, err := api.NewServer(*config, store)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
