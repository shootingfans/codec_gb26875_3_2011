package constant_test

import (
	"strconv"
	"testing"

	"github.com/shootingfans/codec_gb26875_3_2011/constant"

	"github.com/stretchr/testify/assert"
)

func TestAppType_String(t *testing.T) {
	for k, v := range constant.AppTypeNames {
		assert.Equal(t, k.String(), "["+strconv.Itoa(int(k))+"]"+v)
	}
	assert.Equal(t, constant.AppType(200).String(), "[200]")
}
