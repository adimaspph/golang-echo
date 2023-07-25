package routes

import "golang-echo/handler"

func cookieRoute() {

	e.GET("/set-cookie", handler.SetCookieHandler)

	e.GET("/get-cookie", handler.GetCookieHandler)

}
