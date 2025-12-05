package main

func Day6A(io *IO) {
	var s string
	for io.Read(&s) != nil {
		io.Write("%s\n", s)
	}
}

func Day6B(io *IO) {
	var s string
	io.Readln(&s)
	io.Write("%s\n", s)
}
