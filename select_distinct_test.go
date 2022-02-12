package pgxbuilder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuery_Distinct(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		q := From("weather_reports").
			Select("location", "time", "report").
			Distinct()

		assert.Equal(t, "SELECT DISTINCT location, time, report FROM weather_reports", q.String())
	})

	t.Run("distinct on", func(t *testing.T) {
		q := From("weather_reports").
			Select("location", "time", "report").
			DistinctOn("location").
			OrderBy("location").
			OrderBy("time", "DESC")

		assert.Equal(t, "SELECT DISTINCT ON (location) location, time, report FROM weather_reports ORDER BY location, time DESC", q.String())
	})
}
