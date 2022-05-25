package services

import "testing"

func TestHealthCheckService(t *testing.T) {
	
	cases := []struct {
		input string
		expectOutput int //1 for pass, 0 for fail
	}{
		{ "google.com", 1},
		{ "www.google.com", 1},
		{ "http://www.google.com", 1},
		{ "https://www.google.com", 1},
		{ "1234qweasdzxc.com", 0},
		{ "www.1234qweasdzxc.com", 0},
		{ "http://www.1234qweasdzxc.com", 0},
		{ "https://www.1234qweasdzxc.com", 0},
	}

	// output := make(int,0)
	for _,c := range cases {
		ch := make(chan int)
		go HealthCheckService(c.input, ch)
		output := <-ch
		if output != c.expectOutput {
			t.Errorf("incorrect output for `%s`: expected `%d` but got `%d`", c.input, c.expectOutput, output)
		}
	}
}

func TestCheckURL(t *testing.T) {
	cases := []struct {
		input string
		expectOutput string
	}{
		{ "google.com", "http://google.com"},
		{ "www.google.com", "http://www.google.com"},
		{ "http://www.google.com", "http://www.google.com"},
		{ "https://www.google.com", "https://www.google.com"},
	}

	for _, c := range cases {
		output := checkURL(c.input)
		if output != c.expectOutput {
			t.Errorf("incorrect output for `%s`: expected `%s` but got `%s`", c.input, c.expectOutput, output)
		}
	}
}