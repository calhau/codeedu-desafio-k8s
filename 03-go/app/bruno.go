package main

import(
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func greeting(text string) string {
	data := fmt.Sprintf("<b>%s</b>",text)
	return data
}

func BrunoFunc(w http.ResponseWriter, r *http.Request){
	mensagem := greeting("Bruno Leal - Code.education Rocks!")
	//w.Write([]byte(mensagem))
	fmt.Fprintf(w,mensagem)
}

func main(){
	route := mux.NewRouter()
	route.HandleFunc("/", BrunoFunc)
	http.ListenAndServe(":8000", route)
}