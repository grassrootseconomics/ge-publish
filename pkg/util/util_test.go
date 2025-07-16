package util

import (
	"testing"
)

func TestCalculateDecayLevel(t *testing.T) {
	// Test parameters: 2% demurrage rate and 1 week (10080 minutes)
	demurrageRate := int64(2)
	redistributionPeriod := int64(10080)

	decayLevel, err := CalculateDecayLevel(demurrageRate, redistributionPeriod)
	if err != nil {
		t.Fatalf("CalculateDecayLevel failed: %v", err)
	}

	t.Logf("Decay level for 2%% demurrage rate with 1 week redistribution period: %s", decayLevel.String())

}
