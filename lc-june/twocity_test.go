package lcjune

import "testing"

func TestTwoCity1(t *testing.T) {
	var a [][]int
	var cost int
	a = [][]int{
		{10, 20},
		{30, 200},
		{400, 50},
		{30, 20},
	}
	cost = twoCitySchedCost(a)
	if cost != 110 {
		t.Error(cost, "should be", 110)
	}
}

func TestTwoCity2(t *testing.T) {
	var a [][]int
	var cost int
	a = [][]int{
		{70, 311}, {74, 927}, {732, 711},
		{126, 583}, {857, 118}, {97, 928},
		{975, 843}, {175, 221}, {284, 929},
		{816, 602}, {689, 863}, {721, 888},
	}

	cost = twoCitySchedCost(a)
	if cost != 4723 {
		t.Error(cost, "should be", 4723)
	}
}
