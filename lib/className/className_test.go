package className

import (
	"fmt"
	"testing"
)

func Test_NewClassName(t *testing.T) {
	var tests = []struct {
		input  string
		output ClassName
	}{
		{"block", ClassName{
			block: "block",
		}},
		{"block__element", ClassName{
			block:   "block",
			element: "element",
		}},
		{"block_modificator", ClassName{
			block:       "block",
			modificator: "modificator",
		}},
		{"block_modificator_value", ClassName{
			block:       "block",
			modificator: "modificator",
			value:       "value",
		}},
		{"block__element_modificator_value", ClassName{
			block:       "block",
			element:     "element",
			modificator: "modificator",
			value:       "value",
		}},
		{"block__element_modificator", ClassName{
			block:       "block",
			element:     "element",
			modificator: "modificator",
		}},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%v, %+v", tt.input, tt.output)
		t.Run(testName, func(t *testing.T) {
			got := *NewClassName(tt.input)
			if tt.output != got {
				t.Errorf("got %+v, want %+v", got, tt.output)
			}
		})
	}
}

func Test__assembleDirPath(t *testing.T) {
	var tests = []struct {
		input  ClassName
		output string
	}{
		{
			input: ClassName{
				block: "block",
			},
			output: "blocks/block",
		},
		{
			input: ClassName{
				block:   "block",
				element: "element",
			},
			output: "blocks/block/__element",
		},
		{
			input: ClassName{
				block:       "block",
				element:     "element",
				modificator: "modificator",
			},
			output: "blocks/block/__element/_modificator",
		},
		{
			input: ClassName{
				block:       "block",
				element:     "element",
				modificator: "modificator",
				value:       "value",
			},
			output: "blocks/block/__element/_modificator",
		},
		{
			input: ClassName{
				block:       "block",
				modificator: "modificator",
			},
			output: "blocks/block/_modificator",
		},
		{
			input: ClassName{
				block:       "block",
				modificator: "modificator",
				value:       "value",
			},
			output: "blocks/block/_modificator",
		},
	}
	for _, tt := range tests {
		testName := fmt.Sprintf("%v", tt.input)
		t.Run(testName, func(t *testing.T) {
			got := tt.input.GetDirPath()
			if got != tt.output {
				t.Errorf("got %v, expected %v", got, tt.output)
			}
		})
	}
}

func Test__getFileName(t *testing.T) {
	var tests = []struct {
		input  ClassName
		output string
	}{
		{
			input: ClassName{
				block: "block",
			},
			output: "block.css",
		},
		{
			input: ClassName{
				block:   "block",
				element: "element",
			},
			output: "block__element.css",
		},
		{
			input: ClassName{
				block:       "block",
				element:     "element",
				modificator: "modificator",
			},
			output: "block__element_modificator.css",
		},
		{
			input: ClassName{
				block:       "block",
				element:     "element",
				modificator: "modificator",
				value:       "value",
			},
			output: "block__element_modificator_value.css",
		},
		{
			input: ClassName{
				block:       "block",
				modificator: "modificator",
			},
			output: "block_modificator.css",
		},
		{
			input: ClassName{
				block:       "block",
				modificator: "modificator",
				value:       "value",
			},
			output: "block_modificator_value.css",
		},
	}
	for _, tt := range tests {
		testName := fmt.Sprintf("%v", tt.input)
		t.Run(testName, func(t *testing.T) {
			got := tt.input.GetFileName()
			if got != tt.output {
				t.Errorf("got %v, expected %v", got, tt.output)
			}
		})
	}
}

func Test__GetFilePath(t *testing.T) {
	var tests = []struct {
		input  ClassName
		output string
	}{
		{
			input: ClassName{
				block: "block",
			},
			output: "blocks/block/block.css",
		},
		{
			input: ClassName{
				block:   "block",
				element: "element",
			},
			output: "blocks/block/__element/block__element.css",
		},
		{
			input: ClassName{
				block:       "block",
				element:     "element",
				modificator: "modificator",
			},
			output: "blocks/block/__element/_modificator/block__element_modificator.css",
		},
		{
			input: ClassName{
				block:       "block",
				element:     "element",
				modificator: "modificator",
				value:       "value",
			},
			output: "blocks/block/__element/_modificator/block__element_modificator_value.css",
		},
		{
			input: ClassName{
				block:       "block",
				modificator: "modificator",
			},
			output: "blocks/block/_modificator/block_modificator.css",
		},
		{
			input: ClassName{
				block:       "block",
				modificator: "modificator",
				value:       "value",
			},
			output: "blocks/block/_modificator/block_modificator_value.css",
		},
	}
	for _, tt := range tests {
		testName := fmt.Sprintf("%v", tt.input)
		t.Run(testName, func(t *testing.T) {
			got := tt.input.GetFilePath()
			if got != tt.output {
				t.Errorf("got %v, expected %v", got, tt.output)
			}
		})
	}
}

func Test__GetCssTemplate(t *testing.T) {
	var tests = []struct {
		input  ClassName
		output string
	}{
		{
			input: ClassName{
				block: "block",
			},
			output: ".block {\n    \n}\n",
		},
		{
			input: ClassName{
				block:   "block",
				element: "element",
			},
			output: ".block__element {\n    \n}\n",
		},
		{
			input: ClassName{
				block:       "block",
				element:     "element",
				modificator: "modificator",
			},
			output: ".block__element_modificator {\n    \n}\n",
		},
		{
			input: ClassName{
				block:       "block",
				element:     "element",
				modificator: "modificator",
				value:       "value",
			},
			output: ".block__element_modificator_value {\n    \n}\n",
		},
		{
			input: ClassName{
				block:       "block",
				modificator: "modificator",
			},
			output: ".block_modificator {\n    \n}\n",
		},
		{
			input: ClassName{
				block:       "block",
				modificator: "modificator",
				value:       "value",
			},
			output: ".block_modificator_value {\n    \n}\n",
		},
	}
	for _, tt := range tests {
		testName := fmt.Sprintf("%v", tt.input)
		t.Run(testName, func(t *testing.T) {
			got := tt.input.GetCssTemplate()
			if got != tt.output {
				t.Errorf("got %v, expected %v", got, tt.output)
			}
		})
	}
}

func Test__GetImportPath(t *testing.T) {
	var tests = []struct {
		input  ClassName
		output string
	}{
		{
			input: ClassName{
				block: "block",
			},
			// importing to css from pages folder/
			output: "../blocks/block/block.css",
		},
		{
			input: ClassName{
				block:   "block",
				element: "element",
			},
			// importing to block css/
			output: "__element/block__element.css",
		},
		{
			input: ClassName{
				block:       "block",
				element:     "element",
				modificator: "modificator",
			},
			output: "_modificator/block__element_modificator.css",
		},
		{
			input: ClassName{
				block:       "block",
				element:     "element",
				modificator: "modificator",
				value:       "value",
			},
			output: "_modificator/block__element_modificator_value.css",
		},
		{
			input: ClassName{
				block:       "block",
				modificator: "modificator",
			},
			output: "_modificator/block_modificator.css",
		},
		{
			input: ClassName{
				block:       "block",
				modificator: "modificator",
				value:       "value",
			},
			output: "_modificator/block_modificator_value.css",
		},
	}
	for _, tt := range tests {
		testName := fmt.Sprintf("%v", tt.input)
		t.Run(testName, func(t *testing.T) {
			got := tt.input.GetImportPath()
			if got != tt.output {
				t.Errorf("got %v, expected %v", got, tt.output)
			}
		})
	}
}
