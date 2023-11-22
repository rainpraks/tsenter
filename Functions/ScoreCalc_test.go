package functions

import (
	"container/list"
	"fmt"
	"reflect"
	"testing"
)

func TestScoreCalc(t *testing.T) {
	type args struct {
		inputKon []string
		inputTeh []string
	}
	tests := []struct {
		name  string
		args  args
		want  map[string]int
		want1 map[string]int
		want2 *list.List
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ScoreCalc(tt.args.inputKon, tt.args.inputTeh)
			fmt.Print(tt.want2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ScoreCalc() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertStars(t *testing.T) {
	type args struct {
		con  int
		tech int
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
	}{
		{
			name: "Test case 1: 500 score",
			args: struct {
				con  int
				tech int
			}{
				con:  500,
				tech: 500,
			},
			want:  "★★★★★",
			want1: "★★★★★",
		},
		{
			name: "Test case 2: 400 score",
			args: struct {
				con  int
				tech int
			}{
				con:  400,
				tech: 400,
			},
			want:  "★★★★☆",
			want1: "★★★★☆",
		},
		{
			name: "Test case 3: 300 score",
			args: struct {
				con  int
				tech int
			}{
				con:  300,
				tech: 300,
			},
			want:  "★★★☆☆",
			want1: "★★★☆☆",
		},
		{
			name: "Test case 4: 200 score",
			args: struct {
				con  int
				tech int
			}{
				con:  200,
				tech: 200,
			},
			want:  "★★☆☆☆",
			want1: "★★☆☆☆",
		},
		{
			name: "Test case 5: 100 score",
			args: struct {
				con  int
				tech int
			}{
				con:  100,
				tech: 100,
			},
			want:  "★☆☆☆☆",
			want1: "★☆☆☆☆",
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ConvertStars(tt.args.con, tt.args.tech)
			if got != tt.want {
				t.Errorf("ConvertStars() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ConvertStars() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
