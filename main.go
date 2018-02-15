package main

import (
	"encoding/csv"
	"os"
	"strings"
)

// CsvIn Estrutura de recebimento do arquivo
type CsvIn struct {
	IDEntrevista string
	IDQuestao    string
	IMEI         string
	Questao      string
	Resposta     string
}

var (
	origin      = "*"
	destination = "*"
)

func main() {

	f, err := os.Open(origin)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	reader := csv.NewReader(f)
	reader.Comma = ';'

	lines, err := reader.ReadAll()

	if err != nil {
		panic(err)
	}

	var dataIn []CsvIn

	for _, line := range lines {

		line := CsvIn{
			IDEntrevista: line[0],
			IMEI:         line[1],
			IDQuestao:    line[2],
			Questao:      line[3],
			Resposta:     line[4],
		}

		dataIn = append(dataIn, line)
	}

	var dataOut [][]string

	for e := range dataIn {

		outLine := []string{
			dataIn[e].IDEntrevista,
			dataIn[e].IDQuestao,
			dataIn[e].IMEI,
			dataIn[e].Questao}

		var resposta = strings.Split(dataIn[e].Resposta, "&")

		for x := range resposta {
			outLine = append(outLine, resposta[x])
		}

		dataOut = append(dataOut, outLine)
	}

	fileOut, err := os.Create(destination)
	defer fileOut.Close()
	writer := csv.NewWriter(fileOut)
	writer.Comma = ';'
	defer writer.Flush()

	for _, value := range dataOut {
		writer.Write(value)
	}

}
