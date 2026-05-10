package platform

type Arch int

const (
	X64 Arch = iota
	ARM64
	X32
)

func (a Arch) String() string {
	switch a {
	case ARM64:
		return "aarch64"
	case X32:
		return "x32"
	default:
		return "x64"
	}
}
