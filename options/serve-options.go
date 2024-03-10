package options

type ServeOptions struct {
	ServePort *int
}

func Serve() *ServeOptions {
	return &ServeOptions{}
}

func (leo *ServeOptions) SetHttpPort(p int) *ServeOptions {
	leo.ServePort = &p
	return leo
}

func MergeServeOptions(opts ...*ServeOptions) *ServeOptions {
	uOpts := Serve()
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
