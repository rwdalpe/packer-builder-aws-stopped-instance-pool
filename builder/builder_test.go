package builder

import (
	"github.com/hashicorp/packer/packer"
	"os"
	"testing"
)

func testConfig() map[string]interface{} {
	return map[string]interface{}{
		"access_key":             "foo",
		"secret_key":             "bar",
		"source_ami":             "foo",
		"instance_type":          "foo",
		"region":                 "us-east-1",
		"ssh_username":           "root",
		"ami_name":               "foo",
		"instance_pool_min_size": 1,
		"ssh_keypair_name":       "kp",
		"ssh_private_key_file":   ".gotest_stoppedinstancepool_pk",
	}
}

func TestMain(m *testing.M) {
	f, err := os.OpenFile(".gotest_stoppedinstancepool_pk", os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		panic(err)
	}

	// Found this key from a random website; hope it's not important
	_, err = f.WriteString("-----BEGIN RSA PRIVATE KEY-----\nMIIBOwIBAAJBAJv8ZpB5hEK7qxP9K3v43hUS5fGT4waKe7ix4Z4mu5UBv+cw7WSF\nAt0Vaag0sAbsPzU8Hhsrj/qPABvfB8asUwcCAwEAAQJAG0r3ezH35WFG1tGGaUOr\nQA61cyaII53ZdgCR1IU8bx7AUevmkFtBf+aqMWusWVOWJvGu2r5VpHVAIl8nF6DS\nkQIhAMjEJ3zVYa2/Mo4ey+iU9J9Vd+WoyXDQD4EEtwmyG1PpAiEAxuZlvhDIbbce\n7o5BvOhnCZ2N7kYb1ZC57g3F+cbJyW8CIQCbsDGHBto2qJyFxbAO7uQ8Y0UVHa0J\nBO/g900SAcJbcQIgRtEljIShOB8pDjrsQPxmI1BLhnjD1EhRSubwhDw5AFUCIQCN\nA24pDtdOHydwtSB5+zFqFLfmVZplQM/g5kb4so70Yw==\n-----END RSA PRIVATE KEY-----")

	if err != nil {
		panic(err)
	}

	runResult := m.Run()

	err = f.Close()

	if err != nil {
		panic(err)
	}

	err = os.Remove(f.Name())

	if err != nil {
		panic(err)
	}

	os.Exit(runResult)
}

func TestBuilder_ImplementsBuilder(t *testing.T) {
	var raw interface{}
	raw = &StoppedInstancePoolBuilder{}
	if _, ok := raw.(packer.Builder); !ok {
		t.Fatalf("Builder should be a builder")
	}
}

func TestBuilder_Prepare_BadType(t *testing.T) {
	b := &StoppedInstancePoolBuilder{}
	c := map[string]interface{}{
		"access_key": []string{},
	}

	warnings, err := b.Prepare(c)
	if len(warnings) > 0 {
		t.Fatalf("bad: %#v", warnings)
	}
	if err == nil {
		t.Fatalf("prepare should fail")
	}
}

func TestBuilderPrepare_AMIName(t *testing.T) {
	var b StoppedInstancePoolBuilder
	config := testConfig()

	// Test good
	config["ami_name"] = "foo"
	warnings, err := b.Prepare(config)
	if len(warnings) > 0 {
		t.Fatalf("bad: %#v", warnings)
	}
	if err != nil {
		t.Fatalf("should not have error: %s", err)
	}

	// Test bad
	config["ami_name"] = "foo {{"
	b = StoppedInstancePoolBuilder{}
	warnings, err = b.Prepare(config)
	if len(warnings) > 0 {
		t.Fatalf("bad: %#v", warnings)
	}
	if err == nil {
		t.Fatal("should have error")
	}

	// Test bad
	delete(config, "ami_name")
	b = StoppedInstancePoolBuilder{}
	warnings, err = b.Prepare(config)
	if len(warnings) > 0 {
		t.Fatalf("bad: %#v", warnings)
	}
	if err == nil {
		t.Fatal("should have error")
	}
}

