package database

import "sync"

type DB struct {
	Path string
	Mux  *sync.RWMutex
}