package platform

type Platform struct {
	OS   OS
	Arch Arch
}

func Detect() Platform {
	return Platform{
		OS:   DetectOS(),
		Arch: DetectArch(),
	}
}
