package main

import (
  "github.com/hoisie/web"
  "github.com/hoisie/mustache"
  "os"
  "strings"
  "regexp"
  "net/http"
)

var mongo = os.Getenv("MONGO_HOST")
var mongo_db = os.Getenv("MONGO_DB")
var projects = os.Getenv("PROJECTS")
var host = os.Getenv("HOST")
var kiosk = os.Getenv("KIOSK_MODE")

var emailMatcher, _ = regexp.Compile("^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\\.[A-Za-z]{2,6}$")


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

func handlePing() string{
  return "pong."
}

func main(){
  port := os.Getenv("PORT")
  if port == "" {
    port = "5000"
  }

  if host == "" {
    host = "http://localhost:"+port
  }

  if mongo == "" {
    mongo = "localhost"
  }

  if mongo_db == "" {
    mongo_db = "signup_dev"
  }

  if projects == "" {
    projects = "test"
  }


  web.Get("/", handlePing)
  web.Get("/script/(.+)/(.+)", handleScriptInjection)
  web.Post("/(.+)", handleSignup)
  web.Run("0.0.0.0:"+port)
}

