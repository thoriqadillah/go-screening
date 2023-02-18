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
		_, err := os.Stat(targetpath + sourcefile.Path + "/" + sourcefile.Name)

		if err != nil && sourcefile.CreatedAt.Equal(sourcefile.ModifiedAt) {
			fmt.Println(d.path + sourcefile.Path + "/" + sourcefile.Name + " NEW")
		}

		if sourcefile.ModifiedAt.After(sourcefile.CreatedAt) {
			fmt.Println(d.path + sourcefile.Path + "/" + sourcefile.Name + " MODIFIED")
		}
	}

	// Jika file tidak ada di source tapi ada di target berikan keterangan DELETED
	for _, targetfile := range targetfiles {
		// fmt.Println(targetfile.Path, "index", i)
		_, err := os.Stat(d.path + targetfile.Path + "/" + targetfile.Name)

		if err != nil && !targetfile.CreatedAt.Equal(targetfile.ModifiedAt) {
			fmt.Println(d.path + targetfile.Path + "/" + targetfile.Name + " DELETED")
		}

		if targetfile.ModifiedAt.After(targetfile.CreatedAt) {
			fmt.Println(targetpath + targetfile.Path + "/" + targetfile.Name + " MODIFIED")
		}
	}
}

func (d *directory) Scan() []file {
	split := strings.Split(d.path, "/")
	d.readDir(d.path, split, split[len(split)-1])

	return d.files
}

func (d *directory) readDir(path string, dir []string, root string) {
	files, err := os.ReadDir(path)
	if err != nil {
		return
	}

	for _, file := range files {
		if string(file.Name()[0]) == "." {
			continue
		}

		directory := root + "/" + dir[len(dir)-1]

		if file.IsDir() {
			newpath := path + "/" + file.Name()
			d.readDir(newpath, strings.Split(newpath, "/"), directory)
			continue
		}

		var st syscall.Stat_t
		if err := syscall.Stat(path+"/"+file.Name(), &st); err != nil {
			panic(err)
		}

		path := strings.Join(strings.Split(directory, "/")[2:], "/")
		createdAt := time.Unix(st.Ctim.Sec, 0)

		f, _ := file.Info()
		modifiedAt := time.Unix(f.ModTime().Unix(), 0)
		d.files = append(d.files, NewFile(file.Name(), "/"+path, createdAt, modifiedAt))
	}
}
