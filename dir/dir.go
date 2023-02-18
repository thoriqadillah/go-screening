package dir

import (
	"os"
	"strings"
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

func (d *directory) Compare(dir *directory) {
	// ====UNTUK DELETED=====
	// read data pada source dan target dir kemudian simpan (initial)
	// mulai hapus file pada source, kemudian run program untuk mengetahui file mana yang DELETED pada source
	// load saved source dan bandingkan ulang dengan yang baru
	// jika ada yang nil pada scan file, maka update deletedAt pada elemen dari array of file pada source dan kemudian simpan ulang

	// ====UNTUK NEW=====
	// mulai tambah file pada source, kemudian run program untuk mengetahui file mana yang NEW pada source
	// load saved source dan bandingkan ulang dengan yang baru per index
	// jika pada nama dalam suatu elemen berbeda, print file tersebut, kemudian overwrite masing2 index mulai dari situ
	// simpan ulang ke persistent

	// ====UNTUK MODIFIED=====
	// mulai tambah ubah content pada source, kemudian run program untuk mengetahui file mana yang MODIFIED pada source
	// load saved source dan bandingkan ulang dengan yang baru per index
	// jika pada suatu elemen file yang disimpan modifiedAt nya berbeda dengan file baru yang discan, maka jadikan dia modified
	// simpan ulang ke persistent
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

		f, _ := file.Info()
		time := f.ModTime()
		d.files = append(d.files, NewFile(dir, time))
	}
}

func (d *directory) updateDir(path string, i int) {
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

		if _, err := os.Stat(path + d.files[i].Name); err != nil {
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
