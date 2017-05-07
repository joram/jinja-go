package jinja_go

type INode interface {
	append(*INode)
	isClosed() bool
	toString() string
	close()
}
