package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/sandjuarezg/http-lol/struct_json"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/champion", champion)

	fmt.Println("Listening on localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func champion(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	defer fmt.Printf("Response from %s\n", r.URL.RequestURI())

	//get data from form of index
	r.ParseForm()
	var champ string = r.FormValue("champion")

	champ = strings.ToLower(champ)
	if strings.Contains(champ, "'") {
		champ = strings.Replace(champ, "'", "", -1)
		champ = strings.Title(champ)
	} else {
		champ = strings.Title(champ)
		champ = strings.Replace(champ, ".", "", -1)
		champ = strings.Replace(champ, " ", "", -1)
	}

	//make request to api
	var url string = fmt.Sprintf("http://ddragon.leagueoflegends.com/cdn/11.15.1/data/en_US/champion/%s.json", champ)

	resp, err := http.Get(url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		w.WriteHeader(resp.StatusCode)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	championData, err := struct_json.GetChampion(body, champ)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	t, err := template.ParseFiles("./static/html/champion.html")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	championData.Name = strings.ToUpper(championData.Name)
	championData.Title = strings.Title(championData.Title)

	t.Execute(w, championData)

}
