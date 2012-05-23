package main

import (
	"github.com/hoisie/mustache"
	"github.com/hoisie/web"
	"net/http"
	"os"
	"regexp"
	"strings"
)


func env(key, def string) (res string) {
	if res = os.Getenv(key); res == "" {
		res = def
	}
	return
}



var (
	port		= env("PORT", "5000")
	mongo		= env("MONGO_HOST", "localhost")
	mongo_db	= env("MONGO_DB", "signup_dev")
	projects	= env("PROJECTS", "test")
	host		= env("HOST", "http://localhost:"+port)
	kiosk		= env("KIOSK_MODE", "")
)

var emailMatcher = regexp.MustCompile("^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\\.[A-Za-z]{2,6}$")

func handleSignup(ctx *web.Context, project string) {
	ctx.Header().Set("Access-Control-Allow-Origin", "*")
	ctx.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	ctx.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, X-Prototype-Version")

	email := strings.TrimSpace(ctx.Params["email"])
	project = strings.TrimSpace(project)

	if email != "" && project != "" && strings.Contains(projects, project) && emailMatcher.MatchString(email) {
		c := Contact{}
		if kiosk == "" {
			c.Create(project, email)
		}
		ctx.WriteString("OK")
		return
	}

	ctx.Abort(http.StatusNotAcceptable, "")
}

func handleScriptInjection(ctx *web.Context, project string, form_id string) string {
	ctx.Header().Set("Content-Type", "application/javascript")
	project = strings.TrimSpace(project)
	form_id = strings.TrimSpace(form_id)

	return mustache.RenderFile("inject.js.mustache", map[string]string{"project": project, "form_id": form_id, "host": host})
}

func handlePing() string {
	return "pong."
}

func main() {
	web.Get("/", handlePing)
	web.Get("/script/(.+)/(.+)", handleScriptInjection)
	web.Post("/(.+)", handleSignup)
	web.Run("0.0.0.0:" + port)
}
