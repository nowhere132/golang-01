package type_system

import "fmt"

type SubscriptionType string

const (
	SubscriptionTypeMasterListing SubscriptionType = "master_listing"
	SubscriptionTypeAuction       SubscriptionType = "auction"
)

func TestSubscriptionType() {
	allTypes := []SubscriptionType{
		SubscriptionTypeMasterListing, SubscriptionTypeAuction,
	}

	if contains(allTypes, SubscriptionTypeMasterListing) {
		fmt.Println("CONTAINED master_listing")
	}

	//x :=
}

func contains[T SubscriptionType | string](arr []T, val T) bool {
	for _, elem := range arr {
		if elem == val {
			return true
		}
	}
	return false
}
