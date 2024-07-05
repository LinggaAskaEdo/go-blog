package common

import (
	"strings"

	"github.com/foolin/mixer"

	x "github.com/linggaaskaedo/go-blog/stdlib/errors/entity"
)

const (
	key = "3m35eS"
)

func MixerEncode(data int64) string {
	encodePaddingData := mixer.EncodeNumber(key, data)
	return strings.ToLower(encodePaddingData)

	// return strings.ToLower(mixer.EncodeID(key, data))
}

func MixerDecode(data string) (int64, error) {
	decodePaddingData, err := mixer.DecodeNumber(key, strings.ToUpper(data))
	if err != nil {
		return decodePaddingData, x.Wrap(err, "decode_id")
	}

	return decodePaddingData, nil

	// result, err := mixer.DecodeID(key, strings.ToUpper(data))
	// if err != nil {
	// 	return result, err
	// }

	// return result, nil
}
