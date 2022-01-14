package main

import "strings"

type Const struct {
	Name  string
	Type  string
	Value string
}

type Type struct {
	Name  string
	Type  string
	Const []Const
}

type Field struct {
	Name      string
	NameJSON  string
	NameInput string
	Type      string
	IsArray   bool
	Const     string
}

func (f Field) Default(s string) Field {
	f.Const = s
	return f
}

type Struct struct {
	Doc       string
	Name      string
	Required  []Field
	Optional  []Field
	Implement []string
}

type File struct {
	Type    []Type
	Package string
	Struct  []Struct
}

func F(name string, t string) Field {
	isArray := false
	if strings.HasPrefix(t, "[]") {
		isArray = true
		t = strings.TrimPrefix(t, "[]")
	}
	return Field{
		Name:      name,
		NameJSON:  snake(name),
		NameInput: lower(name),
		Type:      t,
		IsArray:   isArray,
	}
}

func FS(s string) Field {
	ss := strings.Split(s, ",")
	f := F(ss[0], ss[1])
	if len(ss) == 3 {
		f = f.Default(ss[2])
	}
	return f
}

func FSS(ss ...string) []Field {
	fs := make([]Field, 0, len(ss))
	for _, s := range ss {
		fs = append(fs, FS(s))
	}
	return fs
}

func TT(t ...Type) []Type {
	return t
}

func T(s string, vs ...string) Type {
	ss := strings.Split(s, ",")
	t := Type{
		Name:  ss[0],
		Type:  ss[1],
		Const: nil,
	}
	for _, v := range vs {
		vv := strings.Split(v, ",")
		t.Const = append(t.Const, Const{
			Name:  vv[0],
			Type:  t.Type,
			Value: vv[1],
		})
	}

	return t
}
