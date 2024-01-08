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
        fmt.Fprintf(w, "안녕하세요 %q ", res)
    })

    log.Fatal(http.ListenAndServe(":8081", nil))

}
