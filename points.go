package main

import (
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func calculatePoints(receipt Receipt) int {
	points := 0

	// 1 point for every alphanumeric character in the retailer name.
	points += len(regexp.MustCompile(`[a-zA-Z0-9]`).FindAllString(receipt.Retailer, -1))

	// 50 points if the total is a round dollar amount with no cents.
	if total, err := strconv.ParseFloat(receipt.Total, 64); err == nil {
		if total == math.Floor(total) {
			points += 50
		}
	}

	// 25 points if the total is a multiple of 0.25.
	if total, err := strconv.ParseFloat(receipt.Total, 64); err == nil {
		if math.Mod(total, 0.25) == 0 {
			points += 25
		}
	}

	// 5 points for every two items on the receipt.
	points += (len(receipt.Items) / 2) * 5

	// Points for item description length
	for _, item := range receipt.Items {
		description := strings.TrimSpace(item.ShortDescription)
		if len(description)%3 == 0 {
			if price, err := strconv.ParseFloat(item.Price, 64); err == nil {
				points += int(math.Ceil(price * 0.2))
			}
		}
	}

	// 6 points if the day in the purchase date is odd.
	if date, err := time.Parse("2006-01-02", receipt.PurchaseDate); err == nil {
		if date.Day()%2 != 0 {
			points += 6
		}
	}

	// 10 points if the time of purchase is after 2:00pm and before 4:00pm.
	if t, err := time.Parse("15:04", receipt.PurchaseTime); err == nil {
		if t.Hour() == 14 {
			points += 10
		}
	}

	return points
}
