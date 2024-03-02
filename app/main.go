package main

// imports
import (
	"fmt"

	"github.com/inancgumus/screen"
)

// Struct do usuario

// Funcao main do programa
func main() {
	// inicializando variaveis
	var opcao int
	// inicializando medicos
	medicos := make([]Medico, 0)
	medicos = append(medicos, Medico{Nome: "Maria", Sobrenome: "Silva", CRM: 123456, Especialidade: "Cardiologia"})
	medicos = append(medicos, Medico{Nome: "Hans", Sobrenome: "Gruber", CRM: 123123, Especialidade: "Dermatologia"})
	medicos = append(medicos, Medico{Nome: "Patrick", Sobrenome: "Carlos", CRM: 88827, Especialidade: "Neurologia"})
	medicos = append(medicos, Medico{Nome: "Sofia", Sobrenome: "Pera", CRM: 92827, Especialidade: "Ortopedia"})

	//inicializando pacientes
	pacientes := make([]Paciente, 0)

	//inicializando consultas
	consultas := make([]Consulta, 0)
	consultas = append(consultas, Consulta{Data: "20/03/2024", Hora: []string{"10:00", "10:30", "11:00", "11:30", "12:00", "12:30"}})
	consultas = append(consultas, Consulta{Data: "22/03/2024", Hora: []string{"10:00", "10:30", "11:00", "11:30", "12:00", "12:30"}})
	consultas = append(consultas, Consulta{Data: "21/03/2024", Hora: []string{"10:00", "10:30", "11:00", "11:30", "12:00", "12:30"}})

	usuarios := make([]Usuario, 0)
	cirurgias := make([]Cirurgia, 0)
	exames := make([]Exame, 0)
	tratamentos := make([]Tratamento, 0)

	for opcao != 3 {
	loginpart:
		screen.Clear()
		opcao = menuPrincipal() //switch para a area de login
		//fmt.Scanln(&opcao)
		screen.Clear()
		switch opcao {
		case 1:
			usuario := novoUsuario()
			usuarios = append(usuarios, usuario)
		case 2:
			login(usuarios[0])
			goto espaco
		case 3:
			fmt.Println("Saindo...")
		default:
			fmt.Println("Opção inválida")
			goto loginpart
		}
	}
	// goto para o espaco do usuario

espaco:
	opcao = espacoUsuario()
	//fmt.Scanln(&opcao)
	screen.Clear()
	switch opcao { // switch para as opcoes do espaco do usuario
	case 1:
		screen.Clear()
		paciente := cadastroPaciente()
		pacientes = append(pacientes, paciente)
		goto espaco
	case 2:
		screen.Clear()
		medico := cadastroMedico()
		medicos = append(medicos, medico)
		goto espaco
	case 3:
		screen.Clear()
		consulta := marcarConsulta(medicos, pacientes, consultas)
		consultas = append(consultas, consulta)
		goto espaco
	case 4:
		screen.Clear()
		exame := marcarExame(pacientes, exames)
		exames = append(exames, exame)
		goto espaco
	case 5:
		screen.Clear()
		cirurgia := marcarCirurgia()
		cirurgias = append(cirurgias, cirurgia)
		goto espaco
	case 6:
		screen.Clear()
		tratamento := marcarTratamento()
		tratamentos = append(tratamentos, tratamento)
		goto espaco
	case 7:
		screen.Clear()
		listarPacientes(pacientes)
		goto espaco
	case 8:
		screen.Clear()
		listarMedicos(medicos)
		goto espaco
	case 9:
		screen.Clear()
		imprimirConsultas(consultas)
		goto espaco
	case 10:
		screen.Clear()
		imprimirExames(exames)
		goto espaco
	case 11:
		screen.Clear()
		imprimirCirurgias(cirurgias)
		goto espaco
	case 12:
		screen.Clear()
		imprimirTratamentos(tratamentos)
		goto espaco
	case 0:
		fmt.Println("Saindo...")
		return // Sai do programa
	default:
		fmt.Println("Opção inválida")
		goto espaco
	}

}
