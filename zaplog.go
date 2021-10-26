package goutils

import (
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitZapLog(v *viper.Viper) (*zap.Logger, error) {
	zapCfg := zap.NewProductionConfig()

	if v != nil {
		if err := v.Unmarshal(&zapCfg, viper.DecodeHook(mapstructure.TextUnmarshallerHookFunc())); err != nil {
			return nil, err
		}

		// Configure EncodreConfig. This doesnot work with viper & unmarshall with unmarshallText.
		// So we hardcode them here.
		if v.GetBool("color") {
			zapCfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		}
	}

	zapCfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	return zapCfg.Build()
}
