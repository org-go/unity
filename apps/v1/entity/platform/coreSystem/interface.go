package coreSystem

import (
	"context"
)

type (
	IcorePlatform interface {
		corePlatform
	}

	corePlatform interface {

		//  Payment
		Payment(ctx context.Context) (pay string, err error)

	}


)
