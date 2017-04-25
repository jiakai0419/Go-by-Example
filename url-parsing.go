package main

import "fmt"
import "net"
import "net/url"

func main() {
	P := fmt.Println

	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	P(u.Scheme)

	P(u.User)
	P(u.User.Username())
	p, _ := u.User.Password()
	P(p)

	P(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	P(host)
	P(port)

	P(u.Path)
	P(u.Fragment)

	P(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	P(m)
	P(m["k"][0])
}
