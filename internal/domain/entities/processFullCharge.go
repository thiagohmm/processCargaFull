package entities

import (
	"time"
)

type ProcessFullCharge struct {
	IdProcessoCargaFull int       `json:"id_processo_carga_full" gorm:"primaryKey;autoIncrement"`
	CodigoIbm           string    `json:"codigo_ibm"`
	DataInicioProcesso  time.Time `json:"data_inicio_processo"`
	Processando         string    `json:"processando,omitempty"`
	DataFimProcesso     time.Time `json:"data_fim_processo,omitempty"`
	Sucesso             string    `json:"sucesso,omitempty"`
	Cancelado           string    `json:"cancelado,omitempty"`
}
