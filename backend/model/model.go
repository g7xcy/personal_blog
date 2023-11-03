package model

import (
	"context"
	"fmt"
	"log"
	"time"

	"example.com/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

type User struct {
	User_id     int64
	Email       string
	Name        string
	Password    string
	Create_time time.Time
	Update_time time.Time
}

type Blog struct {
	Blog_id     int64
	User_id     int64
	User        string
	Title       string
	Content     string
	Create_time time.Time
	Update_time time.Time
}

func InitDB() error {
	connString := "postgres://" + util.Env["DB_USER"] + ":" + util.Env["DB_PWD"] + "@" + util.Env["DB_URL"] + ":" + util.Env["DB_PORT"] + "/" + util.Env["DB"]
	connConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		log.Fatal(err)
		return fmt.Errorf("InitDB: %v", err)
	}
	connConfig.MaxConns = 5
	db, err := pgxpool.NewWithConfig(context.Background(), connConfig)
	if err != nil {
		log.Fatal(err)
		return fmt.Errorf("InitDB: %v", err)
	}

	pingErr := db.Ping(context.Background())
	if pingErr != nil {
		log.Fatal(pingErr)
		return fmt.Errorf("InitDB: %v", err)
	}

	log.Println("Connected!")
	DB = db
	DB.Exec(context.Background(), `
	CREATE TABLE IF NOT EXISTS users (
		user_id SERIAL PRIMARY KEY NOT NULL,
		email VARCHAR(50) NOT NULL,
		name VARCHAR(28) NOT NULL,
		password VARCHAR(448) NOT NULL,
		create_time TIMESTAMPTZ NOT NULL,
		update_time TIMESTAMPTZ NOT NULL
		)
	`)
	DB.Exec(context.Background(), `
	CREATE TABLE IF NOT EXISTS blogs (
		blog_id SERIAL PRIMARY KEY NOT NULL,
		user_id INTEGER REFERENCES users,
		title VARCHAR(50) NOT NULL,
		content TEXT NOT NULL,
		create_time TIMESTAMPTZ NOT NULL,
		update_time TIMESTAMPTZ NOT NULL
	)
`)
	return nil
}

func HealthCheck() error {
	return DB.Ping(context.Background())
}
