package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		println("give an argument of source file path")
		os.Exit(1)
	}
	source := os.Args[1]
	fmt.Println("splitting the file : " + source)
	err := split(source)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
}

type layer struct {
	Name      string
	Converter func(color.Color) color.Gray16
}

func split(source string) error {
	f, err := os.Open(source)
	if err != nil {
		return err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return err
	}

	// split
	layers := []layer{
		{"red", rgbaRedToGray},
		{"green", rgbaGreenToGray},
		{"blue", rgbaBlueToGray},
		{"alpha", rgbaAlphaToGray},
	}
	for _, v := range layers {
		outFilePath := outputPath(source, v.Name)
		fmt.Println("writing the file : " + outFilePath)
		outFile, err := os.Create(outFilePath)
		if err != nil {
			return err
		}
		outImage := extractLayer(img, v.Converter)
		err = png.Encode(outFile, outImage)
		outFile.Close()
		if err != nil {
			return err
		}
	}

	return nil
}

func outputPath(source string, version string) string {
	// e.g. "~/hoge/test.png" => "~/hoge/test-red.png"
	dir := filepath.Dir(source)
	ext := filepath.Ext(source)
	base := strings.TrimSuffix(filepath.Base(source), ext)
	outputFile := base + "-" + version + ext
	outputPath := filepath.Join(dir, outputFile)
	return outputPath
}

func extractLayer(img image.Image, convFunc func(color.Color) color.Gray16) *image.Gray16 {
	bounds := img.Bounds()
	outImage := image.NewGray16(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			outImage.Set(x, y, convFunc(img.At(x, y)))
		}
	}
	return outImage
}

func rgbaRedToGray(rgba color.Color) color.Gray16 {
	r, _, _, _ := rgba.RGBA() // uint32
	var gray color.Gray16
	gray.Y = uint16(r)
	return gray
}
func rgbaGreenToGray(rgba color.Color) color.Gray16 {
	_, g, _, _ := rgba.RGBA() // uint32
	var gray color.Gray16
	gray.Y = uint16(g)
	return gray
}
func rgbaBlueToGray(rgba color.Color) color.Gray16 {
	_, _, b, _ := rgba.RGBA() // uint32
	var gray color.Gray16
	gray.Y = uint16(b)
	return gray
}
func rgbaAlphaToGray(rgba color.Color) color.Gray16 {
	_, _, _, a := rgba.RGBA() // uint32
	var gray color.Gray16
	gray.Y = uint16(a)
	return gray
}
