package main

import (
	"fmt"
	"log"
	"net/http"
	"server/pkg/setting"
	"server/routers"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}


	if setting.RunMode == "release" {
		s.ListenAndServeTLS(setting.HTTPS_CRT, setting.HTTPS_KEY)
	} else if setting.RunMode == "debug" {
		s.ListenAndServe()
	} else {
		log.Fatalln("Unknown RunMode")
	}
}
