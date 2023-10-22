/*
package main

import (

	_"fmt"
	"image"
	"math"
	"math/rand"
	"os"
	"time"

	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"

)

	func loadPicture(path string) (pixel.Picture, error) {
		file, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		defer file.Close()
		img, _, err := image.Decode(file)
		if err != nil {
			return nil, err
		}
		return pixel.PictureDataFromImage(img), nil
	}

	func run() {
		cfg := pixelgl.WindowConfig{
			Title:  "Pixel Rocks!",
			Bounds: pixel.R(0, 0, 1024, 768),
		}
		win, err := pixelgl.NewWindow(cfg)
		if err != nil {
			panic(err)
		}

		spritesheet, err := loadPicture("trees.png")
		if err != nil {
			panic(err)
		}

		batch := pixel.NewBatch(&pixel.TrianglesData{}, spritesheet)

		var treesFrames []pixel.Rect
		for x := spritesheet.Bounds().Min.X; x < spritesheet.Bounds().Max.X; x += 32 {
			for y := spritesheet.Bounds().Min.Y; y < spritesheet.Bounds().Max.Y; y += 32 {
				treesFrames = append(treesFrames, pixel.R(x, y, x+32, y+32))
			}
		}

		var (
			camPos       = pixel.ZV
			camSpeed     = 500.0
			camZoom      = 1.0
			camZoomSpeed = 1.2
		)

		last := time.Now()
		for !win.Closed() {
			dt := time.Since(last).Seconds()
			last = time.Now()

			cam := pixel.IM.Scaled(camPos, camZoom).Moved(win.Bounds().Center().Sub(camPos))
			win.SetMatrix(cam)

			if win.Pressed(pixelgl.MouseButtonLeft) {
				tree := pixel.NewSprite(spritesheet, treesFrames[rand.Intn(len(treesFrames))])
				mouse := cam.Unproject(win.MousePosition())
				tree.Draw(batch, pixel.IM.Scaled(pixel.ZV, 4).Moved(mouse))
			}
			if win.Pressed(pixelgl.KeyLeft) {
				camPos.X -= camSpeed * dt
			}
			if win.Pressed(pixelgl.KeyRight) {
				camPos.X += camSpeed * dt
			}
			if win.Pressed(pixelgl.KeyDown) {
				camPos.Y -= camSpeed * dt
			}
			if win.Pressed(pixelgl.KeyUp) {
				camPos.Y += camSpeed * dt
			}
			camZoom *= math.Pow(camZoomSpeed, win.MouseScroll().Y)

			win.Clear(colornames.Forestgreen)
			batch.Draw(win)
			win.Update()
		}
	}

	func main() {
		pixelgl.Run(run)
	}
*/
package main

import (
	_"fmt"
	_ "fmt"
	"image"
	"math/rand"
	"os"

	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Parking Simulator!!!",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	picParking, err := loadPicture("./assets/parking.png")
	if err != nil {
		panic(err)
	}

	picStreet, err := loadPicture("./assets/street.png")
	if err != nil {
		panic(err)
	}

	picCars, err := loadPicture("./assets/cars.png")
	if err != nil {
		panic(err)
	}

	parking := pixel.NewSprite(picParking, picParking.Bounds())
	street := pixel.NewSprite(picStreet, picStreet.Bounds())
	
	var carsFrames []pixel.Rect
	for x := picCars.Bounds().Min.X; x < picCars.Bounds().Max.X; x += 50 {
		for y := picCars.Bounds().Min.Y; y < picCars.Bounds().Max.Y; y += 50 {
			carsFrames = append(carsFrames, pixel.R(x, y, x+50, y+50))
		}
	}

	var (
		camPos       = pixel.ZV
		camZoom      = 1.0
		cars        []*pixel.Sprite
		matrices     []pixel.Matrix
	)


	for !win.Closed() {
		parking.Draw(win, pixel.IM.Moved(pixel.Vec{X: 520, Y: 500}))
		street.Draw(win, pixel.IM.Moved(pixel.Vec{X: 0, Y: -300}))

		cam := pixel.IM.Scaled(camPos, camZoom).Moved(win.Bounds().Center().Sub(camPos))
		win.SetMatrix(cam)

		if win.JustPressed(pixelgl.MouseButtonLeft) {
			car := pixel.NewSprite(picCars, carsFrames[rand.Intn(len(carsFrames))])
			cars = append(cars, car)
			mouse := cam.Unproject(win.MousePosition())
			matrices = append(matrices, pixel.IM.Scaled(pixel.ZV, 4).Moved(mouse))
		}

		for i, car := range cars {
			car.Draw(win, matrices[i])
		}

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}