package main

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/structs"
	"github.com/leecombs/destinymanifesttrimmer/models"
	"io/ioutil"
	"log"
	_ "os"
	_ "reflect"
)

func refl(mini interface{}, whole interface{}) {
	miniS := structs.New(mini)
	wholeS := structs.New(whole)

	for _, e := range miniS.Names() {
		// fmt.Println(e)
		miniS.Field(e).Set(wholeS.Field(e).Value())
	}

}

func main() {
	fmt.Println("Running trimmer")

	// Read and load the local manifest file
	var manifest models.DestinyManifest
	file, fileErr := ioutil.ReadFile("DestinyManifest.json")
	if fileErr != nil {
		log.Fatal("fileErr", fileErr)
		return
	}
	json.Unmarshal(file, &manifest)

	// Initialize the output vars
	miniMani := make(map[string]interface{})

	// Start going through the manifest data, transform it into our structs, and build the output
	fmt.Printf("Activity Definitions... ")
	mdadMap := make(map[int64]models.MiniDestinyActivityDefinition)
	for _, e := range manifest.Manifest[0].DestinyActivityDefinition {
		var mdad models.MiniDestinyActivityDefinition
		refl(&mdad, &e)
		mdadMap[e.Hash] = mdad
	}
	miniMani["DestinyActivityDefinition"] = mdadMap
	fmt.Printf("Done: %d\n", len(mdadMap))

	fmt.Printf("Activity Type Definitions... ")
	mdatdMap := make(map[int64]models.MiniDestinyActivityTypeDefinition)
	for _, e := range manifest.Manifest[1].DestinyActivityTypeDefinition {
		var mdatd models.MiniDestinyActivityTypeDefinition
		refl(&mdatd, &e)
		mdatdMap[e.Hash] = mdatd
	}
	miniMani["DestinyActivityTypeDefinition"] = mdatdMap
	fmt.Printf("Done: %d\n", len(mdatdMap))

	fmt.Printf("Class Definitions... ")
	mdcdMap := make(map[int64]models.MiniDestinyClassDefinition)
	for _, e := range manifest.Manifest[2].DestinyClassDefinition {
		var mdcd models.MiniDestinyClassDefinition
		refl(&mdcd, &e)
		mdcdMap[e.Hash] = mdcd
	}
	miniMani["DestinyClassDefinition"] = mdcdMap
	fmt.Printf("Done: %d\n", len(mdcdMap))

	fmt.Printf("Gender Definitions... ")
	mdgdMap := make(map[int64]models.MiniDestinyGenderDefinition)
	for _, e := range manifest.Manifest[3].DestinyGenderDefinition {
		var mdgd models.MiniDestinyGenderDefinition
		refl(&mdgd, &e)
		mdgdMap[e.Hash] = mdgd
	}
	miniMani["DestinyGenderDefinition"] = mdgdMap
	fmt.Printf("Done: %d\n", len(mdgdMap))

	fmt.Printf("Inventory Bucket Definitions... ")
	mibdMap := make(map[int64]models.MiniDestinyInventoryBucketDefinition)
	for _, e := range manifest.Manifest[4].DestinyInventoryBucketDefinition {
		var mibd models.MiniDestinyInventoryBucketDefinition
		refl(&mibd, &e)
		mibdMap[e.Hash] = mibd
	}
	miniMani["DestinyInventoryBucketDefinition"] = mibdMap
	fmt.Printf("Done: %d\n", len(mibdMap))

	fmt.Printf("Inventory Item Definitions... ")
	mdiibMap := make(map[int64]models.MiniDestinyInventoryItemDefinition)
	for _, e := range manifest.Manifest[5].DestinyInventoryItemDefinition {
		var mdiib models.MiniDestinyInventoryItemDefinition

		mdiib.ItemName = e.ItemName
		mdiib.ItemDescription = e.ItemDescription
		mdiib.HasIcon = e.HasIcon
		mdiib.Icon = e.Icon
		mdiib.SecondaryIcon = e.SecondaryIcon
		mdiib.TierType = e.TierType
		mdiib.TierTypeName = e.TierTypeName
		mdiib.ItemType = e.ItemType
		mdiib.ItemTypeName = e.ItemTypeName
		mdiib.SpecialItemType = e.SpecialItemType
		mdiib.ClassType = e.ClassType
		mdiib.BucketTypeHash = int64(e.BucketTypeHash)
		mdiib.PrimaryBaseStatHash = int64(e.PrimaryBaseStatHash)
		mdiib.Stats = e.Stats
		mdiib.Exclusive = e.Exclusive

		mdiibMap[e.Hash] = mdiib
	}
	miniMani["DestinyInventoryItemDefinition"] = mdiibMap
	fmt.Printf("Done: %d\n", len(mdiibMap))

	fmt.Printf("Progression Definitions... ")
	mdpdMap := make(map[int64]models.MiniDestinyProgressionDefinition)
	for _, e := range manifest.Manifest[6].DestinyProgressionDefinition {
		var mdpd models.MiniDestinyProgressionDefinition

		mdpd.Name = e.Name
		mdpd.Scope = e.Scope
		// mdpd.Steps = e.Steps
		mdpd.Visible = e.Visible
		mdpd.Icon = e.Icon
		mdpd.Description = e.Description
		mdpd.Source = e.Source

		mdpdMap[e.Hash] = mdpd
	}
	miniMani["DestinyProgressionDefinition"] = mdpdMap
	fmt.Printf("Done: %d\n", len(mdpdMap))

	fmt.Printf("Race Definitions... ")
	mdrdMap := make(map[int64]models.MiniDestinyRaceDefinition)
	for _, e := range manifest.Manifest[7].DestinyRaceDefinition {
		var mdrd models.MiniDestinyRaceDefinition
		refl(&mdrd, &e)
		mdrdMap[e.Hash] = mdrd
	}
	miniMani["DestinyRaceDefinition"] = mdrdMap
	fmt.Printf("Done: %d\n", len(mdrdMap))

	fmt.Printf("Talent Grid Definitions... ")
	mtgdMap := make(map[int64]models.MiniDestinyTalentGridDefinition)
	for _, e := range manifest.Manifest[8].DestinyTalentGridDefinition {
		var mtdg models.MiniDestinyTalentGridDefinition
		refl(&mtdg, &e)
		mtgdMap[e.Hash] = mtdg
	}
	miniMani["DestinyTalentGridDefinition"] = mtgdMap
	fmt.Printf("Done: %d\n", len(mtgdMap))

	fmt.Printf("Unlock Flag Definitions... ")
	mdufdMap := make(map[int64]models.MiniDestinyUnlockFlagDefinition)
	for _, e := range manifest.Manifest[9].DestinyUnlockFlagDefinition {
		var mdufd models.MiniDestinyUnlockFlagDefinition
		refl(&mdufd, &e)
		mdufdMap[e.Hash] = mdufd
	}
	miniMani["DestinyTalentGridDefinition"] = mdufdMap
	fmt.Printf("Done: %d\n", len(mdufdMap))

	fmt.Printf("Vendor Definitions... ")
	mdvdMap := make(map[int64]models.MiniDestinyVendorDefinition)
	for _, e := range manifest.Manifest[10].DestinyVendorDefinition {
		var mdvd models.MiniDestinyVendorDefinition

		mdvd.Summary.VendorName = e.Summary.VendorName
		mdvd.Summary.VendorDescription = e.Summary.VendorDescription
		mdvd.Summary.VendorIcon = e.Summary.VendorIcon
		mdvd.Summary.FactionHash = e.Summary.FactionHash
		mdvd.Summary.Visible = e.Summary.Visible
		mdvd.Summary.VendorPortrait = e.Summary.VendorPortrait
		mdvd.Summary.VendorBanner = e.Summary.VendorBanner
		mdvd.Summary.VendorCategoryHash = e.Summary.VendorCategoryHash
		mdvd.Summary.VendorSubcategoryHash = e.Summary.VendorSubcategoryHash

		mdvdMap[e.Hash] = mdvd
	}
	miniMani["DestinyVendorDefinition"] = mdvdMap
	fmt.Printf("Done: %d\n", len(mdvdMap))

	fmt.Printf("Historical Stats Definitions... ")
	mdhsdMap := make(map[string]models.MiniDestinyHistoricalStatsDefinition)
	for _, e := range manifest.Manifest[11].DestinyHistoricalStatsDefinition {
		var mdhsd models.MiniDestinyHistoricalStatsDefinition
		refl(&mdhsd, &e)
		mdhsdMap[e.StatID] = mdhsd
	}
	miniMani["DestinyHistoricalStatsDefinition"] = mdhsdMap
	fmt.Printf("Done: %d\n", len(mdhsdMap))

	fmt.Printf("Director Book Definitions... ")
	mddbdMap := make(map[int64]models.MiniDestinyDirectorBookDefinition)
	for _, e := range manifest.Manifest[12].DestinyDirectorBookDefinition {
		var mddbd models.MiniDestinyDirectorBookDefinition
		refl(&mddbd, &e)
		mddbdMap[e.Hash] = mddbd
	}
	miniMani["DestinyDirectorBookDefinition"] = mddbdMap
	fmt.Printf("Done: %d\n", len(mddbdMap))

	fmt.Printf("Stat Definitions... ")
	mdsdMap := make(map[int64]models.MiniDestinyStatDefinition)
	for _, e := range manifest.Manifest[13].DestinyStatDefinition {
		var mdsd models.MiniDestinyStatDefinition
		refl(&mdsd, &e)
		mdsdMap[e.Hash] = mdsd
	}
	miniMani["DestinyStatDefinition"] = mdsdMap
	fmt.Printf("Done: %d\n", len(mdsdMap))

	fmt.Printf("Sandbox Perk Definitions... ")
	mdspdMap := make(map[int64]models.MiniDestinySandboxPerkDefinition)
	for _, e := range manifest.Manifest[14].DestinySandboxPerkDefinition {
		var mdspd models.MiniDestinySandboxPerkDefinition

		mdspd.DisplayName = e.DisplayName
		mdspd.DisplayDescription = e.DisplayDescription
		mdspd.DisplayIcon = e.DisplayIcon
		mdspd.IsDisplayable = e.IsDisplayable
		mdspd.PerkIdentifer = e.PerkIdentifier

		mdspd.PerkGroups.WeaponPerformance = e.PerkGroups.WeaponPerformance
		mdspd.PerkGroups.ImpactEffects = e.PerkGroups.ImpactEffects
		mdspd.PerkGroups.GuardianAttributes = e.PerkGroups.GuardianAttributes
		mdspd.PerkGroups.LightAbilities = e.PerkGroups.LightAbilities
		mdspd.PerkGroups.DamageTypes = e.PerkGroups.DamageTypes

		mdspdMap[e.Hash] = mdspd
	}
	miniMani["DestinySandboxPerkDefinition"] = mdspdMap
	fmt.Printf("Done: %d\n", len(mdspdMap))

	fmt.Printf("Destination Definitions... ")
	mdddMap := make(map[int64]models.MiniDestinyDestinationDefinition)
	for _, e := range manifest.Manifest[15].DestinyDestinationDefinition {
		var mddd models.MiniDestinyDestinationDefinition
		refl(&mddd, &e)
		mdddMap[e.Hash] = mddd
	}
	miniMani["DestinyDestinationDefinition"] = mdddMap
	fmt.Printf("Done: %d\n", len(mdddMap))

	fmt.Printf("Place Definitions... ")
	mindefPlaceMap := make(map[int64]models.MiniDestinyPlaceDefinition)
	for _, e := range manifest.Manifest[16].DestinyPlaceDefinition {
		var mindefPlace models.MiniDestinyPlaceDefinition
		refl(&mindefPlace, &e)
		mindefPlaceMap[e.Hash] = mindefPlace
	}
	miniMani["DestinyPlaceDefinition"] = mindefPlaceMap
	fmt.Printf("Done: %d\n", len(mindefPlaceMap))

	fmt.Printf("Activity Bundle Definitions... ")
	mindefActBunMap := make(map[int64]models.MiniDestinyActivityBundleDefinition)
	for _, e := range manifest.Manifest[17].DestinyActivityBundleDefinition {
		var mindefActBun models.MiniDestinyActivityBundleDefinition
		refl(&mindefActBun, &e)
		mindefActBunMap[e.Hash] = mindefActBun
	}
	miniMani["DestinyActivityBundleDefinition"] = mindefActBunMap
	fmt.Printf("Done: %d\n", len(mindefActBunMap))

	fmt.Printf("Stat Group Definitions... ")
	mindefStatGrpMap := make(map[int64]models.MiniDestinyStatGroupDefinition)
	for _, e := range manifest.Manifest[18].DestinyStatGroupDefinition {
		var mindefStatGrp models.MiniDestinyStatGroupDefinition
		refl(&mindefStatGrp, &e)
		mindefStatGrpMap[e.Hash] = mindefStatGrp
	}
	miniMani["DestinyStatGroupDefinition"] = mindefStatGrpMap
	fmt.Printf("Done: %d\n", len(mindefStatGrpMap))

	fmt.Printf("Special Event Definitions... ")
	mindefSpEvtMap := make(map[int64]models.MiniDestinySpecialEventDefinition)
	for _, e := range manifest.Manifest[19].DestinySpecialEventDefinition {
		var mindefSpEvt models.MiniDestinySpecialEventDefinition
		refl(&mindefSpEvt, &e)
		mindefSpEvtMap[e.Hash] = mindefSpEvt
	}
	miniMani["DestinySpecialEventDefinition"] = mindefSpEvtMap
	fmt.Printf("Done: %d\n", len(mindefSpEvtMap))

	fmt.Printf("Faction Definitions... ")
	mindefFactionMap := make(map[int64]models.MiniDestinyFactionDefinition)
	for _, e := range manifest.Manifest[20].DestinyFactionDefinition {
		var mindefFaction models.MiniDestinyFactionDefinition
		refl(&mindefFaction, &e)
		mindefFactionMap[e.Hash] = mindefFaction
	}
	miniMani["DestinyFactionDefinition"] = mindefFactionMap
	fmt.Printf("Done: %d\n", len(mindefFactionMap))

	fmt.Printf("Vendor Category Definitions... ")
	mindefVenCatMap := make(map[int64]models.MiniDestinyVendorCategoryDefinition)
	for _, e := range manifest.Manifest[21].DestinyVendorCategoryDefinition {
		var mindefVenCat models.MiniDestinyVendorCategoryDefinition
		refl(&mindefVenCat, &e)
		mindefVenCatMap[e.Hash] = mindefVenCat
	}
	miniMani["DestinyVendorCategoryDefinition"] = mindefVenCatMap
	fmt.Printf("Done: %d\n", len(mindefVenCatMap))

	fmt.Printf("Enemy Race Definitions... ")
	mderdMap := make(map[int64]models.MiniDestinyEnemyRaceDefinition)
	for _, e := range manifest.Manifest[22].DestinyEnemyRaceDefinition {
		var mderd models.MiniDestinyEnemyRaceDefinition
		refl(&mderd, &e)
		mderdMap[e.Hash] = mderd
	}
	miniMani["DestinyEnemyRaceDefinition"] = mderdMap
	fmt.Printf("Done: %d\n", len(mderdMap))

	fmt.Printf("Scripted Skull Definitions... ")
	mdssdMap := make(map[int64]models.MiniDestinyScriptedSkullDefinition)
	for _, e := range manifest.Manifest[23].DestinyScriptedSkullDefinition {
		var mdssd models.MiniDestinyScriptedSkullDefinition
		refl(&mdssd, &e)
		mdssdMap[e.Hash] = mdssd
	}
	miniMani["DestinyScriptedSkullDefinition"] = mdssdMap
	fmt.Printf("Done: %d\n", len(mdssdMap))

	fmt.Printf("Triumph Set Definitions... ")
	mdtsMap := make(map[int64]models.MiniDestinyTriumphSetDefinition)
	for _, e := range manifest.Manifest[24].DestinyTriumphSetDefinition {
		var mdts models.MiniDestinyTriumphSetDefinition
		refl(&mdts, &e)
		mdtsMap[e.Hash] = mdts
	}
	miniMani["DestinyTriumphSetDefinition"] = mdtsMap
	fmt.Printf("Done: %d\n", len(mdtsMap))

	fmt.Printf("Item Category Definitions... ")
	mdicdMap := make(map[int64]models.MiniDestinyItemCategoryDefinition)
	for _, e := range manifest.Manifest[25].DestinyItemCategoryDefinition {
		var mdicd models.MiniDestinyItemCategoryDefinition
		refl(&mdicd, &e)
		mdicdMap[e.Hash] = mdicd
	}
	miniMani["DestinyItemCategoryDefinition"] = mdicdMap
	fmt.Printf("Done: %d\n", len(mdicdMap))

	fmt.Printf("Reward Source Definitions... ")
	mdrsdMap := make(map[int64]models.MiniDestinyRewardSourceDefinition)
	for _, e := range manifest.Manifest[26].DestinyRewardSourceDefinition {
		var mdrsd models.MiniDestinyRewardSourceDefinition
		refl(&mdrsd, &e)
		mdrsdMap[e.Hash] = mdrsd
	}
	miniMani["DestinyRewardSourceDefinition"] = mdrsdMap
	fmt.Printf("Done: %d\n", len(mdrsdMap))

	fmt.Printf("Objective Definitions... ")
	mdodMap := make(map[int64]models.MiniDestinyObjectiveDefinition)
	for _, e := range manifest.Manifest[27].DestinyObjectiveDefinition {
		var mdod models.MiniDestinyObjectiveDefinition
		refl(&mdod, &e)
		mdodMap[e.Hash] = mdod
	}
	miniMani["DestinyObjectiveDefinition"] = mdodMap
	fmt.Printf("Done: %d\n", len(mdodMap))

	fmt.Printf("Damage Type Definitions... ")
	mddtdMap := make(map[int64]models.MiniDestinyDamageTypeDefinition)
	for _, e := range manifest.Manifest[28].DestinyDamageTypeDefinition {
		var mddtd models.MiniDestinyDamageTypeDefinition
		refl(&mddtd, &e)
		mddtdMap[e.Hash] = mddtd
	}
	miniMani["DestinyDamageTypeDefinition"] = mddtdMap
	fmt.Printf("Done: %d\n", len(mddtdMap))

	fmt.Printf("Combatant Definitions... ")
	mdcombatdMap := make(map[int64]models.MiniDestinyCombatantDefinition)
	for _, e := range manifest.Manifest[29].DestinyCombatantDefinition {
		var mdcombatd models.MiniDestinyCombatantDefinition
		refl(&mdcombatd, &e)
		mdcombatdMap[e.Hash] = mdcombatd
	}
	miniMani["DestinyCombatantDefinition"] = mdcombatdMap
	fmt.Printf("Done: %d\n", len(mdcombatdMap))

	fmt.Printf("Activity Category Definitions... ")
	mdacdMap := make(map[int64]models.MiniDestinyActivityCategoryDefinition)
	for _, e := range manifest.Manifest[30].DestinyActivityCategoryDefinition {
		var mdacd models.MiniDestinyActivityCategoryDefinition
		refl(&mdacd, &e)
		mdacdMap[e.Hash] = mdacd
	}
	miniMani["DestinyActivityCategoryDefinition"] = mdacdMap
	fmt.Printf("Done: %d\n", len(mdacdMap))

	fmt.Printf("Record Definitions... ")
	mdrecordMap := make(map[int64]models.MiniDestinyRecordDefinition)
	for _, e := range manifest.Manifest[31].DestinyRecordDefinition {
		var mdrecord models.MiniDestinyRecordDefinition
		refl(&mdrecord, &e)
		mdrecordMap[e.Hash] = mdrecord
	}
	miniMani["DestinyRecordDefinition"] = mdrecordMap
	fmt.Printf("Done: %d\n", len(mdrecordMap))

	fmt.Printf("Record Book Definitions... ")
	mdrecordbookMap := make(map[int64]models.MiniDestinyRecordBookDefinition)
	for _, e := range manifest.Manifest[32].DestinyRecordBookDefinition {
		var mdrecordbook models.MiniDestinyRecordBookDefinition
		refl(&mdrecordbook, &e)
		mdrecordbookMap[e.Hash] = mdrecordbook
	}
	miniMani["DestinyRecordBookDefinition"] = mdrecordbookMap
	fmt.Printf("Done: %d\n", len(mdrecordbookMap))

	fmt.Printf("Bond Definitions... ")
	mdbondMap := make(map[int64]models.MiniDestinyBondDefinition)
	for _, e := range manifest.Manifest[33].DestinyBondDefinition {
		var mdbond models.MiniDestinyBondDefinition
		refl(&mdbond, &e)
		mdbondMap[e.Hash] = mdbond
	}
	miniMani["DestinyBondDefinition"] = mdbondMap
	fmt.Printf("Done: %d\n", len(mdbondMap))

	fmt.Printf("Location Definitions... ")
	mdlocMap := make(map[int64]models.MiniDestinyLocationDefinition)
	for _, e := range manifest.Manifest[34].DestinyLocationDefinition {
		var mdloc models.MiniDestinyLocationDefinition
		refl(&mdloc, &e)
		mdlocMap[e.Hash] = mdloc
	}
	miniMani["DestinyLocationDefinition"] = mdlocMap
	fmt.Printf("Done: %d\n", len(mdlocMap))

	fmt.Printf("Grimoire Definitions... ")
	mdgrimMap := make(map[int64]models.MiniDestinyGrimoireDefinition)
	for i, e := range manifest.Manifest[35].DestinyGrimoireDefinition {
		var mdgrim models.MiniDestinyGrimoireDefinition
		refl(&mdgrim, &e)
		mdgrimMap[int64(i)] = mdgrim
	}
	miniMani["DestinyGrimoireDefinition"] = mdgrimMap
	fmt.Printf("Done: %d\n", len(mdgrimMap))

	fmt.Printf("Grimoire Card Definitions... ")
	mdgrimcardMap := make(map[int64]models.MiniDestinyGrimoireCardDefinition)
	for _, e := range manifest.Manifest[36].DestinyGrimoireCardDefinition {
		var mdgrimcard models.MiniDestinyGrimoireCardDefinition
		refl(&mdgrimcard, &e)
		mdgrimcardMap[int64(e.CardID)] = mdgrimcard
	}
	miniMani["DestinyGrimoireCardDefinition"] = mdgrimcardMap
	fmt.Printf("Done: %d\n", len(mdgrimcardMap))

	// Convert the new manifest to .json and write the file
	// b, _ := json.Marshal(miniMani)
	b, _ := json.MarshalIndent(miniMani, "", "    ")
	ioutil.WriteFile("MiniMani-NewConversion.json", b, 0644)
}
