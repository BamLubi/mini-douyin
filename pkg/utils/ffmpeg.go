package utils

import (
	"os/exec"
)

func GetVideoCover(videoPath string, coverPath string) error {
	cmd := exec.Command("ffmpeg", "-i", videoPath, "-ss", "00:00:01.000", "-vframes", "1", coverPath)
	return cmd.Run()
}
