package ewords

import (
	"flag"
	"fmt"
)

type FlagDef struct {
	Flag  string
	Field string
	Help  string
	Value interface{}
}

type Opts interface {
	SetInt(opt string, value *int)
	SetBool(opt string, value *bool)
	SetString(field string, value *string)
}

func ParseArgs(flags []*FlagDef, s Opts) error {
	err := initFlag(s, flags)
	if err != nil {
		return fmt.Errorf("error while parsing arguments: %v", err)
	}
	flag.Parse()
	err = setFlag(s, flags)
	if err != nil {
		return fmt.Errorf("error while inserting parsed arguments: %v", err)
	}
	return nil
}

func initFlag(o Opts, flags []*FlagDef) error {
	for _, f := range flags {
		switch f.Value.(type) {
		case int:
			f.Value = flag.Int(f.Flag, f.Value.(int), f.Help)
		case bool:
			f.Value = flag.Bool(f.Flag, f.Value.(bool), f.Help)
		case string:
			f.Value = flag.String(f.Flag, f.Value.(string), f.Help)
		default:
			return fmt.Errorf("can't parse flag: %v, it has unknown type: %T", f, f)
		}
	}
	return nil
}

func setFlag(o Opts, flags []*FlagDef) error {
	for _, f := range flags {
		switch f.Value.(type) {
		case *int:
			o.SetInt(f.Field, f.Value.(*int))
		case *bool:
			o.SetBool(f.Field, f.Value.(*bool))
		case *string:
			o.SetString(f.Field, f.Value.(*string))
		default:
			return fmt.Errorf("can't insert flag: %v, it has unknown type: %T", f, f)
		}
	}
	return nil
}
