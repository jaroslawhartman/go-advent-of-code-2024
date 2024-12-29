package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"log"
	"os"
	"regexp"
	"strconv"
	"sync"

	"github.com/fogleman/gg"
)

var totalCost int

type robot struct {
	px int
	py int
	vx int
	vy int
}

var robots []robot

var maxX, maxY int

var images []*image.Paletted
var delays []int
var disposals []byte

var palette color.Palette = color.Palette{
	image.Transparent,
	image.Black,
	image.White,
	color.RGBA{0xEE, 0xEE, 0xEE, 255},
	color.RGBA{0xCC, 0xCC, 0xCC, 255},
	color.RGBA{0x99, 0x99, 0x99, 255},
	color.RGBA{0x66, 0x66, 0x66, 255},
	color.RGBA{0x33, 0x33, 0x33, 255},
}

func (r robot) String() string {
	return fmt.Sprintf("Pos [%d,%d] V: [%d,%d]", r.px, r.py, r.vx, r.vy)
}

func Atoi(n string) int {
	if i, err := strconv.Atoi(n); err != nil {
		return 0
	} else {
		return i
	}
}

func findRobotCount(x, y int) int {
	count := 0
	for _, r := range robots {
		if r.px == x && r.py == y {
			count += 1
		}
	}
	return count
}

func findRobot(x, y int) int {
	for i, r := range robots {
		if r.px == x && r.py == y {
			return i
		}
	}
	return -1
}

func moveRobots() {
	var wg sync.WaitGroup

	for i := range robots {
		wg.Add(1)

		go func() {
			defer wg.Done()

			robots[i].px = (robots[i].px + robots[i].vx)
			robots[i].py = (robots[i].py + robots[i].vy)

			if robots[i].px >= maxX {
				robots[i].px = robots[i].px % (maxX)
			}

			if robots[i].py >= maxY {
				robots[i].py = robots[i].py % (maxY)
			}

			if robots[i].px < 0 {
				robots[i].px = robots[i].px + maxX
			}

			if robots[i].py < 0 {
				robots[i].py = robots[i].py + maxY
			}
		}()
	}

	wg.Wait()
}

func countRobots(qx, qy int) int {
	var startX, startY, stopX, stopY int

	if qx == 0 {
		startX = 0
		stopX = maxX / 2
	} else if qx == 1 {
		startX = maxX/2 + 1
		stopX = maxX
	} else {
		return -1
	}

	if qy == 0 {
		startY = 0
		stopY = maxY / 2
	} else if qy == 1 {
		startY = maxY/2 + 1
		stopY = maxY
	} else {
		return -1
	}

	count := 0

	for y := startY; y < stopY; y++ {
		for x := startX; x < stopX; x++ {
			count += findRobotCount(x, y)
		}
	}
	return count
}

func drawFrame(i int) {
	dc := gg.NewContext(maxX, maxY)

	dc.Fill()

	dc.SetRGB255(255, 0, 0)
	dc.DrawString(fmt.Sprintf("Frame %d", i), 0.0, 9.0)

	dc.SetRGBA(0, 0, 0, 1)

	for y := range maxY {
		for x := range maxX {
			count := findRobotCount(x, y)

			if count > 0 {
				dc.SetPixel(x, y)
			}
		}
	}

	img := dc.Image()
	bounds := img.Bounds()

	dst := image.NewPaletted(bounds, palette)
	draw.Draw(dst, bounds, img, bounds.Min, draw.Src)
	images = append(images, dst)
	delays = append(delays, 10)
	disposals = append(disposals, gif.DisposalBackground)
}

func testFrame() bool {
	for y := range maxY {
		lineCount := 0

		for x := range maxX {
			count := findRobot(x, y)

			if count != -1 {
				lineCount += 1

				if lineCount > 9 {
					return true
				}
			} else {
				lineCount = 0
			}
		}

	}
	return false
}

func readInput(file string) {
	f, err := os.Open(file)

	if err != nil {
		log.Fatal("Can't open file: " + err.Error())
	}

	s := bufio.NewScanner(f)

	s.Split(bufio.ScanLines)

	reRobot := regexp.MustCompile(`p=(\d+),(\d+) v=(-?\d+),(-?\d+)`)

	for s.Scan() {
		line := s.Text()
		res := reRobot.FindStringSubmatch(line)
		r := robot{
			px: Atoi(res[1]),
			py: Atoi(res[2]),
			vx: Atoi(res[3]),
			vy: Atoi(res[4]),
		}

		robots = append(robots, r)
		fmt.Println(r)
	}
}

func writeGif() {
	f, err := os.OpenFile("rgb.gif", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	gif.EncodeAll(f, &gif.GIF{
		Image:    images,
		Delay:    delays,
		Disposal: disposals,
	})
}

func run(file string, x, y int) int {
	maxX = x
	maxY = y
	readInput(file)

	for i := range 100000 {
		moveRobots()
		// drawMap(i + 1)
		if testFrame() {
			drawFrame(i + 1)
			fmt.Println("Found !!", i+1)
			break
		}

		if i%10000 == 0 {
			fmt.Println(i)
		}
	}

	total := countRobots(0, 0) * countRobots(1, 0) * countRobots(0, 1) * countRobots(1, 1)

	writeGif()

	return total
}

func main() {
	total := run("input.txt", 101, 103)

	fmt.Println("Total: ", total)
}
