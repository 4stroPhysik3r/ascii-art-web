package main

import (
	"fmt"
	"html/template"
	"net"
	"net/http"
	"os"
	"strings"
)

var tmpl *template.Template

func main() {
	if len(os.Args) != 1 {
		fmt.Println("Usage: go run .") // user instruction
		return
	}
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/result", ResultHandler)
	tmpl = template.Must(template.ParseGlob("templates/*.html"))

	l, err := net.Listen("tcp", ":3030")
	if err != nil {
		fmt.Println("ERROR: Port is already in use")
		return
	}
	fmt.Println("Starting the server on: http://localhost:3030")
	err = http.Serve(l, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func ResultHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	input := r.FormValue("input")
	banner := r.FormValue("banner")
	if banner != "shadow" && banner != "standard" && banner != "thinkertoy" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	input = AsciiArt(input, banner) //Gets the ascii art value from our code
	FormData := struct {
		TextValue  string
		BannerType string
	}{
		TextValue:  input,
		BannerType: banner,
	}
	tmpl.ExecuteTemplate(w, "result.html", FormData)
}
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" { //404 page
		w.WriteHeader(404)
		tmpl.ExecuteTemplate(w, "404.html", nil)
		return
	}
	tmpl.ExecuteTemplate(w, "index.html", nil)
}

func AsciiArt(input string, banner string) string { // AsciiArt function
	var output string
	if input == "\\n" || input == "" { // convert "\n" to newline
		output += "\n"
	} else if input == "\\n\\n" { // handle "\n\n" input
		output = "\n\n"
	} else if input != "" {
		bannerBytes, err := os.ReadFile("banners/" + banner + ".txt") // read Font File
		if err != nil {
			output += input + "Error reading font file" + banner // error handling
			return output
		}
		lines := strings.Split(string(bannerBytes), "\n") // split Font File
		input = strings.ReplaceAll(input, "\\n", "\n")    // handle "\n" input
		input = strings.ReplaceAll(input, "\r", "\n")     // handle newline in input
		parts := strings.Split(input, "\n")               // split Input
		for _, line := range parts {
			for i := 1; i < 9; i++ {
				for _, ascii := range line {
					if ascii < 32 || ascii > 126 { // handle invalid characters
						output = "Invalid input."
						return output
					}
					output += lines[(ascii-32)*9+rune(i)]
				}
				if line == "" { // handle empty line
					output += "\n"
					break
				}
				output += "\n"
			}
		}
	} else {
		return "Please type something\nin the box above."
	}
	return output
}
