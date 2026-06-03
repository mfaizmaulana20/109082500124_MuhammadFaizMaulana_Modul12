# <h1 align="center">Laporan Praktikum Modul 12 - ... </h1>
<p align="center">[MuhammadFaizMaulana] - [109082500124]</p>

## Unguided 

### 1. [Soal]
#### soal1.go

```go
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
	fmt.Println()

	for calon := 1; calon <= 20; calon++ {
		jumlah := hitungSuara(T, n, calon)
		if jumlah > 0 {
			fmt.Printf("%d: %d\n", calon, jumlah)
		}
	}
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/mfaizmaulana20/109082500124_MuhammadFaizMaulana_Modul9/blob/main/modul%209/output/soal%201.png)
[Program ini membaca data suara hingga ditemukan angka 0 sebagai penanda akhir input, kemudian menghitung jumlah seluruh suara yang masuk dan memvalidasi suara yang sah, yaitu suara dengan nomor calon antara 1 sampai 20. Semua suara sah disimpan ke dalam array, lalu untuk setiap calon dari nomor 1 hingga 20 program menggunakan fungsi hitungSuara() yang melakukan sequential search untuk menghitung berapa kali nomor calon tersebut muncul dalam array. Setelah itu program menampilkan jumlah suara masuk, jumlah suara sah, serta daftar calon yang memperoleh suara beserta jumlah suara yang diperolehnya.]

### 2. [Soal]
#### soal2.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_2](https://github.com/mfaizmaulana20/109082500124_MuhammadFaizMaulana_Modul9/blob/main/modul%209/output/soal%202.png)
[Program ini membaca data suara hingga ditemukan angka 0 sebagai penanda akhir input, kemudian menghitung jumlah seluruh suara yang masuk dan jumlah suara sah, yaitu suara dengan nomor calon antara 1 sampai 20. Suara yang sah disimpan ke dalam array, lalu fungsi hitungSuara() digunakan untuk menghitung jumlah suara yang diperoleh setiap calon dengan metode sequential search. Setelah seluruh suara dihitung, program menentukan ketua RT sebagai calon dengan jumlah suara terbanyak dan wakil ketua sebagai calon dengan jumlah suara terbanyak kedua. Jika terdapat beberapa calon dengan jumlah suara yang sama, calon dengan nomor yang lebih kecil akan dipilih terlebih dahulu sebagai ketua, kemudian nomor terkecil berikutnya sebagai wakil. Terakhir, program menampilkan jumlah suara masuk, jumlah suara sah, nomor ketua RT, dan nomor wakil ketua RT yang terpilih.]

### 3. [Soal]
#### soal3.go

```go
package main

import "fmt"

const NMAX = 1000000

var data [NMAX]int

func main() {
	var n, k int

	fmt.Scan(&n, &k)

	isiArray(n)

	hasil := posisi(n, k)

	if hasil == -1 {
		fmt.Println("TIDAK ADA")
	} else {
		fmt.Println(hasil)
	}
}

func isiArray(n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}
}

func posisi(n, k int) int {
	kiri := 0
	kanan := n - 1

	for kiri <= kanan {
		tengah := (kiri + kanan) / 2

		if data[tengah] == k {
			return tengah
		} else if k < data[tengah] {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}

	return -1
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_2](https://github.com/mfaizmaulana20/109082500124_MuhammadFaizMaulana_Modul9/blob/main/modul%209/output/soal%204.png)
[Program ini membaca dua buah input, yaitu n sebagai jumlah data dan k sebagai nilai yang akan dicari, kemudian mengisi array data dengan n bilangan yang sudah terurut membesar melalui prosedur isiArray(). Selanjutnya fungsi posisi() melakukan pencarian menggunakan metode Binary Search, yaitu dengan membandingkan nilai k dengan elemen tengah array dan mempersempit daerah pencarian ke kiri atau ke kanan hingga data ditemukan atau tidak ada lagi bagian array yang perlu diperiksa. Jika nilai k ditemukan, fungsi mengembalikan indeks posisinya yang dihitung mulai dari 0, sedangkan jika tidak ditemukan fungsi mengembalikan -1. Program kemudian menampilkan posisi data tersebut atau mencetak "TIDAK ADA" jika nilai k tidak terdapat dalam array.]

