package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
)

const (
	IN  = "IN"
	OUT = "OUT"
)

type Port struct {
	Number int
	Type   string
	Value  int
}
type PortController struct {
	mu    *sync.Mutex
	ports []Port
}

func New(inNumber, outNumber int) *PortController {
	return &PortController{
		mu:    new(sync.Mutex),
		ports: InitializePort(inNumber, outNumber),
	}
}
func InitializePort(inNumber, outNumber int) []Port {
	ports := make([]Port, inNumber+outNumber)
	for i := 0; i < inNumber+outNumber; i++ {
		ports[i] = Port{
			Number: i,
		}
		if i < inNumber {
			ports[i].Type = IN
			ports[i].Value = rand.Intn(2)
			continue
		}
		ports[i].Type = OUT
	}
	return ports
}
func (c *PortController) Read(portN int) (int, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	port := c.ports[portN]
	if port.Type != IN {
		return 0, fmt.Errorf("Port type must be IN not %s", port.Type)
	}
	if port.Value != 0 && port.Value != 1 {
		return 0, fmt.Errorf("Port value must be 0 or 1 not %d", port.Value)
	}
	return port.Value, nil
}
func (c *PortController) Write(portN int, value int) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	port := c.ports[portN]
	if port.Type != OUT {
		return fmt.Errorf("Port type must be OUT not %s", port.Type)
	}
	port.Value = value
	fmt.Println(port.Value)
	return nil
}
func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Printf("Must be 2 arguments")
		os.Exit(1)
	}
	var inNumber, outNumber int
	inNumber, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("It must be a number")
		os.Exit(1)
	}
	outNumber, err = strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("It must be a number")
		os.Exit(1)
	}
	ports := New(inNumber, outNumber)
	var wg sync.WaitGroup
	for n := 0; n < inNumber+outNumber; n++ {
		wg.Add(2)

		go func(n int) {
			defer wg.Done()
			value, err := ports.Read(n)
			if err != nil {
				fmt.Printf("Error reading port %d. Error: %s\n", n, err.Error())
			} else {
				fmt.Printf("Read %d from %d port\n", value, n)
			}
		}(n)
		go func(n int) {
			defer wg.Done()

			value := rand.Intn(2)
			err = ports.Write(n, value)
			if err != nil {
				fmt.Printf("Error writing port %d. Error: %s\n", n, err.Error())
			} else {
				fmt.Printf("Wrote %d to %d port\n", value, n)
			}
		}(n)
	}
	wg.Wait()
}
