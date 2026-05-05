package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	signal := Signal{Demand: 79, Capacity: 72, Latency: 22, Risk: 21, Weight: 13}
	if got := Score(signal); got != 128 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "review" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 63, Capacity: 90, Latency: 26, Risk: 20, Weight: 10}
	if got := Score(signal); got != 104 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "review" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 85, Capacity: 70, Latency: 12, Risk: 22, Weight: 9}
	if got := Score(signal); got != 146 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "review" { t.Fatalf("decision = %s", got) }
}
