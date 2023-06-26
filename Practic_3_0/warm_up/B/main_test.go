package main

import (
	"reflect"
	"testing"
	"time"
)

func Test_parseTasks(t *testing.T) {
	type args struct {
		date  time.Time
		lines []string
	}
	tests := []struct {
		name    string
		args    args
		want    []Task
		wantErr bool
	}{

	test

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseTasks(tt.args.date, tt.args.lines)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseTasks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseTasks() got = %v, want %v", got, tt.want)
			}
		})
	}
}
