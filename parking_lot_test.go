package main

import (
	"bytes"
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

var ParkingLotTestVar = &Slot{}

// assert fails the test if the condition is false.
func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}

func TestCreatingParkingLot(t *testing.T) {

	maxSize := 6
	resp1 := ParkingLotTestVar.Create(maxSize)
	exptString := fmt.Sprintf("Created parking lot with %d slots", maxSize)
	equals(t, exptString, resp1)
	t.Log("Created parking lot and string matches")

	resp2 := ParkingLotTestVar.Create(0)
	equals(t, "Sorry, parking lot is not created", resp2)

}

func TestParkingCarInParkingLot(t *testing.T) {

	resp1 := ParkingLotTestVar.Park("KA-01-HH-1234", "White")
	equals(t, "Sorry, parking lot is not created", resp1)

	maxSize := 3
	ParkingLotTestVar.Create(maxSize)
	for i := 0; i < maxSize; i++ {
		resp2 := ParkingLotTestVar.Park("KA-01-HH-1234", "White")
		exptString := fmt.Sprintf("Allocated slot number: %d", i+1)
		equals(t, exptString, resp2)
	}

	resp3 := ParkingLotTestVar.Park("KA-01-HH-1234", "White")
	equals(t, "Sorry, parking lot is full", resp3)

}

func TestLeaveCarFromParkingLot(t *testing.T) {

	ParkingLotTestVar.Create(0)
	resp1 := ParkingLotTestVar.Leave(2)
	equals(t, "Sorry, parking lot is not created", resp1)

	maxSize := 3
	ParkingLotTestVar.Create(maxSize)
	ParkingLotTestVar.Park("KA-01-HH-1234", "White")
	ParkingLotTestVar.Park("KA-01-HH-9999", "White")
	ParkingLotTestVar.Park("KA-01-BB-0001", "Black")

	slotNo := 2
	resp2 := ParkingLotTestVar.Leave(slotNo)
	exptString := fmt.Sprintf("Slot number %d is free", slotNo)
	equals(t, exptString, resp2)

	resp3 := ParkingLotTestVar.Leave(slotNo)
	equals(t, "Sorry, parking lot is already free", resp3)

	resp4 := ParkingLotTestVar.Leave(maxSize + 2)
	equals(t, "Sorry, parking slot is not available", resp4)

}

func TestGetRegNoByCarColor(t *testing.T) {

	colorLower := "white"
	colorUpper := "White"

	ParkingLotTestVar.Create(0)
	resp1 := ParkingLotTestVar.GetRegNoByCarColor(colorUpper)
	equals(t, "Sorry, parking lot is not created", resp1)

	maxSize := 3
	ParkingLotTestVar.Create(maxSize)
	resp2 := ParkingLotTestVar.GetRegNoByCarColor(colorUpper)
	equals(t, "Sorry, parking slot is empty", resp2)

	ParkingLotTestVar.Park("KA-01-HH-1234", "White")
	ParkingLotTestVar.Park("KA-01-HH-9999", "White")
	ParkingLotTestVar.Park("KA-01-BB-0001", "Black")

	//Test UpperCase
	resp3 := ParkingLotTestVar.GetRegNoByCarColor(colorUpper)
	equals(t, "KA-01-HH-1234, KA-01-HH-9999", resp3)

	//Test LowerCase
	resp4 := ParkingLotTestVar.GetRegNoByCarColor(colorLower)
	equals(t, "KA-01-HH-1234, KA-01-HH-9999", resp4)

	resp5 := ParkingLotTestVar.GetRegNoByCarColor("Green")
	equals(t, "Not Found", resp5)

}

func TestGetSlotNoByCarColor(t *testing.T) {

	colorLower := "white"
	colorUpper := "White"

	ParkingLotTestVar.Create(0)
	resp1 := ParkingLotTestVar.GetSlotNoByCarColor(colorUpper)
	equals(t, "Sorry, parking lot is not created", resp1)

	maxSize := 3
	ParkingLotTestVar.Create(maxSize)
	resp2 := ParkingLotTestVar.GetSlotNoByCarColor(colorUpper)
	equals(t, "Sorry, parking slot is empty", resp2)

	ParkingLotTestVar.Park("KA-01-HH-1234", "White")
	ParkingLotTestVar.Park("KA-01-HH-9999", "White")
	ParkingLotTestVar.Park("KA-01-BB-0001", "Black")

	//Test UpperCase
	resp3 := ParkingLotTestVar.GetSlotNoByCarColor(colorUpper)
	equals(t, "1, 2", resp3)

	//Test LowerCase
	resp4 := ParkingLotTestVar.GetSlotNoByCarColor(colorLower)
	equals(t, "1, 2", resp4)

	resp5 := ParkingLotTestVar.GetSlotNoByCarColor("Green")
	equals(t, "Not Found", resp5)

}

