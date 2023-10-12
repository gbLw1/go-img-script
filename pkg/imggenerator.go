package pkg

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func GenerateImage(width, height int) {
	url := fmt.Sprintf("https://picsum.photos/%d/%d", width, height)

	fmt.Printf("Generating %d x %d image...\n", width, height)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Error:", err)
	}
	defer resp.Body.Close()

	fileName := fmt.Sprintf("%dx%d.jpg", width, height)

	saveImage(resp.Body, fileName)

	fmt.Printf("Image '%s' generated successfully!\n", fileName)
	fmt.Println("You can find it in the .\\img folder")
}

func saveImage(data io.Reader, fileName string) {
	file, err := os.Create(".\\img\\" + fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = io.Copy(file, data)
	if err != nil {
		log.Fatal(err)
	}
}
