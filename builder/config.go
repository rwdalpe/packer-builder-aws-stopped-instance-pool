package builder

import (
	"fmt"
	amazonEbsBuilder "github.com/hashicorp/packer/builder/amazon/ebs"
	"github.com/hashicorp/packer/packer"
	"github.com/hashicorp/packer/template/interpolate"
)

type StoppedInstancePoolBuilderConfig struct {
	amazonEbsBuilder.Config `mapstructure:",squash"`
	InstancePoolMinSize     int64 `mapstructure:"instance_pool_min_size"`

	ctx interpolate.Context
}

func (c *StoppedInstancePoolBuilderConfig) Prepare(ctx *interpolate.Context) []error {
	var errs *packer.MultiError

	if c.InstancePoolMinSize <= 0 {
		errs = packer.MultiErrorAppend(errs, fmt.Errorf("Stopped instance pool min size must be greater than 0"))
	}

	errs = packer.MultiErrorAppend(errs, c.Config.AccessConfig.Prepare(&c.ctx)...)
	errs = packer.MultiErrorAppend(errs,
		c.Config.AMIConfig.Prepare(&c.Config.AccessConfig, &c.ctx)...)
	errs = packer.MultiErrorAppend(errs, c.Config.BlockDevices.Prepare(&c.ctx)...)
	errs = packer.MultiErrorAppend(errs, c.Config.RunConfig.Prepare(&c.ctx)...)

	return errs.Errors
}
