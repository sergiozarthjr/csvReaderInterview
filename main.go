package main

import (
	"encoding/csv"
	"os"
	"strings"
)

type CsvIn struct {
	IdEntrevista string
	IdEntrevistador string
	IMEI string
	Questao string
	Resposta string
}

func main() {

	filename := "*"

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	reader := csv.NewReader(f)
	reader.Comma = ';'

	lines, _ := reader.ReadAll()

	if err != nil {
		panic(err)
	}

	var dataIn []CsvIn

	for _, line := range lines {
		line := CsvIn{
			IdEntrevista:    line[0],
			IdEntrevistador: line[1],
			IMEI:            line[2],
			Questao:         line[3],
			Resposta:        line[4],
		}
		dataIn = append(dataIn, line)
	}

	var dataOut [][]string

	for e := range dataIn {
		var resposta = strings.Split(dataIn[e].Resposta, "&")

		resposta1 := ""
		resposta2 := ""
		resposta3 := ""
		resposta4 := ""
		resposta5 := ""
		resposta6 := ""
		resposta7 := ""
		resposta8 := ""
		resposta9 := ""
		resposta10 := ""
		resposta11:= ""

		if len(resposta) > 0{
			resposta1 = resposta[0]
		}

		if len(resposta) > 1{
			resposta2 = resposta[1]
		}

		if len(resposta) > 2{
			resposta3 = resposta[2]
		}

		if len(resposta) > 3{
			resposta4 = resposta[3]
		}

		if len(resposta) > 4{
			resposta5 = resposta[4]
		}

		if len(resposta) > 5{
			resposta6 = resposta[5]
		}

		if len(resposta) > 6{
			resposta7 = resposta[6]
		}

		if len(resposta) > 7{
			resposta8 = resposta[7]
		}

		if len(resposta) > 8{
			resposta9 = resposta[8]
		}

		if len(resposta) > 9{
			resposta10 = resposta[9]
		}

		if len(resposta) > 10 {
			resposta11 = resposta[10]
		}

		outLine := []string{
			dataIn[e].IdEntrevista,
			dataIn[e].IdEntrevistador,
			dataIn[e].IMEI,
			dataIn[e].Questao,
			resposta1,
			resposta2,
			resposta3,
			resposta4,
			resposta5,
			resposta6,
			resposta7,
			resposta8,
			resposta9,
			resposta10,
			resposta11,
		}

		dataOut = append(dataOut, outLine)
	}

	fileOut, err := os.Create("/home/zarthjr/Downloads/exportacao_entrevistas_corrigido_pos_processamento_2.csv")
	defer fileOut.Close()
	writer := csv.NewWriter(fileOut)
	writer.Comma = ';'
	defer writer.Flush()

	for _, value := range dataOut {
		writer.Write(value)
	}

}
	//Código usado para descobrir o número de colunas necessárias
	/*
	r := regexp.MustCompile("&")
	var control = 0
	var str string
	for e := range data {
		var count = len(r.FindAllStringIndex(data[e].Resposta, -1))
		if count > control{
			control = count
			str = data[e].Resposta
		}
	}

	fmt.Print(control, " , ", str)
	*/
