package database

import "fmt"

type Password struct {
	Id       int
	AppName  string `db:"appname"`
	Password string
}

func CreatePassword(p *Password) error {
	if p.Password == "" || p.AppName == "" {
		return fmt.Errorf("password or appname is empty")
	}
	_, err := db.Exec("INSERT INTO password (appname, password) VALUES ($1, $2)", p.AppName, p.Password)
	return err
}

func GetPasswordByAppName(appName string) (*Password, error) {
	row := db.QueryRow("SELECT * FROM password WHERE appname = $1", appName)

	var id int
	var appname string
	var password string
	err := row.Scan(&id, &appname, &password)
	if err != nil {
		return nil, err
	}

	return &Password{
		Id:       id,
		AppName:  appName,
		Password: password,
	}, nil
}
