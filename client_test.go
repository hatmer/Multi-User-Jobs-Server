package main

import "testing"

func TestAdd(t *testing.T){

    got := start("ls README.md")
    want := "README.md"

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}
