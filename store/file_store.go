package store

import (
    _ "fmt"
    "crypto/md5"
    "io"
    "encoding/base64"
    "os"
    "launchpad.net/gommap"
    "io/ioutil"
    "strings"
    "path/filepath"
)

type mappedFile struct {
    File *os.File
    Map gommap.MMap
}

type FileStore struct {
    Dir string
    files map[string]mappedFile
}

func NewFileStore(dir string) (s *FileStore) {
    s = &FileStore { dir, make(map[string]mappedFile) }
    if s.exists() {
        s.loadKeys()
    }
    return
}

func (s *FileStore) Add(key, value string) (err error) {
    h := hash(key)
    if existing, ok := s.files[h]; ok {
        existing.Map.UnsafeUnmap()
    }
    file, err := osFile(s, h, true)
    if err != nil {
        return
    }
    if _, err := file.WriteString(value); err != nil {
        file.Sync()
    }
    if err != nil {
        return
    }
    if f, err := mapped(file); err == nil {
        s.files[h] = f
    }
    return
}

func (s *FileStore) Get(key string) (string, bool) {
    h := hash(key)
    f, ok := s.files[h]
    return string(f.Map), ok
}

func (s *FileStore) exists() bool {
    _, err := os.OpenFile(s.Dir, 644, os.ModeDir)
    if err != nil {
        if err = os.MkdirAll(s.Dir, os.ModePerm); err != nil {
            panic(err)
        }
    }
    return true
}

func (s *FileStore) loadKeys() {
    files, err := ioutil.ReadDir(s.Dir)
    if err != nil {
        panic(err)
    }

    for _, child := range files {
        h := child.Name()
        file, err := osFile(s, h, false)
        if err != nil {
            panic(err)
        }
        f, err := mapped(file)
        if err != nil {
            panic(err)
        }
        s.files[h] = f
    }
}

func hash(key string) string {
    h := md5.New()
    io.WriteString(h, key)
    return strings.Replace(base64.StdEncoding.EncodeToString(h.Sum(nil)), "/", "_", -1)
}

func osFile(s *FileStore, hash string, create bool) (f *os.File, err error) {
    path := s.Dir + string(filepath.Separator) + hash
    if (create) {
        f, err = os.Create(path)
    } else {
        f, err = os.Open(path)
    }
    return
}

func mapped(file *os.File) (f mappedFile, err error) {
    mm, err := gommap.Map(file.Fd(), gommap.PROT_READ | gommap.PROT_WRITE, gommap.MAP_PRIVATE)
    f = mappedFile { file, mm }
    return
}
