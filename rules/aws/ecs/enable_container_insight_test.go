package ecs

import (
	"testing"

	"github.com/aquasecurity/defsec/provider/aws/ecs"
	"github.com/aquasecurity/defsec/rules"
	"github.com/aquasecurity/defsec/state"
	"github.com/stretchr/testify/assert"
)

func TestCheckEnableContainerInsight(t *testing.T) {
	t.SkipNow()
	tests := []struct {
		name     string
		input    ecs.ECS
		expected bool
	}{
		{
			name:     "positive result",
			input:    ecs.ECS{},
			expected: true,
		},
		{
			name:     "negative result",
			input:    ecs.ECS{},
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var testState state.State
			testState.AWS.ECS = test.input
			results := CheckEnableContainerInsight.Evaluate(&testState)
			var found bool
			for _, result := range results {
				if result.Status() != rules.StatusPassed && result.Rule().LongID() == CheckEnableContainerInsight.Rule().LongID() {
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