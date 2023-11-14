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
