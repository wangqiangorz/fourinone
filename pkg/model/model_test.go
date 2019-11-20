package model

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	ketchup := NewKetchup()
	batchtestingfunc := func(t *testing.T, got, want interface{}) {
		t.Helper()
		fmt.Println(got, want)
		// if got != want {
		// 	t.Errorf("got '%s' want '%s'", got, want)
		// }
	}
	t.Run("int", func(t *testing.T) {
		got := 1
		ketchup.SetValue(1, 1)
		want, _ := ketchup.GetInt(1)
		batchtestingfunc(t, got, want)
	})

	t.Run("float", func(t *testing.T) {
		var tmp float32
		tmp = 2.1
		got := tmp
		ketchup.SetValue(2, tmp)
		want, err := ketchup.GetFloat32(2)
		if err != nil {
			fmt.Println(err)
		}
		batchtestingfunc(t, got, want)
	})

	t.Run("string", func(t *testing.T) {
		var tmp string
		tmp = "test"
		got := tmp
		ketchup.SetValue(tmp, tmp)
		want, err := ketchup.GetString(tmp)
		if err != nil {
			fmt.Println(err)
		}
		batchtestingfunc(t, got, want)
	})

	t.Run("any", func(t *testing.T) {
		var tmp []interface{}
		tmp = make([]interface{}, 0)
		tmp = append(tmp, "1234")
		tmp = append(tmp, 1)
		got := tmp
		ketchup.SetValue("test", tmp)
		want, err := ketchup.GetAnyList("test")
		if err != nil {
			fmt.Println(err)
		}
		batchtestingfunc(t, got, want)
	})

	t.Run("stringList", func(t *testing.T) {
		var tmp []string
		tmp = make([]string, 0)
		tmp = append(tmp, "1234")
		tmp = append(tmp, "123")
		got := tmp
		ketchup.SetValue("test1", tmp)
		want, err := ketchup.GetStringList("test1")
		if err != nil {
			fmt.Println(err)
		}
		batchtestingfunc(t, got, want)
	})

	t.Run("stringList", func(t *testing.T) {
		var tmp []string
		tmp = make([]string, 0)
		tmp = append(tmp, "1234")
		tmp = append(tmp, "123")
		got := tmp
		ketchup.SetValue("test", tmp)
		want, err := ketchup.GetStringList("tes2")
		if err != nil {
			fmt.Println(err)
		}
		batchtestingfunc(t, got, want)
	})

}
