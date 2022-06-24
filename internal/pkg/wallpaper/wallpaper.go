package wallpaper

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/tsivinsky/walle/internal/pkg/config"
)

func SetImage(imagePath string) error {
	var err error = nil

	err = exec.Command("killall", "swaybg").Run()

	err = exec.Command("swaybg", "-i", imagePath).Run()

	return err
}

func GetImageFromHTTPUri(uri string) ([]byte, string, error) {
	resp, err := http.Get(uri)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	contentType := resp.Header.Get("Content-Type")
	s := strings.Split(contentType, "/")
	ext := s[1]

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, "", err
	}

	return data, ext, nil
}

func SaveHTTPWallpaper(data []byte, ext string) (string, error) {
	configDir, err := config.GetConfigDir()
	if err != nil {
		return "", err
	}

	filePath := fmt.Sprintf("%s/wallpaper-from-internet.%s", configDir, ext)

	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return "", err
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		return "", err
	}

	return filePath, nil
}
