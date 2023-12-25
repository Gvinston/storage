package main

import (
	"fmt"
	"github.com/Gvinston/storage/internal/storage"
	"log"
)

func main() {
	st := storage.NewStorage()

	file, err := st.Upload("test.txt", []byte("hello"))

	if err != nil {
		log.Fatal(err)
	}

	restoredFile, err := st.GetById(file.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("it is restored", restoredFile)
}
