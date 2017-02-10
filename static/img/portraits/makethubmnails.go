package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/nfnt/resize"
)

func main() {
	files := getAllImagesInDir(filepath.Join("."))
	fmt.Println("found ", len(files), " files")
	for _, filename := range files {
		fmt.Println("resizing file ", filename)
		file, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}

		// decode png into image.Image
		var img image.Image
		if path.Ext(filename) == ".jpg" {
			img, err = jpeg.Decode(file)
		} else {
			img, err = png.Decode(file)
		}
		if err != nil {
			log.Fatal(err)
		}
		file.Close()

		// resize to width 350 using Lanczos resampling
		// and preserve aspect ratio
		m := resize.Resize(300, 0, img, resize.Lanczos3)

		out, err := os.Create(filepath.Join("resized", filename))
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		// write new image to file
		if path.Ext(filename) == ".jpg" {
			jpeg.Encode(out, m, &jpeg.Options{Quality: 80})
		} else {
			png.Encode(out, m)
		}
	}

}

func getAllImagesInDir(dirname string) []string {
	var files []string
	d, err := os.Open(dirname)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer d.Close()
	fi, err := d.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, fi := range fi {
		if fi.Mode().IsRegular() {
			switch path.Ext(fi.Name()) {
			case ".png":
				files = append(files, fi.Name())
			case ".jpg":
				files = append(files, fi.Name())
			default:
				continue
			}
		}
	}
	return files
}
