package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/abyss-works/used-book/aladin"
	"github.com/abyss-works/used-book/handler"
)

//go:embed frontend/dist/*
var frontendFS embed.FS

func main() {
	// 환경변수
	port := env("PORT", "8080")
	ttbKey := env("ALADIN_TTB_KEY", "")

	// 알라딘 클라이언트
	aladinClient := aladin.NewClient(ttbKey)

	// API 핸들러
	h := handler.New(aladinClient)
	mux := http.NewServeMux()
	h.Register(mux)

	// 프론트엔드 정적 파일
	subFS, err := fs.Sub(frontendFS, "frontend/dist")
	if err != nil {
		log.Printf("frontend dist not found, serving API only: %v", err)
	} else {
		mux.Handle("/", http.FileServer(http.FS(subFS)))
	}

	addr := ":" + port
	log.Printf("used-book server starting on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("server error: %v", err)
	}
}

func env(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
