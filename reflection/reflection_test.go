package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Rob"},
			[]string{"Rob"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Rob", "Warrenton"},
			[]string{"Rob", "Warrenton"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Rob", 47},
			[]string{"Rob"},
		},
		{
			"nested fields",
			Person{
				"Rob",
				Profile{47, "Warrenton"},
			},
			[]string{"Rob", "Warrenton"},
		},
		{
			"pointers to things",
			&Person{
				"Rob",
				Profile{47, "Warrenton"},
			},
			[]string{"Rob", "Warrenton"},
		},
		{
			"handle slices",
			[]Profile{
				{47, "Rob"},
				{52, "Jen"},
				{16, "Hannah"},
			},
			[]string{"Rob", "Jen", "Hannah"},
		},
		{
			"handling arrays",
			[2]Profile{
				{47, "Rob"},
				{52, "Jen"},
			},
			[]string{"Rob", "Jen"},
		}}
	for _, test := range cases {
		t.Run("test.Name", func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{47, "Rob"}
			aChannel <- Profile{52, "Jen"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Rob", "Jen"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{47, "Rob"}, Profile{52, "Jen"}
		}

		var got []string
		want := []string{"Rob", "Jen"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}

	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}
