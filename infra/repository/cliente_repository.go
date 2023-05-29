package repository

import (
	"ApiGoLang/core/repository/contracts"
	"ApiGoLang/domain/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ClienteRepository struct {
	contracts.IRepositoryBase[entities.Cliente]
	db gorm.DB
}

func NewClienteRepository(base contracts.IRepositoryBase[entities.Cliente], db gorm.DB) *ClienteRepository {
	obj := new(ClienteRepository)
	obj.IRepositoryBase = base
	obj.db = db
	return obj
}

func (o *ClienteRepository) Exists(entity entities.Cliente) (error, *entities.Cliente) {
	var obj entities.Cliente
	result := o.db.Find(&obj, "cpf = ?", entity.Cpf)

	if result.Error != nil {
		return result.Error, nil
	}

	if result.RowsAffected == 0 {
		return nil, nil
	}

	return nil, &obj
}

func (o *ClienteRepository) GetClienteEmailSenhaLogin(email string, senha string) (error, *entities.Cliente) {
	var obj entities.Cliente
	result := o.db.Find(&obj, "Email = ? and Senha = ?", email, senha)

	if result.Error != nil {
		return result.Error, nil
	}

	if result.RowsAffected == 0 {
		return nil, nil
	}

	return nil, &obj
}

func (o *ClienteRepository) AtualizarStatus(id uuid.UUID, ativar bool) error {
	var obj entities.Cliente
	var status = "A"

	err := o.db.Find(&obj, "id = ?", id)

	if err.Error != nil {
		return err.Error
	}

	if !ativar {
		status = "I"
	}

	o.db.Model(&obj).Update("status", status)

	return nil
}
