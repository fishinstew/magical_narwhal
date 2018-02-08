package grifts

import (
	"github.com/fishinstew/magical_narwhal/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
