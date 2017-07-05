package main

import (
	"encoding/json"
	"fmt"
	"github.com/leecombs/destinymanifesttrimmer/models"
	"io/ioutil"
	"log"
	"strconv"
	"sync"
)

// Attempt to write a given file to disk, with goroutine support
func writeToFile(filename string, data []byte, wg *sync.WaitGroup) {
	fmt.Println("Writing File:", filename)
	defer wg.Done()

	err := ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		log.Fatal("Write error:", err)
	}

	fmt.Println("Done writing file", filename)
}

func main() {
	fmt.Println("Running trimmer")
	fmt.Println("Reading Manifest.json...")

	// Read and load the local manifest file
	var manifest models.DestinyManifest
	file, fileErr := ioutil.ReadFile("DestinyManifest.json")
	if fileErr != nil {
		log.Fatal("fileErr", fileErr)
		return
	}
	json.Unmarshal(file, &manifest)

	// Initialize the output var
	miniMani := make(map[string]interface{})

	// Set up goroutines to read through the manifest data, then minify specific entries individually
	maniWaitGroup := sync.WaitGroup{}
	maniWaitGroup.Add(len(manifest.Manifest))
	for i, e := range manifest.Manifest {
		switch i {
		case 0: // Activity
			maniEntry := e.DestinyActivityDefinition
			go func() {
				fmt.Println("Building Activity Definitions")
				defer maniWaitGroup.Done()

				miniActMap := make(map[int64]models.MiniDestinyActivityDefinition)
				for _, e := range maniEntry {
					var miniAct models.MiniDestinyActivityDefinition

					miniAct.ActivityName = e.ActivityName
					miniAct.ActivityDescription = e.ActivityDescription
					miniAct.Icon = e.Icon
					miniAct.ActivityLevel = e.ActivityLevel
					miniAct.DestinationHash = int64(e.DestinationHash)
					miniAct.PlaceHash = int64(e.PlaceHash)
					miniAct.ActivityTypeHash = int64(e.ActivityTypeHash)
					miniAct.Rewards = e.Rewards
					miniAct.Skulls = e.Skulls

					miniActMap[e.Hash] = miniAct
				}
				miniMani["MiniActivityDefinition"] = miniActMap
				fmt.Println("Done Activity Definitions, Count:", len(miniActMap))
			}()
		case 1: // Activity Type
			maniEntry := e.DestinyActivityTypeDefinition
			go func() {
				fmt.Println("Building Activity Type Definitions")
				defer maniWaitGroup.Done()

				miniActTypeMap := make(map[int64]models.MiniDestinyActivityTypeDefinition)
				for _, e := range maniEntry {
					var miniActType models.MiniDestinyActivityTypeDefinition

					miniActType.Identifier = e.Identifier
					miniActType.ActivityTypeName = e.ActivityTypeName
					miniActType.ActivityTypeDescription = e.ActivityTypeDescription
					miniActType.Icon = e.Icon

					miniActTypeMap[e.Hash] = miniActType
				}
				miniMani["MiniActivityTypeDefinition"] = miniActTypeMap
				fmt.Println("Done Activity Type Definitions, Count:", len(miniActTypeMap))
			}()
		case 2: // Class
			maniEntry := e.DestinyClassDefinition
			go func() {
				fmt.Println("Building Class Definitions")
				defer maniWaitGroup.Done()

				miniClassMap := make(map[int64]models.MiniDestinyClassDefinition)
				for _, e := range maniEntry {
					var miniClass models.MiniDestinyClassDefinition

					miniClass.ClassName = e.ClassName
					miniClass.ClassType = e.ClassType

					miniClassMap[e.Hash] = miniClass
				}
				miniMani["MiniClassDefinition"] = miniClassMap
				fmt.Println("Done Class Definitions, Count:", len(miniClassMap))
			}()
		case 3: // Gender
			maniEntry := e.DestinyGenderDefinition
			go func() {
				fmt.Println("Building Gender Definitions")
				defer maniWaitGroup.Done()

				miniGenderMap := make(map[int64]models.MiniDestinyGenderDefinition)
				for _, e := range maniEntry {
					var miniGender models.MiniDestinyGenderDefinition

					miniGender.GenderName = e.GenderName
					miniGender.GenderType = e.GenderType

					miniGenderMap[e.Hash] = miniGender
				}
				miniMani["MiniGenderDefinition"] = miniGenderMap
				fmt.Println("Done Gender Definitions, Count:", len(miniGenderMap))
			}()
		case 4: // Inv Bucket
			go func() {
				defer maniWaitGroup.Done()
			}()
		case 5: // Inv Item
			maniEntry := e.DestinyInventoryItemDefinition
			go func() {
				fmt.Println("Building Item Definitions")
				defer maniWaitGroup.Done()

				miniItemMap := make(map[int64]models.MiniDestinyInventoryItemDefinition)
				for _, e := range maniEntry {
					var miniItem models.MiniDestinyInventoryItemDefinition

					miniItem.ItemName = e.ItemName
					miniItem.ItemDescription = e.ItemDescription
					miniItem.HasIcon = e.HasIcon
					miniItem.Icon = e.Icon
					miniItem.SecondaryIcon = e.SecondaryIcon
					miniItem.TierType = e.TierType
					miniItem.TierTypeName = e.TierTypeName
					miniItem.ItemType = e.ItemType
					miniItem.ItemTypeName = e.ItemTypeName
					miniItem.SpecialItemType = e.SpecialItemType
					miniItem.ClassType = e.ClassType
					miniItem.BucketTypeHash = int64(e.BucketTypeHash)
					miniItem.PrimaryBaseStatHash = int64(e.PrimaryBaseStatHash)
					miniItem.Stats = e.Stats
					miniItem.Exclusive = e.Exclusive

					miniItemMap[e.Hash] = miniItem
				}
				miniMani["MiniItemDefinition"] = miniItemMap
				fmt.Println("Done Item Definitions, Count:", len(miniItemMap))
			}()
		case 6: // Progression
			go func() {
				defer maniWaitGroup.Done()
			}()
		case 7: // Race
			maniEntry := e.DestinyRaceDefinition
			go func() {
				fmt.Println("Building Race Definitions")
				defer maniWaitGroup.Done()

				miniRaceMap := make(map[int64]models.MiniDestinyRaceDefinition)
				for _, e := range maniEntry {
					var miniRace models.MiniDestinyRaceDefinition

					miniRace.RaceType = e.RaceType
					miniRace.RaceName = e.RaceName
					miniRace.RaceNameMale = e.RaceNameMale
					miniRace.RaceNameFemale = e.RaceNameFemale
					miniRace.RaceDescription = e.RaceDescription

					miniRaceMap[e.Hash] = miniRace
				}
				miniMani["MiniRaceDefinition"] = miniRaceMap
				fmt.Println("Done Race Definitions, Count:", len(miniRaceMap))
			}()
		case 8: // Talent Grid
			go func() {
				defer maniWaitGroup.Done()
			}()
		case 9: // Unlock Flag
			go func() {
				defer maniWaitGroup.Done()
			}()
		case 10: // Vendor
			go func() {
				defer maniWaitGroup.Done()
			}()
		case 11: // Historical Stats
			go func() {
				defer maniWaitGroup.Done()
			}()
		case 12: // Director Book
			go func() {
				defer maniWaitGroup.Done()
			}()
		case 13: // Stat
			maniEntry := e.DestinyStatDefinition
			go func() {
				fmt.Println("Building Stat Definitions")
				defer maniWaitGroup.Done()

				miniStatMap := make(map[int64]models.MiniDestinyStatDefinition)
				for _, e := range maniEntry {
					var miniStat models.MiniDestinyStatDefinition

					miniStat.StatName = e.StatName
					miniStat.StatDescription = e.StatDescription
					miniStat.Icon = e.Icon

					miniStatMap[e.Hash] = miniStat
				}
				miniMani["MiniStatDefinition"] = miniStatMap
				fmt.Println("Done Stat Definitions, Count:", len(miniStatMap))
			}()
		case 14: // Sandbox Perk
			go func() {
				defer maniWaitGroup.Done()
			}()
		case 15: // Destination
			go func() {
				defer maniWaitGroup.Done()
			}()
		case 16: // Place
			maniEntry := e.DestinyPlaceDefinition
			go func() {
				fmt.Println("Building Place Definitions")
				defer maniWaitGroup.Done()

				miniPlaceMap := make(map[int64]models.MiniDestinyPlaceDefinition)
				for _, e := range maniEntry {
					var miniPlace models.MiniDestinyPlaceDefinition

					miniPlace.PlaceName = e.PlaceName
					miniPlace.PlaceDescription = e.PlaceDescription
					miniPlace.Icon = e.Icon

					miniPlaceMap[e.Hash] = miniPlace
				}
				miniMani["MiniPlaceDefinition"] = miniPlaceMap
				fmt.Println("Done Place Definitions, Count:", len(miniPlaceMap))
			}()
		case 17: // Activity Bundle
			go func() {
				defer maniWaitGroup.Done()
			}()
		case 18: // Stat Group
			go func() {
				defer maniWaitGroup.Done()
			}()
		case 19: // Special Event
			go func() {
				defer maniWaitGroup.Done()
			}()
		case 20: // Faction
			maniEntry := e.DestinyFactionDefinition
			go func() {
				fmt.Println("Building Faction Definitions")
				defer maniWaitGroup.Done()

				miniFactMap := make(map[int64]models.MiniDestinyFactionDefinition)
				for _, e := range maniEntry {
					var miniFact models.MiniDestinyFactionDefinition

					miniFact.FactionName = e.FactionName
					miniFact.FactionDescription = e.FactionDescription
					miniFact.FactionIcon = e.FactionIcon
					miniFact.ProgressionHash = e.ProgressionHash

					miniFactMap[e.Hash] = miniFact
				}
				miniMani["MiniFactionDefinition"] = miniFactMap
				fmt.Println("Done Faction Definitions, Count:", len(miniFactMap))
			}()
		case 21: // Vendor Category
			go func() {
				defer maniWaitGroup.Done()
			}()
		case 22: // Enemy Race
			maniEntry := e.DestinyEnemyRaceDefinition
			go func() {
				fmt.Println("Building Enemy Race Definitions")
				defer maniWaitGroup.Done()

				miniEnemyRaceMap := make(map[int64]models.MiniDestinyEnemyRaceDefinition)
				for _, e := range maniEntry {
					var miniEnemyRace models.MiniDestinyEnemyRaceDefinition

					miniEnemyRace.RaceName = e.RaceName
					miniEnemyRace.Description = e.Description
					miniEnemyRace.IconPath = e.IconPath

					miniEnemyRaceMap[e.Hash] = miniEnemyRace
				}
				miniMani["MiniEnemyRaceDefinition"] = miniEnemyRaceMap
				fmt.Println("Done Enemy Race Definitions, Count:", len(miniEnemyRaceMap))
			}()
		case 23: // Scripted Skill
			go func() {
				defer maniWaitGroup.Done()
			}()
		case 24: // Triumph Set
			go func() {
				defer maniWaitGroup.Done()
			}()
		case 25: // Item Category
			go func() {
				defer maniWaitGroup.Done()
			}()
		case 26: // Reward Source
			go func() {
				defer maniWaitGroup.Done()
			}()
		case 27: // Objective
			go func() {
				defer maniWaitGroup.Done()
			}()
		case 28: // Damage Type
			maniEntry := e.DestinyDamageTypeDefinition
			go func() {
				fmt.Println("Building Damage Type Definitions")
				defer maniWaitGroup.Done()

				miniDamTypeMap := make(map[int64]models.MiniDestinyDamageTypeDefinition)
				for _, e := range maniEntry {
					var miniDamType models.MiniDestinyDamageTypeDefinition

					miniDamType.DamageTypeName = e.DamageTypeName
					miniDamType.IconPath = e.IconPath
					miniDamType.TransparentIconPath = e.TransparentIconPath
					miniDamType.ShowIcon = e.ShowIcon
					miniDamType.EnumValue = e.EnumValue

					miniDamTypeMap[e.Hash] = miniDamType
				}
				miniMani["MiniDamageDefinition"] = miniDamTypeMap
				fmt.Println("Done Damage Type Definitions, Count:", len(miniDamTypeMap))
			}()
		case 29: // Combatant
			maniEntry := e.DestinyCombatantDefinition
			go func() {
				fmt.Println("Building Combatant Definitions")
				defer maniWaitGroup.Done()

				miniCombatantMap := make(map[int64]models.MiniDestinyCombatantDefinition)
				for _, e := range maniEntry {
					var miniCombatant models.MiniDestinyCombatantDefinition

					miniCombatant.Icon = e.Icon
					miniCombatant.CombatantName = e.CombatantName
					miniCombatant.Description = e.Description
					miniCombatant.Image = e.Image

					miniCombatantMap[e.Hash] = miniCombatant
				}
				miniMani["MiniCombatantDefinition"] = miniCombatantMap
				fmt.Println("Done Combatant Definitions, Count:", len(miniCombatantMap))
			}()
		case 30: // Activity Category
			go func() {
				defer maniWaitGroup.Done()
			}()
		case 31: // Record
			go func() {
				defer maniWaitGroup.Done()
			}()
		case 32: // Record Book
			go func() {
				defer maniWaitGroup.Done()
			}()
		case 33: // Bond
			go func() {
				defer maniWaitGroup.Done()
			}()
		case 34: // Location
			go func() {
				defer maniWaitGroup.Done()
			}()
		case 35: // Grimoire
			go func() {
				defer maniWaitGroup.Done()
			}()
		case 36: // Grimoire Card
			go func() {
				defer maniWaitGroup.Done()
			}()
		default:
			fmt.Println("Manifest Entry Miss", i)
			maniWaitGroup.Done()
		}
	}
	maniWaitGroup.Wait()
	fmt.Println("Done reading manifest")
	fmt.Println("Writing files...")

	// Build the output .json file
	b, _ := json.MarshalIndent(miniMani, "", "    ")

	// Setup go routines to write files individually
	// Really only done as practice
	// TODO: Add support for writing individual definitions to their own files
	numFiles := 1
	fileWaitGroup := sync.WaitGroup{}
	fileWaitGroup.Add(numFiles)
	for i := 0; i < numFiles; i++ {
		filename := "MiniMani" + strconv.Itoa(i) + ".json"
		go writeToFile(filename, b, &fileWaitGroup)
	}
	fileWaitGroup.Wait()

	fmt.Println("Done writing files. Enter to continue")
	fmt.Scanln()
}
