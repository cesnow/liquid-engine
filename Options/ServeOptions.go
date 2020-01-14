package Options

var DefaultServePort = 8080

type ServeOptions struct {
	ServePort *int
}

func LoadServeEnv() *ServeOptions {
	return &ServeOptions{
		ServePort: &DefaultServePort,
	}
}

func (leo *ServeOptions) SetHttpPort(p int) *ServeOptions {
	leo.ServePort = &p
	return leo
}

func MergeServeOptions(opts ...*ServeOptions) *ServeOptions {
	uOpts := LoadServeEnv()
	for _, uo := range opts {
		if uo == nil {
			continue
		}
		if uo.ServePort != nil {
			uOpts.ServePort = uo.ServePort
		}
	}

	return uOpts
}
