package main

import (
	"Udemy/config"
	"Udemy/pkg/handeler"
	"Udemy/pkg/renderer"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"net/http"
	"time"
)

var session *scs.SessionManager

func main() {
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
		Addr:    ":8080",
		Handler: routes(&app),
	}
	err := srv.ListenAndServe()
	if err != nil {
		return
	}
}
