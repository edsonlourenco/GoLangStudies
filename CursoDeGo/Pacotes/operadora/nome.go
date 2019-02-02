package operadora

import (
	"strconv"

	"github.com/edsonlourenco/CursoDeGo/Pacotes/prefixo"
)

//NomeOperadora representa o nome da empresa de telecom
var NomeOperadora = "XPTO Telecom"

//prefixoDaCapitalOperadora prefixo mais o nome da operadora
var prefixoDaCapitalOperadora = strconv.Itoa(prefixo.Capital) + " " + NomeOperadora
