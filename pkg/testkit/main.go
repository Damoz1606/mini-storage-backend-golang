package testkit

import "testing"

type TestCase[I any, T any, R any] struct {
	Name   string
	Input  func() I
	Setup  func(t *testing.T) T
	Assert func(t *testing.T, result R, err error)
}
