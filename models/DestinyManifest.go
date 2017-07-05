package models

import (
	"time"
)

type ActivityDefinition struct {
	ActivityHash        int64         `json:"activityHash"`
	ActivityName        string        `json:"activityName"`
	ActivityDescription string        `json:"activityDescription"`
	Icon                string        `json:"icon"`
	ReleaseIcon         string        `json:"releaseIcon"`
	ReleaseTime         int           `json:"releaseTime"`
	ActivityLevel       int           `json:"activityLevel"`
	CompletionFlagHash  int64         `json:"completionFlagHash"`
	ActivityPower       float64       `json:"activityPower"`
	MinParty            int           `json:"minParty"`
	MaxParty            int           `json:"maxParty"`
	MaxPlayers          int           `json:"maxPlayers"`
	DestinationHash     int64         `json:"destinationHash"`
	PlaceHash           int64         `json:"placeHash"`
	ActivityTypeHash    int64         `json:"activityTypeHash"`
	Tier                int           `json:"tier"`
	PgcrImage           string        `json:"pgcrImage"`
	Rewards             []interface{} `json:"rewards"`
	Skulls              []interface{} `json:"skulls"`
	IsPlaylist          bool          `json:"isPlaylist"`
	IsMatchmade         bool          `json:"isMatchmade"`
	Hash                int64         `json:"hash"`
	Index               int           `json:"index"`
}

