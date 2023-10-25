package password

import (
	"passwordManager/pkg/database"
	"passwordManager/pkg/encryption"
)

func CreatePassword(p *Password) error {
	pass, err := encryption.Encrypt(p.Password)
	if err != nil {
		return err
	}
	pDb := &database.Password{
		AppName:  p.AppName,
		Password: pass,
	}
	err = database.CreatePassword(pDb)
	return err
}

func GetPassword(appname string) (*Password, error) {
	p, err := database.GetPasswordByAppName(appname)
	if err != nil {
		return nil, err
	}
	pass, err := encryption.Decrypt(p.Password)
	if err != nil {
		return nil, err
	}
	return &Password{
		AppName:  p.AppName,
		Password: pass,
	}, nil
}
