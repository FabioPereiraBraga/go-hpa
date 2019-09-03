package main

import "testing"

func TestHello(t *testing.T) {

	got  := raizQuadrada();
	want := "Code.education Rocks!";
	
	if got != want {
		t.Errorf("resultado da função incorreto %s, resultado correto: %s", got, want)
	}
}