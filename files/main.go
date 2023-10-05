package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"os/exec"
)

func oldmain() {
	names := `package main

import (
	"log"
	"os"
)

func main() {
	names := "Tilak"
	err := os.WriteFile("names.txt", []byte(names), 0600)
	if err != nil {
		log.Fatalln(err)
	}
}
	`
	// err := os.Chmod("notmain.go", 0600)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	err := os.WriteFile("notmain.go", []byte(names), 0600)
	if err != nil {
		log.Fatalln(err)
	}
	cmd := exec.Command("go", "run", "notmain.go")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		log.Fatalln(err)
	}
}

func manyfilesmain() {
	names := []string{"Tilak", "Shreesh", "Sarthak"}
	err := os.MkdirAll("test", 0755)
	if err != nil {
		log.Fatalln(err)
	}
	for i := 0; i < 1000; i++ {
		err := os.WriteFile(fmt.Sprintf("test/%s-%d.txt", names[i%3], i), []byte(names[i%3]), 0600)
		if err != nil {
			log.Fatalln(err)
		}
	}

}

func readfileStatsmain() {
	fs, err := os.Stat("names_expected.txt")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%q", fs.Name())
}

var img = image.NewRGBA(image.Rect(0, 0, 100, 100))
var col color.Color

func main() {
	col = color.RGBA{0, 0, 255, 255} // Red
	HLine(10, 40, 80)
	col = color.RGBA{0, 255, 0, 255} // Green
	Rect(10, 10, 80, 50)

	f, err := os.Create("draw.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, img)
}
func HLine(x1, y, x2 int) {
	for ; x1 <= x2; x1++ {
		img.Set(x1, y, col)
	}
}

// VLine draws a veritcal line
func VLine(x, y1, y2 int) {
	for ; y1 <= y2; y1++ {
		img.Set(x, y1, col)
	}
}

// Rect draws a rectangle utilizing HLine() and VLine()
func Rect(x1, y1, x2, y2 int) {
	HLine(x1, y1, x2)
	HLine(x1, y2, x2)
	VLine(x1, y1, y2)
	VLine(x2, y1, y2)
}
