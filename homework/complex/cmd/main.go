package main

import (
	"context"
	"github.com/exitialis/workshop/homework/complex/app"
	_ "net/http/pprof"
)

// Сервис избранного пользователя. Ручка для получения понравившихся плейлистов пользователя
// Возвращает все понравившиеся пользователю плейлисты вместе с песнями, исполнителями и альбомами.
// В этом сервисе специально допущено несколько проблем с производительностью. Найди и почини их все!
func main() {
	ctx, cancel := context.WithCancel(context.Background())

	app.StartApp(ctx, true)
	cancel()
}

