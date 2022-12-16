package file

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
)

// GetSize get the file size
func GetSize(f multipart.File) (int, error) {
	content, err := ioutil.ReadAll(f)

	return len(content), err
}

// GetExt get the file ext
func GetExt(fileName string) string {
	return path.Ext(fileName)
}

// CheckNotExist check if the file exists
func CheckNotExist(src string) bool {
	_, err := os.Stat(src)

	return os.IsNotExist(err)
}

// CheckPermission check if the file has permission
func CheckPermission(src string) bool {
	_, err := os.Stat(src)

	return os.IsPermission(err)
}

// IsNotExistMkDir create a directory if it does not exist
func IsNotExistMkDir(src string) error {
	if notExist := CheckNotExist(src); notExist == true {
		if err := MkDir(src); err != nil {
			return err
		}
	}

	return nil
}

// MkDir create a directory
func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// Open a file according to a specific mode
func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// MustOpen maximize trying to open the file
func MustOpen(fileName, filePath string) (*os.File, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("os.Getwd err: %v", err)
	}

	src := dir + "/" + filePath
	perm := CheckPermission(src)
	if perm == true {
		return nil, fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	err = IsNotExistMkDir(src)
	if err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", src, err)
	}

	f, err := Open(src+fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, fmt.Errorf("Fail to OpenFile :%v", err)
	}

	return f, nil
}



func Copy(from, to string) error {
	var err error

	f, err := os.Stat(from)
	if err != nil {
		return err
	}

	fn := func(fromFile string) error {
		//复制文件的路径
		rel, err := filepath.Rel(from, fromFile)
		if err != nil {
			return err
		}
		toFile := filepath.Join(to, rel)

		//创建复制文件目录
		if err = os.MkdirAll(filepath.Dir(toFile), 0777); err != nil {
			return err
		}

		//读取源文件
		file, err := os.Open(fromFile)
		if err != nil {
			return err
		}

		defer file.Close()
		bufReader := bufio.NewReader(file)
		// 创建复制文件用于保存
		out, err := os.Create(toFile)
		if err != nil {
			return err
		}

		defer out.Close()
		// 然后将文件流和文件流对接起来
		_, err = io.Copy(out, bufReader)
		return err
	}

	//转绝对路径
	pwd, _ := os.Getwd()
	if !filepath.IsAbs(from) {
		from = filepath.Join(pwd, from)
	}
	if !filepath.IsAbs(to) {
		to = filepath.Join(pwd, to)
	}

	//复制
	if f.IsDir() {
		return filepath.WalkDir(from, func(path string, d fs.DirEntry, err error) error {
			if !d.IsDir() {
				return fn(path)
			} else {
				if err = os.MkdirAll(path, 0777); err != nil {
					return err
				}
			}
			return err
		})
	} else {
		return fn(from)
	}
}
