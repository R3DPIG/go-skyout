package model

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/shopspring/decimal"
)

func TestEntityJSONContracts(t *testing.T) {
	employee, err := json.Marshal(Employee{Password: "private"})
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(employee), "private") || strings.Contains(string(employee), "password") {
		t.Fatalf("employee password leaked in JSON")
	}

	dish, err := json.Marshal(Dish{Price: decimal.RequireFromString("12.50")})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(dish), `"price":"12.5"`) {
		t.Fatalf("unexpected price JSON: %s", dish)
	}
}
