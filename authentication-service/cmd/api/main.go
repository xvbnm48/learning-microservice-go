package main

import (
	"authentication/data"
	"database/sql"
)

const webPort = "8081"

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {

}
