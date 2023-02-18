package dir

import (
	"fmt"
	"os"
	"strings"
	"syscall"
	"time"
)

type Vesrionable interface {
	Commit()
	Fetch()
}

type directory struct {
	path     string
	scanTime time.Time
	files    []file
}

func NewDirectory(path string) directory {
	dir, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	return directory{
		path:     path,
		scanTime: time.Now(),
		files:    make([]file, 0, len(dir)),
	}
}

func (d *directory) Compare(target *directory) {
	sourcefiles := d.Scan()

	targetpath := target.path
	targetfiles := target.Scan()

	// Jika file ada di source tapi tidak ada di target berikan keterangan NEW
	for _, sourcefile := range sourcefiles {
		_, err := os.Stat(targetpath + sourcefile.Name)

		if err != nil && sourcefile.CreatedAt.Equal(sourcefile.ModifiedAt) {
			fmt.Println(d.path + sourcefile.Name + " NEW")
		}

		if sourcefile.ModifiedAt.After(sourcefile.CreatedAt) {
			fmt.Println(d.path + sourcefile.Name + " MODIFIED")
		}
	}

	// Jika file tidak ada di source tapi ada di target berikan keterangan DELETED
	for _, targetfile := range targetfiles {
		_, err := os.Stat(d.path + targetfile.Name)

		if err != nil && !targetfile.CreatedAt.Equal(targetfile.ModifiedAt) {
			fmt.Println(d.path + targetfile.Name + " DELETED")
		}

		if targetfile.ModifiedAt.After(targetfile.CreatedAt) {
			fmt.Println(targetpath + targetfile.Name + " MODIFIED")
		}
	}
}

func (d *directory) Scan() []file {
	d.readDir(d.path)

	return d.files
}

func (d *directory) readDir(path string) {
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

		var st syscall.Stat_t
		if err := syscall.Stat(path+"/"+file.Name(), &st); err != nil {
			panic(err)
		}

		createdAt := time.Unix(st.Ctim.Sec, 0)

		f, _ := file.Info()
		modifiedAt := time.Unix(f.ModTime().Unix(), 0)
		d.files = append(d.files, NewFile(dir, createdAt, modifiedAt))
	}
}
