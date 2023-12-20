package entity

type Process struct {
	Command string
	PID     int
	User    string
	FD      string
	Type    string
	Device  string
	Size    string
	Node    string
	Name    string
}
