package components

import "github.com/a-h/templ"

type NavbarParams struct {
	ID         string
	Title      string
	HxTarget   string
	UserID     string
	PageTarget string
}

type NavbarBuilder NavbarParams

func NewNavbarBuilder() *NavbarBuilder {
	return &NavbarBuilder{}
}

func (b *NavbarBuilder) WithID(v string) *NavbarBuilder {
	b.ID = v
	return b
}

func (b *NavbarBuilder) WithTitle(v string) *NavbarBuilder {
	b.Title = v
	return b
}

func (b *NavbarBuilder) WithHxTarget(v string) *NavbarBuilder {
	b.HxTarget = v
	return b
}

func (b *NavbarBuilder) WithUserID(v string) *NavbarBuilder {
	b.UserID = v
	return b
}

func (b *NavbarBuilder) WithPageTarget(v string) *NavbarBuilder {
	b.PageTarget = v
	return b
}

func (b NavbarBuilder) Build() templ.Component {
	p := NavbarParams(b)
	return Navbar(p)
}
