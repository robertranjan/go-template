package lib

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_FlattenStruct(t *testing.T) {
	type Address struct {
		Street  string
		City    string
		Country string
	}

	type Person struct {
		Name    string
		Age     int
		Address Address
	}
	person := Person{
		Name: "John Doe",
		Age:  30,
		Address: Address{
			Street:  "123 Main Street",
			City:    "New York",
			Country: "USA",
		},
	}

	type args struct {
		obj    interface{}
		prefix string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{
			name: "success case",
			args: args{
				obj: person,
			},
			want: map[string]string{"Address.City": "New York", "Address.Country": "USA", "Address.Street": "123 Main Street", "Age": "30", "Name": "John Doe"},
		},
		{
			name: "with prefix",
			args: args{
				obj:    person,
				prefix: "::", // any string
			},
			want: map[string]string{"::Address.City": "New York", "::Address.Country": "USA", "::Address.Street": "123 Main Street", "::Age": "30", "::Name": "John Doe"},
		},
	}
	for _, tt := range tests {

		flattenedMap := FlattenStruct(person, "")
		for key, value := range flattenedMap {
			fmt.Println(key, ":", value)
		}

		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenStruct(tt.args.obj, tt.args.prefix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("flattenStruct() = %#v, want %v", got, tt.want)
			}
		})
	}
}
