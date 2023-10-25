package database

type Password struct {
	Id       int
	AppName  string `db:"appname"`
	Password string
}

func CreatePassword(p *Password) error {
	_, err := db.Exec("INSERT INTO password (appname, password) VALUES ($1, $2)", p.AppName, p.Password)
	return err
}
