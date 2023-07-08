package repository

import (
	"context"
	"database/sql"
	"github.com/danyukod/go-client-server-api/server/model"
	"time"
)

type CotacaoRepository interface {
	Save(cotacao model.USDBRL) error
}

type cotacaoRepository struct {
	db  *sql.DB
	ctx context.Context
}

func NewCotacaoRepository(db *sql.DB, ctx context.Context) CotacaoRepository {
	return &cotacaoRepository{db, ctx}
}

func (c *cotacaoRepository) Save(cotacao model.USDBRL) error {
	ctx, cancel := context.WithTimeout(c.ctx, 10*time.Millisecond)
	defer cancel()

	stmt, err := c.db.Prepare("INSERT INTO cotacao(code, codein, name, high, low, varBid, pctChange, bid, ask, timestamp, createDate) values(?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, cotacao.Code, cotacao.Codein, cotacao.Name, cotacao.High, cotacao.Low, cotacao.VarBid, cotacao.PctChange, cotacao.Bid, cotacao.Ask, cotacao.Timestamp, cotacao.CreateDate)
	if err != nil {
		return err
	}
	return nil
}
