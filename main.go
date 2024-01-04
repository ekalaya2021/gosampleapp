package main
import (
    "fmt"
    "html"
    "log"
    "net/http"
    "strings"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        res := strings.TrimPrefix(html.EscapeString(r.URL.Path),"/")
        fmt.Fprintf(w, "Hello, %q apa khabar", res)
    })

    log.Fatal(http.ListenAndServe(":8081", nil))

}
