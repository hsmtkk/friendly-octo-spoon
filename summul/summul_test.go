package summul_test

import (
	"testing"

	"github.com/hsmtkk/friendly-octo-spoon/summul"
	"github.com/stretchr/testify/assert"
)

func TestSumMultiples(t *testing.T) {
	assert.Equal(t, 78, summul.SumMultiples(20))
}
