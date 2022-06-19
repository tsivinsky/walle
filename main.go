package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/tsivinsky/walle/internal/pkg/config"
	"github.com/tsivinsky/walle/internal/pkg/wallpaper"
)

// flags
var (
	image               string
	setImageImmediately bool
	restoringWallpaper  bool
)

func main() {
	flag.StringVar(&image, "i", "", "walle -i ./path/to/image.png")
	flag.BoolVar(&restoringWallpaper, "restore", false, "walle --restore")
	flag.BoolVar(&setImageImmediately, "s", false, "walle -i ./path/to/image.png -s")

	flag.Parse()

	err := config.CreateConfigPathIfNotExist()
	if err != nil {
		log.Fatal(err)
	}

	if image == "" && !restoringWallpaper {
		flag.Usage()
		os.Exit(1)
	}

	conf, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	if restoringWallpaper {
		imagePath := conf.ImagePath
		if imagePath == "" {
			fmt.Println("First, you need to call `walle -i ./path/to/image.png` to set some wallpaper")
			os.Exit(1)
		}

		err = wallpaper.SetImage(imagePath)
		if err != nil {
			log.Fatal(err)
		}
	} else if image != "" {
		if !strings.HasPrefix(image, "/") {
			cwd, err := os.Getwd()
			if err != nil {
				log.Fatal(err)
			}

			image = path.Join(cwd, image)
		}

		if setImageImmediately {
			err = wallpaper.SetImage(image)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			conf.ImagePath = image
			err = conf.Save()
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
