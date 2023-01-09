package main

import "fmt"

// coba variabel
// func main() {
// 	var fname, lname string = "Riki", "Widiantoro"
// 	m, n, o := 1, 2, 3
// 	item, price := "Mobil", 123000
// 	fmt.Println("")
// 	fmt.Println(fname, lname)
// 	fmt.Println(o, n, m)
// 	fmt.Println(item, "seharga :", price)
// 	fmt.Println("")
// }

// coba function
// string ucapan
func Helo() {
	ucapan, name := "Selamat Pagi", "Riki Widiantoro"
	fmt.Print(ucapan, ", ", name, "!")
}
// penjumlahan
func penjumlahan(a int, b int) {
	total := 0
	total = a + b
	fmt.Println(total)
}

func main() {
	Helo()
	penjumlahan(5, 9)
}