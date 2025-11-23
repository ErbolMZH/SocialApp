package main

import (
	"Tiago/internal/db"
	"Tiago/internal/env"
	"Tiago/internal/store"
	"fmt"
	"log"
)

func main() {

	addr := env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost:5432/social?sslmode=disable")
	addr2 := env.GetString("DB_ADDR", "ret")
	fmt.Println(addr2)
	fmt.Println("DB:", addr)
	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	store := store.NewStorage(conn)
	db.Seed((store))
}
