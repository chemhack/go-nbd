//go:build linux && !cgo && (amd64 || arm)

package ioctl

const (
	TRANSMISSION_IOCTL_DISCONNECT = 43784
	TRANSMISSION_IOCTL_CLEAR_SOCK = 43780
	TRANSMISSION_IOCTL_CLEAR_QUE  = 43781
)
