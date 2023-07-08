package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/danyukod/go-client-server-api/server/model"
	"github.com/danyukod/go-client-server-api/server/repository"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"time"
)

func main() {

	err := Serve()
	if err != nil {
		log.Println("Error: ", err)
		return
	}

}

func Serve() error {
	http.HandleFunc("/cotacao", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return err
	}
	return nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 200*time.Millisecond)
	defer cancel()
	select {
	case <-ctx.Done():
		log.Println("client canceled the request")
	default:
		cotacaoUsecase(w, ctx)
	}
}

func cotacaoUsecase(w http.ResponseWriter, ctx context.Context) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	usdbrl := model.USDBRLResponse{}
	err = json.NewDecoder(res.Body).Decode(&usdbrl)
	if err != nil {
		panic(err)
	}

	db := createSqliteDatabase(err)
	defer db.Close()

	repo := repository.NewCotacaoRepository(db)

	err = repo.Save(usdbrl.USDBRL)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(usdbrl.USDBRL)
	if err != nil {
		panic(err)
	}
}

func createSqliteDatabase(err error) *sql.DB {
	db, err := sql.Open("sqlite3", "cotacao.db")
	if err != nil {
		panic(err)
	}

	createTableIfNotExists(db)
	return db
}

func createTableIfNotExists(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS cotacao (id INTEGER PRIMARY KEY AUTOINCREMENT, code TEXT, codein TEXT, name TEXT, high TEXT, low TEXT, varBid TEXT, pctChange TEXT, bid TEXT, ask TEXT, timestamp TEXT, createDate TEXT)`)
	if err != nil {
		panic(err)
	}
}
