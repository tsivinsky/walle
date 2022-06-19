package wallpaper

import "os/exec"

func SetImage(imagePath string) error {
	var err error = nil

	err = exec.Command("killall", "swaybg").Run()

	err = exec.Command("swaybg", "-i", imagePath).Run()

	return err
}
