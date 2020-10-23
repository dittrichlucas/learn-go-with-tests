package recursion

import (
	"reflect"
	"testing"
)

func TestRecursion(t *testing.T) {

	tests := []struct {
		Name  string
		Input interface{}
		Want  []string
	}{
		{
			"Struct with a string field",
			struct {
				Name string
			}{"Lucas"},
			[]string{"Lucas"},
		},
		{
			"Struct with two string field",
			struct {
				Name string
				City string
			}{"Lucas", "London"},
			[]string{"Lucas", "London"},
		},
		{
			"Struct with two string field",
			struct {
				Name string
				Age  int
			}{"Lucas", 26},
			[]string{"Lucas"},
		},
		{
			"Nested fields",
			Person{
				"Lucas",
				Profile{26, "London"},
			},
			[]string{"Lucas", "London"},
		},
		{
			"Pointers to things",
			&Person{
				"Lucas",
				Profile{26, "London"},
			},
			[]string{"Lucas", "London"},
		},
		{
			"Slices",
			[]Profile{
				{26, "London"},
				{25, "Havai"},
			},
			[]string{"London", "Havai"},
		},
		{
			"Arrays",
			[2]Profile{
				{33, "London"},
				{32, "Reykjavík"},
			},
			[]string{"London", "Reykjavík"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			var got []string
			recursion(tt.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, tt.Want) {
				t.Errorf("got %v, want %v", got, tt.Want)
			}
		})
	}

	t.Run("Maps", func(t *testing.T) {
		mapA := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var result []string
		recursion(mapA, func(entry string) {
			result = append(result, entry)
		})

		checkContains(t, result, "Bar")
		checkContains(t, result, "Boz")
	})
}

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func checkContains(t *testing.T, myMap []string, key string) {
	contain := false
	for _, x := range myMap {
		if x == key {
			contain = true
		}
	}
	if !contain {
		t.Errorf("")
	}
}
