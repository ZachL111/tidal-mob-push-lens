package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	tests := []struct {
		name         string
		signal       Signal
		wantScore    int
		wantDecision string
	}{
		{name: "case_1", signal: Signal{Demand: 79, Capacity: 72, Latency: 22, Risk: 21, Weight: 13}, wantScore: 128, wantDecision: "review"},
		{name: "case_2", signal: Signal{Demand: 63, Capacity: 90, Latency: 26, Risk: 20, Weight: 10}, wantScore: 104, wantDecision: "review"},
		{name: "case_3", signal: Signal{Demand: 85, Capacity: 70, Latency: 12, Risk: 22, Weight: 9}, wantScore: 146, wantDecision: "review"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Score(tc.signal); got != tc.wantScore {
				t.Fatalf("score = %d, want %d", got, tc.wantScore)
			}
			if got := Classify(tc.signal); got != tc.wantDecision {
				t.Fatalf("decision = %s, want %s", got, tc.wantDecision)
			}
		})
	}
}
