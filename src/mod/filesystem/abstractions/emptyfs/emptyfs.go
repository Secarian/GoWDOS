package emptyfs

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"time"

	wdosfs "imuslab.com/arozos/mod/filesystem/arozfs"
)

/*
	filesystemAbstraction.go

	This file contains all the abstraction funtion of a local file system.

*/

type EmptyFileSystemAbstraction struct {
}

func NewEmptyFileSystemAbstraction() EmptyFileSystemAbstraction {
	return EmptyFileSystemAbstraction{}
}

func (l EmptyFileSystemAbstraction) Chmod(filename string, mode os.FileMode) error {
	return wdosfs.ErrNullOperation
}
func (l EmptyFileSystemAbstraction) Chown(filename string, uid int, gid int) error {
	return wdosfs.ErrNullOperation
}
func (l EmptyFileSystemAbstraction) Chtimes(filename string, atime time.Time, mtime time.Time) error {
	return wdosfs.ErrNullOperation
}
func (l EmptyFileSystemAbstraction) Create(filename string) (wdosfs.File, error) {
	return nil, wdosfs.ErrNullOperation
}
func (l EmptyFileSystemAbstraction) Mkdir(filename string, mode os.FileMode) error {
	return wdosfs.ErrNullOperation
}
func (l EmptyFileSystemAbstraction) MkdirAll(filename string, mode os.FileMode) error {
	return wdosfs.ErrNullOperation
}
func (l EmptyFileSystemAbstraction) Name() string {
	return ""
}
func (l EmptyFileSystemAbstraction) Open(filename string) (wdosfs.File, error) {
	return nil, wdosfs.ErrNullOperation
}
func (l EmptyFileSystemAbstraction) OpenFile(filename string, flag int, perm os.FileMode) (wdosfs.File, error) {
	return nil, wdosfs.ErrNullOperation
}
func (l EmptyFileSystemAbstraction) Remove(filename string) error {
	return wdosfs.ErrNullOperation
}
func (l EmptyFileSystemAbstraction) RemoveAll(path string) error {
	return wdosfs.ErrNullOperation
}
func (l EmptyFileSystemAbstraction) Rename(oldname, newname string) error {
	return wdosfs.ErrNullOperation
}
func (l EmptyFileSystemAbstraction) Stat(filename string) (os.FileInfo, error) {
	return nil, wdosfs.ErrNullOperation
}
func (l EmptyFileSystemAbstraction) Close() error {
	return nil
}

/*
	Abstraction Utilities
*/

func (l EmptyFileSystemAbstraction) VirtualPathToRealPath(subpath string, username string) (string, error) {
	return "", wdosfs.ErrVpathResolveFailed
}

func (l EmptyFileSystemAbstraction) RealPathToVirtualPath(fullpath string, username string) (string, error) {
	return "", wdosfs.ErrRpathResolveFailed
}

func (l EmptyFileSystemAbstraction) FileExists(realpath string) bool {
	return false
}

func (l EmptyFileSystemAbstraction) IsDir(realpath string) bool {
	return false
}

func (l EmptyFileSystemAbstraction) Glob(realpathWildcard string) ([]string, error) {
	return []string{}, wdosfs.ErrNullOperation
}

func (l EmptyFileSystemAbstraction) GetFileSize(realpath string) int64 {
	return 0
}

func (l EmptyFileSystemAbstraction) GetModTime(realpath string) (int64, error) {
	return 0, wdosfs.ErrOperationNotSupported
}

func (l EmptyFileSystemAbstraction) WriteFile(filename string, content []byte, mode os.FileMode) error {
	return wdosfs.ErrNullOperation
}
func (l EmptyFileSystemAbstraction) ReadFile(filename string) ([]byte, error) {
	return []byte(""), wdosfs.ErrOperationNotSupported
}
func (l EmptyFileSystemAbstraction) ReadDir(filename string) ([]fs.DirEntry, error) {
	return []fs.DirEntry{}, wdosfs.ErrOperationNotSupported
}
func (l EmptyFileSystemAbstraction) WriteStream(filename string, stream io.Reader, mode os.FileMode) error {
	return wdosfs.ErrNullOperation
}
func (l EmptyFileSystemAbstraction) ReadStream(filename string) (io.ReadCloser, error) {
	return nil, wdosfs.ErrOperationNotSupported
}

func (l EmptyFileSystemAbstraction) Walk(root string, walkFn filepath.WalkFunc) error {
	return wdosfs.ErrOperationNotSupported
}

func (l EmptyFileSystemAbstraction) Heartbeat() error {
	return nil
}
