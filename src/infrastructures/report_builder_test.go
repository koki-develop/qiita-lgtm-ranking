package infrastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewReportBuilder(t *testing.T) {
	t.Run("return ReportBuilder", func(t *testing.T) {
		b := NewReportBuilder()
		assert.NotNil(t, b)
		assert.IsType(t, &ReportBuilder{}, b)
	})
}
