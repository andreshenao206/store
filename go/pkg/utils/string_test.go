package utils

import (
	"reflect"
	"strings"
	"testing"
)

func TestStringReverse(t *testing.T) {
	st := &String{}

	type args struct {
		s string
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
	}{
		{"Combined parts of FNV hash (long string)", args{s: "3353966794807206467625917882199466832106335076677692108410633507667769210843041696344406181581330956300501893356915755575511849520176"}, "6710259481155755575196533981050036590331851816044436961403480129677667053360148012967766705336012386649912887195267646027084976693533"},
		{"String with same characters", args{s: "fff"}, "fff"},
		{"Empty string", args{s: ""}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := st.StringReverse(tt.args.s); gotResult != tt.wantResult {
				t.Errorf("StringReverse() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestRandomCharacter(t *testing.T) {
	st := &String{}

	type args struct {
		s string
	}

	const TARGETS string = "EM"
	aForRandomTest := args{s: TARGETS}

	tests := []struct {
		name string
		args args
		want string
	}{
		{"Empty string", args{s: ""}, ""},
		{"String with only 1 character", args{s: "a"}, "a"},
		{"String with 2 characters", aForRandomTest, TARGETS},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := st.StringRandomCharacter(tt.args.s)
			if tt.args == aForRandomTest {
				if !strings.Contains(tt.want, got) {
					t.Errorf("StringRandomCharacter() = %v is not in %v", got, tt.want)
				}
			} else {
				if got != tt.want {
					t.Errorf("StringRandomCharacter() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestSplitRecursive(t *testing.T) {
	st := &String{}

	type args struct {
		s    string
		size int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"Empty string", args{s: "ewoqjP7PxY4yr3iLTpLisriqt94hdyDFNgchSxGGztUrTXtNN", size: 9}, []string{"ewoqjP7Px", "Y4yr3iLTp", "Lisriqt94", "hdyDFNgch", "SxGGztUrT", "XtNN"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := st.StringSplitFixedSize(tt.args.s, tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringSplitFixedSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