type DestinyManifest struct {
	Manifest []struct {
		// 0
		DestinyActivityDefinition []ActivityDefinition `json:"DestinyActivityDefinition"`

		// 1
		DestinyActivityTypeDefinition []struct {
			ActivityTypeHash                       int64  `json:"activityTypeHash"`
			Identifier                             string `json:"identifier"`
			ActivityTypeName                       string `json:"activityTypeName"`
			ActivityTypeDescription                string `json:"activityTypeDescription"`
			Icon                                   string `json:"icon"`
			ActiveBackgroundVirtualPath            string `json:"activeBackgroundVirtualPath"`
			CompletedBackgroundVirtualPath         string `json:"completedBackgroundVirtualPath"`
			HiddenOverrideVirtualPath              string `json:"hiddenOverrideVirtualPath"`
			TooltipBackgroundVirtualPath           string `json:"tooltipBackgroundVirtualPath"`
			EnlargedActiveBackgroundVirtualPath    string `json:"enlargedActiveBackgroundVirtualPath"`
			EnlargedCompletedBackgroundVirtualPath string `json:"enlargedCompletedBackgroundVirtualPath"`
			EnlargedHiddenOverrideVirtualPath      string `json:"enlargedHiddenOverrideVirtualPath"`
			EnlargedTooltipBackgroundVirtualPath   string `json:"enlargedTooltipBackgroundVirtualPath"`
			Order                                  int    `json:"order"`
			Hash                                   int64  `json:"hash"`
			Index                                  int    `json:"index"`
			StatGroup                              string `json:"statGroup"`
			FriendlyURLID                          string `json:"friendlyUrlId"`
		} `json:"DestinyActivityTypeDefinition"`

		// 2
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

		// 3
		DestinyGenderDefinition []struct {
			GenderHash        int64  `json:"genderHash"`
			GenderType        int    `json:"genderType"`
			GenderName        string `json:"genderName"`
			GenderDescription string `json:"genderDescription"`
			Hash              int64  `json:"hash"`
			Index             int    `json:"index"`
		} `json:"DestinyGenderDefinition"`

		// 4
		DestinyInventoryBucketDefinition []struct {
			BucketHash             int64  `json:"bucketHash"`
			BucketName             string `json:"bucketName,omitempty"`
			BucketDescription      string `json:"bucketDescription,omitempty"`
			Scope                  int    `json:"scope"`
			Category               int    `json:"category"`
			BucketOrder            int    `json:"bucketOrder"`
			BucketIdentifier       string `json:"bucketIdentifier,omitempty"`
			ItemCount              int    `json:"itemCount"`
			Location               int    `json:"location"`
			HasTransferDestination bool   `json:"hasTransferDestination"`
			Enabled                bool   `json:"enabled"`
			Hash                   int64  `json:"hash"`
			Index                  int    `json:"index"`
		} `json:"DestinyInventoryBucketDefinition"`

		// 5  -- Note: Unsure if this is comprehensive. Generated from a handful of entries
		DestinyInventoryItemDefinition []struct {
			ItemHash            int64  `json:"itemHash"`
			ItemName            string `json:"itemName"`
			ItemDescription     string `json:"itemDescription"`
			Icon                string `json:"icon"`
			HasIcon             bool   `json:"hasIcon"`
			SecondaryIcon       string `json:"secondaryIcon"`
			ActionName          string `json:"actionName,omitempty"`
			HasAction           bool   `json:"hasAction"`
			DeleteOnAction      bool   `json:"deleteOnAction"`
			TierTypeName        string `json:"tierTypeName"`
			TierType            int    `json:"tierType"`
			ItemTypeName        string `json:"itemTypeName"`
			BucketTypeHash      int64  `json:"bucketTypeHash"`
			PrimaryBaseStatHash int64  `json:"primaryBaseStatHash"`
			Stats               struct {
				Num2391494160 struct {
					StatHash int64 `json:"statHash"`
					Value    int   `json:"value"`
					Minimum  int   `json:"minimum"`
					Maximum  int   `json:"maximum"`
				} `json:"2391494160"`
				Num3897883278 struct {
					StatHash int64 `json:"statHash"`
					Value    int   `json:"value"`
					Minimum  int   `json:"minimum"`
					Maximum  int   `json:"maximum"`
				} `json:"3897883278"`
			} `json:"stats"`
			PerkHashes      []interface{} `json:"perkHashes"`
			SpecialItemType int           `json:"specialItemType"`
			TalentGridHash  int64         `json:"talentGridHash"`
			EquippingBlock  struct {
				WeaponSandboxPatternIndex int `json:"weaponSandboxPatternIndex"`
				GearArtArrangementIndex   int `json:"gearArtArrangementIndex"`
				DefaultDyes               []struct {
					ChannelHash int64 `json:"channelHash"`
					DyeHash     int64 `json:"dyeHash"`
				} `json:"defaultDyes"`
				LockedDyes          []interface{} `json:"lockedDyes"`
				CustomDyes          []interface{} `json:"customDyes"`
				CustomDyeExpression struct {
					Steps []struct {
						StepOperator int   `json:"stepOperator"`
						FlagHash     int64 `json:"flagHash"`
						ValueHash    int64 `json:"valueHash"`
						Value        int   `json:"value"`
					} `json:"steps"`
				} `json:"customDyeExpression"`
				WeaponPatternHash int64 `json:"weaponPatternHash"`
				Arrangements      []struct {
					ClassHash               int64 `json:"classHash"`
					GearArtArrangementIndex int   `json:"gearArtArrangementIndex"`
				} `json:"arrangements"`
			} `json:"equippingBlock,omitempty"`
			HasGeometry    bool  `json:"hasGeometry"`
			StatGroupHash  int64 `json:"statGroupHash"`
			ItemLevels     []int `json:"itemLevels"`
			QualityLevel   int   `json:"qualityLevel"`
			Equippable     bool  `json:"equippable"`
			Instanced      bool  `json:"instanced"`
			RewardItemHash int64 `json:"rewardItemHash"`
			Values         struct {
			} `json:"values"`
			ItemType                     int           `json:"itemType"`
			ItemSubType                  int           `json:"itemSubType"`
			ClassType                    int           `json:"classType"`
			ItemCategoryHashes           []int         `json:"itemCategoryHashes"`
			SourceHashes                 []interface{} `json:"sourceHashes"`
			NonTransferrable             bool          `json:"nonTransferrable"`
			Exclusive                    int           `json:"exclusive"`
			MaxStackSize                 int           `json:"maxStackSize"`
			ItemIndex                    int           `json:"itemIndex"`
			SetItemHashes                []interface{} `json:"setItemHashes"`
			TooltipStyle                 string        `json:"tooltipStyle"`
			QuestlineItemHash            int64         `json:"questlineItemHash"`
			NeedsFullCompletion          bool          `json:"needsFullCompletion"`
			ObjectiveHashes              []interface{} `json:"objectiveHashes"`
			AllowActions                 bool          `json:"allowActions"`
			QuestTrackingUnlockValueHash int64         `json:"questTrackingUnlockValueHash"`
			BountyResetUnlockHash        int64         `json:"bountyResetUnlockHash"`
			UniquenessHash               int64         `json:"uniquenessHash"`
			ShowActiveNodesInTooltip     bool          `json:"showActiveNodesInTooltip"`
			Hash                         int64         `json:"hash"`
			Index                        int           `json:"index"`
			Sources                      []struct {
				ExpansionIndex   int `json:"expansionIndex"`
				Level            int `json:"level"`
				MinQuality       int `json:"minQuality"`
				MaxQuality       int `json:"maxQuality"`
				MinLevelRequired int `json:"minLevelRequired"`
				MaxLevelRequired int `json:"maxLevelRequired"`
				Exclusivity      int `json:"exclusivity"`
				ComputedStats    struct {
					Num144602215 struct {
						StatHash int64 `json:"statHash"`
						Value    int   `json:"value"`
						Minimum  int   `json:"minimum"`
						Maximum  int   `json:"maximum"`
					} `json:"144602215"`
					Num1735777505 struct {
						StatHash int64 `json:"statHash"`
						Value    int   `json:"value"`
						Minimum  int   `json:"minimum"`
						Maximum  int   `json:"maximum"`
					} `json:"1735777505"`
					Num2391494160 struct {
						StatHash int64 `json:"statHash"`
						Value    int   `json:"value"`
						Minimum  int   `json:"minimum"`
						Maximum  int   `json:"maximum"`
					} `json:"2391494160"`
					Num3897883278 struct {
						StatHash int64 `json:"statHash"`
						Value    int   `json:"value"`
						Minimum  int   `json:"minimum"`
						Maximum  int   `json:"maximum"`
					} `json:"3897883278"`
					Num4244567218 struct {
						StatHash int64 `json:"statHash"`
						Value    int   `json:"value"`
						Minimum  int   `json:"minimum"`
						Maximum  int   `json:"maximum"`
					} `json:"4244567218"`
				} `json:"computedStats"`
				SourceHashes []int `json:"sourceHashes"`
				SpawnIndexes []int `json:"spawnIndexes"`
			} `json:"sources,omitempty"`
		} `json:"DestinyInventoryItemDefinition"`

		// 6
		DestinyProgressionDefinition []struct {
			ProgressionHash int64  `json:"progressionHash"`
			Name            string `json:"name"`
			Scope           int    `json:"scope"`
			RepeatLastStep  bool   `json:"repeatLastStep"`
			Steps           []struct {
				ProgressTotal int           `json:"progressTotal"`
				RewardItems   []interface{} `json:"rewardItems"`
			} `json:"steps"`
			Visible     bool   `json:"visible"`
			Hash        int64  `json:"hash"`
			Index       int    `json:"index"`
			Icon        string `json:"icon,omitempty"`
			Identifier  string `json:"identifier,omitempty"`
			Description string `json:"description,omitempty"`
			Source      string `json:"source,omitempty"`
		} `json:"DestinyProgressionDefinition"`

		// 7
		DestinyRaceDefinition []struct {
			RaceHash        int64  `json:"raceHash"`
			RaceType        int    `json:"raceType"`
			RaceName        string `json:"raceName"`
			RaceNameMale    string `json:"raceNameMale"`
			RaceNameFemale  string `json:"raceNameFemale"`
			RaceDescription string `json:"raceDescription"`
			Hash            int64  `json:"hash"`
			Index           int    `json:"index"`
		} `json:"DestinyRaceDefinition"`

		// 8 -- Note: Unsure if this is comprehensive. Generated from a handful of entries
		DestinyTalentGridDefinition []struct {
			GridHash               int64         `json:"gridHash"`
			MaxGridLevel           int           `json:"maxGridLevel"`
			GridLevelPerColumn     int           `json:"gridLevelPerColumn"`
			ProgressionHash        int64         `json:"progressionHash"`
			Nodes                  []interface{} `json:"nodes"`
			CalcMaxGridLevel       int           `json:"calcMaxGridLevel"`
			CalcProgressToMaxLevel int           `json:"calcProgressToMaxLevel"`
			ExclusiveSets          []interface{} `json:"exclusiveSets"`
			IndependentNodeIndexes []interface{} `json:"independentNodeIndexes"`
			Hash                   int64         `json:"hash"`
			Index                  int           `json:"index"`
		} `json:"DestinyTalentGridDefinition"`

		// 9
		DestinyUnlockFlagDefinition []struct {
			FlagHash           int64  `json:"flagHash"`
			IsOffer            bool   `json:"isOffer"`
			UnlockType         int    `json:"unlockType"`
			Hash               int64  `json:"hash"`
			Index              int    `json:"index"`
			DisplayName        string `json:"displayName,omitempty"`
			DisplayDescription string `json:"displayDescription,omitempty"`
		} `json:"DestinyUnlockFlagDefinition"`

		// 10
		DestinyVendorDefinition []struct {
			Summary struct {
				VendorHash               int64         `json:"vendorHash"`
				VendorName               string        `json:"vendorName"`
				VendorDescription        string        `json:"vendorDescription"`
				VendorIcon               string        `json:"vendorIcon"`
				VendorOrder              int           `json:"vendorOrder"`
				FactionHash              int64         `json:"factionHash"`
				ResetIntervalMinutes     int           `json:"resetIntervalMinutes"`
				ResetOffsetMinutes       int           `json:"resetOffsetMinutes"`
				VendorIdentifier         string        `json:"vendorIdentifier"`
				PositionX                int           `json:"positionX"`
				PositionY                int           `json:"positionY"`
				TransitionNodeIdentifier string        `json:"transitionNodeIdentifier"`
				Visible                  bool          `json:"visible"`
				ProgressionHash          int64         `json:"progressionHash"`
				SellString               string        `json:"sellString"`
				BuyString                string        `json:"buyString"`
				VendorPortrait           string        `json:"vendorPortrait"`
				VendorBanner             string        `json:"vendorBanner"`
				UnlockFlagHashes         []interface{} `json:"unlockFlagHashes"`
				EnabledUnlockFlagHashes  []interface{} `json:"enabledUnlockFlagHashes"`
				MapSectionIdentifier     string        `json:"mapSectionIdentifier"`
				MapSectionName           string        `json:"mapSectionName"`
				MapSectionOrder          int           `json:"mapSectionOrder"`
				ShowOnMap                bool          `json:"showOnMap"`
				EventHash                int64         `json:"eventHash"`
				VendorCategoryHash       int64         `json:"vendorCategoryHash"`
				VendorCategoryHashes     []int         `json:"vendorCategoryHashes"`
				VendorSubcategoryHash    int64         `json:"vendorSubcategoryHash"`
				InhibitBuying            bool          `json:"inhibitBuying"`
			} `json:"summary"`
			AcceptedItems []interface{} `json:"acceptedItems"`
			Categories    []struct {
				CategoryHash            int64  `json:"categoryHash"`
				DisplayTitle            string `json:"displayTitle"`
				OverlayCurrencyItemHash int64  `json:"overlayCurrencyItemHash"`
				QuantityAvailable       int    `json:"quantityAvailable"`
				ShowUnavailableItems    bool   `json:"showUnavailableItems"`
				HideIfNoCurrency        bool   `json:"hideIfNoCurrency"`
				OverlayTitle            string `json:"overlayTitle"`
				OverlayDescription      string `json:"overlayDescription"`
				OverlayChoice           string `json:"overlayChoice"`
				OverlayIcon             string `json:"overlayIcon"`
				HasOverlay              bool   `json:"hasOverlay"`
				HideFromRegularPurchase bool   `json:"hideFromRegularPurchase"`
			} `json:"categories"`
			FailureStrings  []interface{} `json:"failureStrings"`
			UnlockValueHash int64         `json:"unlockValueHash"`
			Hash            int64         `json:"hash"`
			Index           int           `json:"index"`
		} `json:"DestinyVendorDefinition"`

		// 11
		DestinyHistoricalStatsDefinition []struct {
			StatID          string `json:"statId"`
			Group           int    `json:"group"`
			PeriodTypes     []int  `json:"periodTypes"`
			Modes           []int  `json:"modes"`
			Category        int    `json:"category"`
			StatName        string `json:"statName"`
			StatDescription string `json:"statDescription,omitempty"`
			UnitType        int    `json:"unitType"`
			UnitLabel       string `json:"unitLabel"`
			Weight          int    `json:"weight"`
			IconImage       string `json:"iconImage,omitempty"`
		} `json:"DestinyHistoricalStatsDefinition"`

		// 12
		DestinyDirectorBookDefinition []struct {
			BookHash        int64  `json:"bookHash"`
			BookName        string `json:"bookName,omitempty"`
			BookDescription string `json:"bookDescription"`
			BookNumber      string `json:"bookNumber,omitempty"`
			Nodes           []struct {
				NodeDefinitionHash   int64   `json:"nodeDefinitionHash"`
				StyleHash            int64   `json:"styleHash"`
				PositionX            int     `json:"positionX"`
				PositionY            int     `json:"positionY"`
				PositionZ            int     `json:"positionZ"`
				ActivityBundleHashes []int64 `json:"activityBundleHashes"`
				States               []struct {
					State int `json:"state"`
				} `json:"states"`
				UIModifier int `json:"uiModifier"`
			} `json:"nodes"`
			Connections         []interface{} `json:"connections"`
			Visible             bool          `json:"visible"`
			IsOverview          bool          `json:"isOverview"`
			IsDefaultExpression struct {
				Steps []interface{} `json:"steps"`
			} `json:"isDefaultExpression"`
			IsVisibleExpression struct {
				Steps []struct {
					StepOperator int   `json:"stepOperator"`
					FlagHash     int64 `json:"flagHash"`
					ValueHash    int64 `json:"valueHash"`
					Value        int   `json:"value"`
				} `json:"steps"`
			} `json:"isVisibleExpression"`
			DestinationHash int64 `json:"destinationHash"`
			Hash            int64 `json:"hash"`
			Index           int   `json:"index"`
		} `json:"DestinyDirectorBookDefinition"`

		// 13
		DestinyStatDefinition []struct {
			StatHash        int64  `json:"statHash"`
			StatName        string `json:"statName,omitempty"`
			StatDescription string `json:"statDescription,omitempty"`
			Icon            string `json:"icon"`
			StatIdentifier  string `json:"statIdentifier"`
			Interpolate     bool   `json:"interpolate"`
			Hash            int64  `json:"hash"`
			Index           int    `json:"index"`
		} `json:"DestinyStatDefinition"`

		// 14
		DestinySandboxPerkDefinition []struct {
			PerkHash           int64  `json:"perkHash"`
			DisplayName        string `json:"displayName,omitempty"`
			DisplayDescription string `json:"displayDescription,omitempty"`
			DisplayIcon        string `json:"displayIcon"`
			IsDisplayable      bool   `json:"isDisplayable"`
			Hash               int64  `json:"hash"`
			Index              int    `json:"index"`
			PerkGroups         struct {
				WeaponPerformance  int `json:"weaponPerformance"`
				ImpactEffects      int `json:"impactEffects"`
				GuardianAttributes int `json:"guardianAttributes"`
				LightAbilities     int `json:"lightAbilities"`
				DamageTypes        int `json:"damageTypes"`
			} `json:"perkGroups,omitempty"`
			PerkIdentifier string `json:"perkIdentifier,omitempty"`
		} `json:"DestinySandboxPerkDefinition"`

		// 15
		DestinyDestinationDefinition []struct {
			DestinationHash        int64  `json:"destinationHash"`
			DestinationName        string `json:"destinationName,omitempty"`
			Icon                   string `json:"icon"`
			PlaceHash              int64  `json:"placeHash"`
			DestinationIdentifier  string `json:"destinationIdentifier"`
			Hash                   int64  `json:"hash"`
			Index                  int    `json:"index"`
			DestinationDescription string `json:"destinationDescription,omitempty"`
			LocationIdentifier     string `json:"locationIdentifier,omitempty"`
		} `json:"DestinyDestinationDefinition"`

		// 16
		DestinyPlaceDefinition []struct {
			PlaceHash        int64  `json:"placeHash"`
			PlaceName        string `json:"placeName"`
			PlaceDescription string `json:"placeDescription"`
			Icon             string `json:"icon"`
			Hash             int64  `json:"hash"`
			Index            int    `json:"index"`
		} `json:"DestinyPlaceDefinition"`

		// 17
		DestinyActivityBundleDefinition []struct {
			BundleHash          int64   `json:"bundleHash"`
			ActivityName        string  `json:"activityName,omitempty"`
			ActivityDescription string  `json:"activityDescription,omitempty"`
			Icon                string  `json:"icon"`
			ReleaseIcon         string  `json:"releaseIcon"`
			ReleaseTime         int     `json:"releaseTime"`
			DestinationHash     int64   `json:"destinationHash"`
			PlaceHash           int64   `json:"placeHash"`
			ActivityTypeHash    int64   `json:"activityTypeHash"`
			ActivityHashes      []int64 `json:"activityHashes"`
			Hash                int64   `json:"hash"`
			Index               int     `json:"index"`
		} `json:"DestinyActivityBundleDefinition"`

		// 18
		DestinyStatGroupDefinition []struct {
			StatGroupHash int64 `json:"statGroupHash"`
			MaximumValue  int   `json:"maximumValue"`
			UIPosition    int   `json:"uiPosition"`
			ScaledStats   []struct {
				StatHash             int64 `json:"statHash"`
				MaximumValue         int   `json:"maximumValue"`
				DisplayAsNumeric     bool  `json:"displayAsNumeric"`
				DisplayInterpolation []struct {
					Value  int `json:"value"`
					Weight int `json:"weight"`
				} `json:"displayInterpolation"`
			} `json:"scaledStats"`
			Overrides struct {
				Num144602215 struct {
					StatHash           int64  `json:"statHash"`
					DisplayName        string `json:"displayName"`
					DisplayDescription string `json:"displayDescription"`
					DisplayIcon        string `json:"displayIcon"`
				} `json:"144602215"`
				Num1735777505 struct {
					StatHash           int64  `json:"statHash"`
					DisplayName        string `json:"displayName"`
					DisplayDescription string `json:"displayDescription"`
					DisplayIcon        string `json:"displayIcon"`
				} `json:"1735777505"`
				Num4244567218 struct {
					StatHash           int64  `json:"statHash"`
					DisplayName        string `json:"displayName"`
					DisplayDescription string `json:"displayDescription"`
					DisplayIcon        string `json:"displayIcon"`
				} `json:"4244567218"`
			} `json:"overrides"`
			Hash  int64 `json:"hash"`
			Index int   `json:"index"`
		} `json:"DestinyStatGroupDefinition"`

		// 19
		DestinySpecialEventDefinition []struct {
			EventHash               int64         `json:"eventHash"`
			EventIdentifier         string        `json:"eventIdentifier"`
			BackgroundImageWeb      string        `json:"backgroundImageWeb"`
			Title                   string        `json:"title"`
			Subtitle                string        `json:"subtitle"`
			Description             string        `json:"description"`
			Link                    string        `json:"link"`
			Icon                    string        `json:"icon"`
			ShowNagMessage          bool          `json:"showNagMessage"`
			ReturnInActivityAdvisor bool          `json:"returnInActivityAdvisor"`
			ProgressionHash         int64         `json:"progressionHash"`
			VendorHash              int64         `json:"vendorHash"`
			FactionHash             int64         `json:"factionHash"`
			BackgroundImageMobile   string        `json:"backgroundImageMobile,omitempty"`
			ActiveUnlockValueHash   int64         `json:"activeUnlockValueHash"`
			BountyHashes            []interface{} `json:"bountyHashes"`
			QuestHashes             []interface{} `json:"questHashes"`
			FriendlyIdentifier      string        `json:"friendlyIdentifier"`
			RecruitmentIds          []string      `json:"recruitmentIds"`
			PlaylistActivityHash    int64         `json:"playlistActivityHash,omitempty"`
			UnlockEventHash         int64         `json:"unlockEventHash"`
			Hash                    int64         `json:"hash"`
			Index                   int           `json:"index"`
		} `json:"DestinySpecialEventDefinition"`

		// 20
		DestinyFactionDefinition []struct {
			FactionHash        int64  `json:"factionHash"`
			FactionName        string `json:"factionName"`
			FactionDescription string `json:"factionDescription"`
			FactionIcon        string `json:"factionIcon"`
			ProgressionHash    int64  `json:"progressionHash"`
			Hash               int64  `json:"hash"`
			Index              int    `json:"index"`
		} `json:"DestinyFactionDefinition"`

		// 21
		DestinyVendorCategoryDefinition []struct {
			CategoryHash     int64  `json:"categoryHash"`
			Order            int    `json:"order"`
			CategoryName     string `json:"categoryName"`
			MobileBannerPath string `json:"mobileBannerPath"`
			Identifier       string `json:"identifier"`
			Hash             int64  `json:"hash"`
			Index            int    `json:"index"`
		} `json:"DestinyVendorCategoryDefinition"`

		// 22
		DestinyEnemyRaceDefinition []struct {
			RaceHash    int64  `json:"raceHash"`
			Identifier  string `json:"identifier"`
			RaceName    string `json:"raceName,omitempty"`
			Description string `json:"description,omitempty"`
			IconPath    string `json:"iconPath"`
			Hash        int64  `json:"hash"`
			Index       int    `json:"index"`
		} `json:"DestinyEnemyRaceDefinition"`

		// 23
		DestinyScriptedSkullDefinition []struct {
			SkullHash   int64  `json:"skullHash"`
			Identifier  string `json:"identifier"`
			SkullName   string `json:"skullName"`
			Description string `json:"description"`
			IconPath    string `json:"iconPath"`
			Hash        int64  `json:"hash"`
			Index       int    `json:"index"`
		} `json:"DestinyScriptedSkullDefinition"`

		// 24
		DestinyTriumphSetDefinition []struct {
			TriumphSetHash     int64     `json:"triumphSetHash"`
			Identifier         string    `json:"identifier"`
			Order              int       `json:"order"`
			Title              string    `json:"title"`
			IconPath           string    `json:"iconPath"`
			IncompleteSubtitle string    `json:"incompleteSubtitle"`
			IncompleteDetails  string    `json:"incompleteDetails"`
			CompletedSubtitle  string    `json:"completedSubtitle"`
			CompletedDetails   string    `json:"completedDetails"`
			LockedSubtitle     string    `json:"lockedSubtitle"`
			LockedDetails      string    `json:"lockedDetails"`
			LockdownDate       time.Time `json:"lockdownDate"`
			Triumphs           []struct {
				Identifier      string `json:"identifier"`
				Title           string `json:"title"`
				Subtitle        string `json:"subtitle"`
				IconPath        string `json:"iconPath"`
				HasProgress     bool   `json:"hasProgress"`
				MinimumProgress int    `json:"minimumProgress"`
				MaximumProgress int    `json:"maximumProgress"`
			} `json:"triumphs"`
			IsRecordBook       bool  `json:"isRecordBook"`
			LockdownUnlockHash int64 `json:"lockdownUnlockHash"`
			Hash               int64 `json:"hash"`
			Index              int   `json:"index"`
		} `json:"DestinyTriumphSetDefinition"`

		// 25
		DestinyItemCategoryDefinition []struct {
			ItemCategoryHash     int64  `json:"itemCategoryHash"`
			Identifier           string `json:"identifier"`
			Visible              bool   `json:"visible"`
			Title                string `json:"title"`
			ShortTitle           string `json:"shortTitle"`
			Description          string `json:"description"`
			GrantDestinyItemType int    `json:"grantDestinyItemType"`
			GrantDestinySubType  int    `json:"grantDestinySubType"`
			GrantDestinyClass    int    `json:"grantDestinyClass"`
			Hash                 int64  `json:"hash"`
			Index                int    `json:"index"`
		} `json:"DestinyItemCategoryDefinition"`

		// 26
		DestinyRewardSourceDefinition []struct {
			SourceHash  int64  `json:"sourceHash"`
			Category    int    `json:"category"`
			Identifier  string `json:"identifier"`
			SourceName  string `json:"sourceName"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
			Hash        int64  `json:"hash"`
			Index       int    `json:"index"`
		} `json:"DestinyRewardSourceDefinition"`

		// 27
		DestinyObjectiveDefinition []struct {
			ObjectiveHash                 int64  `json:"objectiveHash"`
			UnlockValueHash               int64  `json:"unlockValueHash"`
			CompletionValue               int    `json:"completionValue"`
			VendorHash                    int64  `json:"vendorHash"`
			VendorCategoryHash            int64  `json:"vendorCategoryHash"`
			DisplayDescription            string `json:"displayDescription,omitempty"`
			LocationHash                  int64  `json:"locationHash"`
			AllowNegativeValue            bool   `json:"allowNegativeValue"`
			AllowValueChangeWhenCompleted bool   `json:"allowValueChangeWhenCompleted"`
			IsCountingDownward            bool   `json:"isCountingDownward"`
			ValueStyle                    int    `json:"valueStyle"`
			Hash                          int64  `json:"hash"`
			Index                         int    `json:"index"`
		} `json:"DestinyObjectiveDefinition"`

		// 28
		DestinyDamageTypeDefinition []struct {
			DamageTypeHash      int64  `json:"damageTypeHash"`
			Identifier          string `json:"identifier"`
			DamageTypeName      string `json:"damageTypeName"`
			IconPath            string `json:"iconPath"`
			TransparentIconPath string `json:"transparentIconPath"`
			ShowIcon            bool   `json:"showIcon"`
			EnumValue           int    `json:"enumValue"`
			Hash                int64  `json:"hash"`
			Index               int    `json:"index"`
		} `json:"DestinyDamageTypeDefinition"`

		// 29
		DestinyCombatantDefinition []struct {
			CombatantHash int64  `json:"combatantHash"`
			Icon          string `json:"icon"`
			Hash          int64  `json:"hash"`
			Index         int    `json:"index"`
			CombatantName string `json:"combatantName,omitempty"`
			Description   string `json:"description,omitempty"`
			Image         string `json:"image,omitempty"`
		} `json:"DestinyCombatantDefinition"`

		// 30
		DestinyActivityCategoryDefinition []struct {
			Title       string `json:"title"`
			Identifier  string `json:"identifier"`
			Description string `json:"description"`
			HelpTitle   string `json:"helpTitle"`
			Image       string `json:"image"`
			Links       []struct {
				Title string `json:"title"`
				URL   string `json:"url"`
			} `json:"links"`
			ParentHashes    []int64 `json:"parentHashes"`
			Order           int     `json:"order"`
			HasHelpContent  bool    `json:"hasHelpContent"`
			ResetIdentifier string  `json:"resetIdentifier"`
			Hash            int64   `json:"hash"`
			Index           int     `json:"index"`
		} `json:"DestinyActivityCategoryDefinition"`

		// 31
		DestinyRecordDefinition []struct {
			DisplayName        string `json:"displayName"`
			Description        string `json:"description"`
			RecordValueUIStyle string `json:"recordValueUIStyle,omitempty"`
			Icon               string `json:"icon,omitempty"`
			Rewards            []struct {
				UIItemHash     int64 `json:"uiItemHash"`
				UIItemQuantity int   `json:"uiItemQuantity"`
				LevelRewarded  int   `json:"levelRewarded"`
			} `json:"rewards,omitempty"`
			Objectives []struct {
				ObjectiveHash int64 `json:"objectiveHash"`
			} `json:"objectives,omitempty"`
			Hash  int64 `json:"hash"`
			Index int   `json:"index"`
		} `json:"DestinyRecordDefinition"`

		// 32
		DestinyRecordBookDefinition []struct {
			DisplayName        string `json:"displayName"`
			DisplayDescription string `json:"displayDescription"`
			UnavailableReason  string `json:"unavailableReason"`
			ProgressionHash    int64  `json:"progressionHash"`
			RecordCount        int    `json:"recordCount"`
			Hash               int64  `json:"hash"`
			Index              int    `json:"index"`
			Pages              []struct {
				DisplayName        string `json:"displayName"`
				DisplayDescription string `json:"displayDescription"`
				DisplayStyle       int    `json:"displayStyle"`
				Records            []struct {
					RecordHash int64 `json:"recordHash"`
					Spotlight  bool  `json:"spotlight"`
				} `json:"records"`
			} `json:"pages,omitempty"`
			Icon     string `json:"icon,omitempty"`
			ItemHash int64  `json:"itemHash,omitempty"`
		} `json:"DestinyRecordBookDefinition"`

		// 33
		DestinyBondDefinition []struct {
			DisplayIcon             string `json:"displayIcon,omitempty"`
			ProvidedUnlockHash      int64  `json:"providedUnlockHash"`
			ProvidedUnlockValueHash int64  `json:"providedUnlockValueHash"`
			ShowInAdvisor           bool   `json:"showInAdvisor"`
			Hash                    int64  `json:"hash"`
			Index                   int    `json:"index"`
		} `json:"DestinyBondDefinition"`

		// 34
		DestinyLocationDefinition []struct {
			LocationHash     int64 `json:"locationHash"`
			LocationReleases []struct {
				DestinationHash       int64 `json:"destinationHash"`
				ActivityHash          int64 `json:"activityHash"`
				DirectorBookHash      int64 `json:"directorBookHash"`
				ActivityGraphHash     int64 `json:"activityGraphHash"`
				ActivityGraphNodeHash int64 `json:"activityGraphNodeHash"`
			} `json:"locationReleases"`
			Hash  int64 `json:"hash"`
			Index int   `json:"index"`
		} `json:"DestinyLocationDefinition"`

		// 35
		DestinyGrimoireDefinition []struct {
			ThemeCollection []struct {
				ThemeID          string `json:"themeId"`
				ThemeName        string `json:"themeName"`
				NormalResolution struct {
					Image struct {
						Rect struct {
							X      int `json:"x"`
							Y      int `json:"y"`
							Height int `json:"height"`
							Width  int `json:"width"`
						} `json:"rect"`
						SheetPath string `json:"sheetPath"`
						SheetSize struct {
							X      int `json:"x"`
							Y      int `json:"y"`
							Height int `json:"height"`
							Width  int `json:"width"`
						} `json:"sheetSize"`
					} `json:"image"`
					SmallImage struct {
						Rect struct {
							X      int `json:"x"`
							Y      int `json:"y"`
							Height int `json:"height"`
							Width  int `json:"width"`
						} `json:"rect"`
						SheetPath string `json:"sheetPath"`
						SheetSize struct {
							X      int `json:"x"`
							Y      int `json:"y"`
							Height int `json:"height"`
							Width  int `json:"width"`
						} `json:"sheetSize"`
					} `json:"smallImage"`
				} `json:"normalResolution"`
				HighResolution struct {
					Image struct {
						Rect struct {
							X      int `json:"x"`
							Y      int `json:"y"`
							Height int `json:"height"`
							Width  int `json:"width"`
						} `json:"rect"`
						SheetPath string `json:"sheetPath"`
						SheetSize struct {
							X      int `json:"x"`
							Y      int `json:"y"`
							Height int `json:"height"`
							Width  int `json:"width"`
						} `json:"sheetSize"`
					} `json:"image"`
					SmallImage struct {
						Rect struct {
							X      int `json:"x"`
							Y      int `json:"y"`
							Height int `json:"height"`
							Width  int `json:"width"`
						} `json:"rect"`
						SheetPath string `json:"sheetPath"`
						SheetSize struct {
							X      int `json:"x"`
							Y      int `json:"y"`
							Height int `json:"height"`
							Width  int `json:"width"`
						} `json:"sheetSize"`
					} `json:"smallImage"`
				} `json:"highResolution"`
				PageCollection []struct {
					PageID           string `json:"pageId"`
					PageName         string `json:"pageName"`
					NormalResolution struct {
						Image struct {
							Rect struct {
								X      int `json:"x"`
								Y      int `json:"y"`
								Height int `json:"height"`
								Width  int `json:"width"`
							} `json:"rect"`
							SheetPath string `json:"sheetPath"`
							SheetSize struct {
								X      int `json:"x"`
								Y      int `json:"y"`
								Height int `json:"height"`
								Width  int `json:"width"`
							} `json:"sheetSize"`
						} `json:"image"`
						SmallImage struct {
							Rect struct {
								X      int `json:"x"`
								Y      int `json:"y"`
								Height int `json:"height"`
								Width  int `json:"width"`
							} `json:"rect"`
							SheetPath string `json:"sheetPath"`
							SheetSize struct {
								X      int `json:"x"`
								Y      int `json:"y"`
								Height int `json:"height"`
								Width  int `json:"width"`
							} `json:"sheetSize"`
						} `json:"smallImage"`
					} `json:"normalResolution"`
					HighResolution struct {
						Image struct {
							Rect struct {
								X      int `json:"x"`
								Y      int `json:"y"`
								Height int `json:"height"`
								Width  int `json:"width"`
							} `json:"rect"`
							SheetPath string `json:"sheetPath"`
							SheetSize struct {
								X      int `json:"x"`
								Y      int `json:"y"`
								Height int `json:"height"`
								Width  int `json:"width"`
							} `json:"sheetSize"`
						} `json:"image"`
						SmallImage struct {
							Rect struct {
								X      int `json:"x"`
								Y      int `json:"y"`
								Height int `json:"height"`
								Width  int `json:"width"`
							} `json:"rect"`
							SheetPath string `json:"sheetPath"`
							SheetSize struct {
								X      int `json:"x"`
								Y      int `json:"y"`
								Height int `json:"height"`
								Width  int `json:"width"`
							} `json:"sheetSize"`
						} `json:"smallImage"`
					} `json:"highResolution"`
					CardBriefs []struct {
						CardID      int `json:"cardId"`
						Points      int `json:"points"`
						TotalPoints int `json:"totalPoints"`
					} `json:"cardBriefs"`
				} `json:"pageCollection"`
				ThemeDescription string `json:"themeDescription,omitempty"`
			} `json:"themeCollection"`
		} `json:"DestinyGrimoireDefinition"`

		// 36 -- Note: Unsure if this is comprehensive. Generated from a handful of entries
		DestinyGrimoireCardDefinition []struct {
			CardID           int    `json:"cardId"`
			CardName         string `json:"cardName"`
			CardIntro        string `json:"cardIntro,omitempty"`
			CardDescription  string `json:"cardDescription"`
			UnlockHowToText  string `json:"unlockHowToText"`
			Rarity           int    `json:"rarity"`
			UnlockFlagHash   int64  `json:"unlockFlagHash"`
			Points           int    `json:"points"`
			NormalResolution struct {
				Image struct {
					Rect struct {
						X      int `json:"x"`
						Y      int `json:"y"`
						Height int `json:"height"`
						Width  int `json:"width"`
					} `json:"rect"`
					SheetPath string `json:"sheetPath"`
					SheetSize struct {
						X      int `json:"x"`
						Y      int `json:"y"`
						Height int `json:"height"`
						Width  int `json:"width"`
					} `json:"sheetSize"`
				} `json:"image"`
				SmallImage struct {
					Rect struct {
						X      int `json:"x"`
						Y      int `json:"y"`
						Height int `json:"height"`
						Width  int `json:"width"`
					} `json:"rect"`
					SheetPath string `json:"sheetPath"`
					SheetSize struct {
						X      int `json:"x"`
						Y      int `json:"y"`
						Height int `json:"height"`
						Width  int `json:"width"`
					} `json:"sheetSize"`
				} `json:"smallImage"`
			} `json:"normalResolution"`
			HighResolution struct {
				Image struct {
					Rect struct {
						X      int `json:"x"`
						Y      int `json:"y"`
						Height int `json:"height"`
						Width  int `json:"width"`
					} `json:"rect"`
					SheetPath string `json:"sheetPath"`
					SheetSize struct {
						X      int `json:"x"`
						Y      int `json:"y"`
						Height int `json:"height"`
						Width  int `json:"width"`
					} `json:"sheetSize"`
				} `json:"image"`
				SmallImage struct {
					Rect struct {
						X      int `json:"x"`
						Y      int `json:"y"`
						Height int `json:"height"`
						Width  int `json:"width"`
					} `json:"rect"`
					SheetPath string `json:"sheetPath"`
					SheetSize struct {
						X      int `json:"x"`
						Y      int `json:"y"`
						Height int `json:"height"`
						Width  int `json:"width"`
					} `json:"sheetSize"`
				} `json:"smallImage"`
			} `json:"highResolution"`
			StatisticCollection []struct {
				StatNumber      int    `json:"statNumber"`
				CardID          int    `json:"cardId"`
				StatName        string `json:"statName"`
				AccumulatorHash int64  `json:"accumulatorHash"`
				RankCollection  []struct {
					Rank           int   `json:"rank"`
					Threshold      int   `json:"threshold"`
					UnlockFlagHash int64 `json:"unlockFlagHash"`
					Points         int   `json:"points"`
				} `json:"rankCollection"`
			} `json:"statisticCollection,omitempty"`
			CardLabel string `json:"cardLabel,omitempty"`
		} `json:"DestinyGrimoireCardDefinition"`
	} `json:"manifest"`
}
