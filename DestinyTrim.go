package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	_ "os"
)

var (
//
)

type Mani struct {
	Manifest []struct {
		DestinyActivityDefinition []struct {
			ActivityHash        int64         `json:"activityHash"`
			ActivityName        string        `json:"activityName,omitempty"`
			ActivityDescription string        `json:"activityDescription,omitempty"`
			Icon                string        `json:"icon"`
			ReleaseIcon         string        `json:"releaseIcon"`
			ReleaseTime         int           `json:"releaseTime"`
			ActivityLevel       int           `json:"activityLevel"`
			CompletionFlagHash  int           `json:"completionFlagHash"`
			ActivityPower       float64       `json:"activityPower"`
			MinParty            int           `json:"minParty"`
			MaxParty            int           `json:"maxParty"`
			MaxPlayers          int           `json:"maxPlayers"`
			DestinationHash     int64         `json:"destinationHash"`
			PlaceHash           int64         `json:"placeHash"`
			ActivityTypeHash    int           `json:"activityTypeHash"`
			Tier                int           `json:"tier"`
			PgcrImage           string        `json:"pgcrImage,omitempty"`
			Rewards             []interface{} `json:"rewards"`
			Skulls              []interface{} `json:"skulls"`
			IsPlaylist          bool          `json:"isPlaylist"`
			IsMatchmade         bool          `json:"isMatchmade"`
			Hash                int64         `json:"hash"`
			Index               int           `json:"index"`
		} `json:"DestinyActivityDefinition"`

		DestinyActivityTypeDefinition []struct {
			ActivityTypeHash                       int64  `json:"activityTypeHash"`
			Identifier                             string `json:"identifier,omitempty"`
			ActivityTypeName                       string `json:"activityTypeName,omitempty"`
			ActivityTypeDescription                string `json:"activityTypeDescription,omitempty"`
			Icon                                   string `json:"icon"`
			ActiveBackgroundVirtualPath            string `json:"activeBackgroundVirtualPath,omitempty"`
			CompletedBackgroundVirtualPath         string `json:"completedBackgroundVirtualPath,omitempty"`
			HiddenOverrideVirtualPath              string `json:"hiddenOverrideVirtualPath,omitempty"`
			TooltipBackgroundVirtualPath           string `json:"tooltipBackgroundVirtualPath,omitempty"`
			EnlargedActiveBackgroundVirtualPath    string `json:"enlargedActiveBackgroundVirtualPath,omitempty"`
			EnlargedCompletedBackgroundVirtualPath string `json:"enlargedCompletedBackgroundVirtualPath,omitempty"`
			EnlargedHiddenOverrideVirtualPath      string `json:"enlargedHiddenOverrideVirtualPath,omitempty"`
			EnlargedTooltipBackgroundVirtualPath   string `json:"enlargedTooltipBackgroundVirtualPath,omitempty"`
			Order                                  int    `json:"order"`
			Hash                                   int64  `json:"hash"`
			Index                                  int    `json:"index"`
			StatGroup                              string `json:"statGroup,omitempty"`
			FriendlyURLID                          string `json:"friendlyUrlId,omitempty"`
		} `json:"DestinyActivityTypeDefinition"`

		DestinyClassDefinition []struct {
			ClassHash              int64  `json:"classHash"`
			ClassType              int    `json:"classType"`
			ClassName              string `json:"className"`
			ClassNameMale          string `json:"classNameMale"`
			ClassNameFemale        string `json:"classNameFemale"`
			ClassIdentifier        string `json:"classIdentifier"`
			MentorVendorIdentifier string `json:"mentorVendorIdentifier"`
			Hash                   int64  `json:"hash"`
			Index                  int    `json:"index"`
		} `json:"DestinyClassDefinition"`

		DestinyGenderDefinition []struct {
		} `json:"DestinyGenderDefinition,omitempty"`
		DestinyInventoryBucketDefinition []struct {
		} `json:"DestinyInventoryBucketDefinition,omitempty"`
		DestinyInventoryItemDefinition []struct {
		} `json:"DestinyInventoryItemDefinition,omitempty"`
		DestinyProgressionDefinition []struct {
		} `json:"DestinyProgressionDefinition,omitempty"`
		DestinyRaceDefinition []struct {
		} `json:"DestinyRaceDefinition,omitempty"`
		DestinyTalentGridDefinition []struct {
		} `json:"DestinyTalentGridDefinition,omitempty"`
		DestinyUnlockFlagDefinition []struct {
		} `json:"DestinyUnlockFlagDefinition,omitempty"`
		DestinyVendorDefinition []struct {
		} `json:"DestinyVendorDefinition,omitempty"`
		DestinyHistoricalStatsDefinition []struct {
		} `json:"DestinyHistoricalStatsDefinition,omitempty"`
		DestinyDirectorBookDefinition []struct {
		} `json:"DestinyDirectorBookDefinition,omitempty"`
		DestinyStatDefinition []struct {
		} `json:"DestinyStatDefinition,omitempty"`
		DestinySandboxPerkDefinition []struct {
		} `json:"DestinySandboxPerkDefinition,omitempty"`
		DestinyDestinationDefinition []struct {
		} `json:"DestinyDestinationDefinition,omitempty"`
		DestinyPlaceDefinition []struct {
		} `json:"DestinyPlaceDefinition,omitempty"`
		DestinyActivityBundleDefinition []struct {
		} `json:"DestinyActivityBundleDefinition,omitempty"`
		DestinyStatGroupDefinition []struct {
		} `json:"DestinyStatGroupDefinition,omitempty"`
		DestinySpecialEventDefinition []struct {
		} `json:"DestinySpecialEventDefinition,omitempty"`
		DestinyFactionDefinition []struct {
		} `json:"DestinyFactionDefinition,omitempty"`
		DestinyVendorCategoryDefinition []struct {
		} `json:"DestinyVendorCategoryDefinition,omitempty"`
		DestinyEnemyRaceDefinition []struct {
		} `json:"DestinyEnemyRaceDefinition,omitempty"`
		DestinyScriptedSkullDefinition []struct {
		} `json:"DestinyScriptedSkullDefinition,omitempty"`
		DestinyTriumphSetDefinition []struct {
		} `json:"DestinyTriumphSetDefinition,omitempty"`
		DestinyItemCategoryDefinition []struct {
		} `json:"DestinyItemCategoryDefinition,omitempty"`
		DestinyRewardSourceDefinition []struct {
		} `json:"DestinyRewardSourceDefinition,omitempty"`
		DestinyObjectiveDefinition []struct {
		} `json:"DestinyObjectiveDefinition,omitempty"`
		DestinyDamageTypeDefinition []struct {
		} `json:"DestinyDamageTypeDefinition,omitempty"`
		DestinyCombatantDefinition []struct {
		} `json:"DestinyCombatantDefinition,omitempty"`
		DestinyActivityCategoryDefinition []struct {
		} `json:"DestinyActivityCategoryDefinition,omitempty"`
		DestinyRecordDefinition []struct {
		} `json:"DestinyRecordDefinition,omitempty"`
		DestinyRecordBookDefinition []struct {
		} `json:"DestinyRecordBookDefinition,omitempty"`
		DestinyBondDefinition []struct {
		} `json:"DestinyBondDefinition,omitempty"`
		DestinyLocationDefinition []struct {
		} `json:"DestinyLocationDefinition,omitempty"`
		DestinyGrimoireDefinition []struct {
		} `json:"DestinyGrimoireDefinition,omitempty"`
		DestinyGrimoireCardDefinition []struct {
		} `json:"DestinyGrimoireCardDefinition,omitempty"`
	} `json:"manifest"`
}

func main() {
	fmt.Println("Running trimmer")

	file, _ := ioutil.ReadFile("DestinyManifest.json")

	var manifest Mani
	json.Unmarshal(file, &manifest)

	fmt.Println("Activity Definitions")
	for i, e := range manifest.Manifest[0].DestinyActivityDefinition {
		fmt.Println(i, e.ActivityName)
	}

	fmt.Println()
	fmt.Println("Activity Type Definition")
	for i, e := range manifest.Manifest[1].DestinyActivityTypeDefinition {
		fmt.Println(i, e.ActivityTypeName)
	}

	fmt.Println()
	fmt.Println("Class Definitions")
	for i, e := range manifest.Manifest[2].DestinyClassDefinition {
		fmt.Println(i, e.ClassName)
	}

}
