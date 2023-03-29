// Copyright 2022 Cascadia Foundation
// This file is part of the Cascadia Network packages.
//
// Cascadia is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Cascadia packages are distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Cascadia packages. If not, see https://github.com/evmos/evmos/blob/main/LICENSE

package utils

import (
	"strings"

	"github.com/cascadiafoundation/cascadia/crypto/ethsecp256k1"

	errorsmod "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/crypto/types/multisig"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errortypes "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	// MainnetChainID defines the Cascadia EIP155 chain ID for mainnet
	MainnetChainID = "cascadia_9001"
	// TestnetChainID defines the Cascadia EIP155 chain ID for testnet
	TestnetChainID = "cascadia_9000"
	// BaseDenom defines the Cascadia mainnet denomination
	BaseDenom = "aCC"
)

// IsMainnet returns true if the chain-id has the Cascadia mainnet EIP155 chain prefix.
func IsMainnet(chainID string) bool {
	return strings.HasPrefix(chainID, MainnetChainID)
}

// IsTestnet returns true if the chain-id has the Cascadia testnet EIP155 chain prefix.
func IsTestnet(chainID string) bool {
	return strings.HasPrefix(chainID, TestnetChainID)
}

// IsSupportedKey returns true if the pubkey type is supported by the chain
// (i.e eth_secp256k1, amino multisig, ed25519).
// NOTE: Nested multisigs are not supported.
func IsSupportedKey(pubkey cryptotypes.PubKey) bool {
	switch pubkey := pubkey.(type) {
	case *ethsecp256k1.PubKey, *ed25519.PubKey:
		return true
	case multisig.PubKey:
		if len(pubkey.GetPubKeys()) == 0 {
			return false
		}

		for _, pk := range pubkey.GetPubKeys() {
			switch pk.(type) {
			case *ethsecp256k1.PubKey, *ed25519.PubKey:
				continue
			default:
				// Nested multisigs are unsupported
				return false
			}
		}

		return true
	default:
		return false
	}
}

// GetCascadiaAddressFromBech32 returns the sdk.Account address of given address,
// while also changing bech32 human readable prefix (HRP) to the value set on
// the global sdk.Config (eg: `cascadia`).
// The function fails if the provided bech32 address is invalid.
func GetCascadiaAddressFromBech32(address string) (sdk.AccAddress, error) {
	bech32Prefix := strings.SplitN(address, "1", 2)[0]
	if bech32Prefix == address {
		return nil, errorsmod.Wrapf(errortypes.ErrInvalidAddress, "invalid bech32 address: %s", address)
	}

	addressBz, err := sdk.GetFromBech32(address, bech32Prefix)
	if err != nil {
		return nil, errorsmod.Wrapf(errortypes.ErrInvalidAddress, "invalid address %s, %s", address, err.Error())
	}

	// safety check: shouldn't happen
	if err := sdk.VerifyAddressFormat(addressBz); err != nil {
		return nil, err
	}

	return sdk.AccAddress(addressBz), nil
}
