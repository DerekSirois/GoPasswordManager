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
