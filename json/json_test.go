package json

import (
	"encoding/json"

	"testing"
)

func TestMarshalStringOption(t *testing.T) {
	{
		str := `{"hoge_int": "12345"}`
		a := struct {
			HogeInt int `json:"hoge_int,string"`
		}{}

		if err := json.Unmarshal([]byte(str), &a); err != nil {
			t.Fatalf("got err: %s", err)
		}

		if a.HogeInt != 12345 {
			t.Fatalf("expected %d, got %d", 12345, a.HogeInt)
		}
	}

	{
		str := `{"int_arr": "[1,2,3,4,5]"}`
		a := struct {
			IntArr []int `json:"int_arr,string"`
		}{}

		if err := json.Unmarshal([]byte(str), &a); err == nil {
			t.Fatal("should be err")
		}
	}

	{
		str := `{"json_obj": "{'str': 'abcd'}"}`
		a := struct {
			JsonObj map[string]string `json:"json_obj,string"`
		}{}
		if err := json.Unmarshal([]byte(str), &a); err == nil {
			t.Fatal("should be err")
		}
	}
}
