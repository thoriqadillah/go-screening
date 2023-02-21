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
	savedSource := make([]file, len(d.files))
	splitSrc := strings.Split(d.path, "/")
	sourceFolder := splitSrc[len(splitSrc)-1]
	if err := load(sourceFolder+".tmp", &savedSource); err != nil {
		save(sourceFolder+".tmp", d.files)
	}

	targetpath := target.path
	targetfiles := target.Scan()

	// Jika file ada di source tapi tidak ada di target berikan keterangan NEW
	for i, sourcefile := range sourcefiles {
		_, err := os.Stat(targetpath + "/" + sourcefile.path + sourcefile.name)

		if err != nil && sourcefile.createdAt.Equal(sourcefile.modifiedAt) {
			fmt.Println(d.path + "/" + sourcefile.path + sourcefile.name + " NEW")
		}

		if !savedSource[i].modifiedAt.Equal(sourcefile.modifiedAt) {
			fmt.Println(d.path + "/" + sourcefile.path + sourcefile.name + " MODIFIED")
		}
	}

	savedTarget := make([]file, len(targetfiles))
	splitTgt := strings.Split(targetpath, "/")
	targetFolder := splitTgt[len(splitTgt)-1]
	if err := load(targetFolder+".tmp", &savedTarget); err != nil {
		save(targetFolder+".tmp", target.files)
	}
	// Jika file tidak ada di source tapi ada di target berikan keterangan DELETED
	for i, targetfile := range targetfiles {
		_, err := os.Stat(d.path + "/" + targetfile.path + targetfile.name)

		if err != nil && !targetfile.createdAt.Equal(targetfile.modifiedAt) {
			fmt.Println(d.path + "/" + targetfile.path + targetfile.name + " DELETED")
		}

		if !savedTarget[i].modifiedAt.Equal(targetfile.modifiedAt) {
			fmt.Println(targetpath + "/" + targetfile.path + targetfile.name + " MODIFIED")
		}
	}
}

func (d *directory) Scan() []file {
	split := strings.Split(d.path, "/")
	d.readDir(d.path, split, split[len(split)-1])

	//===UNTUK MODIFIED===
	//check apakah sudah ada file tmp nya belum
	//kalo belum ada, buat. Kalo sudah, biarin aja, jangan buat lagi
	//mulai modified file yang diinginkan
	//check size nya apakah sama atau tidak. jika berbeda, maka dia modified

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
		if path != "" {
			path = path + "/"
		}

		createdAt := time.Unix(st.Ctim.Sec, 0)

		f, _ := file.Info()
		modifiedAt := time.Unix(f.ModTime().Unix(), 0)
		d.files = append(d.files, NewFile(file.Name(), path, createdAt, modifiedAt))
	}
}
