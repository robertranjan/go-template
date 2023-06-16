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

func TestFixFloatPrecision(t *testing.T) {
	type args struct {
		f float64
		p int
	}
	tests := []struct {
		name    string
		args    args
		wantN   float64
		wantErr bool
	}{
		{
			name: "success case; expected result - rounding 1",
			args: args{
				f: 1.234567890,
				p: 5,
			},
			wantN: 1.23457,
		},
		{
			name: "success case; expected result - rounding 2",
			args: args{
				f: 1.234567890,
				p: 2,
			},
			wantN: 1.23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotN, err := FixFloatPrecision(tt.args.f, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("FixFloatPrecision() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("FixFloatPrecision() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}
