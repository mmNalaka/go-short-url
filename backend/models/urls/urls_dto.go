package urls

type Url struct {
	Id       int64  `json:"id"`
	ShortUrl string `json:"short_url"`
	LongUrl  string `json:"long_url"`
}
