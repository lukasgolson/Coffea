package application

import (
	"github.com/supabase/postgrest-go"
)

type Values struct {
	DiceValue int `json:"dice_value"`
}

func DownloadDiceValues(url, key string, limit, offset int) ([]int, error) {

	client := postgrest.NewClient(url, "public", map[string]string{"apikey": key})

	var values []Values

	_, err := client.From("rolls").Select("dice_value", "exact", false).Order("unix_milliseconds", &postgrest.OrderOpts{
		Ascending: false,
	}).Limit(limit, "").ExecuteTo(&values)
	if err != nil {
		return make([]int, 0), err
	}
	// Extract the dice values and store them in a separate slice.
	diceValues := make([]int, len(values))
	for i, j := 0, len(values)-1; i < len(values); i, j = i+1, j-1 {
		diceValues[i] = values[j].DiceValue
	}

	return diceValues, nil

}
