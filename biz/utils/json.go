package utils

import (
	"github.com/bytedance/gopkg/util/logger"
	"github.com/bytedance/sonic"
)

func Marshal(input interface{}) string {
	if input == nil {
		return ""
	}

	output, err := sonic.MarshalString(input)
	if err != nil {
		logger.Error(err)
		return ""
	}
	return output
}
