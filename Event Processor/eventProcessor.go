package main

import (
	"errors"
	"fmt"
	"sort"
)

type Event struct {
	ID     string
	UserID string
	Amount float64
}

type UserSummary struct {
	UserID      string
	TotalAmount float64
	EventCount  int
}

func SummarizeEvents(events []Event) ([]UserSummary, error) {
	// 1. Edge case check (using || instead of &&)
	if events == nil || len(events) == 0 {
		return nil, errors.New("events is nil or empty")
	}

	// Declare tracking maps OUTSIDE the loop so they persist across iterations
	duEvent := make(map[string]struct{})
	userMap := make(map[string]*UserSummary)

	for _, event := range events {
		// 2. Validation (IDs, UserIDs, and Amounts)
		if event.ID == "" || event.UserID == "" {
			return nil, errors.New("event ID or UserID is empty")
		}
		if event.Amount < 0 {
			return nil, errors.New("event amount cannot be negative")
		}

		// 3. Deduplication check
		_, exists := duEvent[event.ID]
		if exists {
			continue // Skip duplicates
		}
		duEvent[event.ID] = struct{}{}

		// 4. Aggregation per user
		summary, userExists := userMap[event.UserID]
		if !userExists {
			userMap[event.UserID] = &UserSummary{
				UserID:      event.UserID,
				TotalAmount: event.Amount,
				EventCount:  1,
			}
		} else {
			// Update existing summary values using the pointer
			summary.TotalAmount += event.Amount
			summary.EventCount++
		}
	}

	// 5. Extract from map to slice
	var result []UserSummary
	for _, summary := range userMap {
		result = append(result, *summary)
	}
	// 6. Sort by total spending descending, tie-break by UserID ascending
	sort.Slice(result, func(i, j int) bool {
		if result[i].TotalAmount == result[j].TotalAmount {
			return result[i].UserID < result[j].UserID // Tie-breaker: UserID ascending
		}
		return result[i].TotalAmount > result[j].TotalAmount // Primary: Total spending descending
	})

	return result, nil
}

func main() {
	events := []Event{
		{ID: "e1", UserID: "u1", Amount: 100},
		{ID: "e2", UserID: "u2", Amount: 200},
		{ID: "e3", UserID: "u1", Amount: 150},
		{ID: "e1", UserID: "u1", Amount: 100},
	}

	result, err := SummarizeEvents(events)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
