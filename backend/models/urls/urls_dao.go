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

func GetUrl(u string) *Url {
	stmt, err := database.Client.Prepare("SELECT * FROM urls WHERE long_url = $1")
	url := Url{}

	if err != nil {
		log.Println(fmt.Printf("Error when trying to prepare database statements, GetUrl: %s", err))
		return nil
	}
	defer stmt.Close()

	result := stmt.QueryRow(u)
	if err := result.Scan(&url.Id, &url.LongUrl, &url.ShortUrl); err != nil {
		log.Println(fmt.Printf("Error when trying to query database: %s", err))
		return nil
	}
	return &url
}

func CreateUrl(long string) *Url {
	stmt, err := database.Client.Prepare("INSERT INTO urls (long_url, short_url) VALUES ($1, $2) RETURNING id, long_url, short_url")
	if err != nil {
		log.Println(fmt.Printf("Error when trying to prepare database statements, CreateUrl: %s", err))
		return nil
	}
	defer stmt.Close()

	result := stmt.QueryRow(long)
	url := Url{}
	if err := result.Scan(&url.Id, &url.LongUrl, &url.ShortUrl); err != nil {
		log.Println(fmt.Printf("Error when trying to query database: %s", err))
		return nil
	}
	return &url
}
