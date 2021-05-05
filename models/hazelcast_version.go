package models

//The Hazelcast version in order to create a cluster or find upgradeable versions.
type HazelcastVersion struct {
	Version                string   `json:"version"`
	IsEnabledForStarter    bool     `json:"IsEnabledForStarter"`
	IsEnabledForEnterprise bool     `json:"isEnabledForEnterprise"`
	UpgradeableVersions    []string `json:"upgradeableVersions"`
}
