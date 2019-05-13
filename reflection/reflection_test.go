package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Pointers to things",
			&Person{"Chris", Profile{33, "London"}},
			[]string{"Chris", "London"},
		},
	}

	for _, test := range cases {
		var got []string
		walk(test.Input, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, test.ExpectedCalls) {
			t.Errorf("got %v, want %v", got, test.ExpectedCalls)
		}
	}

}

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}
