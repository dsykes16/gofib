package pg_cache_test

import (
	"database/sql"
	"log"
	"strconv"

	. "github.com/onsi/ginkgo"
	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"

	"github.com/dsykes16/gofib/cache"
	"github.com/dsykes16/gofib/cache/pg_cache"

	. "github.com/dsykes16/gofib/cache/shared_tests"
)

var _ = Describe("Postgres Cache Tests", func() {
	// Ginkgo's BeforeSuite and AfterSuite do not work properly
	// in this use case. Ref: https://github.com/onsi/ginkgo/issues/457
	// TODO: Find better workaround for cleanup than expiring containers
	dockerPool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker daemon: %s", err)
	}

	newConfiguredPgCache := func() cache.Cache {
		container, db := startPostgres(dockerPool)
		container.Expire(uint(30))
		return pg_cache.New(db, "bigfib")
	}

	SharedCacheTests(newConfiguredPgCache)
})

func startPostgres(pool *dockertest.Pool) (container *dockertest.Resource, db *sql.DB) {
	var err error

	conn := pg_cache.PgConnection{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "testpass",
		DbName:   "gofib",
		SSL:      false,
	}

	container, err = pool.RunWithOptions(
		&dockertest.RunOptions{
			Repository: "postgres",
			Tag:        "13.3",
			Env: []string{
				"POSTGRES_USER=postgres",
				"POSTGRES_PASSWORD=testpass",
				"POSTGRES_DB=gofib",
			},
		},
		func(config *docker.HostConfig) {
			config.AutoRemove = true
			config.RestartPolicy = docker.RestartPolicy{
				Name: "no",
			}
		},
	)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	conn.Port, err = strconv.Atoi(container.GetPort("5432/tcp"))
	if err != nil {
		log.Fatalf("Could not get postgres port: %s", err)
	}

	if err = pool.Retry(func() error {
		db, err = sql.Open("postgres", conn.ConnectionString())
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to postgres: %s", err)
	}
	return
}
