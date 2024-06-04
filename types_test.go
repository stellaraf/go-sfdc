package sfdc_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stellaraf/go-sfdc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
		require.NoError(t, err)
		assert.Equal(t, rfc3339, s.Timestamp.Format(time.RFC3339))
	})
}

func TestBulkJob_CSV(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		type Data struct {
			String string `json:"string"`
			Int    int    `json:"int"`
			Bool   bool   `json:"bool"`
		}
		data := Data{String: "value", Int: 100, Bool: true}
		job := &sfdc.BulkJob{}
		result, err := job.CSV(data)
		require.NoError(t, err)
		assert.Contains(t, result, "value")
		assert.Contains(t, result, "100")
		assert.Contains(t, result, "true")
	})
}
