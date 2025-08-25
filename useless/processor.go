package useless

type Processor interface {
	Process(input string) (output string, err error)
}
