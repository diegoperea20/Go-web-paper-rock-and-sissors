package handlers

import (
	"encoding/json"

	"log"
	"net/http"
	"strconv"
	"text/template"
	"webgopaper/rps"
)

const(
	templateDir = "templates/"
	templateBase = templateDir + "base.html"
)
func renderTemplate(w http.ResponseWriter, page string,data any) {
	tpl := template.Must(template.ParseFiles( templateBase, templateDir + page))

	err := tpl.ExecuteTemplate(w, "base", data)//enviar datos 
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
type Player struct {
	Name string
}

var player Player
func Index(w http.ResponseWriter, r *http.Request){
	restartValues()
	renderTemplate(w, "index.html", nil)
}

func NewGame(w http.ResponseWriter, r *http.Request){
	restartValues()
	renderTemplate(w, "new-game.html", nil)
}

func Game(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		player.Name = r.Form.Get("name")
	}
	if player.Name == "" {
		http.Redirect(w, r, "/new", http.StatusFound)
		return
		
	}
	renderTemplate(w, "game.html", player)
}

func Play(w http.ResponseWriter, r *http.Request){
	playerChoice , _:= strconv.Atoi(r.URL.Query().Get("c"))
	result := rps.PlayRound( playerChoice)
	
	out , err := json.MarshalIndent(result, "", " ")
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func About(w http.ResponseWriter, r *http.Request){
	restartValues()
	renderTemplate(w, "about.html", nil)
}

func restartValues(){
	player.Name = ""
	rps.ComputerScore = 0
	rps.PlayerScore = 0
}