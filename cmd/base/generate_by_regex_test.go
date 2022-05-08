package base

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestGenerate_generateRegex(t *testing.T) {
	rand.Seed(time.Now().Unix())
	g := &Generate{MaxLength: 12}
	fmt.Println(g.Generate("[a-zA-Z0-9]{4,4}",10))
}

func TestRegexValueGenerate_Generate(t *testing.T) {
	type args struct {
		regex []string
		min   int
		max   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RegexValueGenerate{}
			if got := r.Generate(tt.args.regex, tt.args.min, tt.args.max); got != tt.want {
				t.Errorf("Generate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_appendArray(t *testing.T) {
	type args struct {
		target []string
		begin  string
		end    string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := appendArray(tt.args.target, tt.args.begin, tt.args.end); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("appendArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateWordArray(t *testing.T) {
	type args struct {
		regex string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateWordArray(tt.args.regex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateWordArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getBigBracketsMinMax(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		wantMin int
		wantMax int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMin, gotMax := getBigBracketsMinMax(tt.args.str)
			if gotMin != tt.wantMin {
				t.Errorf("getBigBracketsMinMax() gotMin = %v, want %v", gotMin, tt.wantMin)
			}
			if gotMax != tt.wantMax {
				t.Errorf("getBigBracketsMinMax() gotMax = %v, want %v", gotMax, tt.wantMax)
			}
		})
	}
}

func Test_getEncodeMean(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getEncodeMean(tt.args.word); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getEncodeMean() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_scanLittleContent(t *testing.T) {
	type args struct {
		src []string
		i   int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := scanLittleContent(tt.args.src, tt.args.i)
			if got != tt.want {
				t.Errorf("scanLittleContent() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("scanLittleContent() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_scanMiddleContent(t *testing.T) {
	type args struct {
		src []string
		i   int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := scanMiddleContent(tt.args.src, tt.args.i)
			if got != tt.want {
				t.Errorf("scanMiddleContent() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("scanMiddleContent() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
