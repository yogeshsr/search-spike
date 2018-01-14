package pgsearch_api

import (
	"github.com/yogeshsr/search-spike/golang-api/models"
	"database/sql"
	"fmt"
)

func SearchVouchers(db * sql.DB, search string) []models.Voucher {

	fmt.Println("# Querying")

	//TODO avoid sql injection

	query := `with q as(select to_tsquery('%s') as query),
	ranked as(select title,sponsor,description, ts_rank_cd(tsv, query) as rank
		from vouchers, q where q.query @@ tsv order by rank desc limit 10 )
		select title,sponsor from ranked, q order by rank desc`
	//rows, err := db.Query("SELECT title, sponsor FROM vouchers LIMIT 1")
	rows, err := db.Query(fmt.Sprintf(query, search))
	checkErr(err)

	//rows.NextResultSet()

	var vouchers []models.Voucher
	for rows.Next() {
		var title string
		var sponsor string

		err = rows.Scan(&title, &sponsor)
		checkErr(err)

		fmt.Sprintf(title, sponsor)
		vouchers = append(vouchers, models.Voucher{
			Title: title,
			Sponsor: sponsor,
		})
	}

	return vouchers
}