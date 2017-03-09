package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"html/template"
	"regexp"
	"errors"
	
)

// Caching template
var templates = template.Must(template.ParseFiles("edit.html", "view.html"))
	var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
	// page struct
type Page struct {
	Title string
	Body  []byte
}

func readMongo(){
	session, err := mgo.Dial(url)
	
	c := session.DB(database).C(collection)
	err := c.Find(query).One(&result)
}

// common load page
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// view page
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
    p, err := loadPage(title)
    if err != nil {
        http.Redirect(w, r, "/edit/"+title, http.StatusFound)
        return
    }
    renderTemplate(w, "view", p)
}

// edit page
func editHandler(w http.ResponseWriter, r *http.Request, title string) {
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
    //t, _ := template.ParseFiles("edit.html")
    //t.Execute(w, p)
    renderTemplate(w, "edit", p)
}

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
    m := validPath.FindStringSubmatch(r.URL.Path)
    if m == nil {
        http.NotFound(w, r)
        return "", errors.New("Invalid Page Title")
    }
    return m[2], nil // The title is the second subexpression.
}

// render template hadler 
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    err := templates.ExecuteTemplate(w, tmpl+".html", p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}


// search restarents
func searchHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/search/"):]
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
    fmt.Fprintf(w, "<h1>Search %s</h1>"+
        "<form action=\"/viewall/%s\" method=\"POST\">"+
        "<textarea name=\"body\">%s</textarea><br>"+
        "<input type=\"submit\" value=\"ViewAll\">"+
        "</form>",
        p.Title, p.Title, p.Body)
}

// search restarents
func viewAllHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/viewall/"):]
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
    fmt.Fprintf(w, "<h1>View All the restaurants.</h1>"+
        "<form action=\"/search/%s\" method=\"POST\">"+
        "<div>show data</div><br>"+
        "<input type=\"submit\" value=\"Back\">"+
        "</form>",
        p.Title, p.Title, p.Body)
}
// make common handler and call from each http request
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        m := validPath.FindStringSubmatch(r.URL.Path)
        if m == nil {
            http.NotFound(w, r)
            return
        }
        fn(w, r, m[2])
    }
}

func main() {
	
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
    //http.HandleFunc("/search/", makeHandler(searchHandler))
	http.HandleFunc("/search/", searchHandler)
	http.HandleFunc("/viewall/", viewAllHandler)
	
	http.ListenAndServe(":8080", nil)
}