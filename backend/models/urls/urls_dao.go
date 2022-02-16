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

func GetUrl(u string) string {
	stmt, err := database.Client.Prepare("SELECT * FROM urls WHERE LongUrl = $1")
	if err != nil {
		log.Println(fmt.Printf("Error when trying to prepare database statements: %s", err))
		return ""
	}
	defer stmt.Close()

	result := stmt.QueryRow(u)
	url := Url{}
	if err := result.Scan(&url.Id, &url.LongUrl, &url.ShortUrl); err != nil {
		log.Println(fmt.Printf("Error when trying to query database: %s", err))
		return ""
	}
	return url.ShortUrl
}

func CreateUrl(u string) *Url {
	stmt, err := database.Client.Prepare("INSERT INTO urls (LongUrl) VALUES ($1) RETURNING id, long_url, short_url")
	if err != nil {
		log.Println(fmt.Printf("Error when trying to prepare database statements: %s", err))
		return nil
	}
	defer stmt.Close()

	result := stmt.QueryRow(u)
	url := Url{}
	if err := result.Scan(&url.Id, &url.LongUrl, &url.ShortUrl); err != nil {
		log.Println(fmt.Printf("Error when trying to query database: %s", err))
		return nil
	}
	return &url
}
