package dao

import (
	"cic-project/model"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"log"
	"time"
)

type PgDao struct {
	Host 		string
	User 		string
	Password	string
	DbName  	string
	Db 			*pg.DB
}

func (d *PgDao) GetConnection() {
	db := pg.Connect(&pg.Options{
		User: d.User,
		Password: d.Password,
		Addr: d.Host,
		Database: d.DbName,
	})
	d.Db = db
	_ = createSchema(db)
}

func (d PgDao) InsertForecast(report model.DayReport) {
	_, err := d.Db.Model(&report).
		OnConflict("(local_date) DO UPDATE").
		Insert()
	if err != nil {
		log.Fatal(err)
	}
}

func (d PgDao) FindForecast(date string) []model.DayReport {
	timeLayout := "2006-01-02"

	t, err := time.Parse(timeLayout, date)
	query := t.Format(timeLayout)

	var report []model.DayReport
	if date != ""{
		err = d.Db.Model(&report).Where("local_date=?", query).Select()
	} else {
		err = d.Db.Model(&report).Order("id DESC").Limit(14).Select()
	}
	if err != nil {
		log.Println("Cannot get forecast", err)
	}
	return report
}
func createSchema(db *pg.DB) error {
	for _, model := range []interface{}{(*model.DayReport)(nil)} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
		})
		if err != nil {
			return err
		}
	}
	return nil
}