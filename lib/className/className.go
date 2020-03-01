package className

import (
	"fmt"
	"path"
	"strings"
)

type ClassName struct {
	block, element, modificator, value string
}

func NewClassName(input string) *ClassName {
	var className ClassName

	if !strings.Contains(input, "_") {
		className.block = input
		return &className
	}

	if strings.Contains(input, "__") {
		splitted := strings.Split(input, "__")
		className.block = splitted[0]
		input = splitted[1]
	}

	if !strings.Contains(input, "_") {
		className.element = input
		return &className
	}

	splitted := strings.Split(input, "_")
	switch {
	case len(splitted) == 3:
		if className.block == "" {
			className.block = splitted[0]
			className.modificator = splitted[1]
			className.value = splitted[2]
		} else {
			className.element = splitted[0]
			className.modificator = splitted[1]
			className.value = splitted[2]
		}
	case len(splitted) == 2:
		if className.block == "" {
			className.block = splitted[0]
			className.modificator = splitted[1]
		} else {
			className.element = splitted[0]
			className.modificator = splitted[1]
		}
	case len(splitted) == 1:
		className.modificator = splitted[0]
	}

	return &className
}


func (c *ClassName) GetDirPath() string {
	d := path.Join("blocks", c.block)
	if c.element != "" {
		d = path.Join(d, fmt.Sprintf("__%v", c.element))
	}
	if c.modificator != "" {
		d = path.Join(d, fmt.Sprintf("_%v", c.modificator))
	}
	return d
}

func (c *ClassName) GetClassName() string {
	f := c.block
	if c.element != "" {
		f = fmt.Sprintf("%v__%v", f, c.element)
	}
	if c.modificator != "" {
		f = fmt.Sprintf("%v_%v", f, c.modificator)
	}
	if c.value != "" {
		f = fmt.Sprintf("%v_%v", f, c.value)
	}

	return f
}

func (c *ClassName) GetFileName() string {
	return fmt.Sprintf("%v.css", c.GetClassName())
}

func (c *ClassName) GetFilePath() string {
	return path.Join(c.GetDirPath(), c.GetFileName())
}

func (c *ClassName) GetCssTemplate() string {
	return fmt.Sprintf(".%v {\n    \n}\n", c.GetClassName())
}
