package json

import (
	"testing"
)

func Test_GetTimestamp(t *testing.T) {
	resp := GetTimestamp()
	t.Log(resp.Data.T)
}
func Test_GetTimestampT(t *testing.T) {
	resp := GetTimestampT()
	t.Log(resp)
}
func Test_GetTimestampMap(t *testing.T) {
	resp := GetTimestampMap()
	t.Log(resp["data"].(map[string]interface{})["t"].(string))
}
