package main

import (
	"fmt"
	"os"
	"strings"

	giant "github.com/multiverse-os/cli/framework/text/banner/fonts/giant"
)

// TODO: Start writing tools for generating Go source code files on-the-fly to
// simplify tasks like this and others. This is a very common usecase and saves
// A LOT of time.

type DataType int

const (
	String DataType = iota
	Rune
	Byte
	Int
	Int8
	Int32
	Int64
	Float
	Float64
	Interface
)

func (self DataType) String() string {
	switch self {
	case String:
		return "string"
	case Rune:
		return "rune"
	case Byte:
		return "byte"
	case Int:
		return "int"
	case Int8:
		return "int8"
	case Int32:
		return "int32"
	case Int64:
		return "int64"
	case Float:
		return "float"
	case Float64:
		return "float64"
	default:
		return "interface"
	}
}

type Element struct {
	Type  DataType
	Index int
	Value interface{}
}

type Slice struct {
	ElementType DataType
	Elements    []string
}

func NewSlice(dataType DataType, elements []string) *Slice {
	return &Slice{
		ElementType: dataType,
		Elements:    elements,
	}
}

func (self *Slice) Size() int      { return len(self.Elements) }
func (self *Slice) Length() int    { return self.Size() }
func (self *Slice) LastIndex() int { return self.Size() - 1 }
func (self *Slice) Last() string   { return self.Elements[self.LastIndex()] }
func (self *Slice) First() string  { return self.Elements[0] }

func (self *Slice) String() (output string) {
	// TODO: Consider how we would do nested slices and maps
	output += "[]" + self.ElementType.String() + "{"

	for index, element := range self.Elements {
		output += "\"" + element + "\""
		if self.LastIndex() == index {
			output += "}"
		} else {
			output += ", "
		}
	}
	return output
}

type Map struct {
	KeyType   DataType
	ValueType DataType
}

func (self Map) String() (output string) {
	// TODO: We need to consider how we would do nested maps, nested strings
	output += "map[" + self.KeyType.String() + "]" + self.ValueType.String() + "{"
	output += "  \"keyName\": \"value\","
	output += "}"
	return output
}

type Attribute struct {
	Name     string
	DataType DataType
	Value    string
}

type Struct struct {
	Scope      Scope
	Attributes []Attribute
	Anonymous  bool
}

type Scope int

const (
	Global Scope = iota
	Local
)

type Constant struct {
	Scope Scope
	Name  string
	Value string
}

func (self Constant) String() string { return "const " + self.Name + " = " + self.Value }

type Variable struct {
	Scope Scope
	Type  DataType
	Name  string
	Value string
}

func (self Variable) String() string {
	switch self.Scope {
	case Local:
		return self.Name + " := " + self.Name
	default: // Global
		return "var " + self.Name + " = " + self.Value
	}
}

type Package struct {
	Name      string
	IsLibrary bool
	Files     []*SourceFile
}

func NewPackage(name string) *Package {
	return &Package{
		Name: name,
	}
}

func (self *Package) Library(isLibrary bool) *Package {
	self.IsLibrary = isLibrary
	return self
}

func (self *Package) AddFile(file *SourceFile) *Package {
	self.Files = append(self.Files, file)
	return self
}

func (self Package) String() string { return "package " + self.Name + "\n\n" }

type Function struct {
	Variables  []Variable
	Constants  []Constant
	ReturnType DataType
}

// TODO: For now, for loading a binary as an embedded file in the source code.
// It could also be used for later automatically building a project, signing and
// releasing
type Binary struct {
	Name string
	Path string
	Data []byte
}

// TODO: The idea here is writing multiple files (whole package) together, this
// may be useful
type Source struct {
	Name     string
	Makefile bool
	Binaries []*Binary
	Packages []*Package
}

type SourceFile struct {
	Filepath        string
	Package         *Package
	Imports         []string
	GlobalVariables []Variable
	GlobalConstants []Constant
	Functions       []Function
	File            *os.File
}

func NewSourceFile(p *Package, filepath string) *SourceFile {
	osFile, err := os.Create(filepath)
	if err != nil {
		fmt.Println("failed to open output file:", err)
		os.Exit(1)
	}
	sourceFile := &SourceFile{
		Package: p,
		File:    osFile,
	}
	sourceFile.Package.AddFile(sourceFile)
	return sourceFile
}

func (self *SourceFile) Close()                     { self.File.Close() }
func (self *SourceFile) WriteString(content string) { self.File.WriteString(content) }

func (self *SourceFile) WriteFile() string {
	// TODO: Write package, then imports, then globals (constants/variables), then
	// global structs, then functions.
	return ""
}

func main() {
	fmt.Println("fix font format")
	fmt.Println("===============")
	fmt.Println("fonts are built wrong, so we are going to build a go source file with the corrected format")

	sourceFile := NewSourceFile(NewPackage("giant"), "giant.go")
	defer sourceFile.Close()

	fmt.Println("writing to file:", sourceFile.Filepath, " included in package:", sourceFile.Package.Name)

	sourceFile.WriteString(sourceFile.Package.String())

	sourceFile.WriteString("var Characters = map[string][]string{\n")
	for key, character := range giant.Characters {
		lines := strings.Split(character, "\n")
		if key == `"` {
			key = `\` + key
		}
		element := "  \"" + key + "\": " + NewSlice(String, lines).String() + ", \n"
		sourceFile.WriteString(element)
	}

	sourceFile.WriteString("}\n")

	fmt.Println("Completed writing source files!")
}
