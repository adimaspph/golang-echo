package main

import "golang-echo/routes"

/*
Echo Framework Golang
Echo itu adalah sebuah framework web berbasis bahasa pemrograman Go (Golang). Dikembangkan oleh Labstack, Echo menyediakan alat-alat yang sederhana, ringan, dan efisien untuk membangun aplikasi web dengan cepat dan mudah. Framework ini populer karena kinerjanya yang tinggi, kemudahan penggunaan, dan dukungan yang kuat untuk routing, middleware, pengolahan permintaan, dan respons HTTP.

Penggunaan Echo:
Instalasi:
Sebelum menggunakan Echo, pastikan telah menginstal Go (Golang) di sistem komputer kalian. Setelah itu, dapat menginstal Echo dengan perintah berikut menggunakan Go Modules:
go get github.com/labstack/echo/v4


Contoh Penggunaan:
Berikut ini adalah contoh sederhana untuk membuat server HTTP menggunakan Echo:
package main
*/

func main() {
	// Membuat objek echo
	e := routes.New()

	// Menjalankan server HTTP di port 8080
	e.Logger.Fatal(e.Start(":8080"))
}
