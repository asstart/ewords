package ewords

type TermSource struct {
	Term          string
	Transcription string
	Definition    string
	Example       string
}

type TermDefenition struct {
	Term          string
	Transcription string
	Defenition    string
}

type TermExample struct {
	Term    string
	Example string
}

type Definer interface {
	MakeDefenition(t *TermSource) TermDefenition
}

type Exampler interface {
	MakeExample(t *TermSource) TermExample
}

type DefaultTerm struct{}

func ToSingleDefenetion(d Definer, t *TermSource) TermDefenition {
	return d.MakeDefenition(t)
}

func ToDefenitions(d Definer, t []TermSource) []TermDefenition {
	res := []TermDefenition{}
	for _, v := range t {
		res = append(res, ToSingleDefenetion(d, &v))
	}
	return res
}

func ToSingleExample(e Exampler, t *TermSource) TermExample {
	return e.MakeExample(t)
}

func ToExamples(e Exampler, t []TermSource) []TermExample {
	res := []TermExample{}
	for _, v := range t {
		res = append(res, ToSingleExample(e, &v))
	}
	return res
}

func (d *DefaultTerm) MakeDefenition(t *TermSource) TermDefenition {
	return TermDefenition{
		Term:          t.Term,
		Transcription: t.Transcription,
		Defenition:    t.Definition,
	}
}

func (d *DefaultTerm) MakeExample(t *TermSource) TermExample {
	return TermExample{
		Term:    t.Term,
		Example: t.Example,
	}
}
