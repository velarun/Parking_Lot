package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

//SIZE declaration
var SIZE int

//Car structure
type Car struct {
	RegNo string
	Color string
}

//Slot Array having Car structure as elements
type Slot []Car

//IsFull Checking Slot array is full or not
func (s *Slot) IsFull() bool {

	flag := true
	for i := range *s {
		if (*s)[i] == (Car{}) {
			flag = false
			break
		}
	}

	return flag
}

//IsEmpty Checking Slot array is empty or not
func (s *Slot) IsEmpty() bool {

	flag := true
	for i := range *s {
		if (*s)[i] != (Car{}) {
			flag = false
			break
		}
	}

	return flag
}

//ArrayToString convert array of int to string with delimiter
func ArrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

//Create func is for creating slot array with mentioned size
func (s *Slot) Create(count int) string {
	var st string
	SIZE = count
	if SIZE <= 0 {
		st = fmt.Sprintf("Sorry, parking lot is not created")
	} else {
		*s = make([]Car, SIZE)
		st = fmt.Sprintf("Created parking lot with %d slots", SIZE)
	}

	return st
}

//Park func is for Allocating slot for Car
func (s *Slot) Park(regno string, color string) string {

	var st string
	if SIZE <= 0 {
		st = fmt.Sprintf("Sorry, parking lot is not created")
	} else if s.IsFull() {
		st = fmt.Sprintf("Sorry, parking lot is full")
	} else {
		car := Car{RegNo: regno, Color: color}
		for i := range *s {
			if (*s)[i] == (Car{}) {
				(*s)[i] = car
				st = fmt.Sprintf("Allocated slot number: %d", i+1)
				break
			}
		}
	}

	return st
}

//Leave func is for unallocating slot when car leaving
func (s *Slot) Leave(slot int) string {

	var st string
	if SIZE <= 0 {
		st = fmt.Sprintf("Sorry, parking lot is not created")
	} else if (slot - 1) <= SIZE {
		if (*s)[slot-1] != (Car{}) {
			(*s)[slot-1] = Car{}
			st = fmt.Sprintf("Slot number %d is free", slot)
		} else {
			st = fmt.Sprintf("Sorry, parking lot is already free")
		}
	} else {
		st = fmt.Sprintf("Sorry, parking slot is not available")
	}

	return st
}

//Status func shows status of all the cars parked in slots
func (s *Slot) Status(out io.Writer) {

	if SIZE <= 0 {
		fmt.Fprint(out, "Sorry, parking lot is not created")
	} else if s.IsEmpty() {
		fmt.Fprint(out, "Sorry, parking slot is empty")
	} else {
		fmt.Fprint(out, "Slot No.\tRegistration No\t\t\tColour\n")
		for i := range *s {
			if (*s)[i] != (Car{}) {
				fmt.Fprintf(out, "%d\t\t%s\t\t\t%s\n", i+1, (*s)[i].RegNo, (*s)[i].Color)
			}
		}
	}
}

//GetRegNoByCarColor func is used for getting Regno by given car colour
func (s *Slot) GetRegNoByCarColor(color string) string {

	var st string
	if SIZE <= 0 {
		return "Sorry, parking lot is not created"
	} else if s.IsEmpty() {
		return "Sorry, parking slot is empty"
	} else {
		var regno []string
		for i := range *s {
			if (*s)[i] != (Car{}) && strings.ToLower((*s)[i].Color) == strings.ToLower(color) {
				regno = append(regno, (*s)[i].RegNo)
			}
		}

		if len(regno) <= 0 {
			st = "Not Found"
		} else {
			st = strings.Join(regno, ", ")
		}
	}

	return st
}

//GetSlotNoByCarColor func is used for getting Slot number by given car colour
func (s *Slot) GetSlotNoByCarColor(color string) string {

	var st string
	if SIZE <= 0 {
		return "Sorry, parking lot is not created"
	} else if s.IsEmpty() {
		return "Sorry, parking slot is empty"
	} else {
		var slotno []int
		for i := range *s {
			if (*s)[i] != (Car{}) && strings.ToLower((*s)[i].Color) == strings.ToLower(color) {
				slotno = append(slotno, i+1)
			}
		}

		if len(slotno) <= 0 {
			st = "Not Found"
		} else {
			st = ArrayToString(slotno, ", ")
		}
	}

	return st
}

//GetSlotNoByCarRegNo func is used for getting Slot number by given car registration number
func (s *Slot) GetSlotNoByCarRegNo(regno string) string {

	var st string
	if SIZE <= 0 {
		return "Sorry, parking lot is not created"
	} else if s.IsEmpty() {
		return "Sorry, parking slot is empty"
	} else {
		var flag bool
		for i := range *s {
			if (*s)[i] != (Car{}) && strings.ToLower((*s)[i].RegNo) == strings.ToLower(regno) {
				st = strconv.Itoa(i + 1)
				flag = true
				break
			}
		}

		if !flag {
			return "Not Found"
		}
	}

	return st
}

//RunCommand validate line and choose task on switch case
func (s *Slot) RunCommand(out io.Writer, line string) {

	line = strings.TrimSuffix(line, "\n")
	words := strings.Fields(line)
	switch strings.ToLower(words[0]) {
	case "create_parking_lot":
		if len(words) > 1 {
			count, err := strconv.Atoi(words[1])
			if err != nil {
				fmt.Fprintln(out, "Invalid input")
			}

			fmt.Fprintln(out, s.Create(count))
		} else {
			fmt.Fprintln(out, "Invalid input")
		}
	case "park":
		if len(words) > 2 {
			fmt.Fprintln(out, s.Park(words[1], words[2]))
		} else {
			fmt.Fprintln(out, "Invalid input")
		}
	case "leave":
		if len(words) > 1 {
			count, err := strconv.Atoi(words[1])
			if err != nil {
				fmt.Fprintln(out, "Invalid input")
			}

			fmt.Fprintln(out, s.Leave(count))
		} else {
			fmt.Fprintln(out, "Invalid input")
		}
	case "status":
		s.Status(os.Stdout)
	case "registration_numbers_for_cars_with_colour":
		if len(words) > 1 {
			fmt.Fprintln(out, s.GetRegNoByCarColor(words[1]))
		} else {
			fmt.Fprintln(out, "Invalid input")
		}
	case "slot_numbers_for_cars_with_colour":
		if len(words) > 1 {
			fmt.Fprintln(out, s.GetSlotNoByCarColor(words[1]))
		} else {
			fmt.Fprintln(out, "Invalid input")
		}
	case "slot_number_for_registration_number":
		if len(words) > 1 {
			fmt.Fprintln(out, s.GetSlotNoByCarRegNo(words[1]))
		} else {
			fmt.Fprintln(out, "Invalid input")
		}
	case "exit":
		os.Exit(0)
	default:
		fmt.Fprintln(out, "Invalid input")
	}
}

//ReadLines read file and returns array of string
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {

	s := &Slot{}

	switch len(os.Args) - 1 {
	case 0:
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print("$ ")
			cmdString, err := reader.ReadString('\n')
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}

			if len(cmdString) > 1 {
				s.RunCommand(os.Stdout, cmdString)
			}
		}
	case 1:
		lines, err := ReadLines(os.Args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		for _, line := range lines {
			if len(line) > 0 {
				s.RunCommand(os.Stdout, line)
			}
		}
	default:
		fmt.Fprintln(os.Stdout, "Usage: ./parking_lot.exe <file_input> or ./parking_lot.exe")
	}
}
