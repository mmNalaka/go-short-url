package urls

import (
	"fmt"
	"log"

	"github.com/mmnalaka/go-short-url/platfrom/database"
)

func (u *Url) Get() error {
	stmt, err := database.Client.Prepare("SELECT * FROM urls WHERE LongUrl = $1")
	if err != nil {
		log.Println(fmt.Printf("Error when trying to prepare database statements: %s", err))
		return err
	}
	defer stmt.Close()

	result := stmt.QueryRow(u.LongUrl)
	if err := result.Scan(&u.Id, &u.LongUrl, &u.ShortUrl); err != nil {
		log.Println(fmt.Printf("Error when trying to query database: %s", err))
		return err
	}

	return nil
}
