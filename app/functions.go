package main

import (
	"fmt"

	"github.com/inancgumus/screen"
)

// Imprimir especialidades
func imprimirEspecialidades(medicos []Medico) {
	for i, medico := range medicos {
		fmt.Println("Medico ID: [", i+1, "]\nEspecialidade: ", medico.Especialidade)
	}
}

// Imprimir medicos
func imprimirMedicos(medicos []Medico) {
	for i, medico := range medicos {
		fmt.Println("Medico ID: [", i+1, "]\nNome: ", medico.Nome, "\nSobrenome: ", medico.Sobrenome, "\nCRM: ", medico.CRM)
	}
}

// Funcao para marcar consultas
func marcarConsulta(medicos []Medico, pacientes []Paciente, consultas []Consulta) Consulta {
	if len(pacientes) == 0 {
		fmt.Println("Não há pacientes cadastrados")
		return Consulta{}
	} else if len(medicos) == 0 {
		fmt.Println("Não há medicos cadastrados")
		return Consulta{}
	}
	var paciente Paciente
	// var medico Medico
	var consulta Consulta
	var nome, especialidade, data string
	var cpf int
	fmt.Println("Marcar Consulta")
	fmt.Println("===============")
	fmt.Println("Digite o nome do paciente deseja marcar a consulta:")
	fmt.Scanln(&nome)
	fmt.Println("Digite o CPF do paciente: ")
	fmt.Scanln(&cpf)

	for _, paciente := range pacientes {
		if paciente.Nome == nome && paciente.CPF == cpf {
			fmt.Println("Paciente encontrado")
			fmt.Println()
			if len(medicos) == 0 {
				for _, medico := range medicos {
					fmt.Println("Digite qual especialidade deseja marcar:")
					imprimirEspecialidades(medicos)
					fmt.Scanln(&especialidade)
					if especialidade == medico.Especialidade {
						fmt.Println("Digite um medico para marcar a consulta")
						imprimirMedicos(medicos)
						fmt.Scanln(&medico)
						fmt.Println("Digite a data da consulta")
						imprimirDatasConsultas(consultas)
						fmt.Scanln(&data)
						if data == consulta.Data {
							fmt.Println("Digite a hora da consulta")
							imprimirHorasConsultas(consultas)
							fmt.Scanln(&consulta.Hora)
							fmt.Println("Consulta marcada com sucesso")
							fmt.Scanln()
							screen.Clear()
							return consulta
						}
					}
				}
			}
		}
	}
	if paciente.Nome == nome && paciente.CPF == cpf {

	} else {
		fmt.Println("Paciente não encontrado")

	}
	return consulta
}

// Funcao para marcar exames
func marcarExame(pacientes []Paciente, exames []Exame) Exame {
	if len(pacientes) == 0 {
		fmt.Println("Não há pacientes cadastrados")
		return Exame{}
	}
	var paciente Paciente
	var exame Exame
	var nome, data, exameEscolhido string
	var cpf int
	fmt.Println("Marcar Consulta")
	fmt.Println("===============")
	fmt.Println("Digite o nome do paciente deseja marcar o exame:")
	fmt.Scanln(&nome)
	fmt.Println("Digite o CPF do paciente: ")
	fmt.Scanln(&cpf)

	for _, paciente := range pacientes {
		if paciente.Nome == nome && paciente.CPF == cpf {
			fmt.Println("Paciente encontrado")
			fmt.Println()
			for _, exame := range exames {
				fmt.Println("Digite qual exame deseja marcar:")
				imprimirExames(exames)
				fmt.Scanln(&exame)
				if exameEscolhido == exame.Exame {
					fmt.Println("Digite a data do exame")
					imprimirDatasExames(exames)
					fmt.Scanln(&data)
					if data == exame.Data {
						fmt.Println("Digite a hora do exame")
						imprimirHorasExames(exames)
						fmt.Scanln(&exame.Hora)
						fmt.Println("Exame marcada com sucesso")
						fmt.Scanln()
						screen.Clear()
						return exame
					}
				}
			}

		}
	}
	if paciente.Nome == nome && paciente.CPF == cpf {

	} else {
		fmt.Println("Paciente não encontrado")

	}
	return exame
}

// Imprimir datas consultas
func imprimirDatasConsultas(consultas []Consulta) {
	for i, consulta := range consultas {
		fmt.Println("Consulta ID: [", i+1, "]\nData: ", consulta.Data)
	}
}

// Imprimir horario consulta
func imprimirHorasConsultas(consultas []Consulta) {
	for i, consulta := range consultas {
		fmt.Println("Consulta ID: [", i+1, "]\nHora: ", consulta.Hora)
	}
}

