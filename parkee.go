package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type ParkingLot struct {
	slots []string
}

var needsParking = map[string]bool{
	"park":   true,
	"leave":  true,
	"status": true,
}

func Command(parking *ParkingLot, cmd []string) (*ParkingLot, error) {
	cmd[0] = strings.ToLower(cmd[0])

	if needsParking[cmd[0]] && parking == nil {
		return parking, errors.New("Parking lot has not been created yet")
	}

	switch cmd[0] {
	case "create_parking_lot":
		capacity, _ := strconv.Atoi(cmd[1])
		parking = &ParkingLot{
			slots: make([]string, capacity),
		}
	case "park":
		slot := parking.Park(cmd[1])
		if slot == -1 {
			fmt.Println("Sorry, parking lot is full")
		} else {
			fmt.Printf("Allocated slot number: %d\n", slot)
		}
	case "leave":
		slot, fee := parking.Leave(cmd[1], cmd[2])
		if slot == -1 {
			fmt.Printf("Registration number %s not found\n", cmd[1])
		} else {
			fmt.Printf("Registration number %s with Slot Number %d is free with Charge $%d\n", cmd[1], slot, fee)
		}
	case "status":
		parking.Status()
	default:
		return parking, errors.New("Command Not Found")
	}

	return parking, nil
}

func (p *ParkingLot) Park(plate string) int {
	for i := 0; i < len(p.slots); i++ {
		if p.slots[i] == "" {
			p.slots[i] = plate
			return i + 1
		}
	}
	return -1
}

func (p *ParkingLot) Leave(plate, hours string) (int, int) {
	for i := range p.slots {
		if p.slots[i] == plate {
			p.slots[i] = ""
			fee := p.CalculateFee(hours)
			return i + 1, fee
		}
	}
	return -1, 0
}

func (p *ParkingLot) Status() {
	fmt.Println("Slot No. Registration No.")
	for i, car := range p.slots {
		if car != "" {
			fmt.Println(i+1, car)
		}
	}
}

func (p *ParkingLot) CalculateFee(hours string) int {
	h, _ := strconv.Atoi(hours)
	if h <= 2 {
		return 10
	}
	return 10 + (h-2)*10
}
