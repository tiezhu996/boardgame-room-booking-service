package main

import (
	"fmt"

	"boardgame/internal/config"
	"boardgame/internal/model"
	"boardgame/internal/service"
	"boardgame/internal/store"
	"boardgame/internal/worker"
)

func main() {
	cfg := config.Load()
	st := store.New()
	_ = st.AddRoom(model.Room{ID: "r1", Capacity: 4})
	svc := service.New(st)
	fmt.Printf("%s reserved=%d total=%d\n", cfg.AppName, svc.BatchReserve([]string{"r1", "r1", "r1"}), worker.New(st).Run())
}
