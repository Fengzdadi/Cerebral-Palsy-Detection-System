package Algorithm

import (
	"Cerebral-Palsy-Detection-System/model"
	"testing"
)

func TestStartAlgorithm(t *testing.T) {
	type args struct {
		res *model.Result
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			StartAlgorithm(tt.args.res)
		})
	}
}
