package moysklad

// Token Токен.
// Документация МойСклад: https://dev.moysklad.ru/doc/api/remap/1.2/#mojsklad-json-api-obschie-swedeniq-autentifikaciq-poluchenie-nowogo-tokena
type Token struct {
	AccessToken string `json:"access_token"`
}

func (t Token) String() string {
	return Stringify(t)
}
