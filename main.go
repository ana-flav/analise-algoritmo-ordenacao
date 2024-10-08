package main

import (
	"fmt"
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
			resultMap[k][listSizes[i]] = make(map[string]map[string]interface{}, len(listNames))
			for j := range listNames {
				resultMap[k][listSizes[i]][listNames[j]] = make(map[string]interface{}, 3)
				resultMap[k][listSizes[i]][listNames[j]]["comparacoes"] = 0
				resultMap[k][listSizes[i]][listNames[j]]["trocas"] = 0
				resultMap[k][listSizes[i]][listNames[j]]["tempo"] = time.Now()
			}
		}
	}
	return resultMap
}

func runAlgorithms(l *Lists) map[string]map[string]map[string]map[string]interface{} {
	v := reflect.ValueOf(l).Elem()
	resultMap := initResultMap()

	for i := 0; i < v.NumField(); i++ {
		listas := v.Field(i)
		for j := 0; j < listas.NumField(); j++ {
			listaAny := listas.Field(j).Addr().Interface()
			fmt.Println(listas.Field(j).Cap())
			lista := listaAny.(*[]int)
			var listaCopia = make([]int, len(*lista))
			copy(listaCopia, *lista)

			for k, v := range mapAlgorithms {
				fmt.Println("---", k, "---")
				c, t, d := executeAlgorithm(v, listaCopia)
				resultMap[k][listSizes[i]][listNames[j]]["comparacoes"] = c
				resultMap[k][listSizes[i]][listNames[j]]["trocas"] = t
				duration := float64(d) / float64(time.Millisecond)
				resultMap[k][listSizes[i]][listNames[j]]["tempo"] = duration
				// fmt.Printf("Tipo de lista: %v - %v, Algoritmo: %v -- \nComparações: %v\nTrocas: %v\nDuração: %v\n\n",
				// 	listNames[j], listSizes[i], k, c, t, duration,
				// )
			}
		}
	}

	return resultMap
}

func main() {
	lists := NewListsNumbers()
	res := runAlgorithms(lists)
	// resJson, err := json.Marshal(res)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(string(resJson))
	// var lista = []int{5, 6, 9, 3, 1, 10, 8, 2, 7, 4}
	// a.SelectionSort(lista)
	plotGraphs(res)
}
