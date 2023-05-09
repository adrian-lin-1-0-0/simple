package simple

import (
	"reflect"
	"testing"
)

func initTree() *node {
	root := newNode(nil)
	root.addRoute("/p/:lang/doc", nil)
	root.addRoute("/users/:id", nil)
	root.addRoute("/asserts/*", nil)
	return root
}

func Test_node_getRoute(t *testing.T) {

	tree := initTree()

	type args struct {
		fullPath string
	}
	type want struct {
		path     string
		fullPath string
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			args: args{
				fullPath: "/p/en/doc",
			},
			want: want{
				path:     "doc",
				fullPath: "/p/:lang/doc",
			},
		},
		{
			args: args{
				fullPath: "/users/1",
			},
			want: want{
				path:     ":id",
				fullPath: "/users/:id",
			},
		},
		{
			args: args{
				fullPath: "/asserts/1/2/3",
			},
			want: want{
				path:     "*",
				fullPath: "/asserts/*",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tree.getRoute(tt.args.fullPath)
			if got.path != tt.want.path || got.fullPath != tt.want.fullPath {
				t.Errorf("node.getRoute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fullPathToPaths(t *testing.T) {
	type args struct {
		fullPath string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			args: args{
				fullPath: "/p/:lang/doc",
			},
			want: []string{"p", ":lang", "doc"},
		},
		{
			args: args{
				fullPath: "/",
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fullPathToPaths(tt.args.fullPath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fullPathToPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_node_parseParams(t *testing.T) {

	root := newNode(nil)
	root.addRoute("/users/:userID/homes/:homeID/rooms/:rooms", nil)
	node := root.getRoute("/users/:userID/homes/:homeID/rooms/:rooms")
	params := node.parseParams("/users/userxxx/homes/homexxx/rooms/roomxxx")
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				name: "userID",
			},
			want: "userxxx",
		},
		{
			args: args{
				name: "homeID",
			},
			want: "homexxx",
		},
		{
			args: args{
				name: "rooms",
			},
			want: "roomxxx",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := params[tt.args.name]; got != tt.want {
				t.Errorf("node.parseParams() = %v, want %v", got, tt.want)
			}
		})
	}
}
