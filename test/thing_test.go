package main

import (
	"testing"
)

var zero = Thing(0)

func TestAnd(t *testing.T) {
	True := ThingPredicate(func(x Thing) bool { return true })
	False := ThingPredicate(func(x Thing) bool { return false })

	ff := False.And(False)
	ft := False.And(True)
	tt := True.And(True)
	tf := True.And(False)

	if ff(zero) {
		t.Errorf("And FF should not be true")
	}
	if ft(zero) {
		t.Errorf("And FT should not be true")
	}
	if !tt(zero) {
		t.Errorf("And TT should be true")
	}
	if tf(zero) {
		t.Errorf("And TF should not be true")
	}
}

func TestOr(t *testing.T) {
	True := ThingPredicate(func(x Thing) bool { return true })
	False := ThingPredicate(func(x Thing) bool { return false })

	ff := False.Or(False)
	ft := False.Or(True)
	tt := True.Or(True)
	tf := True.Or(False)

	if ff(zero) {
		t.Errorf("Or FF should not be true")
	}
	if !ft(zero) {
		t.Errorf("Or FT should be true")
	}
	if !tt(zero) {
		t.Errorf("Or TT should be true")
	}
	if !tf(zero) {
		t.Errorf("Or TF should be true")
	}
}

func TestNot(t *testing.T) {
	True := ThingPredicate(func(x Thing) bool { return true })
	False := ThingPredicate(func(x Thing) bool { return false })

	expectFalse := True.Not()
	expectTrue := False.Not()

	if expectFalse(zero) {
		t.Errorf("Not True should not be true")
	}
	if !expectTrue(zero) {
		t.Errorf("Not False should be true")
	}
}
