package block

import (
	"fmt"
	"testing"

	"github.com/ThotaGopichandThota/gn-core3/core/check"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewEmptyMetaBlockCreator(t *testing.T) {
	t.Parallel()

	creator := NewEmptyMetaBlockCreator()
	require.False(t, check.IfNil(creator))
}

func TestEmptyMetaBlockCreator_CreateNewHeader(t *testing.T) {
	t.Parallel()

	creator := NewEmptyMetaBlockCreator()
	header := creator.CreateNewHeader()
	require.False(t, check.IfNil(header))
	assert.Equal(t, "*block.MetaBlock", fmt.Sprintf("%T", header))
}
