package postgres

import (
	"fmt"

	"github.com/ThailanTec/poc-serasa/internal/core/domain"
	erros "github.com/ThailanTec/poc-serasa/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type MessengerPostgresRepository struct {
	db *gorm.DB
}

func NewPG(settings domain.DBSettings) (p *MessengerPostgresRepository) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", settings.Host, settings.User, settings.Password, settings.DBName, settings.Port, settings.SSL, settings.TimeZone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}

	if !settings.DisableAutoMigrate {
		err = db.AutoMigrate(&domain.User{})
		if err != nil {
			return
		}
	}

	p = &MessengerPostgresRepository{db: db}
	return
}

func (ps *MessengerPostgresRepository) CreateUser(user domain.User) (u domain.User, e error) {
	req := ps.db.Create(&user)

	if req.RowsAffected == 0 {
		return domain.User{}, erros.CustomError("Erro ao criar usuário")
	}

	return u, e
}

func (ps *MessengerPostgresRepository) GetUser(id int) (u domain.User, e error) {
	u.ID = id
	ps.db.First(&u, id)

	return u, nil
}

func (ps *MessengerPostgresRepository) DeleteUser(id int) (delete bool, e error) {

	u := domain.User{}
	u.ID = id

	ps.db.Delete(&u)

	return true, e
}
