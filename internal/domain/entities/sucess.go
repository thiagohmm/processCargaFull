package entities

type Success struct {
	TypeValidation string      `json:"typeValidation,omitempty"`
	Message        string      `json:"message"`
	Data           interface{} `json:"data,omitempty"`
	Success        bool        `json:"success"`
	ID             int         `json:"id,omitempty"`
	Processo       int         `json:"processo,omitempty"`
	IbmCode        string      `json:"ibmCode,omitempty"`
	idDealer       int         `json:"idDealer,omitempty"`
}
