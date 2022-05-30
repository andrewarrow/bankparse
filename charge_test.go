package main

import "testing"

func TestChargeParse(t *testing.T) {

	//IN N OUT BURGER 055 2022-05-21
	//IN N OUT BURGER 055 LOS ANGELES CAUSA 2022-05-23

	a := "ROOT INSURANCE 2022-05-22"

	result := ChargeParse(a)
	if result != "ROOT INSURANCE" {
		t.Fatalf("%s", result)
	}

	a = "ROOT INSURANCE           614-915-0703 OHUSA 2022-05-23"

	result = ChargeParse(a)
	if result != "ROOT INSURANCE" {
		t.Fatalf("%s", result)
	}
}
