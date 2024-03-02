package main

// Struct usuario
type Usuario struct {
	Nome        string
	Sobrenome   string
	NomeUsuario string
	Senha       string
}

// Struct do paciente
type Paciente struct {
	Nome      string
	Sobrenome string
	Sexo      string
	Idade     int
	CPF       int
	Email     string
	Telefone  string
}

// Struct do usuario
type Medico struct {
	Nome          string
	Sobrenome     string
	Especialidade string
	CRM           int
}

// Struct das consultas
type Consulta struct {
	Paciente Paciente
	Medico   Medico
	Data     string
	Hora     []string
}

// Struct dos exames
type Exame struct {
	Exame    string
	Paciente Paciente
	Data     string
	Hora     string
}

// Struct dos tratamentos
type Tratamento struct {
	Paciente Paciente
	Medico   Medico
	Data     string
	Hora     string
}

// Struct das cirurgias
type Cirurgia struct {
	Cirurgia string
	Paciente Paciente
	Data     string
	Hora     string
}
