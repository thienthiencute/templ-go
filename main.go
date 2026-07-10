package main

import (
	"fmt"
	"net/http"
	"templ-go/test"

	"github.com/a-h/templ"
)

func main() {
	component := test.Hello("Thien Thien")
	// component.Render(context.Background(), os.Stdout)
	http.Handle("/", templ.Handler(component))

	fmt.Println("Listening on :5000")
	http.ListenAndServe(":5000", nil)
}
