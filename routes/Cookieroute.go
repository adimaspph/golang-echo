package routes

import "golang-echo/handler"

func cookieRoute() {

	// Menambahkan cookie saat mengakses endpoint /set-cookie
	e.GET("/set-cookie", handler.SetCookieHandler)

	// Membaca cookie saat mengakses endpoint /get-cookie
	e.GET("/get-cookie", handler.GetCookieHandler)

}
