package controllers

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var t = time.Now()

func SaveImagen(nombre, imagen string) error {
	//index del string base 64
	coI := strings.Index(imagen, ",")

	//separar el string base 64
	rawImage := imagen[coI+1:]

	bas, _ := base64.StdEncoding.DecodeString(rawImage)

	res := bytes.NewReader(bas)

	//ruta absoluta
	path, _ := os.Getwd()

	date := fmt.Sprintf("%d-%02d-%02d",
		t.Year(), t.Month(), t.Day())

	newpath := filepath.Join(path + "/public/img/servicios/" + date)
	os.MkdirAll(newpath, os.ModePerm)

	err := imagenSave(nombre, imagen, newpath, coI, res)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func imagenSave(nombre, imagen, newpath string, coI int, res *bytes.Reader) error {
	switch strings.TrimSuffix(imagen[5:coI], ";base64") {
	case "image/jpeg":
		{
			jpg, err := jpeg.Decode(res)

			if err != nil {
				return err
			}

			fs, _ := os.OpenFile(newpath+"/"+nombre+".png", os.O_WRONLY|os.O_CREATE, 0777)

			jpeg.Encode(fs, jpg, &jpeg.Options{Quality: 75})
			return nil
		}
	case "image/png":
		{
			pngD, err := png.Decode(res)

			if err != nil {
				return err
			}

			fs, _ := os.OpenFile(newpath+"/"+nombre+".png", os.O_WRONLY|os.O_CREATE, 0777)

			//png.DefaultCompression
			pngEn := png.Encoder{
				CompressionLevel: png.DefaultCompression,
			}

			err = pngEn.Encode(fs, pngD)

			if err != nil {
				return err
			}

			return nil
		}
		// case "image/webp":
		// 	{
		// 		//:= bytes.
		// 		buf := &bytes.Buffer{}

		// 		buf.ReadFrom(res)
		// 		buf
		// 		webpD, err := webp.Decode(res)

		// 		if err != nil {
		// 			return err
		// 		}

		// 		fs, _ := os.OpenFile(newpath+"/"+nombre+".png", os.O_WRONLY|os.O_CREATE, 0777)

		// 		//webp.

		// 		return nil
		// 	}
	default:
		{
			fmt.Println("error")
			err := errors.New("Tipo de archivo incompatible")
			return err
		}
	}
}
