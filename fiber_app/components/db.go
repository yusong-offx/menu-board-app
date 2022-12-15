package components

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-redis/redis/v9"
	_ "github.com/lib/pq"
)

var (
	Postgres *sql.DB
	Redis    *redis.Client
	err      error
)

func PostgresConnect() {
	if Postgres, err = sql.Open("postgres",
		fmt.Sprintf("host=%s port=%d user=%s password=%s  dbname=%s sslmode=disable",
			"postgresql-db",
			5432,
			"postgres",
			"docker123",
			"menu",
		)); err != nil {
		log.Fatal(err)
	}
	if err = Postgres.Ping(); err != nil {
		log.Fatal(err)
	}
}

func RedisConnect() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     "redis-db:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
