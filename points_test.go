package main

import (
	"testing"
)

func TestCalculatePoints(t *testing.T) {
	tests := []struct {
		name     string
		receipt  Receipt
		expected int
	}{
		{
			name: "Example 1",
			receipt: Receipt{
				Retailer:     "Target",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:01",
				Items: []Item{
					{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
					{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
					{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
					{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
					{ShortDescription: "Klarbrunn 12-PK 12 FL OZ", Price: "12.00"},
				},
				Total: "35.35",
			},
			expected: 28,
		},
		{
			name: "Example 2",
			receipt: Receipt{
				Retailer:     "M&M Corner Market",
				PurchaseDate: "2022-03-20",
				PurchaseTime: "14:33",
				Items: []Item{
					{ShortDescription: "Gatorade", Price: "2.25"},
					{ShortDescription: "Gatorade", Price: "2.25"},
					{ShortDescription: "Gatorade", Price: "2.25"},
					{ShortDescription: "Gatorade", Price: "2.25"},
				},
				Total: "9.00",
			},
			expected: 109,
		},
		{
			name: "Round Dollar Amount",
			receipt: Receipt{
				Retailer:     "Walmart",
				PurchaseDate: "2023-05-15",
				PurchaseTime: "11:00",
				Items: []Item{
					{ShortDescription: "Bread", Price: "2.00"},
				},
				Total: "2.00",
			},
			expected: 88,
		},
		{
			name: "Multiple of 0.25",
			receipt: Receipt{
				Retailer:     "Costco",
				PurchaseDate: "2023-05-15",
				PurchaseTime: "11:00",
				Items: []Item{
					{ShortDescription: "Milk", Price: "1.25"},
				},
				Total: "1.25",
			},
			expected: 37,
		},
		{
			name: "Multiple of 0.25 with hour",
			receipt: Receipt{
				Retailer:     "Costco",
				PurchaseDate: "2023-05-15",
				PurchaseTime: "15:22",
				Items: []Item{
					{ShortDescription: "Milk", Price: "1.25"},
				},
				Total: "1.25",
			},
			expected: 47,
		},
		{
			name: "Morning receipt",
			receipt: Receipt{
				Retailer:     "Walgreens",
				PurchaseDate: "2022-01-02",
				PurchaseTime: "08:13",
				Items: []Item{
					{ShortDescription: "Pepsi - 12-oz", Price: "1.25"},
					{ShortDescription: "Dasani", Price: "1.40"},
				},
				Total: "2.65",
			},
			expected: 15,
		},
		{
			name: "Simple receipt",
			receipt: Receipt{
				Retailer:     "Target",
				PurchaseDate: "2022-01-02",
				PurchaseTime: "13:13",
				Items: []Item{
					{ShortDescription: "Pepsi - 12-oz", Price: "1.25"},
				},
				Total: "1.25",
			},
			expected: 31,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculatePoints(tt.receipt)
			if got != tt.expected {
				t.Errorf("calculatePoints() = %d, want %d", got, tt.expected)
			}
		})
	}
}
