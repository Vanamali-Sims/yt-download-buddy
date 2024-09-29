package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func downloadVideo(url string) error {
	cmd := exec.Command("yt-dlp", url)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to download video %v", err)
	}

	return nil
}

func main() {
	videoURL := ""
	fmt.Println("Starting download !")
	err := downloadVideo(videoURL)
	if err != nil {
		log.Fatalf("Error downloading video %v", err)
	}
	fmt.Println("Download completed!")
}
