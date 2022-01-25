package eks

import (
	"testing"

	"github.com/aquasecurity/defsec/provider/aws/eks"
	"github.com/aquasecurity/defsec/rules"
	"github.com/aquasecurity/defsec/state"
	"github.com/stretchr/testify/assert"
)

func TestCheckEncryptSecrets(t *testing.T) {
	t.SkipNow()
	tests := []struct {
		name     string
		input    eks.EKS
		expected bool
	}{
		{
			name:     "positive result",
			input:    eks.EKS{},
			expected: true,
		},
		{
			name:     "negative result",
			input:    eks.EKS{},
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var testState state.State
			testState.AWS.EKS = test.input
			results := CheckEncryptSecrets.Evaluate(&testState)
			var found bool
			for _, result := range results {
				if result.Status() != rules.StatusPassed && result.Rule().LongID() == CheckEncryptSecrets.Rule().LongID() {
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
