package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/lindsshldz/itinerary-cli/itinerary"
	"github.com/manifoldco/promptui"
)

const (
	addTripCmd    = "Add Trip"
	addDetailsCmd = "Add fun details to your trip"
	printTripCmd  = "Print out your trip itinerary"
)

func main() {
	for {
		fmt.Println()

		prompt := promptui.Select{
			Label: "Select Action",
			Items: []string{
				addTripCmd,
				addDetailsCmd,
				printTripCmd,
			},
		}

		_, result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		switch result {
		case addTripCmd:
			err := addTripPrompt()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}

		case addDetailsCmd:
			err := addDetails()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}

		case printTripCmd:
			err := printItinerary
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}
		}

		time.Sleep(500 * time.Millisecond)
	}

}

func addTripPrompt() error {
	locationPrompt := promptui.Prompt{
		Label: "Location",
	}
	location, err := locationPrompt.Run()
	if err != nil {
		return err
	}

	tripStartDate, err := datePromptHelper("Trip Start Date")
	if err != nil {
		return err
	}

	tripEndDate, err := datePromptHelper("Trip End Date")
	if err != nil {
		return err
	}

	tripBudget, err := numberPromptHelper("Overall Budget")
	if err != nil {
		return err
	}
	newTrip := itinerary.Trip{
		Location:  location,
		StartDate: tripStartDate,
		EndDate:   tripEndDate,
		Budget:    tripBudget,
	}

	itinerary.AddTrip(newTrip)

	fmt.Println("Added " + location + " trip!")

	return nil
}

func addDetails() error {
	availableTrips := itinerary.ListTrips()

	if len(availableTrips) == 0 {
		fmt.Println("Create a trip first")
		return nil
	}
	var options []string
	for _, trip := range availableTrips {
		options = append(options, trip.Location)
	}

	selectTripPrompt := promptui.Select{
		Label: "Select Trip",
		Items: options,
	}

	chosenIndex, _, err := selectTripPrompt.Run()
	if err != nil {
		return err
	}

	chosenTrip := availableTrips[chosenIndex]

	fmt.Println(chosenTrip)

	availableDays := chosenTrip.Days

	if len(availableDays) == 0 {
		fmt.Println("Must add start and end dates")
		return nil
	}
	var dayOptions []time.Time
	for _, day := range availableDays {
		dayOptions = append(dayOptions, day.Date)
	}

	selectDayPrompt := promptui.Select{
		Label: "Select Day",
		Items: dayOptions,
	}

	chosenDayIndex, _, err := selectDayPrompt.Run()
	if err != nil {
		return err
	}

	chosenDay := availableDays[chosenDayIndex]

	fmt.Println(chosenDay)

	//	selectDetailsPrompt := promptui.Select{
	// 	Label: "Select details to add",
	// 	Items: []string{
	// 		"Location",
	// 		"Activities",
	// 		"Restaurants",
	// 		"Hotel",
	// 	},
	// }

	dayLocation, err := detailPromptHelper("Location on that day")
	if err != nil {
		return err
	}

	dayActivities, err := detailPromptHelper("Add activities for the day")
	if err != nil {
		return err
	}

	dayRestaurant, err := detailPromptHelper("Where you will eat that day")
	if err != nil {
		return err
	}

	dayHotel, err := detailPromptHelper("Hotel name for that night")
	if err != nil {
		return err
	}

	chosenDay.Location = dayLocation
	chosenDay.Activities = dayActivities
	chosenDay.Restaurants = dayRestaurant
	chosenDay.Hotel = dayHotel

	fmt.Println("Added details to " + chosenDay.Location)

	return nil
}

func printItinerary() error {

	availableTrips := itinerary.ListTrips()

	if len(availableTrips) == 0 {
		fmt.Println("Create a trip first")
		return nil
	}
	var options []string
	for _, trip := range availableTrips {
		options = append(options, trip.Location)
	}

	selectTripPrompt := promptui.Select{
		Label: "Select Trip",
		Items: options,
	}

	chosenIndex, _, err := selectTripPrompt.Run()
	if err != nil {
		return err
	}

	chosenTrip := availableTrips[chosenIndex]

	fmt.Println(chosenTrip.Location)
	fmt.Println(chosenTrip.Budget)
	fmt.Println(chosenTrip.Days)

	return nil

}

func datePromptHelper(label string) (time.Time, error) {

	validate := func(input string) error {
		_, err := time.Parse("01-02-2006", input)
		if err != nil {
			return errors.New("Needs to be in full date format: mm-dd-yyyy")
		}
		return nil
	}
	datePrompt := promptui.Prompt{
		Label:    label,
		Validate: validate,
	}
	dateStr, err := datePrompt.Run()
	if err != nil {
		return time.Time{}, err
	}
	date, err := time.Parse("01-02-2006", dateStr)
	if err != nil {
		return time.Time{}, err
	}

	return date, nil
}

func detailPromptHelper(label string) (string, error) {

	validate := func(input string) error {
		if len(input) > 20 {
			return errors.New("can only have 20 characters")
		}
		return nil
	}
	detailPrompt := promptui.Prompt{
		Label:    label,
		Validate: validate,
	}
	detail, err := detailPrompt.Run()
	if err != nil {
		return detail, err
	}
	return detail, nil
}

func numberPromptHelper(label string) (float64, error) {

	validate := func(input string) error {
		_, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return errors.New("needs to be a number")
		}
		return nil
	}
	numberPrompt := promptui.Prompt{
		Label:    label,
		Validate: validate,
	}
	numberStr, err := numberPrompt.Run()
	if err != nil {
		return 0, err
	}
	number, err := strconv.ParseFloat(numberStr, 64)
	if err != nil {
		return 0, err
	}

	return number, nil
}
