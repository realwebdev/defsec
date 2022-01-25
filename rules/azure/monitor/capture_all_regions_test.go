package monitor

import (
	"testing"

	"github.com/aquasecurity/defsec/provider/azure/monitor"
	"github.com/aquasecurity/defsec/rules"
	"github.com/aquasecurity/defsec/state"
	"github.com/stretchr/testify/assert"
)

func TestCheckCaptureAllRegions(t *testing.T) {
	t.SkipNow()
	tests := []struct {
		name     string
		input    monitor.Monitor
		expected bool
	}{
		{
			name:     "positive result",
			input:    monitor.Monitor{},
			expected: true,
		},
		{
			name:     "negative result",
			input:    monitor.Monitor{},
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var testState state.State
			testState.Azure.Monitor = test.input
			results := CheckCaptureAllRegions.Evaluate(&testState)
			var found bool
			for _, result := range results {
				if result.Status() != rules.StatusPassed && result.Rule().LongID() == CheckCaptureAllRegions.Rule().LongID() {
					found = true
				}
			}
			if test.expected {
				assert.True(t, found, "Rule should have been found")
			} else {
				assert.False(t, found, "Rule should not have been found")
			}
		})
	}
}
