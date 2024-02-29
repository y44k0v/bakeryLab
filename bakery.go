package main

import "fmt"

// total cost calculation
func calculateTotalCost(ingredients, labor, utility, discount float64) (float64, float64) {
	// totals
	totalCost := ingredients + labor + utility
	finalCostAfterDiscount := totalCost - totalCost*discount
	return totalCost, finalCostAfterDiscount
}

func calculateIngredientCost(ingredients ...[]float64) float64 {

	total := 0.0
	for _, in := range ingredients {
		total += in[0] * in[1]
	}
	return total
}

// Recursion function
func calculateBulkOrderDiscount(numItems int) float64 {
	// calculates a 5% discount for every additional 10 items, capped at 20%.
	if numItems <= 10 {
		return 0.0
	} else if numItems >= 50 {
		// Discount capped at 20%
		return 0.20
	}
	return 0.05 + calculateBulkOrderDiscount(numItems-10)
}

func main() {
	// labor cost per hour
	laborCostHR := 25.5
	numEmployees := 5.0

	// product
	//products := map[string]float64{"cake": 25.3, "bread": 2.58, "cookies": 14.6}

	// utilites
	utilities := map[string]float64{"electricity": 250.6, "gas": 693.2, "water": 26.3}

	specialOffer := func(items int, product string, baseCost float64) float64 {
		discount := 0.0
		if product == "cake" {
			if items > 100 {
				discount = 0.08
			}
		} else if product == "bread" {
			if items > 80 {
				discount = 0.1
			}
		} else if product == "cookies" {
			if items > 250 {
				discount = 0.15
			}
		} else {
			discount = 0.0
		}
		return (1 - discount) * baseCost
	}

	///// EXAMPLE ORDER /////
	// amounts of goods
	qtyBread := 100
	qtyCookies := 88
	qtyCake := 35
	// ingredients quantity and cost
	flour := []float64{1.0, 12.36}
	salt := []float64{2.2, 2.99}
	sugar := []float64{5.0, 95.6}
	yeast := []float64{2.3, 98.7}

	// cost per month   40 hours week * 4 weeks per month = 160
	laborTotal := numEmployees * 160.0 * laborCostHR
	// cost of ingredients
	ingredientsCost := calculateIngredientCost(flour, sugar, salt, yeast)
	// cost of utilities
	utilityTotal := func(utilities map[string]float64) float64 {
		total := 0.0
		for _, v := range utilities {
			total += v
		}
		return total
	}
	// discount based on number of items
	discount := calculateBulkOrderDiscount(qtyBread + qtyCake + qtyCookies)
	//println(discount)

	totalCost, finalCost := calculateTotalCost(ingredientsCost, laborTotal, utilityTotal(utilities), discount)

	// product selection special based on quantities
	product := func(qtyBread, qtyCake, qtyCookies int) string {
		if qtyBread > qtyCake && qtyBread > qtyCookies {
			return "bread"
		} else if qtyCake > qtyBread && qtyCake > qtyCookies {
			return "cake"
		} else {
			return "cookies"
		}
	}
	//  total cost if spacial discount applies
	special := specialOffer(qtyBread, product(qtyBread, qtyCake, qtyCookies), finalCost)

	fmt.Printf("Total before discount  %0.2f\n", totalCost)
	fmt.Printf("Final cost after bulk discount:   %0.2f\n", finalCost)
	fmt.Printf("Final cost after special offer  %0.2f\n", special)
}
