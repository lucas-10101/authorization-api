package utils

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"strings"
)

func LoadEnv() {
	var err error
	var fd *os.File
	var scanner *bufio.Scanner
	var lnNum int
	if fd, err = os.OpenFile(".env", os.O_RDONLY, 0644); err != nil {
		slog.Error("Failed to open .env file", "error", err)
		return
	}
	defer fd.Close()

	for scanner, lnNum = bufio.NewScanner(fd), 1; scanner.Scan(); lnNum++ {
		line := strings.TrimSpace(scanner.Text())

		if !strings.Contains(line, "=") {
			continue
		}

		if strings.Contains(line, "#") {
			line = strings.TrimSpace(strings.SplitN(line, "#", 2)[0])
		}

		lineParts := strings.SplitN(line, "=", 2)
		if len(lineParts) != 2 {
			slog.Error("Invalid line format in .env file", "line", lnNum)
			continue
		}

		err = os.Setenv(strings.TrimSpace(lineParts[0]), strings.TrimSpace(lineParts[1]))
		if err != nil {
			slog.Error("Failed to set environment variable", "key", lineParts[0], "error", err)
			continue
		}
	}

	if err != nil {
		panic(fmt.Sprintf("All environment variables must loaded from .env file, last error is: %v", err))
	}
}
