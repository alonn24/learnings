package greetings

import "testing"

func TestHello(t *testing.T) {
	ans, err := Hello("alony")
	if ans != "Hi, alony. Welcome!" {
		t.Errorf("Hello(alony) = %v; want \"Hi, alony. Welcome!\"", ans)
	}
	if err != nil {
		t.Error("Hello(alony) = expected err not to be nil")
	}
}

func TestHelloEmpty(t *testing.T) {
	ans, err := Hello("")
	if ans != "" {
		t.Errorf("Hello(\"\") = %v; expected to return \"\"", ans)
	}
	if err == nil {
		t.Errorf("Hello(\"\") = %v; to return error", ans)
	}
}

func TestHellos(t *testing.T) {
	ans, err := Hellos("alony")
	if err != nil {
		t.Error("Hellos(alony), expected err not to be nil")
	}
	if ans["alony"] != "Hi, alony. Welcome!" {
		t.Errorf("Hellos(alony) = %v, wanted \"Hi, alony. Welcome!\"", ans)
	}
}
