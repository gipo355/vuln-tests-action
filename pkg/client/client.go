package client

type Client interface {
	Scan() error
	WriteToFile() error
	WriteToStdout() error
	CreteSarif() error
}
