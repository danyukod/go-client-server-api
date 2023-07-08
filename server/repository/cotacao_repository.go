package repository

import (
	"database/sql"
	"github.com/danyukod/go-client-server-api/server/model"
)

type CotacaoRepository interface {
	Save(cotacao model.USDBRL) error
}

type cotacaoRepository struct {
	db *sql.DB
}

func NewCotacaoRepository(db *sql.DB) CotacaoRepository {
	return &cotacaoRepository{db}
}

func (c *cotacaoRepository) Save(cotacao model.USDBRL) error {
	stmt, err := c.db.Prepare("INSERT INTO cotacao(code, codein, name, high, low, varBid, pctChange, bid, ask, timestamp, createDate) values(?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(cotacao.Code, cotacao.Codein, cotacao.Name, cotacao.High, cotacao.Low, cotacao.VarBid, cotacao.PctChange, cotacao.Bid, cotacao.Ask, cotacao.Timestamp, cotacao.CreateDate)
	if err != nil {
		return err
	}
	return nil
}
