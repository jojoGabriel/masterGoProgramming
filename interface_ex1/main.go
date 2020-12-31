package main

import "fmt"

type vehicle interface {
	License() string
	Name() string
}

type car struct {
	licenceNo string
	brand     string
}

func (c car) License() string {
	return c.licenceNo
}

func (c car) Name() string {
	return c.brand
}

func main() {

	vehicle := car{licenceNo: "IL 2345", brand: "Ford"}

	fmt.Printf("license: %s\n", vehicle.License())

	fmt.Printf("brand: %s\n", vehicle.Name())

}
