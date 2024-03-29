//go:build !windows

package ptest

import (
	"io/fs"
	"os"
	"syscall"
)

func checkUid(uid int, stat fs.FileInfo) bool {
	source := stat.Sys()
	if source == nil {
		return false
	}
	info := source.(*syscall.Stat_t)
	return int(info.Uid) == os.Getuid()
}

func checkGid(gid int, stat fs.FileInfo) bool {
	source := stat.Sys()
	if source == nil {
		return false
	}
	info := source.(*syscall.Stat_t)
	return int(info.Gid) == os.Getgid()
}
