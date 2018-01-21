package builder

import (
	amazonEbsBuilder "github.com/hashicorp/packer/builder/amazon/ebs"
	"github.com/hashicorp/packer/template/interpolate"
)

type StoppedInstancePoolBuilderConfig struct {
	amazonEbsBuilder.Config `mapstructure:",squash"`
	InstancePoolMinSize     int64 `mapstructure:"instance_pool_min_size"`

	ctx interpolate.Context
}

func (s *StoppedInstancePoolBuilderConfig) Prepare(ctx *interpolate.Context) []error {
	return make([]error, 0)
}
