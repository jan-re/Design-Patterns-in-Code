package main

import (
	"fmt"
	"sync"
)

type abstractDbConnection string

type db struct {
	connection abstractDbConnection
}

var lock = &sync.Mutex{}

var dbInstance *db

func getConnection() *db {
	if dbInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if dbInstance == nil {
			fmt.Println("Establishing DB connection...")
			dbInstance = &db{
				connection: "psql/abstract:imagined",
			}
		} else {
			fmt.Println("DB connection already instantiated.")
		}
	} else {
		fmt.Println("DB connection already instantiated.")
	}

	return dbInstance
}

/*
// Alternatively

package main

import (
    "fmt"
    "sync"
)

var once sync.Once

type single struct {
}

var singleInstance *single

func getInstance() *single {
    if singleInstance == nil {
        once.Do(
            func() {
                fmt.Println("Creating single instance now.")
                singleInstance = &single{}
            })
    } else {
        fmt.Println("Single instance already created.")
    }

    return singleInstance
}

*/
