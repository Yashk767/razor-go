package utils

import "razor/core"

func getStakeManagerAddress() string {
	addresses := core.AssignAddressesFromJSON()
	return addresses.StakeManagerAddress
}

func GetAssetManagerAddress() string {
	addresses := core.AssignAddressesFromJSON()
	return addresses.AssetManagerAddress
}

func getVoteManagerAddress() string {
	addresses := core.AssignAddressesFromJSON()
	return addresses.VoteManagerAddress
}

func getRAZORAddress() string {
	addresses := core.AssignAddressesFromJSON()
	return addresses.RAZORAddress
}

func getRandomAddress() string {
	addresses := core.AssignAddressesFromJSON()
	return addresses.RandomClientAddress
}

func getBlockManagerAddress() string {
	addresses := core.AssignAddressesFromJSON()
	return addresses.BlockManagerAddress
}

func getParametersAddress() string {
	addresses := core.AssignAddressesFromJSON()
	return addresses.ParametersAddress
}
