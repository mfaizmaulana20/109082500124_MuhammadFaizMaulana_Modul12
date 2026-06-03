package main

import "fmt"

const NMAX = 1000

type tabInt [NMAX]int

func hitungSuara(T tabInt, n int, calon int) int {
	count := 0

	for i := 0; i < n; i++ {
		if T[i] == calon {
			count++
		}
	}

	return count
}

func main() {
	var T tabInt
	var x int
	var n int
	var suaraMasuk int
	var suaraSah int

	for {
		fmt.Scan(&x)

		if x == 0 {
			break
		}

		suaraMasuk++

		if x >= 1 && x <= 20 {
			T[n] = x
			n++
			suaraSah++
		}
	}

	fmt.Println("Suara masuk:", suaraMasuk)
	fmt.Println("Suara sah:", suaraSah)

	ketua := 1
	wakil := 2

	for calon := 1; calon <= 20; calon++ {

		if hitungSuara(T, n, calon) > hitungSuara(T, n, ketua) {
			wakil = ketua
			ketua = calon

		} else if calon != ketua &&
			(hitungSuara(T, n, calon) > hitungSuara(T, n, wakil) ||
				wakil == ketua) {

			wakil = calon
		}
	}

	if hitungSuara(T, n, wakil) > hitungSuara(T, n, ketua) {
		ketua, wakil = wakil, ketua
	}

	fmt.Println("Ketua RT:", ketua)
	fmt.Println("Wakil ketua:", wakil)
}