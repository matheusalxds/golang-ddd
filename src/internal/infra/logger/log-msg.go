package logger

import (
	"fmt"

	"go.uber.org/zap"
)

func MsgStart(msgParts []string, data map[string]interface{}) (string, zap.Field) {
	return fmt.Sprintf("%s - %s - Start", msgParts[0], msgParts[1]), zap.Any("data", data)
}

func MsgEnd(msgParts []string, data map[string]interface{}) (string, zap.Field) {
	return fmt.Sprintf("%s - %s - End", msgParts[0], msgParts[1]), zap.Any("data", data)
}
