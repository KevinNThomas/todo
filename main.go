package main

import (
	"log"

	"github.com/namsral/flag"
	"github.com/prologic/bitcask"
)

var (
	db *bitcask.Bitcask
)

func main() {
	var (
		dbpath           string
		bind             string
		MAX_ITEMS        int
		MAX_TITLE_LENGTH int
	)

	flag.StringVar(&dbpath, "dbpath", "todo.db", "Database path")
	flag.StringVar(&bind, "bind", "0.0.0.0:8000", "[int]:<port> to bind to")
	flag.IntVar(&MAX_ITEMS, "MAX_ITEMS", 100, "maximum number of items allowed in the todo list")
	flag.IntVar(&MAX_TITLE_LENGTH, "MAX_TITLE_LENGTH", 100, "maximum valid length of a todo item's title")
	flag.Parse()

	var err error
	db, err = bitcask.Open(dbpath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	newServer(bind, MAX_ITEMS, MAX_TITLE_LENGTH).listenAndServe()
}
