package builder

import (
	"github.com/hashicorp/packer/packer"
)

type StoppedInstancePoolBuilder struct {
}

func (s *StoppedInstancePoolBuilder) Prepare(...interface{}) ([]string, error) {
	return []string{}, nil
}

func (s *StoppedInstancePoolBuilder) Run(ui packer.Ui, hook packer.Hook, cache packer.Cache) (packer.Artifact, error) {
	return *new(packer.Artifact), nil
}

func (s *StoppedInstancePoolBuilder) Cancel() {

}
