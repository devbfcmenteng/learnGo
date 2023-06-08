package main

import _ "bfcTest/packages/database" //akan langsung mengeksekui function init(),( _ ) didepan nama packageuntuk memperbolehkan function tidak digunakan

func main() {
	/*
		+	Kadang kita hanya ingin menjalankan funciton init di package tanpa harus mengeksekusi function lain yg ada di package
		+	Secara default Go-Lang akan memberikan komplain / error untuk package yang di import namun tidak digunakan,
			Untuk menangani hal tersebut, kita bisa menggunakan blank indetifier atau underscore ( _ ) sebelum nama package ketika melakukan import lihat contoh di paling atas
	*/

}
