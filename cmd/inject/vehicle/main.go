package main

import (
	"fmt"
	"strings"

	"github.com/facebookgo/inject"
)

type ILog interface {
	Log(string)
}

//------------------------------------------------------------------------------

// IVehicle describes ...
type IVehicle interface {
	MinimumAge() int
	Wheels() int
}

type ILicenseCheck interface {
	CanDrive(Person, IVehicle) bool
}

type DefaultLicenseCheck struct{}

func (d *DefaultLicenseCheck) CanDrive(p Person, v IVehicle) bool {
	return p.Age >= v.MinimumAge()
}

//------------------------------------------------------------------------------

// Car is ...
type Car struct {
	Log ILog `inject:""`
	// engine string
	wheels int
}

// NewCar creates ...
func NewCar() *Car {
	return &Car{wheels: 4}
}

// Wheels return ...
func (c *Car) Wheels() int {
	c.Log.Log("Car wheels")
	return c.wheels
}

func (c *Car) MinimumAge() int {
	return 16
}

//------------------------------------------------------------------------------

// Bike is ...
type Bike struct {
	Log    ILog `inject:""`
	wheels int
}

// NewBike creates ...
func NewBike() *Bike {
	return &Bike{wheels: 2}
}

// Wheels return ...
func (b *Bike) Wheels() int {
	b.Log.Log("Bike wheels")
	return b.wheels
}

func (b *Bike) MinimumAge() int {
	return 4
}

//------------------------------------------------------------------------------

// Person is ...
type Person struct {
	Vehicle      IVehicle      `inject:""`
	Log          ILog          `inject:""`
	LicenseCheck ILicenseCheck `inject:""`
	Age          int
}

// NewPerson creates ...
func NewPerson(age int) *Person {
	return &Person{
		Age: age,
	}
}

func (person *Person) Drive() {
	if !person.LicenseCheck.CanDrive(*person, person.Vehicle) {
		person.Log.Log("Waah! I can't drive this yet!")
	} else {
		person.Log.Log(fmt.Sprintf("Driving around on %d wheels", person.Vehicle.Wheels()))
	}
}

//------------------------------------------------------------------------------

type ConsoleLog struct {
}

func (consoleLog *ConsoleLog) Log(info string) {
	fmt.Println(info)
}

//==============================================================================

var (
	person = NewPerson(12)
	car    = NewCar()
	// bike   = NewBike()
)

func init() {
	log := &ConsoleLog{}
	licenseCheck := &DefaultLicenseCheck{}
	_ = inject.Populate(person, log, car, licenseCheck)
}

func main() {
	person.Drive()
	mockTest()
}

type MockILicenseCheck struct {
	timesCalled  int
	WhatToReturn bool
}

func (m *MockILicenseCheck) CanDrive(p Person, v IVehicle) bool {
	m.timesCalled++
	return m.WhatToReturn
}

func (m *MockILicenseCheck) Reset() {
	m.timesCalled = 0
}

type MockILog struct {
	timesCalled int
	LastMessage string
}

func (m *MockILog) Log(value string) {
	m.timesCalled++
	m.LastMessage = value
}

func (m *MockILog) Reset() {
	m.timesCalled = 0
	m.LastMessage = ""
}

func mockTest() {
	log := &MockILog{}
	licenseCheck := &MockILicenseCheck{WhatToReturn: true}
	p := NewPerson(12)
	b := NewBike()
	_ = inject.Populate(log, licenseCheck, b, p)
	p.Drive()
	fmt.Printf("Expect 1, actual %d\n", licenseCheck.timesCalled)
	fmt.Printf("Expect false, actual %t\n", strings.HasPrefix(log.LastMessage, "Waah"))
	log.Reset()
	licenseCheck.Reset()
	licenseCheck.WhatToReturn = false
	p.Drive()
	fmt.Printf("Expect true, actual %t\n", strings.HasPrefix(log.LastMessage, "Waah"))
}
