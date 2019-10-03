// Package goverlayfs is a very simple library for creating an overlay
// filesystem mount.
package goverlayfs

import (
	"github.com/pkg/errors"
	"golang.org/x/sys/unix"
)

type OverlayMount struct {
	upperPath, lowerPath, mergedPath, workdirPath string
}

func New(upperPath, lowerPath, mergedPath, workdirPath string) *OverlayMount {
	return &OverlayMount{upperPath, lowerPath, mergedPath, workdirPath}
}

// Init makes an overlayfs using the specified paths for upper, lower, merged, and workdir paths
// XXX: requires root/cap_sys_admin
func (o *OverlayMount) Init() (err error) {

	if o == nil {
		return errors.New("nil OverlayMount type")
	}

	args := "lowerdir=" + o.lowerPath + ",upperdir=" + o.upperPath + ",workdir=" + o.workdirPath
	err = unix.Mount(
		o.lowerPath,
		o.mergedPath,
		"overlay",
		0,
		args,
	)
	if err != nil {
		return errors.Wrap(err, "could not create overlay mount")
	}

	return nil
}

// Remove unmounts the overlay mount's directory
func (o *OverlayMount) Remove() error {
	return unix.Unmount(o.mergedPath, 0)
}
