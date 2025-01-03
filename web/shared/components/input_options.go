package components

import (
	"fmt"

	"github.com/a-h/templ"
)

type _inputState struct {
	id          string
	label       string
	name        string
	_type       string
	class       string
	step        string
	value       string
	placeholder string
	required    string
	hxGet       string
	hxPost      string
	hxTarget    string
	hxSwap      string
	hxTrigger   string
	hxOn        map[string]string
	options     []InputOption
}

type _inputStateOpt func(s interface{})

func newInputState(opts ..._inputStateOpt) *_inputState {
	s := &_inputState{
		hxOn: make(map[string]string),
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func WithId(v string) _inputStateOpt {
	return func(s interface{}) {
		switch s := s.(type) {
		case *_inputState:
			s.id = v
		}
	}
}
func WithLabel(v string) _inputStateOpt {
	return func(s interface{}) {
		switch s := s.(type) {
		case *_inputState:
			s.label = v
		}
	}
}
func WithName(v string) _inputStateOpt {
	return func(s interface{}) {
		switch s := s.(type) {
		case *_inputState:
			s.name = v
		}
	}
}
func WithValue(v string) _inputStateOpt {
	return func(s interface{}) {
		switch s := s.(type) {
		case *_inputState:
			s.value = v
		}
	}
}
func WithType(v string) _inputStateOpt {
	return func(s interface{}) {
		switch s := s.(type) {
		case *_inputState:
			s._type = v
		}
	}
}
func WithClass(v string) _inputStateOpt {
	return func(s interface{}) {
		switch s := s.(type) {
		case *_inputState:
			s.class = v
		}
	}
}
func WithStep(v string) _inputStateOpt {
	return func(s interface{}) {
		switch s := s.(type) {
		case *_inputState:
			s.step = v
		}
	}
}
func WithRequired(v string) _inputStateOpt {
	return func(s interface{}) {
		switch s := s.(type) {
		case *_inputState:
			s.required = v
		}
	}
}
func WithPlaceholder(v string) _inputStateOpt {
	return func(s interface{}) {
		switch s := s.(type) {
		case *_inputState:
			s.placeholder = v
		}
	}
}
func WithHxGet(v string) _inputStateOpt {
	return func(s interface{}) {
		switch s := s.(type) {
		case *_inputState:
			s.hxGet = v
		}
	}
}
func WithHxPost(v string) _inputStateOpt {
	return func(s interface{}) {
		switch s := s.(type) {
		case *_inputState:
			s.hxPost = v
		}
	}
}
func WithHxTarget(v string) _inputStateOpt {
	return func(s interface{}) {
		switch s := s.(type) {
		case *_inputState:
			s.hxTarget = v
		}
	}
}
func WithHxSwap(v string) _inputStateOpt {
	return func(s interface{}) {
		switch s := s.(type) {
		case *_inputState:
			s.hxSwap = v
		}
	}
}
func WithHxTrigger(v string) _inputStateOpt {
	return func(s interface{}) {
		switch s := s.(type) {
		case *_inputState:
			s.hxTrigger = v
		}
	}
}

func WithHxOn(k, v string) _inputStateOpt {
	return func(s interface{}) {
		switch s := s.(type) {
		case *_inputState:
			s.hxOn[k] = v
		}
	}
}

func ParseHxOn(s *_inputState) templ.Attributes {

	attrs := templ.Attributes{}

	for k, v := range s.hxOn {
		key := fmt.Sprintf("hx-on:%s", k)
		attrs[key] = v
	}

	return attrs
}

type InputOption [2]string

func WithOptions(v ...InputOption) _inputStateOpt {
	return func(s interface{}) {
		switch s := s.(type) {
		case *_inputState:
			if s.options == nil {
				s.options = make([]InputOption, 0)
			}

			s.options = append(s.options, v...)
		}
	}
}
