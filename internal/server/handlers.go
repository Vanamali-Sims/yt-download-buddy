package server

import (
	"youtube-downloader/internal/config"
	"youtube-downloader/internal/downloader"

	"github.com/gofiber/fiber/v2"
)

func DownloadHandler(c *fiber.Ctx) error {
	videoURL := c.Query("url")
	if videoURL == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Video URL is required")
	}
	cfg := config.LoadConfig()

	err := downloader.DownloadVideo(videoURL, cfg.OutputPath, cfg.VideoFormat)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error downloading video: " + err.Error())
	}
	return c.SendString("Video downloaded successfully!")
}
