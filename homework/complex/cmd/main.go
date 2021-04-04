package main

import (
	"fmt"
	"github.com/exitialis/workshop/homework/complex/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Сервис админки??
// При подгрузке пользователя, идет запрос в микросервис конфигов пользователя
// который в ответ возвращает настройки пользователя в сервисе
// К примеру уровень громкости, а также запрос в сервис понравившегося, который возвращает список плейлистов


// В этом сервисе специально допущено несколько проблем с производительностью. Найди и почини их все!
func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	r := mux.NewRouter()

	h := handlers.ProfileHandler{}

	r.HandleFunc("/", h.Handle)

	err := http.ListenAndServe(":8890", r)
	if err != nil {
		log.Fatal(err)
	}
}
