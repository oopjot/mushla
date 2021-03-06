package entities

import (
	"errors"
	"fmt"
	"strings"
)

type dir struct {
	name string
	parent Dir
	files []File
	dirs []Dir
}

type Dir interface {
	Name() string
	Exists(name string) bool
	AddFile(f File) error
	AddDir(d Dir) error
	Remove(name string) error
	Rename(name string) error
	ChangeParent(to Dir)
	List() string
	IsRoot() bool
	Parent() (Dir, error)
	FindDir(name string) (Dir, error)
	FindFile(name string) (File, error)
}

func NewDir(name string, parent Dir) (Dir, error) {
	newDir := &dir{
		name: name,
		parent: parent,
		files: []File{},
		dirs: []Dir{},
	}
	err := parent.AddDir(newDir)
	return newDir, err
}

func (d *dir) Name() string {
	return d.name
}

func (d *dir) Exists(name string) bool {
	for _, _d := range d.dirs {
		if _d.Name() == name {
			return true
		}
	}
	for _, f := range d.files {
		if f.Name() == name {
			return true
		}
	}
	return false
}

func (d *dir) AddFile(newFile File) error {
	if d.Exists(newFile.Name()) {
		msg := fmt.Sprintf("'%s': already exists", newFile.Name())
		return errors.New(msg)
	}
	d.files = append(d.files, newFile)
	return nil
}

func (d *dir) AddDir(newDir Dir) error {
	if d.Exists(newDir.Name()) {
		msg := fmt.Sprintf("'%s': already exists", newDir.Name())
		return errors.New(msg)
	}
	d.dirs = append(d.dirs, newDir)
	return nil
}

func (d *dir) Remove(name string) error {
	for i, f := range d.files {
		if f.Name() == name {
			bp := len(d.files) - 1
			d.files[i] = d.files[bp]
			d.files = d.files[:bp]
			return nil
		}
	}
	for i, _d := range d.dirs {
		if _d.Name() == name {
			bp := len(d.dirs) - 1
			d.dirs[i] = d.dirs[bp]
			d.dirs = d.dirs[:bp]
			return nil
		}
	}

	msg := fmt.Sprintf("'%s': no such file or directory", name)
	return errors.New(msg)
}

func (d *dir) ChangeParent(to Dir) {
	d.parent = to
}

func (d *dir) Rename(name string) error {
	if name == d.Name() {
		return nil
	}
	if len(name) == 0 {
		return errors.New("name cannot be empty")
	}
	if strings.ContainsRune(name, '/') {
		return errors.New("name cannot contain '/'")
	}
	if (d.parent.Exists(name)) {
		msg := fmt.Sprintf("'%s': already exists", name)
		return errors.New(msg)
	}
	d.name = name
	return nil
}

func (d *dir) List() string {
	result := ""
	for _, _d := range d.dirs {
		result = result + "[d] " + _d.Name() + string("\n")
	}
	for _, f := range d.files {
		result = result + "[f] " + f.Name() + string("\n")
	}
	return result
}

func (d *dir) IsRoot() bool {
	return false
}

func (d *dir) Parent() (Dir, error) {
	return d.parent, nil
}

func (d *dir) FindDir(name string) (Dir, error) {
	if name == ".." {
		return d.Parent()
	}
	if name == "." {
		return d, nil
	}
	for _, _d := range d.dirs {
		if _d.Name() == name {
			return _d, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("'%s': does not exist", name))
}

func (d *dir) FindFile(name string) (File, error) {
	for _, f := range d.files {
		if f.Name() == name {
			return f, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("'%s': does not exist", name))
}

