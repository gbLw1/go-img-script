package main

import (
	"flag"
	"fmt"
	"os"

	"go-img-script/pkg"
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

	pkg.GenerateImage(*widthPtr, *heightPtr)
}
