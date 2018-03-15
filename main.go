package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    "golang.org/x/crypto/bcrypt"
    "os"
    "strconv"
)

type bpass struct {
    Password string
    Bpassword string
}

func AppendRainbow(password, bpassword string){
    text := fmt.Sprintf("%s, %s\n", password, bpassword)
    _, err := os.OpenFile("./rainbow", os.O_RDONLY|os.O_CREATE, 0666)
    if err != nil {
        panic(err)
    }
    f, err := os.OpenFile("./rainbow", os.O_APPEND|os.O_WRONLY, 0600)
    if err != nil {
        panic(err)
    }

    defer f.Close()

    if _, err = f.WriteString(text); err != nil {
        panic(err)
    }
}

func HashPassword(password string, bcost int) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcost)
	return string(bytes), err
}

func web(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method) //get request method
    if r.Method == "GET" {
        t, _ := template.ParseFiles("get_hash.gtpl")
        t.Execute(w, nil)
    } else {
        r.ParseForm()
        var bcost int
        bcost, err := strconv.Atoi(r.Form["bcost"][0])
        if err != nil {
            bcost = 4
        }
        Bpass, err := HashPassword(r.Form["password"][0], bcost)
        if err != nil {
            fmt.Println(err)
        }
        t, _ := template.ParseFiles("post_hash.gtpl")
        data := bpass{
            Password: r.Form["password"][0],
            Bpassword: Bpass,
        }
        AppendRainbow(r.Form["password"][0], Bpass)
        t.Execute(w, data)
    }
}

func safe(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        t, _ := template.ParseFiles("safe_get_hash.gtpl")
        t.Execute(w, nil)
    } else {
        r.ParseForm()
        var bcost int
        bcost, err := strconv.Atoi(r.Form["bcost"][0])
        if err != nil {
            bcost = 4
        }
        Bpass, err := HashPassword(r.Form["password"][0], bcost)
        if err != nil {
            fmt.Println(err)
        }
        t, _ := template.ParseFiles("safe_post_hash.gtpl")
        data := bpass{
            Password: r.Form["password"][0],
            Bpassword: Bpass,
        }
        t.Execute(w, data)
    }
}

func main() {
    http.HandleFunc("/", web)
    http.HandleFunc("/safe", safe)
    http.HandleFunc("/rainbow", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "./rainbow")
    })
    err := http.ListenAndServe(":9090", nil) // setting listening port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
