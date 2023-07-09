package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/cotacao", nil)
	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	var usdbrl string
	err = json.NewDecoder(res.Body).Decode(&usdbrl)

	fileName := "cotacao.txt"

	f, err := os.Create(fileName)
	if err != nil {
		log.Println("Error creating file: ", err)
		return
	}
	defer f.Close()

	content := "Dólar: " + usdbrl + "\n"

	_, err = f.WriteString(content)
	if err != nil {
		log.Println("Error writing file: ", err)
		return
	}

	fmt.Println("File", fileName, "created successfully.")

}
