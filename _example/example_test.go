package main

import (
	"database/sql"
	"fmt"
	"github.com/ToQoz/gopwt/assert"
	"reflect"
	"testing"
)

func TestBasicLit(t *testing.T) {
	assert.OK(t, "a" == "b")
	assert.OK(t, 1 == 2)

	a := 1
	b := 2
	c := 3
	assert.OK(t, a+c == b)
	assert.OK(t, (a+c)+a == b)
	assert.OK(t, `foo
bar` == "bar")
}

func TestMapType(t *testing.T) {
	k := "b--------key"
	v := "b------value"
	assert.OK(t, reflect.DeepEqual(map[string]string{}, map[string]string{
		"a": "a",
		k:   v,
	}))
}

func TestArrayType(t *testing.T) {
	index := 1
	assert.OK(t, []int{
		1,
		2,
	}[index] == 3)
}

func TestStructType(t *testing.T) {
	foox := "foo------x"
	assert.OK(t, reflect.DeepEqual(struct{ Name string }{foox}, struct{ Name string }{"foo"}))
	assert.OK(t, reflect.DeepEqual(struct{ Name string }{Name: foox}, struct{ Name string }{Name: "foo"}))
}

func TestNestedCallExpr(t *testing.T) {
	rev := func(a bool) bool {
		return !a
	}

	assert.OK(t, rev(rev(rev(true))))
}

func TestCallWithNonIdempotentFunc(t *testing.T) {
	i := 0
	incl := func() int {
		i++
		return i
	}

	assert.OK(t, incl()+incl() == incl()+incl())
	assert.OK(t, incl() == incl())
	assert.OK(t, incl() == incl())
	assert.OK(t, incl() == incl())
	assert.OK(t, incl() == incl())
	assert.OK(t, (incl() == incl()) != (incl() == incl()))

	i2 := 0
	incl2 := func(i3 int) int {
		i2 += i3
		return i2
	}

	assert.OK(t, incl2(incl2(2)) == 10)
}

func TestPkgValue(t *testing.T) {
	assert.OK(t, sql.ErrNoRows == fmt.Errorf("error"))
}
