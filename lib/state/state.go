package state

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
)

const stateFileName = "data.json"

var state = make(map[int64]string)

func SetState(key int64, value string) {
	state[key] = value
	updateCache()
}

func GetState(key int64) (string, bool) {
	val, ok := state[key]
	return val, ok
}

func RemoveFromState(key int64) {
	delete(state, key)
	updateCache()
}

func updateCache() {
	go func() {
		data, _ := json.MarshalIndent(state, "", " ")
		file, err := os.OpenFile(stateFileName, os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			log.Printf("Failed to open file: %v", err)
			return
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				log.Printf("Failed to write to file: %v", err)
			}
		}(file)

		_, err = file.Write(data)
		if err != nil {
			log.Printf("Failed to write to file: %v", err)
		}
	}()
}

func PopulateFromCache() {
	file, err := os.OpenFile(stateFileName, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&state)
	if err != nil {
		if errors.Is(err, io.EOF) {
			state = make(map[int64]string)
		}
	}
}
