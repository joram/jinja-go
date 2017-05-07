package jinja_go

type Configuration struct {
	StartStartingBytes []byte
	endStartingBytes   []byte

	BlockStartString    string // The string marking the beginning of a block.
	BlockEndString      string // The string marking the end of a block.
	VariableStartString string // The string marking the beginning of a print statement.
	VariableEndString   string // The string marking the end of a print statement.
	CommentStartString  string // The string marking the beginning of a comment.
	CommentEndString    string // The string marking the end of a comment.
	// TODO: add more Configuration from here http://jinja.pocoo.org/docs/2.9/api/#basics
}

type MarkerStringPair struct {
	Start string
	End   string
}

func NewDefaultConfig() Configuration {
	return Configuration{
		[]byte{"{"[0]},
		[]byte{"%"[0], "}"[0], "#"[0]},
		"{%",
		"%}",
		"{{",
		"}}",
		"{#",
		"#}",
	}
}

func (config *Configuration) MarkerStringPairs() []MarkerStringPair {
	return []MarkerStringPair{
		{config.BlockStartString, config.BlockEndString},
		{config.VariableStartString, config.VariableEndString},
		{config.CommentStartString, config.CommentEndString},
	}
}
