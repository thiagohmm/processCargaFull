package entities

import (
	"time"
)

type Dealer struct {
	IdRevendedor           int       `json:"idRevendedor"`
	CodigoIBM              string    `json:"codigoIbm"`
	NomeRevendedor         string    `json:"nomeRevendedor"`
	StatusLoja             string    `json:"statusLoja"`
	UF                     string    `json:"uf"`
	Selcon                 int       `json:"selcon"`
	DataAbertura           time.Time `json:"dataAbertura"`
	DataFechamento         time.Time `json:"dataFechamento"`
	PitStop                int       `json:"pitStop"`
	Email                  string    `json:"email"`
	SiteCode               string    `json:"siteCode"`
	IdSegmento             int       `json:"idSegmento"`
	TipoLoja               string    `json:"tipoLoja"`
	CNPJ                   string    `json:"cnpj"`
	InscricaoEstadual      string    `json:"inscricaoEstadual"`
	Endereco               string    `json:"endereco"`
	Bairro                 string    `json:"bairro"`
	Cidade                 string    `json:"cidade"`
	Area                   int       `json:"area"`
	ContatoLoja            string    `json:"contatoLoja"`
	NomeConsultorLoja      string    `json:"nomeConsultorLoja"`
	NomeConsultorAbertura  string    `json:"nomeConsultorAbertura"`
	DataUltimaAtualizacao  time.Time `json:"dataUltimaAtualizacao"`
	Carencia               int       `json:"carencia"`
	DataInicioCarencia     time.Time `json:"dataInicioCarencia"`
	CEP                    string    `json:"cep"`
	EmailConsultorLoja     string    `json:"emailConsultorLoja"`
	EmailConsultorAbertura string    `json:"emailConsultorAbertura"`
	EnquadramentoFiscal    string    `json:"enquadramentoFiscal"`
	PermitirCargaFull      int       `json:"permitirCargaFull"`
	CNPJCompra             string    `json:"cnpjCompra"`
	DataInicioTransmissao  time.Time `json:"dataInicioTransmissao"`
	DataTerminoTransmissao time.Time `json:"dataTerminoTransmissao"`
	DataInicioReforma      time.Time `json:"dataInicioReforma"`
	DataTerminoReforma     time.Time `json:"dataTerminoReforma"`
	EmailUserSystem        string    `json:"emailUserSystem"`
	ModeloLoja             string    `json:"modeloLoja"`
}
