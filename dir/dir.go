package dir

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Vesrionable interface {
	Commit()
	Fetch()
}

type Directory struct {
	path     string
	scanTime time.Time
	files    []file
}

func NewDirectory(path string) Directory {
	dir, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	return Directory{
		path:     path,
		scanTime: time.Now(),
		files:    make([]file, 0, len(dir)),
	}
}

// intial read: baca directory kemudian simpan dalam bentuk file
// jadikan struct menjadi json
// simpan data dalam bentuk json ke file
func (d *Directory) Commit() {
	d.readDir(d.path)
	for i := range d.files {
		fmt.Println(d.files[i].name, d.files[i].CreatedAt, d.files[i].ModifiedAt)
	}

	fmt.Println("===========")

	d.updateDir(d.path, 0)
	for i := range d.files {
		fmt.Println(d.files[i].name, d.files[i].CreatedAt, d.files[i].ModifiedAt)
	}
}

func (d *Directory) readDir(path string) {
	files, err := os.ReadDir(path)
	if err != nil {
		return
	}

	for _, file := range files {
		if string(file.Name()[0]) == "." {
			continue
		}

		if file.IsDir() {
			d.readDir(path + "/" + file.Name())
			continue
		}

		dir := strings.Join(strings.Split(path, "/")[6:], "/")
		dir = dir + "/" + file.Name()
		dir = "/" + strings.Join(strings.Split(dir, "/")[1:], "/")

		f, _ := file.Info()
		time := f.ModTime()
		d.files = append(d.files, NewFile(dir, time))
	}
}

func (d *Directory) updateDir(path string, i int) {
	files, err := os.ReadDir(path)
	if err != nil {
		return
	}

	if i == len(d.files) {
		return
	}

	for _, file := range files {
		if string(file.Name()[0]) == "." {
			continue
		}

		if file.IsDir() {
			d.updateDir(path+"/"+file.Name(), i+1)
			continue
		}

		if _, err := os.Stat(path + d.files[i].name); err != nil {
			d.files[i].DeletedAt = time.Now()
		}

		f, err := file.Info()
		if err != nil {
			panic(err)
		}

		if f.ModTime().Nanosecond() != d.files[i].CreatedAt.Nanosecond() {
			d.files[i].ModifiedAt = time.Now()
		}
	}
}
