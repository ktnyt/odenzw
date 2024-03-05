package main

import (
	"bytes"
	"fmt"
	"github.com/go-gts/flags"
	"os"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	cases := []struct {
		args       []string
		expects    []string
		notExpects []string
	}{
		{
			args:       []string{"denzow"},
			expects:    []string{"denzow", "odenzw", "wezdon"},
			notExpects: []string{"ddenzo"},
		},
	}

	for i, tt := range cases {
		tt := tt

		t.Run(fmt.Sprintf("case %v", i), func(t *testing.T) {
			r, w, _ := os.Pipe()
			os.Stdin = r
			os.Stdout = w

			ctx := &flags.Context{}
			ctx.Args = tt.args

			err := run(ctx)
			if err != nil {
				t.Errorf("Error running the app: %v", err)
			}
			_ = w.Close()

			var buffer bytes.Buffer
			if _, err := buffer.ReadFrom(r); err != nil {
				t.Fatalf("fail read buf: %v", err)
			}

			actual := buffer.String()
			for _, expect := range tt.expects {
				if !strings.Contains(actual, expect) {
					t.Fatalf("unexpected output: expect=%v, actual=%v", expect, actual)
				}
			}
			for _, notExpect := range tt.notExpects {
				if strings.Contains(actual, notExpect) {
					t.Fatalf("unexpected output: notExpect=%v, actual=%v", notExpect, actual)
				}
			}
		})
	}
}
