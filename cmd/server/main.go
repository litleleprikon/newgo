package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/litleleprikon/newgo/internal/server"
	"github.com/pkg/profile"
)

func main() {
	profilePath := fmt.Sprintf("./pprof/%s", time.Now().Format("15:04:05"))
	defer profile.Start(
		profile.CPUProfile,
		profile.ProfilePath(profilePath)).
		Stop()

	mux := server.New()
	log.Println("Starting server")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
