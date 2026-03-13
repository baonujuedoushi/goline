package main

import (
	"encoding/json/v2"
	"fmt"
	"time"
)

// GOEXPERIMENT=jsonv2
type user struct {
	Name         string `json:"name"`
	password     string
	Age          uint8     `json:"age"`
	RegisterTime time.Time `json:"register_time"`
}

func jsonV2() {
	u := &user{
		Name:         "test",
		Age:          20,
		RegisterTime: time.Now().AddDate(-3, 0, 0),
	}
	timeMarshaler := json.MarshalFunc(func(t time.Time) ([]byte, error) {
		return []byte(`"` + t.Format("2006-01-02") + `"`), nil
	})
	opts := json.WithMarshalers(timeMarshaler)
	uJson, _ := json.Marshal(u, opts)
	fmt.Println(string(uJson))
}
