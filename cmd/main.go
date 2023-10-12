package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	widthPtr := flag.Int("width", 0, "Largura da imagem")
	heightPtr := flag.Int("height", 0, "Altura da imagem")

	flag.Parse()

	if *widthPtr == 0 || *heightPtr == 0 {
		fmt.Println("The -width and -height parameters are required.")
		flag.PrintDefaults()
		os.Exit(1)
	}

	url := fmt.Sprintf("https://picsum.photos/%d/%d", *widthPtr, *heightPtr)

	fmt.Printf("Generating %d x %d image...\n", *widthPtr, *heightPtr)

	resp, err := http.Get(url)

	if err != nil {
		log.Fatal("Error:", err)
	}

	defer resp.Body.Close()

	fileName := fmt.Sprintf("%dx%d.jpg", *widthPtr, *heightPtr)
	fullPath := fmt.Sprintf(".\\img\\%s", fileName)
	file, err := os.Create(fullPath)
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(file, resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	file.Close()

	fmt.Printf("Image '%s' generated successfully!\n", fileName)
	fmt.Println("You can find it in the .\\img folder")
}
