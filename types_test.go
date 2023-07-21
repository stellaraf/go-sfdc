package sfdc_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stellaraf/go-sfdc"
	"github.com/stretchr/testify/assert"
)

func Test_Time(t *testing.T) {
	type Struct struct {
		Timestamp *sfdc.Time `json:"timestamp"`
	}
	rfc3339 := "2023-07-21T20:46:33Z"
	data := []byte(`{"timestamp":"2023-07-21T20:46:33.000+0000"}`)
	t.Run("unmarshal", func(t *testing.T) {
		var s *Struct
		err := json.Unmarshal(data, &s)
		assert.NoError(t, err)
		assert.Equal(t, rfc3339, s.Timestamp.Format(time.RFC3339))
	})
}
