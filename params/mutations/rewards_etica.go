// Copyright 2019 The multi-geth Authors
// This file is part of the multi-geth library.
//
// The multi-geth library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The multi-geth library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the multi-geth library. If not, see <http://www.gnu.org/licenses/>.
package mutations

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params/types/ctypes"
	"github.com/ethereum/go-ethereum/params/vars"
)

func etip1017BlockReward(config ctypes.ChainConfigurator, header *types.Header, uncles []*types.Header) (*big.Int, []*big.Int) {
	blockReward := vars.ETIP1017BlockReward // EGAZ tail emission, 2 EGAZ per block

	
	// Accumulate the rewards for the miner and any included uncles
	uncleRewards := make([]*big.Int, len(uncles))
	reward := new(big.Int).Set(blockReward)
	r := new(big.Int)
	for i, uncle := range uncles {
		r.Add(uncle.Number, big8)
		r.Sub(r, header.Number)
		r.Mul(r, blockReward)
		r.Div(r, big8)

		ur := new(big.Int).Set(r)
		uncleRewards[i] = ur

		r.Div(blockReward, big32)
		reward.Add(reward, r)
	}

	return reward, uncleRewards
}
