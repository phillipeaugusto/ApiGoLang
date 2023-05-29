package output

import "github.com/google/uuid"

type DadosLoginOutPutDto struct {
	ClienteId uuid.UUID
	Token     string
}
