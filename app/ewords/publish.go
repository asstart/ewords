package ewords

type ExamplePublishier interface {
	publishExample(te []TermExample) error
}

type DefenitionPublisher interface {
	publishDefenition(td []TermDefenition) error
}
