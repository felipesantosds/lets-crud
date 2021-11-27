package response

type CustomerResponse struct {
	Id        int64  `json:"id"`
	CPF       string `json:"cpf"`
	Name      string `json:"name"`
	BirthDate string `json:"birthDate,omitempty"`
}
