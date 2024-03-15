package options

import "github.com/aws/aws-sdk-go/aws"

var DefaultHttpSupport = false

type CommandOptions struct {
	HttpSupport *bool
}

func Command() *CommandOptions {
	return &CommandOptions{
		HttpSupport: &DefaultHttpSupport,
	}
}

func (opts *CommandOptions) WithHttpSupport() *CommandOptions {
	opts.HttpSupport = aws.Bool(true)
	return opts
}

// MergeCommandOptions combines the given *CommandOptions into one *CommandOptions in a last one wins fashion.
func MergeCommandOptions(opts ...*CommandOptions) *CommandOptions {
	uOpts := Command()
	for _, uo := range opts {
		if uo == nil {
			continue
		}
		if uo.HttpSupport != nil {
			uOpts.HttpSupport = uo.HttpSupport
		}
	}

	return uOpts
}
