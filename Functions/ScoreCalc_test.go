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
	}{}
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

func TestConvertScore(t *testing.T) {
	type args struct {
		tech int
		con  int
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
	}{
		{
			name: "Test 1: max score",
			args: args{
				con:  200,
				tech: 500,
			},
			want:  "5/5",
			want1: "5/5",
		},

		{
			name: "Test 2: half score",
			args: args{
				con:  100,
				tech: 250,
			},
			want:  "2.5/5",
			want1: "2.5/5",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ConvertScore(tt.args.tech, tt.args.con)
			if got != tt.want {
				t.Errorf("ConvertScore() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ConvertScore() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_roundFloat(t *testing.T) {
	type args struct {
		val       float64
		precision uint
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test 1",
			args: args{
				val:       5.00000,
				precision: 1,
			},
			want: 5.0,
		},
	}
	// TODO: Add test cases.

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := roundFloat(tt.args.val, tt.args.precision); got != tt.want {
				t.Errorf("roundFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}
