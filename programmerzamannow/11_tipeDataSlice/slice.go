package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}
	//[...] -> jumlah array flexible

	var slice1 = months[4:7]
	//keterangan
	/*
		months[4:7]
		4 itu pointer awal, array dengan index 4
		7 adalah pointer akhir, array data ke 7 / array dengan index n-1 atau 7 - 1 = 6

		data yang diambil adalah dari data array dengan index 4 sampai data ke 7 / data dengan index 6 yaitu adar 3 data dan
		dengan kapacity dari index 4 sampai ke index akhir 12 yaitu 8
	*/

	/* Function Slice
	len(slice) 							-> untuk mendapatkan panjang slice
	cap(slice) 							-> untuk mendapatkan kapasitas slice
	append(slice, data) 				-> membuat slice baru dengan menambahkan data ke posisi terakhir slice, jika kapasitas sudah penuh maka akan membuat array baru
	make([]TypeData, length, capacity) 	-> membuat slice baru
	copy(destination, source) 			-> menyalin slice dari source ke destination
	*/

	/* Membuat Slice Dari Array
	array[low:high] -> Membuat Slice dari Array dimulai index low sampai index (high - 1)
	array[low:]		-> Membuat Slice dari array dimulai index low sampai index akhir array
	array[:high]	-> Membuat Slice dari array dimulai index 0 sampai index (high - 1)
	array[:]		-> Membuat Slice dari array dimulai index 0 sampai index akhir array
	*/

	fmt.Println("months = ", months)
	fmt.Println("slice = ", slice1)
	fmt.Println("Lenght slice =", len(slice1))
	fmt.Println("Capacity slice =", cap(slice1))
	fmt.Println("============================")
	fmt.Println()

	var slice2 = make([]string, 2, 2)
	slice2[0] = "Rio"
	slice2[1] = "Stefanus"
	fmt.Println(slice2)
	fmt.Println("Length slice2 =", len(slice2))
	fmt.Println("Cap slice2 =", cap(slice2))

	fmt.Println("============================")
	slice2 = append(slice2, "Antonius")
	fmt.Println(slice2)
	fmt.Println("Length slice2 =", len(slice2))
	fmt.Println("Cap slice2 =", cap(slice2))

	/*
		Ketika kapasitas slice sudah full, ketika diappend, maka kapasitas akan didouble dari kapasitas sebelumnya (2 * n) dan membuat array baru
		cth kapasitas max sebelumnya 2 maka kapasitas slice akan berubah menjadi (2 * n) ->  2 * 2 = 4
	*/

	fmt.Println("============================")
	slice2 = append(slice2, "Jozef")
	fmt.Println(slice2)
	fmt.Println("Length slice2 =", len(slice2))
	fmt.Println("Cap slice2 =", cap(slice2))
	/*
		disiini kapasitasn slice 4, dan value slice belum melebihi batas, jadi kapasitas belum berubah
	*/

	fmt.Println("============================")
	slice2 = append(slice2, "Jozef2")
	fmt.Println(slice2)
	fmt.Println("Length slice2 =", len(slice2))
	fmt.Println("Cap slice2 =", cap(slice2))

	/*
		disini kapasitas sudah melewati batas max sebelumnya yaitu 4 maka kapasitas max slice berubah (2 * n) ->  2 * 4 = 8
	*/
	fmt.Println("=================================")

	Days := [...]string{
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Jumat",
		"Sabtu",
		"Minggu",
	}

	fmt.Println(Days)

	SliceDay := Days[2:4]
	fmt.Println("len SliceDay =", len(SliceDay))
	fmt.Println("Cap SliceDay =", cap(SliceDay))
	fmt.Println(SliceDay)
	fmt.Println("=================================")
	SliceDay = append(SliceDay, "JumatTes")
	SliceDay[0] = "RabuTes"
	fmt.Println("len SliceDay after append =", len(SliceDay))
	fmt.Println("Cap SliceDay after append =", cap(SliceDay))
	fmt.Println("SliceDay after append =", SliceDay)
	fmt.Println("Days after append = ", Days)

	fmt.Println("=================================")
	SliceDay2 := Days[5:]
	fmt.Println("len SliceDay2 =", len(SliceDay2))
	fmt.Println("Cap SliceDay2 =", cap(SliceDay2))
	fmt.Println(Days)
	fmt.Println(SliceDay2)
	fmt.Println("=================================")
	SliceDay2 = append(SliceDay2, "SeninNext")
	SliceDay2[0] = "SabtuTest"
	fmt.Println("len SliceDay2 after append =", len(SliceDay2))
	fmt.Println("Cap SliceDay2 after append =", cap(SliceDay2))
	fmt.Println("SliceDay2 after append =", SliceDay2)
	fmt.Println("Days after append = ", Days)
	fmt.Println("=================================")
	newSlice := make([]string, len(SliceDay2), cap(SliceDay2))
	copy(newSlice, SliceDay2)
	fmt.Println("Copy Slice =", newSlice)
	/*
		Perbedaan pembuatan Array dan Slice
		Array
			[n]TypeData
			[...]TypeData
		Slice
			[]TypeData
	*/

}
