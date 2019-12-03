package interceptors

import (
	"context"
	"regexp"

	"github.com/MixinNetwork/supergroup.mixin.one/session"
)

func CheckQuest(ctx context.Context, data []byte) (bool, error) {
	match, err := regexp.MatchString(`^([1-9]|10)\/10\s`, string(data))
	if err != nil {
		session.Logger(ctx).Errorf("CheckSex NewImageAnnotatorClient ERROR: %+v", err)
		return false, nil
	}

	if match {
		return true, nil;
	}

	return false, nil
}
