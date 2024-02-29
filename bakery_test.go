package main

import "testing"

func TestBakery(t *testing.T) {
	// Test for bulk discount
	t.Run("Discount 7 items", func(t *testing.T) {
		got := calculateBulkOrderDiscount(7)
		want := 0.00
		assertCorrectResult1(t, got, want)
	})
	t.Run("Discount 15 items", func(t *testing.T) {
		got := calculateBulkOrderDiscount(15)
		want := 0.05
		assertCorrectResult1(t, got, want)
	})
	t.Run("Discount 25 items", func(t *testing.T) {
		got := calculateBulkOrderDiscount(25)
		want := 0.10
		assertCorrectResult1(t, got, want)
	})
	t.Run("Discount 35 items", func(t *testing.T) {
		got := calculateBulkOrderDiscount(35)
		want := 0.15
		assertCorrectResult1(t, got, want)
	})
	t.Run("Discount 45 items", func(t *testing.T) {
		got := calculateBulkOrderDiscount(45)
		want := 0.20
		assertCorrectResult1(t, got, want)
	})
	t.Run("Discount 55 items", func(t *testing.T) {
		got := calculateBulkOrderDiscount(55)
		want := 0.20
		assertCorrectResult1(t, got, want)
	})

	// Test for totals with different % of discount
	t.Run("Totals with 0% Discount", func(t *testing.T) {
		got1, got2 := calculateTotalCost(250.2, 500, 249.8, 0.0)
		want1, want2 := 1000.0, 1000.0
		assertCorrectResult2(t, got1, got2, want1, want2)

	})
	t.Run("Totals with 15% Discount", func(t *testing.T) {
		got1, got2 := calculateTotalCost(250.2, 500, 249.8, 0.15)
		want1, want2 := 1000.0, 850.0
		assertCorrectResult2(t, got1, got2, want1, want2)

	})
	t.Run("Totals with 100% Discount", func(t *testing.T) {
		got1, got2 := calculateTotalCost(250.2, 500, 249.8, 1.0)
		want1, want2 := 1000.0, 0.0
		assertCorrectResult2(t, got1, got2, want1, want2)

	})

	// test for product selection basewd on quantities for special offer
	// cannot test anonymous functions on main
	product := func(qtyBread, qtyCake, qtyCookies int) string {
		if qtyBread > qtyCake && qtyBread > qtyCookies {
			return "bread"
		} else if qtyCake > qtyBread && qtyCake > qtyCookies {
			return "cake"
		} else {
			return "cookies"
		}
	}
	t.Run("Bread selecion", func(t *testing.T) {
		got := product(5, 3, 2)
		want := "bread"
		assertCorrectProduct(t, got, want)
	})
	t.Run("Cake selecion", func(t *testing.T) {
		got := product(2, 5, 3)
		want := "cake"
		assertCorrectProduct(t, got, want)
	})
	t.Run("Cookie selecion", func(t *testing.T) {
		got := product(2, 3, 5)
		want := "cookies"
		assertCorrectProduct(t, got, want)
	})
	// Test for special offer discount

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
	t.Run("Discount 60 breads", func(t *testing.T) {
		got := specialOffer(60, "bread", 1000.0)
		want := 1000.0
		assertCorrectResult1(t, got, want)
	})
	t.Run("Discount 100 breads", func(t *testing.T) {
		got := specialOffer(100, "bread", 1000.0)
		want := 900.0
		assertCorrectResult1(t, got, want)
	})
	t.Run("Discount 150 cakes pieces", func(t *testing.T) {
		got := specialOffer(150, "cake", 1000.0)
		want := 920.0
		assertCorrectResult1(t, got, want)
	})
	t.Run("Discount 300 Cookiess", func(t *testing.T) {
		got := specialOffer(300, "cookies", 1000.0)
		want := 850.0
		assertCorrectResult1(t, got, want)
	})
	t.Run("Discount 100 cookies", func(t *testing.T) {
		got := specialOffer(100, "cookies", 1000.0)
		want := 1000.0
		assertCorrectResult1(t, got, want)
	})

}

// helper functions
func assertCorrectResult2(t testing.TB, got1, got2, want1, want2 float64) {
	t.Helper()
	if got1 != want1 {
		t.Errorf("got %0.2f want %0.2f", got1, want1)
	}
	if got2 != want2 {
		t.Errorf("got %0.2f want %0.2f", got2, want2)
	}
}

func assertCorrectResult1(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %0.2f want %0.2f", got, want)
	}
}

func assertCorrectProduct(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
