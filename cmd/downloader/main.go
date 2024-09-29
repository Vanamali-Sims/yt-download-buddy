package main

import (
	"flag"
	"fmt"
	"os"
	"youtube-downloader/internal/config"
	"youtube-downloader/internal/downloader"
	"youtube-downloader/internal/server"
)

func main() {
	// Check if we're running as a web server or command-line tool
	serverMode := flag.Bool("server", false, "Start in web server mode")
	url := flag.String("url", "", "YouTube video URL to download (used in CLI mode)")
	flag.Parse()

	if *serverMode {
		// Start the web server
		server.StartServer()
	} else {
		// Run command-line downloader
		if *url == "" {
			fmt.Println("YouTube URL is required in CLI mode. Use -url to specify the video URL.")
			os.Exit(1)
		}

		// Load config and download the video
		cfg := config.LoadConfig()
		err := downloader.DownloadVideo(*url, cfg.OutputPath, cfg.VideoFormat)
		if err != nil {
			fmt.Printf("Error downloading video: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Download completed successfully!")
	}
}
