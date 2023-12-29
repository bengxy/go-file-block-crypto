package args

import (
	"sync"
)

var once sync.Once

type singleton struct {
}

var instance *singleton

func GetParserInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}

	})
	return instance
}
