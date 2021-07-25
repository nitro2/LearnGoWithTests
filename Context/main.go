package main

import "context_pkg"

type WebStore struct {
	response string
}

func (s *WebStore) Fetch() string {
	return s.response
}

func main() {
	ser := context_pkg.Server(&WebStore)

}
