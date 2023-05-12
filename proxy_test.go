package simple

import "testing"

func assertPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	f()
}

func TestProxy_AddRoute_Panic(t *testing.T) {

	type args struct {
		proxyPath string
		target    string
		handlers  []HandlerFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{
				proxyPath: "/api",
				target:    "htt://localhost:8080",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := DefaultProxy()
			assertPanic(t, func() {
				p.GET(tt.args.proxyPath, tt.args.target, tt.args.handlers...)
			})
		})
	}
}

func TestProxy_AddRoute(t *testing.T) {

	type args struct {
		proxyPath string
		target    string
		handlers  []HandlerFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{
				proxyPath: "/api",
				target:    "http://localhost:8080",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := DefaultProxy()
			p.GET(tt.args.proxyPath, tt.args.target, tt.args.handlers...)
		})
	}
}
