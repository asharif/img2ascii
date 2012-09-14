package main

import (
	"fmt"
	"os"
	"log"
	"image"
	_"image/jpeg"
	_"image/gif"
	_"image/png"
	//"math"
	"flag"

)

func main() {

	var pic_path  = flag.String( "p", "", "the path to the picture (jpeg, gif, png) to create art from")
	var num_col = flag.Int( "c", 50, "the number of columns for the ascii art.")
	flag.Parse()

	if *pic_path == "" {

		fmt.Printf("no picture specified.  specify path to picture with -p\n");
		return
	}


	file, err := os.Open(*pic_path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	m, _, err := image.Decode(file)

	if err != nil {

		log.Fatal(err)
	}

	bounds := m.Bounds()
	fmt.Printf( "%d , %d\n", bounds.Max.Y, bounds.Max.X)

	var y_div int = int(bounds.Max.Y / *num_col)
	var x_div int = int(bounds.Max.X / *num_col)

	for y := bounds.Min.Y; y < bounds.Max.Y/y_div; y++ {

		for x:= bounds.Min.X; x < bounds.Max.X/x_div; x++ {

			var s_y int = y * y_div
			var s_x int = x * x_div

			var avg_lum float32 = 0

			for dy := s_y; dy < s_y + y_div; dy++ {

				for dx := s_x; dx < s_x + x_div; dx++  {

					r, g, b, _ := m.At(dx, dy).RGBA()

					luminance := (float32(0.299)*(float32(r)/float32(255.0)))
					luminance += (float32(0.587)*(float32(g)/float32(255.0)))
					luminance += (float32(0.114)*(float32(b)/float32(255.0) ))
					luminance = luminance/float32(255.0)

					avg_lum += luminance
				}


			}

			avg_lum /= float32(y_div*x_div)

			if avg_lum > 0.5 {

				fmt.Printf("-")
			}

			if avg_lum <= 0.5 && avg_lum > 0.25 {

				fmt.Printf("G")

			}

			if avg_lum <= 0.25 {

				fmt.Printf("@")


			}



		}

		fmt.Printf( "\n")

	}

}

