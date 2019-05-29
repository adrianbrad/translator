// +build db

package dao

import (
	"database/sql"
	"fmt"

	"golang.org/x/text/language"

	_ "github.com/lib/pq"
)

func connectDB(dbHost, dbPort, dbUser, dbPass, dbName string) (db *sql.DB, err error) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s "+
			"user=%s password=%s "+
			"dbname=%s "+
			"sslmode=disable",
		dbHost, dbPort,
		dbUser, dbPass,
		dbName)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return
	}
	err = db.Ping()
	if err != nil {
		return
	}
	return
}

type PSQLDao struct {
	db *sql.DB
}

func NewPSQLDao(dbHost, dbPort, dbUser, dbPass, dbName string) *PSQLDao {
	db, err := connectDB(dbHost, dbPort, dbUser, dbPass, dbName)
	if err != nil {
		panic(err)
	}

	return &PSQLDao{db}
}

func (p *PSQLDao) StoreTranslation(languageFrom language.Tag, textFrom string, languageTo language.Tag, textTo string) (err error) {
	var id int64
	tx, err := p.db.Begin()
	if err != nil {
		return
	}

	tx.QueryRow(
		`
SELECT translation_id
FROM translations_text
WHERE language_tag=$1 
	AND text=$2;
		`,
		languageFrom.String(),
		textFrom).Scan(&id)

	if id == 0 {
		p.db.QueryRow(
			`
INSERT INTO translations 
VALUES (DEFAULT) RETURNING translation_id;
			`).Scan(&id)
		res, e := p.db.Exec(
			`
INSERT INTO translations_text(
	translation_id, language_tag, text)
	VALUES ($1, $2, $3);			
			`, id, languageFrom.String(), textFrom)
		if e != nil {
			err = e
			tx.Rollback()
			return
		}

		ra, _ := res.RowsAffected()
		if ra != 1 {
			err = fmt.Errorf("No rows affected by insert query")
			tx.Rollback()
			return
		}
	}
	res, e := p.db.Exec(
		`
INSERT INTO translations_text(
translation_id, language_tag, text)
VALUES ($1, $2, $3);			
		`, id, languageTo.String(), textTo)

	if e != nil {
		if e.Error() != "pq: duplicate key value violates unique constraint \"translations_text_translation_id_language_tag_key\"" {
			err = e
		}
		tx.Rollback()
		return
	}
	ra, _ := res.RowsAffected()
	if ra != 1 {
		err = fmt.Errorf("No rows affected by insert query")
		tx.Rollback()
		return
	}

	tx.Commit()
	return
}

func (p *PSQLDao) GetTranslation(languageFrom language.Tag, textFrom string, languageTo language.Tag) (wordTo string, err error) {
	p.db.QueryRow(
		`
SELECT text
FROM translations_text
WHERE language_tag=$1
	AND translation_id = 	(SELECT translation_id
							FROM translations_text
							WHERE language_tag=$2
								AND text=$3)
		`,
		languageTo.String(),
		languageFrom.String(),
		textFrom).Scan(&wordTo)
	if wordTo == "" {
		err = fmt.Errorf("Word not found")
	}
	return
}
