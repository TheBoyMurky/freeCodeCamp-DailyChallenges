package main

import (
	"fmt"
	"strconv"
	"strings"
)

func tipCalculator(mealPrice, customTip string) []string {
	cleanMealPrice := strings.ReplaceAll(mealPrice, "$", "")
	parsedMealPrice, err := strconv.ParseFloat(cleanMealPrice, 64)
	if err != nil {
		fmt.Println("Failed to convert:", err)
	}
	cleanCustomTip := strings.ReplaceAll(customTip, "%", "")
	parsedCustomTip, err := strconv.ParseFloat(cleanCustomTip, 64)
	if err != nil {
		fmt.Println("Failed to convert:", err)
	}
	percentCustomTip := parsedCustomTip / 100

	calculatedTips := []string{}
	calculatedTips = append(
		calculatedTips,
		fmt.Sprintf("$%.2f", parsedMealPrice*0.15),
		fmt.Sprintf("$%.2f", parsedMealPrice*0.20),
		fmt.Sprintf("$%.2f", parsedMealPrice*percentCustomTip),
	)
	return calculatedTips
}

func main() {
	mealPrice := "$10.00"
	customTip := "25%"
	tips := tipCalculator(mealPrice, customTip)
	t15 := tips[0]
	t20 := tips[1]
	tCustom := tips[2]
	fmt.Printf(
		"Tip values:\n 15%% : %s\n 20%% : %s\n %s : %s\n",
		t15,
		t20,
		customTip,
		tCustom,
	)
}
