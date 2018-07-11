package itinerary

import (
	"time"
)

var (
	trips []*Trip
	days  []*Day
)

type Trip struct {
	Location  string
	StartDate time.Time
	EndDate   time.Time
	Days      []*Day
	Budget    float64
}

type Day struct {
	Date        time.Time
	Location    string
	Activities  string
	Restaurants string
	Hotel       string
}

func AddTrip(trip Trip) {

	trip.Days = countDays(trip.StartDate, trip.EndDate)
	trips = append(trips, &trip)
}

func ListTrips() []*Trip {
	return trips
}

func countDays(startDate time.Time, endDate time.Time) []*Day {

	var days []*Day
	currentDate := startDate

	for !currentDate.After(endDate) {

		newDay := &Day{
			Date: currentDate,
		}

		days = append(days, newDay)

		currentDate = currentDate.Add(24 * time.Hour)
	}
	return days
}
