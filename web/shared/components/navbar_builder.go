package components

import (
	"github.com/a-h/templ"

	"github.com/zafir-co-ao/onna-narciso/internal/auth"
)

type NavbarParams struct {
	ID         string
	Title      string
	HxTarget   string
	User       auth.UserOutput
	PageTarget string
}

type NavbarBuilder NavbarParams

func NewNavbarBuilder() *NavbarBuilder {
	return &NavbarBuilder{}
}

func (b *NavbarBuilder) WithID(id string) *NavbarBuilder {
	b.ID = id
	return b
}

func (b *NavbarBuilder) WithTitle(t string) *NavbarBuilder {
	b.Title = t
	return b
}

func (b *NavbarBuilder) WithHxTarget(t string) *NavbarBuilder {
	b.HxTarget = t
	return b
}

func (b *NavbarBuilder) WithUser(u auth.UserOutput) *NavbarBuilder {
	b.User = u
	return b
}

func (b *NavbarBuilder) WithPageTarget(t string) *NavbarBuilder {
	b.PageTarget = t
	return b
}

func (b NavbarBuilder) Build() templ.Component {
	p := NavbarParams(b)
	return Navbar(p)
}
