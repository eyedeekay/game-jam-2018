package systems

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

import "engo.io/engo"

func (m *MapSystem) Build(xstart, ystart int) {
	stop := false
	fmt.Println("Running level generator")
	x, y := m.injectMapInfo(xstart, ystart)
	for cycles := 0; cycles <= m.cycleLimit; cycles++ {
		x, y = m.injectMapInfo(x, y)
		if ret, err := m.executeAnko(); err == nil {
			switch v := ret.(type) {
			case nil:
				fmt.Println("x is nil") // here v has type interface{}
				stop = true
				return
			case int:
				fmt.Println("x is int", v) // here v has type int
				stop = true
				return
			case bool:
				if ret.(bool) == true {
					fmt.Println("x is bool", v) // here v has type bool
				}
				stop = true
				return
			case string:
				var e error
				b := true
				x, y, b, e = m.insertMapElement(ret.(string))
				if e != nil {
					stop = true
					return
				}
				if !b {
					stop = true
					return
				}
			default:
				fmt.Println("X is unknown")
			}
		} else if stop {
			return
		} else {
			log.Fatal(err.Error())
		}
		m.tidyMapInfo()
	}
}

func (m MapSystem) injectMapInfo(x, y int) (int, int) {
	m.vmEnv.Define("cx", x)
	m.vmEnv.Define("cy", y)
	return x, y
}

func (m MapSystem) tidyMapInfo() (int, int) {
	m.vmEnv.Delete("cx")
	m.vmEnv.Delete("cy")
	return 0, 0
}

func (m MapSystem) insertMapElement(tags string) (int, int, bool, error) {
	fmt.Printf("x is string %s\n", tags) // here v has type string

	sliceSplit := strings.SplitN(strings.Replace(tags, ",", " ", -1), " ", 4)
	if len(sliceSplit) < 3 {
		return 0, 0, false, fmt.Errorf("Erroneous script line in map generator%s", tags)
	}
	x, e := strconv.Atoi(sliceSplit[0])
	if e != nil {
		return 0, 0, false, e
	}
	y, e := strconv.Atoi(sliceSplit[1])
	if e != nil {
		return 0, 0, false, e
	}
	if sliceSplit[3] != "" {
		m.LoadEntity(x, y, sliceSplit[3])
	}
	fmt.Printf("Next x: %d, y: %d, cursor moved position\n", x, y)
	b := true
	if sliceSplit[2] == "false" {
		b = false
		fmt.Printf("Halting map generation\n")
	}
	return x, y, b, nil
}

func (m *MapSystem) executeAnko() (interface{}, error) {
	if ret, err := m.vmEnv.Execute(m.script); err == nil {
		return ret, err
	} else {
		return nil, err
	}
}

func (m *MapSystem) Width() int {
	return int(engo.GameWidth())
}

func (m *MapSystem) Height() int {
	return int(engo.GameHeight())
}

func (m *MapSystem) returnFormat(x, y int, c bool, e string) string {
	var d string
	if c {
		d = "true"
	} else {
		d = "false"
	}
	if x > int(engo.GameWidth()) {
		d = "false"
	}
	if y > int(engo.GameHeight()) {
		d = "false"
	}
	return strconv.Itoa(x) + " " + strconv.Itoa(y) + " " + d + " " + e
}