// Imprimir datas exames
func imprimirDatasExames(exames []Exame) {
	for i, exame := range exames {
		fmt.Println("Exame ID: [", i+1, "]\nData: ", exame.Data)
	}
}

// Imprimir horario exame
func imprimirHorasExames(exames []Exame) {
	for i, exame := range exames {
		fmt.Println("Exame ID: [", i+1, "]\nHora: ", exame.Hora)
	}
}

// Imprimir todas as consultas marcadas
func imprimirConsultas(consulta []Consulta) {

	if len(consulta) <= 0 {
		fmt.Println("Nao ha consultas marcadas")
	} else {
		fmt.Println("Consultas marcadas")
		fmt.Println("=================")
		for i := 0; i < len(consulta); i++ {
			fmt.Println("Paciente: ", consulta[i].Paciente)
			fmt.Println("Medico: ", consulta[i].Medico)
			fmt.Println("Data: ", consulta[i].Data)
			fmt.Println("Hora: ", consulta[i].Hora)
		}
	}
}

// Imprimir exames marcados
func imprimirExames(exame []Exame) {
	if len(exame) == 0 {
		fmt.Println("Nao ha exames marcados")
	} else {
		fmt.Println("Exames marcados")
		fmt.Println("==============")
		for i := 0; i < len(exame); i++ {
			fmt.Println("Paciente: ", exame[i].Paciente)
			fmt.Println("Exame: ", exame[i].Exame)
			fmt.Println("Data: ", exame[i].Data)
			fmt.Println("Hora: ", exame[i].Hora)
		}
	}
}

// Imprimir cirurgia marcadas
func imprimirCirurgias(cirurgia []Cirurgia) {
	if len(cirurgia) == 0 {
		fmt.Println("Nao ha cirurgias marcadas")
	} else {
		fmt.Println("Cirurgias marcadas")
		fmt.Println("================")
		for i := 0; i < len(cirurgia); i++ {
			fmt.Println("Paciente: ", cirurgia[i].Paciente)
			fmt.Println("Cirurgia: ", cirurgia[i].Cirurgia)
			fmt.Println("Data: ", cirurgia[i].Data)
			fmt.Println("Hora: ", cirurgia[i].Hora)
		}
	}
}

// Imprimir tratamentos marcados
func imprimirTratamentos(tratamento []Tratamento) {
	if len(tratamento) <= 0 {
		fmt.Println("Nao ha tratamentos marcados")
	} else {
		fmt.Println("Tratamentos marcados")
		fmt.Println("====================")
		for i := 0; i < len(tratamento); i++ {
			fmt.Println("Paciente: ", tratamento[i].Paciente)
			fmt.Println("Medico: ", tratamento[i].Medico)
			fmt.Println("Data: ", tratamento[i].Data)
			fmt.Println("Hora: ", tratamento[i].Hora)
		}
	}
}

// Funcao para cadsatrar novo usuario
func novoUsuario() Usuario {
	var usuario Usuario

	fmt.Println("CADASTRO DE USUARIO")
	fmt.Println("===================")

	fmt.Println("Digite o nome do usuario")
	fmt.Scanln(&usuario.Nome)

	fmt.Println("Digite o sobrenome do usuario")
	fmt.Scanln(&usuario.Sobrenome)

	fmt.Println("Digite o Nome de usuario")
	fmt.Scanln(&usuario.NomeUsuario)

	fmt.Println("Digite a senha")
	fmt.Scanln(&usuario.Senha)
	return usuario
}

// Menu para funcoes do usuario
func espacoUsuario() int {
	var opcao int
	fmt.Println("Digite o que deseja fazer")
	fmt.Println("[1] Cadastrar Paciente")
	fmt.Println("[2] Cadastrar Médico")
	fmt.Println("[3] Marcar consulta")
	fmt.Println("[4] Marcar exame")
	fmt.Println("[5] Marcar cirurgia")
	fmt.Println("[6] Marcar tratamento")
	fmt.Println("[7] Listar pacientes")
	fmt.Println("[8] Listar medicos")
	fmt.Println("[9] Consultas Marcadas")
	fmt.Println("[10] Exames Marcados")
	fmt.Println("[11] Cirurgias Marcadas")
	fmt.Println("[12] Tratamentos Marcados")
	fmt.Println("[0] Sair")
	fmt.Scanln(&opcao)
	return opcao
}

