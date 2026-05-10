package platform

type OS int

const (
	Linux OS = iota
	Macos
	Windows
)

func (o OS) String() string {
	switch o {
	case Windows:
		return "windows"
	case Macos:
		return "mac"
	default:
		return "linux"
	}
}
