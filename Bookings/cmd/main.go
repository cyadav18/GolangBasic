package main

import (
	"Bookings/config"
	"Bookings/pkg/constants"
	"Bookings/pkg/handeler"
	"Bookings/pkg/renderer"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"net/http"
	"time"
)

var session *scs.SessionManager

func main() {
	fmt.Println("Booking")
	var app config.AppConfig
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false
	app.Session = session
	app.TemplateCache, _ = renderer.CreateTemplateCash()
	fmt.Println(app.TemplateCache)
	renderer.NewTemplate(&app)
	repo := handeler.SetRepo(&app)
	handeler.NewHandlers(repo)
	srv := &http.Server{
		Addr:    constants.IP+constants.PORT,
		Handler: routes(&app),
	}
	err := srv.ListenAndServe()
	if err != nil {
		return
	}

}
