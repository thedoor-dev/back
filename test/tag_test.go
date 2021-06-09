package test

import (
	"encoding/json"
	"testing"

	"github.com/thedoor-dev/back/models"
)

func TestTags(t *testing.T) {

	var ts models.TagArr
	ts = append(ts, models.Tag{
		ID:   0,
		PID:  0,
		Name: "avc",
	})
	ts = append(ts, models.Tag{
		ID:   0,
		PID:  0,
		Name: "abv",
	})
	t.Log(ts)
	a, _ := json.Marshal(ts)
	t.Logf("%s\n", a)
}
