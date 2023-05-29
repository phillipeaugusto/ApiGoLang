package entities

import (
	"crypto/md5"
	"encoding/hex"
)
import _ "crypto/md5"

type Cliente struct {
	EntidadeBase[Cliente, Cliente, Cliente]
	Nome    string `gorm:"type:varchar(200);not null"`
	Email   string `gorm:"type:varchar(100);not null"`
	Cpf     string `gorm:"type:varchar(20);not null"`
	Celular string `gorm:"type:varchar(30);not null"`
	Senha   string `gorm:"type:varchar(500);not null"`
}

func NewClienteLogin(email string, senha string, criptografarSenha bool) *Cliente {
	u := new(Cliente)
	u.SetDefaultValues()
	u.Email = email
	u.Senha = senha
	if criptografarSenha {
		u.EncryptPassword()
	}
	return u
}

func NewCliente(nome string, email string, cpf string, celular string, senha string, criptografarSenha bool) *Cliente {
	u := new(Cliente)
	u.SetDefaultValues()
	u.Nome = nome
	u.Cpf = cpf
	u.Celular = celular
	u.Email = email
	u.Senha = senha
	if criptografarSenha {
		u.EncryptPassword()
	}

	return u
}

func (u *Cliente) EncryptPassword() {
	hasher := md5.New()
	hasher.Write([]byte(u.Senha))
	u.Senha = hex.EncodeToString(hasher.Sum(nil))
}
