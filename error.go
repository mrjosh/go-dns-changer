package gdc

type Error struct {
	error
	cmdOUT string
}

func (e Error) CommandOutput() string {
	return e.cmdOUT
}