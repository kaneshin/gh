package gh

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Client(t *testing.T) {
	assert := assert.New(t)

	c1 := NewClient("token1", "owner1", "repo1")
	c2 := NewClient("token2", "owner2", "repo2")
	c3 := NewClient("token1", "owner3", "repo3")

	assert.NotEqual(c1.client, c2.client)
	assert.Equal(c1.client, c3.client)
	assert.NotEqual(c2.client, c3.client)
	assert.Equal([]string{"owner1", "repo1"}, []string{c1.owner, c1.repo})
	assert.Equal([]string{"owner2", "repo2"}, []string{c2.owner, c2.repo})
	assert.Equal([]string{"owner3", "repo3"}, []string{c3.owner, c3.repo})
}
