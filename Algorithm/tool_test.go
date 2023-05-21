package Algorithm

import "testing"

func TestStartAlgorithm(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			StartAlgorithm()
		})
	}
}

func Test_findPrediction(t *testing.T) {
	type args struct {
		data string
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
			if got := findPrediction(tt.args.data); got != tt.want {
				t.Errorf("findPrediction() = %v, want %v", got, tt.want)
			}
		})
	}
}
