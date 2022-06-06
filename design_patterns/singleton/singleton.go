package main

import (
	"log"
	"sync"
	"time"
)

type Database struct{}

var db *Database
var lock sync.Mutex

func (Database) CreateSingleConnection() {
	log.Println("Creating Singleton Connection Database")
	time.Sleep(2 * time.Second)
	log.Println("Done")
}

func getInstanceDatabase() *Database {
	lock.Lock()
	defer lock.Unlock()
	if db == nil {
		log.Println("Creating database for Singleton")
		db = &Database{}
		db.CreateSingleConnection()
	} else {
		log.Println("Connection dabase ALREADY created")
	}

	return db
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			getInstanceDatabase()
		}()
	}

	wg.Wait()
}
