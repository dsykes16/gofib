package pg_cache

import (
	"database/sql"
	"fmt"
	"log"
	"math/big"

	_ "github.com/lib/pq"

	"github.com/dsykes16/gofib/cache"
)

type PgConnection struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
	SSL      bool
}

func (conn *PgConnection) ConnectionString() string {
	sslmode := "enable"
	if !conn.SSL {
		sslmode = "disable"
	}

	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		conn.User,
		conn.Password,
		conn.Host,
		conn.Port,
		conn.DbName,
		sslmode,
	)
}

type PgCache struct {
	db *sql.DB
}

func New(db *sql.DB) cache.Cache {
	sqlStatement := `
	CREATE TABLE IF NOT EXISTS bigfib(
		index integer primary key,
		result varchar(100));
	`
	_, err := db.Exec(sqlStatement)
	if err != nil {
		log.Fatalf("could not create postgres table(bigfib): %s", err)
	}
	return &PgCache{db: db}
}

func (c *PgCache) Add(index uint64, result *big.Int) (err error) {
	sqlStatement := `
		INSERT INTO bigfib (index, result)
		VALUES ($1, $2)
		ON CONFLICT (index) DO NOTHING;
	`
	c.CheckConnection()
	_, err = c.db.Exec(sqlStatement, index, result.String())
	if err != nil {
		log.Fatalf("unable to insert fib(%d): %s", index, err)
	}
	return
}

func (c *PgCache) Clear() (err error) {
	sqlStatement := `
		TRUNCATE TABLE bigfib;
	`
	c.CheckConnection()

	_, err = c.db.Exec(sqlStatement)
	return
}

func (c *PgCache) Get(index uint64) (result *big.Int, hit bool) {
	c.CheckConnection()

	sqlStatement := `
		SELECT result FROM bigfib
		WHERE index=$1;
	`

	row := c.db.QueryRow(sqlStatement, index)

	hit = true
	var resstr string
	if err := row.Scan(&resstr); err != nil {
		if err == sql.ErrNoRows {
			hit = false
		} else {
			log.Fatalf("unable to get: %d from postgres: %s", index, err)
		}
	}

	result = big.NewInt(0)
	result.SetString(resstr, 0)
	return
}

func (c *PgCache) GetRange(start, end uint64) (vals []*big.Int, err error) {
	c.CheckConnection()

	sqlStatement := `
		SELECT (result) FROM bigfib
		WHERE (index) BETWEEN $1 AND $2;
	`

	rows, err := c.db.Query(sqlStatement, start, end)
	if err != nil {
		log.Fatalf("unable to get range of results from postgres: %s", err)
	}
	defer rows.Close()

	var resstr string
	for rows.Next() {
		if err := rows.Scan(&resstr); err != nil {
			log.Fatalf("unable to scan postgres row: %s", err)
		}
		res := big.NewInt(0)
		res.SetString(resstr, 10)
		vals = append(vals, res)
	}
	return
}

func (c *PgCache) Size() (size uint64, err error) {
	c.CheckConnection()

	row := c.db.QueryRow(`SELECT COUNT(*) FROM bigfib;`)
	err = row.Scan(&size)
	return
}

func (c *PgCache) CheckConnection() {
	err := c.db.Ping()
	if err != nil {
		log.Fatalf("could not connect to postgres: %s", err)
	}
}
