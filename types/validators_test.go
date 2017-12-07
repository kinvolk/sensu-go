package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateHandlerType(t *testing.T) {
	assert.Error(t, validateHandlerType(""))
	assert.Error(t, validateHandlerType("foo"))
	assert.NoError(t, validateHandlerType("pipe"))
	assert.NoError(t, validateHandlerType("tcp"))
	assert.NoError(t, validateHandlerType("udp"))
	assert.NoError(t, validateHandlerType("transport"))
	assert.NoError(t, validateHandlerType("set"))
}

func TestValidateName(t *testing.T) {
	assert.Error(t, ValidateName(""))
	assert.Error(t, ValidateName("foo bar"))
	assert.Error(t, ValidateName("foo@bar"))
	assert.NoError(t, ValidateName("foo-bar"))
}

func TestValidateNameStrict(t *testing.T) {
	assert.Error(t, ValidateNameStrict(""))
	assert.Error(t, ValidateNameStrict("foo bar"))
	assert.Error(t, ValidateNameStrict("foo@bar"))
	assert.Error(t, ValidateNameStrict("FOO-bar"))
	assert.NoError(t, ValidateNameStrict("foo-bar_2"))
}

func TestValidateSubscriptionName(t *testing.T) {
	assert.Error(t, ValidateSubscriptionName(""))
	assert.Error(t, ValidateSubscriptionName("foo bar"))
	assert.Error(t, ValidateSubscriptionName("foo@bar"))
	assert.Error(t, ValidateSubscriptionName("entity:foo:bar"))
	assert.NoError(t, ValidateSubscriptionName("entity:foo"))
	assert.NoError(t, ValidateSubscriptionName("foo-bar_2"))
}
