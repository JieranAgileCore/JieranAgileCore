package cdrType

import "github.com/JieranAgileCore/chf/cdr/asn"

// Need to import "goJieranAgileCore/lib/aper" if it uses "aper"

type SubscriptionID struct { /* Set Type */
	SubscriptionIDType SubscriptionIDType `ber:"tagNum:0"`
	SubscriptionIDData asn.UTF8String     `ber:"tagNum:1"`
}
