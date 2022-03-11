package main

import (
	"reflect"
	"testing"
)

func Test_averageNumber(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Success",
			args: args{
				str: "3-abw-4-cab-2-as",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := averageNumber(tt.args.str); got != tt.want {
				t.Errorf("averageNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_storyStats(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name         string
		args         args
		wantShortest string
		wantLongest  string
		wantAvg      float64
		wantResult   []string
	}{
		{
			name: "Success",
			args: args{
				str: "3-nabla-4-cab-2-a",
			},
			wantShortest: "a",
			wantLongest:  "nabla",
			wantAvg:      3,
			wantResult: []string{
				"cab",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotShortest, gotLongest, gotAvg, gotResult := storyStats(tt.args.str)
			if gotShortest != tt.wantShortest {
				t.Errorf("storyStats() gotShortest = %v, want %v", gotShortest, tt.wantShortest)
			}
			if gotLongest != tt.wantLongest {
				t.Errorf("storyStats() gotLongest = %v, want %v", gotLongest, tt.wantLongest)
			}
			if gotAvg != tt.wantAvg {
				t.Errorf("storyStats() gotAvg = %v, want %v", gotAvg, tt.wantAvg)
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("storyStats() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_testValidity(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Success",
			args: args{
				str: "3-nabla-4-cab-2-a",
			},
			want: true,
		},
		{
			name: "Fail",
			args: args{
				str: "3-nabla-4-cab-2-a-",
			},
			want: false,
		},
		{
			name: "Fail",
			args: args{
				str: "3-nabla-4-cab-2",
			},
			want: false,
		},
		{
			name: "Fail",
			args: args{
				str: "nabla-4-cab-2-as",
			},
			want: false,
		},
		{
			name: "Fail",
			args: args{
				str: "1-1",
			},
			want: false,
		},
		{
			name: "Fail",
			args: args{
				str: "nabla--4-cab-2-as",
			},
			want: false,
		},
		{
			name: "Fail",
			args: args{
				str: "",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := testValidity(tt.args.str); got != tt.want {
				t.Errorf("testValidity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_wholeStory(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Success",
			args: args{
				str: "3-nabla-4-cab-2-a",
			},
			want: "nabla cab a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wholeStory(tt.args.str); got != tt.want {
				t.Errorf("wholeStory() = %v, want %v", got, tt.want)
			}
		})
	}
}