// Funcao para listar pacientes
func listarPacientes(paciente []Paciente) {
	fmt.Println("LISTA DE PACIENTES")
	fmt.Println("==================")
	for i, paciente := range paciente {
		fmt.Println("Paciente ID: [", i+1, "]\nNome:", paciente.Nome, paciente.Sobrenome, "\nIdade:", paciente.Idade, "\nCPF:", paciente.CPF, "\nTelefone:", paciente.Telefone, "\nEmail:", paciente.Email)
		fmt.Println()
	}
}

// Funcao para listar medicos
func listarMedicos(medicos []Medico) {
	fmt.Println("LISTA DE MÉDICOS")
	fmt.Println("================")
	for i, medico := range medicos {
		fmt.Println("Medico ID: [", i+1, "]\nNome:", medico.Nome, medico.Sobrenome, "\nCRM:", medico.CRM, "\nEspecialidade: ", medico.Especialidade)
		fmt.Println()
	}
}

// Funcao para cadastrar paciente
func cadastroPaciente() Paciente {
	var paciente Paciente
	fmt.Println("Digite o nome do paciente")
	fmt.Scanln(&paciente.Nome)

	fmt.Println("Digite o sobrenome do paciente")
	fmt.Scanln(&paciente.Sobrenome)

	fmt.Println("Digite o sexo do paciente")
	fmt.Scanln(&paciente.Sexo)

	fmt.Println("Digite a idade do paciente")
	fmt.Scanln(&paciente.Idade)

	fmt.Println("Digite o CPF do paciente")
	fmt.Scanln(&paciente.CPF)

	fmt.Println("Digite o email do paciente")
	fmt.Scanln(&paciente.Email)

	fmt.Println("Digite o telefone do paciente")
	fmt.Scanln(&paciente.Telefone)

	fmt.Println("Cadastro efetuado com sucesso")
	fmt.Println()
	return paciente
}

// Funcao para fazer o login do usuario
func login(usuario Usuario) {
	var lnome, lsenha string
login:
	for {
		fmt.Println("LOGIN")
		fmt.Println("=====")
		fmt.Println("Digite seu nome de usuário:")
		fmt.Scanln(&lnome)
		fmt.Println("Digite sua senha:")
		fmt.Scanln(&lsenha)

		if lnome == usuario.NomeUsuario && lsenha == usuario.Senha {
			fmt.Println("Login efetuado com sucesso")
			break
		} else {
			fmt.Println("Usuário ou senha inválidos")
			goto login
		}
	}
	println()
}

// Funcao para cadastro do medico
func cadastroMedico() Medico {
	var medico Medico
	fmt.Println("CADASTRO DE MEDICOS")
	fmt.Println("==================")
	fmt.Println("Digite o nome do medico")
	fmt.Scanln(&medico.Nome)
	fmt.Println("Digite o sobrenome do medico")
	fmt.Scanln(&medico.Sobrenome)
	fmt.Println("Digite o CRM do medico")
	fmt.Scanln(&medico.CRM)
	fmt.Println("Digite a especialidade do medico")
	fmt.Scanln(&medico.Especialidade)
	fmt.Println("Cadastro realizado com sucesso")
	return medico
}

// Funcao do Menu Principal
func menuPrincipal() int {
	var opcao int
	fmt.Println("[1] Cadastro")
	fmt.Println("[2] Login")
	fmt.Println("[3] Sair")
	fmt.Scanln(&opcao)
	return opcao
}

// Funcao para marcar tratamentos
func marcarTratamento() Tratamento {
	var tratamento Tratamento

	fmt.Println("Marcar Tratamento")
	fmt.Println("=================")
	fmt.Println("Qual paciente deseja marcar o tratamento:")
	fmt.Scanln(&tratamento.Paciente)
	fmt.Println("Medico:")
	fmt.Scanln(&tratamento.Medico)
	fmt.Println("Data:")
	fmt.Scanln(&tratamento.Data)
	fmt.Println("Hora:")
	fmt.Scanln(&tratamento.Hora)
	fmt.Println("Tratamento marcado com sucesso")
	fmt.Scanln()
	screen.Clear()
	return tratamento
}

// Funcao para marcar cirurgias
func marcarCirurgia() Cirurgia {
	var cirurgia Cirurgia

	fmt.Println("Marcar Cirurgia")
	fmt.Println("===============")
	fmt.Println("Qual paciente deseja marcar a cirurgia:")
	fmt.Scanln(&cirurgia.Paciente)
	fmt.Println("Cirurgia:")
	fmt.Scanln(&cirurgia.Cirurgia)
	fmt.Println("Data:")
	fmt.Scanln(&cirurgia.Data)
	fmt.Println("Hora:")
	fmt.Scanln(&cirurgia.Hora)
	fmt.Println("Cirurgia marcada com sucesso")

	return cirurgia
}
