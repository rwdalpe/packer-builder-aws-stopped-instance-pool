package builder

import (
	"fmt"
	awscommon "github.com/hashicorp/packer/builder/amazon/common"
	"github.com/hashicorp/packer/common"
	"github.com/hashicorp/packer/helper/config"
	"github.com/hashicorp/packer/packer"
	"github.com/hashicorp/packer/template/interpolate"
	"github.com/mitchellh/multistep"
	"log"
)

type StoppedInstancePoolBuilder struct {
	config StoppedInstancePoolBuilderConfig
	runner multistep.Runner
}

func (b *StoppedInstancePoolBuilder) Prepare(raws ...interface{}) ([]string, error) {
	b.config.ctx.Funcs = awscommon.TemplateFuncs
	err := config.Decode(&b.config, &config.DecodeOpts{
		Interpolate:        true,
		InterpolateContext: &b.config.ctx,
		InterpolateFilter: &interpolate.RenderFilter{
			Exclude: []string{
				"ami_description",
				"run_tags",
				"run_volume_tags",
				"snapshot_tags",
				"tags",
			},
		},
	}, raws...)
	if err != nil {
		return nil, err
	}

	if b.config.Config.PackerConfig.PackerForce {
		b.config.Config.AMIForceDeregister = true
	}

	// Accumulate any errors
	var errs *packer.MultiError
	errs = packer.MultiErrorAppend(errs, b.config.Prepare(&b.config.ctx)...)

	if b.config.Config.IsSpotInstance() && (b.config.Config.AMIENASupport || b.config.Config.AMISriovNetSupport) {
		errs = packer.MultiErrorAppend(errs,
			fmt.Errorf("Spot instances do not support modification, which is required "+
				"when either `ena_support` or `sriov_support` are set. Please ensure "+
				"you use an AMI that already has either SR-IOV or ENA enabled."))
	}

	if errs != nil && len(errs.Errors) > 0 {
		return nil, errs
	}

	log.Println(common.ScrubConfig(b.config, b.config.Config.AccessKey, b.config.Config.SecretKey, b.config.Config.Token))
	return nil, nil
}

func (b *StoppedInstancePoolBuilder) Run(ui packer.Ui, hook packer.Hook, cache packer.Cache) (packer.Artifact, error) {
	return *new(packer.Artifact), nil
}

func (b *StoppedInstancePoolBuilder) Cancel() {
	if b.runner != nil {
		log.Println("Cancelling the step runner...")
		b.runner.Cancel()
	}
}
