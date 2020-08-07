package gee

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArgsInit(t *testing.T) {
	assert.NotPanics(t, func() {
		ArgsInit()
		CliArgs = []string{}
	})
}

func Test_parseArgs(t *testing.T) {
	ExtCliArgs = make(map[string]string)
	type args struct {
		args     []string
		excludes []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "yes",
			args: args{
				args:     []string{"./app", "http", "--port=:8080", "--config=2", "--debug=3"},
				excludes: []string{"config", "debug"},
			},
			want: []string{"http", "--port=:8080"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseArgs(tt.args.args, tt.args.excludes...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}