func TestBuilderPrepare_InvalidKey(t *testing.T) {
	var b StoppedInstancePoolBuilder
	config := testConfig()

	// Add a random key
	config["i_should_not_be_valid"] = true
	warnings, err := b.Prepare(config)
	if len(warnings) > 0 {
		t.Fatalf("bad: %#v", warnings)
	}
	if err == nil {
		t.Fatal("should have error")
	}
}

func TestBuilderPrepare_InvalidShutdownBehavior(t *testing.T) {
	var b StoppedInstancePoolBuilder
	config := testConfig()

	// Test good
	config["shutdown_behavior"] = "terminate"
	warnings, err := b.Prepare(config)
	if len(warnings) > 0 {
		t.Fatalf("bad: %#v", warnings)
	}
	if err != nil {
		t.Fatalf("should not have error: %s", err)
	}

	// Test good
	config["shutdown_behavior"] = "stop"
	warnings, err = b.Prepare(config)
	if len(warnings) > 0 {
		t.Fatalf("bad: %#v", warnings)
	}
	if err != nil {
		t.Fatalf("should not have error: %s", err)
	}

	// Test bad
	config["shutdown_behavior"] = "foobar"
	warnings, err = b.Prepare(config)
	if len(warnings) > 0 {
		t.Fatalf("bad: %#v", warnings)
	}
	if err == nil {
		t.Fatal("should have error")
	}
}

func TestBuilderPrepare_InstancePoolMinSize(t *testing.T) {
	var b StoppedInstancePoolBuilder
	config := testConfig()

	// Test good
	config["instance_pool_min_size"] = 1
	warnings, err := b.Prepare(config)
	if len(warnings) > 0 {
		t.Fatalf("bad: %#v", warnings)
	}
	if err != nil {
		t.Fatalf("should not have error: %s", err)
	}

	// Test bad
	config["instance_pool_min_size"] = 0
	b = StoppedInstancePoolBuilder{}
	warnings, err = b.Prepare(config)
	if len(warnings) > 0 {
		t.Fatalf("bad: %#v", warnings)
	}
	if err == nil {
		t.Fatal("should have error")
	}

	// Test bad
	config["instance_pool_min_size"] = -1
	b = StoppedInstancePoolBuilder{}
	warnings, err = b.Prepare(config)
	if len(warnings) > 0 {
		t.Fatalf("bad: %#v", warnings)
	}
	if err == nil {
		t.Fatal("should have error")
	}

	// Test bad
	delete(config, "instance_pool_min_size")
	b = StoppedInstancePoolBuilder{}
	warnings, err = b.Prepare(config)
	if len(warnings) > 0 {
		t.Fatalf("bad: %#v", warnings)
	}
	if err == nil {
		t.Fatal("should have error")
	}
}

func TestBuilderPrepare_SSHKeys(t *testing.T) {
	var b StoppedInstancePoolBuilder
	config := testConfig()

	// Test good
	config["ssh_keypair_name"] = "kp"
	config["ssh_private_key_file"] = ".gotest_stoppedinstancepool_pk"
	warnings, err := b.Prepare(config)
	if len(warnings) > 0 {
		t.Fatalf("bad: %#v", warnings)
	}
	if err != nil {
		t.Fatalf("should not have error: %s", err)
	}

	// Test bad
	config["ssh_keypair_name"] = ""
	config["ssh_private_key_file"] = ".gotest_stoppedinstancepool_pk"
	b = StoppedInstancePoolBuilder{}
	warnings, err = b.Prepare(config)
	if len(warnings) > 0 {
		t.Fatalf("bad: %#v", warnings)
	}
	if err == nil {
		t.Fatal("should have error")
	}

	// Test bad
	config["ssh_keypair_name"] = "kp"
	config["ssh_private_key_file"] = ""
	b = StoppedInstancePoolBuilder{}
	warnings, err = b.Prepare(config)
	if len(warnings) > 0 {
		t.Fatalf("bad: %#v", warnings)
	}
	if err == nil {
		t.Fatal("should have error")
	}
}
