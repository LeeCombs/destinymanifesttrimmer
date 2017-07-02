package main

import (
	"encoding/json"
	"fmt"
	"github.com/leecombs/destinymanifesttrimmer/models"
	"io/ioutil"
	"log"
	_ "os"
)

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

		mdad.ActivityName = e.ActivityName
		mdad.ActivityDescription = e.ActivityDescription
		mdad.Icon = e.Icon
		mdad.ActivityLevel = e.ActivityLevel
		mdad.DestinationHash = int64(e.DestinationHash)
		mdad.PlaceHash = int64(e.PlaceHash)
		mdad.ActivityTypeHash = int64(e.ActivityTypeHash)
		mdad.Rewards = e.Rewards
		mdad.Skulls = e.Skulls

		mdadMap[e.Hash] = mdad
	}
	miniMani["DestinyActivityDefinition"] = mdadMap
	fmt.Printf("Done: %d\n", len(mdadMap))

	fmt.Printf("Activity Type Definitions... ")
	mdatdMap := make(map[int64]models.MiniDestinyActivityTypeDefinition)
	for _, e := range manifest.Manifest[1].DestinyActivityTypeDefinition {
		var mdatd models.MiniDestinyActivityTypeDefinition

		mdatd.Identifier = e.Identifier
		mdatd.ActivityTypeName = e.ActivityTypeName
		mdatd.ActivityTypeDescription = e.ActivityTypeDescription
		mdatd.Icon = e.Icon

		mdatdMap[e.Hash] = mdatd
	}
	miniMani["DestinyActivityTypeDefinition"] = mdatdMap
	fmt.Printf("Done: %d\n", len(mdatdMap))

	fmt.Printf("Class Definitions... ")
	mdcdMap := make(map[int64]models.MiniDestinyClassDefinition)
	for _, e := range manifest.Manifest[2].DestinyClassDefinition {
		var mdcd models.MiniDestinyClassDefinition

		mdcd.ClassName = e.ClassName
		mdcd.ClassType = e.ClassType

		mdcdMap[e.Hash] = mdcd
	}
	miniMani["DestinyClassDefinition"] = mdcdMap
	fmt.Printf("Done: %d\n", len(mdcdMap))

	fmt.Printf("Gender Definitions... ")
	mdgdMap := make(map[int64]models.MiniDestinyGenderDefinition)
	for _, e := range manifest.Manifest[3].DestinyGenderDefinition {
		var mdgd models.MiniDestinyGenderDefinition

		mdgd.GenderName = e.GenderName
		mdgd.GenderType = e.GenderType

		mdgdMap[e.Hash] = mdgd
	}
	miniMani["DestinyGenderDefinition"] = mdgdMap
	fmt.Printf("Done: %d\n", len(mdgdMap))

	fmt.Printf("Inventory Bucket Definitions... ")
	mibdMap := make(map[int64]models.MiniDestinyInventoryBucketDefinition)
	for _, e := range manifest.Manifest[4].DestinyInventoryBucketDefinition {
		var mibd models.MiniDestinyInventoryBucketDefinition

		mibd.BucketName = e.BucketName
		mibd.BucketDescription = e.BucketDescription
		mibd.ItemCount = e.ItemCount
		mibd.Category = e.Category
		mibd.Location = e.Location
		mibd.HasTransferDestination = e.HasTransferDestination
		mibd.Enabled = e.Enabled

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

		mdrd.RaceType = e.RaceType
		mdrd.RaceName = e.RaceName
		mdrd.RaceNameMale = e.RaceNameMale
		mdrd.RaceNameFemale = e.RaceNameFemale
		mdrd.RaceDescription = e.RaceDescription

		mdrdMap[e.Hash] = mdrd
	}
	miniMani["DestinyRaceDefinition"] = mdrdMap
	fmt.Printf("Done: %d\n", len(mdrdMap))

	fmt.Printf("Talent Grid Definitions... ")
	mtgdMap := make(map[int64]models.MiniDestinyTalentGridDefinition)
	for _, e := range manifest.Manifest[8].DestinyTalentGridDefinition {
		var mtdg models.MiniDestinyTalentGridDefinition

		mtdg.ProgressionHash = int64(e.ProgressionHash)
		mtdg.Nodes = e.Nodes

		mtgdMap[e.Hash] = mtdg
	}
	miniMani["DestinyTalentGridDefinition"] = mtgdMap
	fmt.Printf("Done: %d\n", len(mtgdMap))

	fmt.Printf("Unlock Flag Definitions... ")
	mdufdMap := make(map[int64]models.MiniDestinyUnlockFlagDefinition)
	for _, e := range manifest.Manifest[9].DestinyUnlockFlagDefinition {
		var mdufd models.MiniDestinyUnlockFlagDefinition

		mdufd.DisplayName = e.DisplayName
		mdufd.DisplayDescription = e.DisplayDescription

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

		mdhsd.StatID = e.StatID
		mdhsd.StatName = e.StatName
		mdhsd.StatDescription = e.StatDescription
		mdhsd.Group = e.Group
		mdhsd.Category = e.Category
		mdhsd.UnitType = e.UnitType
		mdhsd.UnitLabel = e.UnitLabel
		mdhsd.Weight = e.Weight
		mdhsd.IconImage = e.IconImage

		mdhsdMap[e.StatID] = mdhsd
	}
	miniMani["DestinyHistoricalStatsDefinition"] = mdhsdMap
	fmt.Printf("Done: %d\n", len(mdhsdMap))

	fmt.Printf("Director Book Definitions... ")
	mddbdMap := make(map[int64]models.MiniDestinyDirectorBookDefinition)
	for _, e := range manifest.Manifest[12].DestinyDirectorBookDefinition {
		var mddbd models.MiniDestinyDirectorBookDefinition

		mddbd.BookName = e.BookName
		mddbd.BookDescription = e.BookDescription
		mddbd.BookNumber = e.BookNumber
		mddbd.Visible = e.Visible
		mddbd.IsOverview = e.IsOverview
		mddbd.DestinationHash = e.DestinationHash

		mddbdMap[e.Hash] = mddbd
	}
	miniMani["DestinyDirectorBookDefinition"] = mddbdMap
	fmt.Printf("Done: %d\n", len(mddbdMap))

	fmt.Printf("Stat Definitions... ")
	mdsdMap := make(map[int64]models.MiniDestinyStatDefinition)
	for _, e := range manifest.Manifest[13].DestinyStatDefinition {
		var mdsd models.MiniDestinyStatDefinition

		mdsd.StatName = e.StatName
		mdsd.StatDescription = e.StatDescription
		mdsd.Icon = e.Icon

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

		mddd.DestinationName = e.DestinationName
		mddd.DestinationIdentifier = e.DestinationIdentifier
		mddd.DestinationDescription = e.DestinationDescription
		mddd.Icon = e.Icon
		mddd.PlaceHash = e.PlaceHash
		mddd.LocationIndentifier = e.LocationIdentifier

		mdddMap[e.Hash] = mddd
	}
	miniMani["DestinyDestinationDefinition"] = mdddMap
	fmt.Printf("Done: %d\n", len(mdddMap))

	fmt.Printf("Place Definitions... ")
	mindefPlaceMap := make(map[int64]models.MiniDestinyPlaceDefinition)
	for _, e := range manifest.Manifest[16].DestinyPlaceDefinition {
		var mindefPlace models.MiniDestinyPlaceDefinition

		mindefPlace.PlaceName = e.PlaceName
		mindefPlace.PlaceDescription = e.PlaceDescription
		mindefPlace.Icon = e.Icon

		mindefPlaceMap[e.Hash] = mindefPlace
	}
	miniMani["DestinyPlaceDefinition"] = mindefPlaceMap
	fmt.Printf("Done: %d\n", len(mindefPlaceMap))

	fmt.Printf("Activity Bundle Definitions... ")
	mindefActBunMap := make(map[int64]models.MiniDestinyActivityBundleDefinition)
	for _, e := range manifest.Manifest[17].DestinyActivityBundleDefinition {
		var mindefActBun models.MiniDestinyActivityBundleDefinition

		mindefActBun.ActivityName = e.ActivityName
		mindefActBun.ActivityDescription = e.ActivityDescription
		mindefActBun.ActivityTypeHash = e.ActivityTypeHash
		mindefActBun.ActivityHashes = e.ActivityHashes
		mindefActBun.Icon = e.Icon
		mindefActBun.DestinationHash = e.DestinationHash
		mindefActBun.PlaceHash = e.PlaceHash

		mindefActBunMap[e.Hash] = mindefActBun
	}
	miniMani["DestinyActivityBundleDefinition"] = mindefActBunMap
	fmt.Printf("Done: %d\n", len(mindefActBunMap))

	fmt.Printf("Stat Group Definitions... ")
	mindefStatGrpMap := make(map[int64]models.MiniDestinyStatGroupDefinition)
	for _, e := range manifest.Manifest[18].DestinyStatGroupDefinition {
		var mindefStatGrp models.MiniDestinyStatGroupDefinition

		mindefStatGrp.MaximumValue = e.MaximumValue

		mindefStatGrpMap[e.Hash] = mindefStatGrp
	}
	miniMani["DestinyStatGroupDefinition"] = mindefStatGrpMap
	fmt.Printf("Done: %d\n", len(mindefStatGrpMap))

	fmt.Printf("Special Event Definitions... ")
	mindefSpEvtMap := make(map[int64]models.MiniDestinySpecialEventDefinition)
	for _, e := range manifest.Manifest[19].DestinySpecialEventDefinition {
		var mindefSpEvt models.MiniDestinySpecialEventDefinition

		mindefSpEvt.Title = e.Title
		mindefSpEvt.Subtitle = e.Subtitle
		mindefSpEvt.Description = e.Description
		mindefSpEvt.Link = e.Link
		mindefSpEvt.Icon = e.Icon
		mindefSpEvt.BackgroundImageWeb = e.BackgroundImageWeb
		mindefSpEvt.BackgroundImageMobile = e.BackgroundImageMobile
		mindefSpEvt.ProgressionHash = e.ProgressionHash
		mindefSpEvt.VendorHash = e.VendorHash
		mindefSpEvt.FactionHash = e.FactionHash
		mindefSpEvt.ActiveUnlockValueHash = e.ActiveUnlockValueHash
		mindefSpEvt.PlaylistActivityHash = e.PlaylistActivityHash
		mindefSpEvt.UnlockEventHash = e.UnlockEventHash
		mindefSpEvt.BountyHashes = e.BountyHashes
		mindefSpEvt.QuestHashes = e.QuestHashes

		mindefSpEvtMap[e.Hash] = mindefSpEvt
	}
	miniMani["DestinySpecialEventDefinition"] = mindefSpEvtMap
	fmt.Printf("Done: %d\n", len(mindefSpEvtMap))

	fmt.Printf("Faction Definitions... ")
	mindefFactionMap := make(map[int64]models.MiniDestinyFactionDefinition)
	for _, e := range manifest.Manifest[20].DestinyFactionDefinition {
		var mindefFaction models.MiniDestinyFactionDefinition

		mindefFaction.FactionName = e.FactionName
		mindefFaction.FactionDescription = e.FactionDescription
		mindefFaction.FactionIcon = e.FactionIcon
		mindefFaction.ProgressionHash = e.ProgressionHash

		mindefFactionMap[e.Hash] = mindefFaction
	}
	miniMani["DestinyFactionDefinition"] = mindefFactionMap
	fmt.Printf("Done: %d\n", len(mindefFactionMap))

	fmt.Printf("Vendor Category Definitions... ")
	mindefVenCatMap := make(map[int64]models.MiniDestinyVendorCategoryDefinition)
	for _, e := range manifest.Manifest[21].DestinyVendorCategoryDefinition {
		var mindefVenCat models.MiniDestinyVendorCategoryDefinition

		mindefVenCat.CategoryName = e.CategoryName
		mindefVenCat.MobileBannerPath = e.MobileBannerPath
		mindefVenCat.Identifier = e.Identifier

		mindefVenCatMap[e.Hash] = mindefVenCat
	}
	miniMani["DestinyVendorCategoryDefinition"] = mindefVenCatMap
	fmt.Printf("Done: %d\n", len(mindefVenCatMap))

	fmt.Printf("Enemy Race Definitions... ")
	mderdMap := make(map[int64]models.MiniDestinyEnemyRaceDefinition)
	for _, e := range manifest.Manifest[22].DestinyEnemyRaceDefinition {
		var mderd models.MiniDestinyEnemyRaceDefinition

		mderd.RaceName = e.RaceName
		mderd.Description = e.Description
		mderd.IconPath = e.IconPath

		mderdMap[e.Hash] = mderd
	}
	miniMani["DestinyEnemyRaceDefinition"] = mderdMap
	fmt.Printf("Done: %d\n", len(mderdMap))

	fmt.Printf("Scripted Skull Definitions... ")
	mdssdMap := make(map[int64]models.MiniDestinyScriptedSkullDefinition)
	for _, e := range manifest.Manifest[23].DestinyScriptedSkullDefinition {
		var mdssd models.MiniDestinyScriptedSkullDefinition

		mdssd.SkullName = e.SkullName
		mdssd.Description = e.Description
		mdssd.IconPath = e.IconPath

		mdssdMap[e.Hash] = mdssd
	}
	miniMani["DestinyScriptedSkullDefinition"] = mdssdMap
	fmt.Printf("Done: %d\n", len(mdssdMap))

	fmt.Printf("Triumph Set Definitions... ")
	mdtsMap := make(map[int64]models.MiniDestinyTriumphSetDefinition)
	for _, e := range manifest.Manifest[24].DestinyTriumphSetDefinition {
		var mdts models.MiniDestinyTriumphSetDefinition

		mdts.Title = e.Title
		mdts.IconPath = e.IconPath
		mdts.IncompleteSubtitle = e.IncompleteSubtitle
		mdts.IncompleteDetails = e.IncompleteDetails
		mdts.CompletedSubtitle = e.CompletedSubtitle
		mdts.CompletedDetails = e.CompletedDetails
		mdts.LockedSubtitle = e.LockedSubtitle
		mdts.LockedDetails = e.LockedDetails
		mdts.LockdownDate = e.LockdownDate
		mdts.LockdownUnlockHash = e.LockdownUnlockHash

		mdtsMap[e.Hash] = mdts
	}
	miniMani["DestinyTriumphSetDefinition"] = mdtsMap
	fmt.Printf("Done: %d\n", len(mdtsMap))

	fmt.Printf("Item Category Definitions... ")
	mdicdMap := make(map[int64]models.MiniDestinyItemCategoryDefinition)
	for _, e := range manifest.Manifest[25].DestinyItemCategoryDefinition {
		var mdicd models.MiniDestinyItemCategoryDefinition

		mdicd.Visible = e.Visible
		mdicd.Title = e.Title
		mdicd.ShortTitle = e.ShortTitle
		mdicd.GrantDestinyItemType = e.GrantDestinyItemType
		mdicd.GrantDestinySubType = e.GrantDestinySubType
		mdicd.GrantDestinyClass = e.GrantDestinyClass

		mdicdMap[e.Hash] = mdicd
	}
	miniMani["DestinyItemCategoryDefinition"] = mdicdMap
	fmt.Printf("Done: %d\n", len(mdicdMap))

	fmt.Printf("Reward Source Definitions... ")
	mdrsdMap := make(map[int64]models.MiniDestinyRewardSourceDefinition)
	for _, e := range manifest.Manifest[26].DestinyRewardSourceDefinition {
		var mdrsd models.MiniDestinyRewardSourceDefinition

		mdrsd.Category = e.Category
		mdrsd.SourceName = e.SourceName
		mdrsd.Description = e.Description
		mdrsd.Icon = e.Icon

		mdrsdMap[e.Hash] = mdrsd
	}
	miniMani["DestinyRewardSourceDefinition"] = mdrsdMap
	fmt.Printf("Done: %d\n", len(mdrsdMap))

	fmt.Printf("Objective Definitions... ")
	mdodMap := make(map[int64]models.MiniDestinyObjectiveDefinition)
	for _, e := range manifest.Manifest[27].DestinyObjectiveDefinition {
		var mdod models.MiniDestinyObjectiveDefinition

		mdod.DisplayDescription = e.DisplayDescription
		mdod.UnlockValueHash = e.UnlockValueHash
		mdod.VendorHash = e.VendorHash
		mdod.VendorCategoryHash = e.VendorCategoryHash
		mdod.LocationHash = e.LocationHash

		mdodMap[e.Hash] = mdod
	}
	miniMani["DestinyObjectiveDefinition"] = mdodMap
	fmt.Printf("Done: %d\n", len(mdodMap))

	fmt.Printf("Damage Type Definitions... ")
	mddtdMap := make(map[int64]models.MiniDestinyDamageTypeDefinition)
	for _, e := range manifest.Manifest[28].DestinyDamageTypeDefinition {
		var mddtd models.MiniDestinyDamageTypeDefinition

		mddtd.DamageTypeName = e.DamageTypeName
		mddtd.IconPath = e.IconPath
		mddtd.TransparentIconPath = e.TransparentIconPath
		mddtd.ShowIcon = e.ShowIcon
		mddtd.EnumValue = e.EnumValue

		mddtdMap[e.Hash] = mddtd
	}
	miniMani["DestinyDamageTypeDefinition"] = mddtdMap
	fmt.Printf("Done: %d\n", len(mddtdMap))

	fmt.Printf("Combatant Definitions... ")
	mdcombatdMap := make(map[int64]models.MiniDestinyCombatantDefinition)
	for _, e := range manifest.Manifest[29].DestinyCombatantDefinition {
		var mdcombatd models.MiniDestinyCombatantDefinition

		mdcombatd.Icon = e.Icon
		mdcombatd.CombatantName = e.CombatantName
		mdcombatd.Description = e.Description
		mdcombatd.Image = e.Image

		mdcombatdMap[e.Hash] = mdcombatd
	}
	miniMani["DestinyCombatantDefinition"] = mdcombatdMap
	fmt.Printf("Done: %d\n", len(mdcombatdMap))

	fmt.Printf("Activity Category Definitions... ")
	mdacdMap := make(map[int64]models.MiniDestinyActivityCategoryDefinition)
	for _, e := range manifest.Manifest[30].DestinyActivityCategoryDefinition {
		var mdacd models.MiniDestinyActivityCategoryDefinition

		mdacd.Title = e.Title
		mdacd.HelpTitle = e.HelpTitle
		mdacd.Description = e.Description
		mdacd.Image = e.Image
		mdacd.ParentHashes = e.ParentHashes

		mdacdMap[e.Hash] = mdacd
	}
	miniMani["DestinyActivityCategoryDefinition"] = mdacdMap
	fmt.Printf("Done: %d\n", len(mdacdMap))

	// Continue Here
	//
	//
	//
	//
	//
	//
	//
	//
	//

	fmt.Printf("Record Definitions... ")
	mdrecordMap := make(map[int64]models.MiniDestinyRecordDefinition)
	for _, e := range manifest.Manifest[31].DestinyRecordDefinition {
		var mdrecord models.MiniDestinyRecordDefinition

		mdrecord.DisplayName = e.DisplayName
		mdrecord.Description = e.Description
		mdrecord.Icon = e.Icon
		mdrecord.RecordValueUIStyle = e.RecordValueUIStyle

		mdrecordMap[e.Hash] = mdrecord
	}
	miniMani["DestinyRecordDefinition"] = mdrecordMap
	fmt.Printf("Done: %d\n", len(mdrecordMap))

	fmt.Printf("Record Book Definitions... ")
	mdrecordbookMap := make(map[int64]models.MiniDestinyRecordBookDefinition)
	for _, e := range manifest.Manifest[32].DestinyRecordBookDefinition {
		var mdrecordbook models.MiniDestinyRecordBookDefinition

		mdrecordbook.DisplayName = e.DisplayName
		mdrecordbook.DisplayDescription = e.DisplayDescription
		mdrecordbook.UnavailableReason = e.UnavailableReason
		mdrecordbook.ProgressionHash = e.ProgressionHash
		mdrecordbook.RecordCount = e.RecordCount
		mdrecordbook.Icon = e.Icon
		mdrecordbook.ItemHash = e.ItemHash

		mdrecordbookMap[e.Hash] = mdrecordbook
	}
	miniMani["DestinyRecordBookDefinition"] = mdrecordbookMap
	fmt.Printf("Done: %d\n", len(mdrecordbookMap))

	fmt.Printf("Bond Definitions... ")
	mdbondMap := make(map[int64]models.MiniDestinyBondDefinition)
	for _, e := range manifest.Manifest[33].DestinyBondDefinition {
		var mdbond models.MiniDestinyBondDefinition

		mdbond.DisplayIcon = e.DisplayIcon
		mdbond.ProvidedUnlockHash = e.ProvidedUnlockHash
		mdbond.ProvidedUnlockValueHash = e.ProvidedUnlockValueHash
		mdbond.ShowInAdvisor = e.ShowInAdvisor

		mdbondMap[e.Hash] = mdbond
	}
	miniMani["DestinyBondDefinition"] = mdbondMap
	fmt.Printf("Done: %d\n", len(mdbondMap))

	fmt.Printf("Location Definitions... ")
	mdlocMap := make(map[int64]models.MiniDestinyLocationDefinition)
	for _, e := range manifest.Manifest[34].DestinyLocationDefinition {
		var mdloc models.MiniDestinyLocationDefinition

		mdloc.LocationHash = e.LocationHash

		mdlocMap[e.Hash] = mdloc
	}
	miniMani["DestinyLocationDefinition"] = mdlocMap
	fmt.Printf("Done: %d\n", len(mdlocMap))

	fmt.Printf("Grimoire Definitions... ")
	mdgrimMap := make(map[int64]models.MiniDestinyGrimoireDefinition)
	for i, _ := range manifest.Manifest[35].DestinyGrimoireDefinition {
		var mdgrim models.MiniDestinyGrimoireDefinition

		// Nothing for now

		mdgrimMap[int64(i)] = mdgrim
	}
	miniMani["DestinyGrimoireDefinition"] = mdgrimMap
	fmt.Printf("Done: %d\n", len(mdgrimMap))

	fmt.Printf("Grimoire Card Definitions... ")
	mdgrimcardMap := make(map[int64]models.MiniDestinyGrimoireCardDefinition)
	for _, e := range manifest.Manifest[36].DestinyGrimoireCardDefinition {
		var mdgrimcard models.MiniDestinyGrimoireCardDefinition

		mdgrimcard.CardID = e.CardID
		mdgrimcard.CardName = e.CardName
		mdgrimcard.CardIntro = e.CardIntro
		mdgrimcard.CardDescription = e.CardDescription
		mdgrimcard.CardLabel = e.CardLabel
		mdgrimcard.UnlockHowToText = e.UnlockHowToText
		mdgrimcard.UnlockFlagHash = e.UnlockFlagHash
		mdgrimcard.Points = e.Points

		mdgrimcardMap[int64(e.CardID)] = mdgrimcard
	}
	miniMani["DestinyGrimoireCardDefinition"] = mdgrimcardMap
	fmt.Printf("Done: %d\n", len(mdgrimcardMap))

	// Convert the new manifest to .json and write the file
	// b, _ := json.Marshal(miniMani)
	b, _ := json.MarshalIndent(miniMani, "", "    ")
	ioutil.WriteFile("MiniMani.json", b, 0644)
}
