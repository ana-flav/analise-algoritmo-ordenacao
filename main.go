package main

import (
	"encoding/json"
	"fmt"
	"log"
	a "main/algoritms"
	"reflect"
	"time"
)

func executeAlgorithm(a a.SortAlg, list []int) (int, int, time.Duration) {
	return a(list)
}

func initResultMap() map[string]map[string]map[string]map[string]interface{} {
	var resultMap = make(map[string]map[string]map[string]map[string]interface{}, len(mapAlgorithms))
	for k, v := range mapAlgorithms {
		v = v
		resultMap[k] = make(map[string]map[string]map[string]interface{}, len(listSizes))
		for i := range listSizes {
			resultMap[k][listSizes[i]] = make(map[string]map[string]interface{}, len(listDistributions))
			for j := range listDistributions {
				resultMap[k][listSizes[i]][listDistributions[j]] = make(map[string]interface{}, 3)
				resultMap[k][listSizes[i]][listDistributions[j]]["comparacoes"] = 0
				resultMap[k][listSizes[i]][listDistributions[j]]["trocas"] = 0
				resultMap[k][listSizes[i]][listDistributions[j]]["tempo"] = time.Now()
			}
		}
	}
	return resultMap
}

func runAlgorithms(l *Lists) map[string]map[string]map[string]map[string]interface{} {
	v := reflect.ValueOf(l).Elem()
	resultMap := initResultMap()

	for k, alg := range mapAlgorithms {
		for i := 0; i < v.NumField(); i++ {
			listas := v.Field(i)
			for j := 0; j < listas.NumField(); j++ {
				listaAny := listas.Field(j).Addr().Interface()
				lista := listaAny.(*[]int)
				var listaCopia = make([]int, len(*lista))
				copy(listaCopia, *lista)

				c, t, d := executeAlgorithm(alg, listaCopia)

				resultMap[k][listSizes[i]][listDistributions[j]]["comparacoes"] = c
				resultMap[k][listSizes[i]][listDistributions[j]]["trocas"] = t
				resultMap[k][listSizes[i]][listDistributions[j]]["tempo"] = float64(d) / float64(time.Millisecond)
			}
		}
	}

	return resultMap
}

func main() {

	lists := NewListsNumbers()
	res := runAlgorithms(lists)
	resJson, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(resJson))
	plotGraphs(res, "Ordem aleatória", "tempo", "Qtd. de Trocas", "trocas-a")
	plotGraphs(res, "Ordem crescente", "trocas", "Qtd. de Trocas", "trocas-c")
	plotGraphs(res, "Ordem decrescente", "trocas", "Qtd. de Trocas", "trocas-d")
	plotGraphs(res, "Ordem aleatória", "comparacoes", "Qtd. de Comparacoes", "comp-a")
	plotGraphs(res, "Ordem crescente", "comparacoes", "Qtd. de Comparacoes", "comp-c")
	plotGraphs(res, "Ordem decrescente", "comparacoes", "Qtd. de Comparacoes", "comp-d")
	plotGraphs(res, "Ordem decrescente", "tempo", "Tempo de Execução", "tempo")

}
