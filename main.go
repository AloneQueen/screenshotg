package main

import (
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
)

func main() {
	log.Fatal(http.ListenAndServe(":8000", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Query().Get("url")
		args := []string{
			"--headless", "--screenshot", "shot.png",
		}
		windowSize := r.URL.Query().Get("window_size")
		if windowSize != "" {
			args = append(args, "--window-size", windowSize)
		}
		args = append(args, url)
		cmd := exec.Command("firefox", args...)

		go func() {
			time.Sleep(time.Second * 10)
			cmd.Process.Kill()
		}()
		_, err := cmd.CombinedOutput()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}

		img, err := os.ReadFile("shot.png")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}

		w.Header().Add("content-type", "image/png")
		w.Write(img)
	})))
}
