package downloader

import (
	"fmt"
	"os"
	"os/exec"
)

func DownloadVideo(url, outputPath, format string) error {
	cmd := exec.Command("yt-dlp", "-f", format, "--output", outputPath+"%(title)s.%(ext)s", url)

	// Connect output to stdout for debugging or real-time logging
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to download video: %w", err)
	}

	return nil
}
