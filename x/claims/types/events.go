// Copyright 2024 Tabi Foundation
// This file is part of the Tabi Network packages.
//
// Tabi is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Tabi packages are distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.

package types

// claim module event types
const (
	EventTypeClaim              = "claim"
	EventTypeMergeClaimsRecords = "merge_claims_records"

	AttributeKeyActionType             = "action"
	AttributeKeyRecipient              = "recipient"
	AttributeKeyClaimedCoins           = "claimed_coins"
	AttributeKeyFundCommunityPoolCoins = "fund_community_pool_coins"
)
