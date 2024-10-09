package main

import (
	"math/rand"
	"sort"
	"time"
)

type Lists struct {
	List1000   ListVariations
	List10000  ListVariations
	List50000  ListVariations
	List100000 ListVariations
}

type ListVariations struct {
	RandomList  []int
	SortedList  []int
	ReverseList []int
}

func NewListsNumbers() *Lists {
	rand.Seed(time.Now().UnixNano())

	return &Lists{
		List1000:   createListVariations(1000),
		List10000:  createListVariations(10000),
		List50000:  createListVariations(50000),
		List100000: createListVariations(100000),
	}
}

func createListVariations(size int) ListVariations {

	listaAleatoria := make([]int, size)
	for i := range listaAleatoria {
		listaAleatoria[i] = rand.Intn(size * 10)
	}

	listaOrdenada := make([]int, size)
	copy(listaOrdenada, listaAleatoria)
	sort.Ints(listaOrdenada)

	listaInversa := make([]int, size)
	copy(listaInversa, listaOrdenada)
	sort.Sort(sort.Reverse(sort.IntSlice(listaInversa)))

	return ListVariations{
		RandomList:  listaAleatoria,
		SortedList:  listaOrdenada,
		ReverseList: listaInversa,
	}
}
