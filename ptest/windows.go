//go:build windows

package ptest

import "io/fs"

func checkUid(uid int, stat fs.FileInfo) bool {
	return false
}

func checkGid(gid int, stat fs.FileInfo) bool {
	return false
}