func TestGetSlotNoByCarRegNo(t *testing.T) {

	colorLower := "ka-01-hH-1234"
	colorUpper := "KA-01-HH-1234"

	ParkingLotTestVar.Create(0)
	resp1 := ParkingLotTestVar.GetSlotNoByCarRegNo(colorUpper)
	equals(t, "Sorry, parking lot is not created", resp1)

	maxSize := 3
	ParkingLotTestVar.Create(maxSize)
	resp2 := ParkingLotTestVar.GetSlotNoByCarRegNo(colorUpper)
	equals(t, "Sorry, parking slot is empty", resp2)

	ParkingLotTestVar.Park("KA-01-HH-1234", "White")
	ParkingLotTestVar.Park("KA-01-HH-9999", "White")
	ParkingLotTestVar.Park("KA-01-BB-0001", "Black")

	//Test UpperCase
	resp3 := ParkingLotTestVar.GetSlotNoByCarRegNo(colorUpper)
	equals(t, "1", resp3)

	//Test LowerCase
	resp4 := ParkingLotTestVar.GetSlotNoByCarRegNo(colorLower)
	equals(t, "1", resp4)

	resp5 := ParkingLotTestVar.GetSlotNoByCarRegNo("hij-998")
	equals(t, "Not Found", resp5)

}

func TestParkingLotStatus(t *testing.T) {

	ParkingLotTestVar.Create(0)

	buffer := &bytes.Buffer{}
	ParkingLotTestVar.Status(buffer)
	resp1 := buffer.String()
	equals(t, "Sorry, parking lot is not created", resp1)

	buffer = &bytes.Buffer{}
	maxSize := 2
	ParkingLotTestVar.Create(maxSize)
	ParkingLotTestVar.Status(buffer)
	resp2 := buffer.String()
	equals(t, "Sorry, parking slot is empty", resp2)

	buffer = &bytes.Buffer{}
	ParkingLotTestVar.Park("KA-01-HH-1234", "White")
	ParkingLotTestVar.Park("KA-01-BB-0001", "Black")
	ParkingLotTestVar.Status(buffer)
	resp3 := buffer.String()

	exptString := fmt.Sprintf("Slot No.\tRegistration No\t\t\tColour\n1\t\tKA-01-HH-1234\t\t\tWhite\n2\t\tKA-01-BB-0001\t\t\tBlack\n")
	equals(t, exptString, resp3)
}

func TestRunCommandTask(t *testing.T) {

	buffer := &bytes.Buffer{}
	ParkingLotTestVar.RunCommand(buffer, "create_parking_lot 6")
	resp1 := buffer.String()
	equals(t, "Created parking lot with 6 slots\n", resp1)

	buffer = &bytes.Buffer{}
	ParkingLotTestVar.RunCommand(buffer, "create_parking_lot")
	resp2 := buffer.String()
	equals(t, "Invalid input\n", resp2)

	buffer = &bytes.Buffer{}
	ParkingLotTestVar.RunCommand(buffer, "park KA-01-HH-1234 White")
	resp3 := buffer.String()
	equals(t, "Allocated slot number: 1\n", resp3)

	buffer = &bytes.Buffer{}
	ParkingLotTestVar.RunCommand(buffer, "park KA-01-HH-1234")
	resp4 := buffer.String()
	equals(t, "Invalid input\n", resp4)

	ParkingLotTestVar.RunCommand(buffer, "park KA-01-HH-9999 White")
	ParkingLotTestVar.RunCommand(buffer, "park KA-01-BB-0001 Black")

	buffer = &bytes.Buffer{}
	ParkingLotTestVar.RunCommand(buffer, "leave 3")
	resp5 := buffer.String()
	equals(t, "Slot number 3 is free\n", resp5)

	buffer = &bytes.Buffer{}
	ParkingLotTestVar.RunCommand(buffer, "leave")
	resp6 := buffer.String()
	equals(t, "Invalid input\n", resp6)

	buffer = &bytes.Buffer{}
	ParkingLotTestVar.RunCommand(buffer, "registration_numbers_for_cars_with_colour White")
	resp7 := buffer.String()
	equals(t, "KA-01-HH-1234, KA-01-HH-9999\n", resp7)

	buffer = &bytes.Buffer{}
	ParkingLotTestVar.RunCommand(buffer, "registration_numbers_for_cars_with_colour")
	resp8 := buffer.String()
	equals(t, "Invalid input\n", resp8)

	buffer = &bytes.Buffer{}
	ParkingLotTestVar.RunCommand(buffer, "slot_numbers_for_cars_with_colour White")
	resp9 := buffer.String()
	equals(t, "1, 2\n", resp9)

	buffer = &bytes.Buffer{}
	ParkingLotTestVar.RunCommand(buffer, "slot_numbers_for_cars_with_colour")
	resp10 := buffer.String()
	equals(t, "Invalid input\n", resp10)

	buffer = &bytes.Buffer{}
	ParkingLotTestVar.RunCommand(buffer, "slot_number_for_registration_number KA-01-HH-1234")
	resp11 := buffer.String()
	equals(t, "1\n", resp11)

	buffer = &bytes.Buffer{}
	ParkingLotTestVar.RunCommand(buffer, "slot_number_for_registration_number")
	resp12 := buffer.String()
	equals(t, "Invalid input\n", resp12)

	buffer = &bytes.Buffer{}
	ParkingLotTestVar.RunCommand(buffer, "slot")
	resp13 := buffer.String()
	equals(t, "Invalid input\n", resp13)

}
