package models

import "time"

type MiniDestinyActivityDefinition struct {
	ActivityName        string
	ActivityDescription string
	Icon                string
	ActivityLevel       int
	DestinationHash     int64
	PlaceHash           int64
	ActivityTypeHash    int64
	Rewards             []interface{}
	Skulls              []interface{}
}

type MiniDestinyActivityTypeDefinition struct {
	Identifier              string
	ActivityTypeName        string
	ActivityTypeDescription string
	Icon                    string
	// Maybe?
	// ActiveBackgroundVirtualPath    string
	// CompletedBackgroundVirtualPath string
	// TooltipBackgroundVirtualPath   string
}

type MiniDestinyClassDefinition struct {
	ClassType int
	ClassName string
}

type MiniDestinyGenderDefinition struct {
	GenderType int
	GenderName string
}

type MiniDestinyInventoryBucketDefinition struct {
	BucketName             string
	BucketDescription      string
	ItemCount              int
	Category               int // Figure out what this is for
	Location               int
	HasTransferDestination bool
	Enabled                bool
}

// Note -- This would be very subject to change, depending one what data you want to extract
// Note -- Unsure if this is comprehensive. Generated from a handful of entries
type MiniDestinyInventoryItemDefinition struct {
	ItemName            string
	ItemDescription     string
	HasIcon             bool
	Icon                string
	SecondaryIcon       string
	TierType            int
	TierTypeName        string // May be unnecessary
	ItemType            int
	ItemSubType         int
	ItemTypeName        string // May be unnecessary
	SpecialItemType     int
	ClassType           int
	BucketTypeHash      int64
	PrimaryBaseStatHash int64
	Stats               interface{}
	Exclusive           int
	// Maybe?
	// ItemCategoryHashes  []int64
	// SourceHashes        []int64
}

type MiniDestinyProgressionDefinition struct {
	Name        string
	Scope       int // Figure out what this is for
	Steps       []interface{}
	Visible     bool
	Icon        string
	Description string
	Source      string
}

type MiniDestinyRaceDefinition struct {
	RaceType        int
	RaceName        string
	RaceNameMale    string
	RaceNameFemale  string
	RaceDescription string
}

// Note -- No clue about this one, so it's pretty empty for now
// Note -- Unsure if this is comprehensive. Generated from a handful of entries
type MiniDestinyTalentGridDefinition struct {
	ProgressionHash int64
	Nodes           []interface{}
}

// Note -- No clue about this one, so it's pretty empty for now
type MiniDestinyUnlockFlagDefinition struct {
	DisplayName        string
	DisplayDescription string
}

type MiniDestinyVendorDefinition struct {
	Summary struct {
		VendorName            string
		VendorDescription     string
		VendorIcon            string
		FactionHash           int64
		Visible               bool
		VendorPortrait        string
		VendorBanner          string
		VendorCategoryHash    int64
		VendorSubcategoryHash int64
		// Maybe?
		// VendorCategoryHashes  []int64
		// MapSectionName        string
		// EventHash             int64
	}
}

// Note -- This is not referenced to using a hash, but instead StatID
type MiniDestinyHistoricalStatsDefinition struct {
	StatID          string
	StatName        string
	StatDescription string
	Group           int
	Category        int
	UnitType        int
	UnitLabel       string
	Weight          int
	IconImage       string
	// Maybe?
	// PeriodTypes []int
	// Modes       []int
}

type MiniDestinyDirectorBookDefinition struct {
	BookName        string
	BookDescription string
	BookNumber      string
	Visible         bool
	IsOverview      bool
	DestinationHash int64

	// Maybe? Figure these out
	/*
	   Connections []interface{}
	   Nodes []struct {
	       NodeDefinitionHash   int64
	       StyleHash            int64
	       PositionX            int
	       PositionY            int
	       PositionZ            int
	       ActivityBundleHashes []int64
	       States               []struct {
	           State int
	       }
	       UIModifier int

	*/
}

type MiniDestinyStatDefinition struct {
	StatName        string
	StatDescription string
	Icon            string
}

type MiniDestinySandboxPerkDefinition struct {
	DisplayName        string
	DisplayDescription string
	DisplayIcon        string
	IsDisplayable      bool
	PerkIdentifer      string
	PerkGroups         struct {
		WeaponPerformance  int
		ImpactEffects      int
		GuardianAttributes int
		LightAbilities     int
		DamageTypes        int
	}
}

type MiniDestinyDestinationDefinition struct {
	DestinationName        string
	DestinationIdentifier  string
	DestinationDescription string
	Icon                   string
	PlaceHash              int64
	LocationIndentifier    string
}

type MiniDestinyPlaceDefinition struct {
	PlaceName        string
	PlaceDescription string
	Icon             string
}

type MiniDestinyActivityBundleDefinition struct {
	ActivityName        string
	ActivityDescription string
	ActivityTypeHash    int64
	ActivityHashes      []int64
	Icon                string
	DestinationHash     int64
	PlaceHash           int64
}

// Note - No clue what this is for, needs research
type MiniDestinyStatGroupDefinition struct {
	MaximumValue int
	// TODO -- Figure out what ScaleStats and Overrides is used for
}

type MiniDestinySpecialEventDefinition struct {
	Title                 string
	Subtitle              string
	Description           string
	Link                  string
	Icon                  string
	BackgroundImageWeb    string
	BackgroundImageMobile string
	ProgressionHash       int64
	VendorHash            int64
	FactionHash           int64
	ActiveUnlockValueHash int64
	PlaylistActivityHash  int64
	UnlockEventHash       int64
	BountyHashes          []interface{}
	QuestHashes           []interface{}
}

type MiniDestinyFactionDefinition struct {
	FactionName        string
	FactionDescription string
	FactionIcon        string
	ProgressionHash    int64
}

type MiniDestinyVendorCategoryDefinition struct {
	CategoryName     string
	MobileBannerPath string // ?
	Identifier       string
}

type MiniDestinyEnemyRaceDefinition struct {
	RaceName    string
	Description string
	IconPath    string
}

type MiniDestinyScriptedSkullDefinition struct {
	SkullName   string
	Description string
	IconPath    string
}

type MiniDestinyTriumphSetDefinition struct {
	Title              string
	IconPath           string
	IncompleteSubtitle string
	IncompleteDetails  string
	CompletedSubtitle  string
	CompletedDetails   string
	LockedSubtitle     string
	LockedDetails      string
	LockdownDate       time.Time
	LockdownUnlockHash int64
	// Triumphs []struct{}
}

type MiniDestinyItemCategoryDefinition struct {
	Visible     bool
	Title       string
	ShortTitle  string
	Description string
	// What is Grant?
	GrantDestinyItemType int
	GrantDestinySubType  int
	GrantDestinyClass    int
}

type MiniDestinyRewardSourceDefinition struct {
	Category    int
	SourceName  string
	Description string
	Icon        string
}

// No clue
type MiniDestinyObjectiveDefinition struct {
	DisplayDescription string
	UnlockValueHash    int64
	VendorHash         int64
	VendorCategoryHash int64
	LocationHash       int64
}

type MiniDestinyDamageTypeDefinition struct {
	DamageTypeName      string
	IconPath            string
	TransparentIconPath string
	ShowIcon            bool
	EnumValue           int
}

type MiniDestinyCombatantDefinition struct {
	Icon          string
	CombatantName string
	Description   string
	Image         string
}

type MiniDestinyActivityCategoryDefinition struct {
	Title        string
	HelpTitle    string
	Description  string
	Image        string
	ParentHashes []int64
	// Links []struct{Title string,URL string}
}

type MiniDestinyRecordDefinition struct {
	DisplayName        string
	Description        string
	Icon               string
	RecordValueUIStyle string
	// Rewards []struct{}
	// Objectives []struct{}
}

type MiniDestinyRecordBookDefinition struct {
	DisplayName        string
	DisplayDescription string
	UnavailableReason  string
	PRogressionHash    int64
	RecordCount        int
	Icon               string
	ItemHash           int64
	// Pages []struct{}
}

type MiniDestinyBondDefinition struct {
	DisplayIcon             string
	ProvidedUnlockHash      int64
	ProvidedUnlockValueHash int64
	ShowInAdvisor           bool
}

// No clue
type MiniDestinyLocationDefinition struct {
	LocationHash int64
	/*
		LocationReleases []struct {
			DestinationHash       int64
			ActivityHash          int64
			DirectorBookHash      int64
			ActivityGraphHash     int64
			ActivityGraphNodeHash int64
		}
	*/
}

// No clue
type MiniDestinyGrimoireDefinition struct {
	// ThemeCollection []struct{}
}

// 36 -- Note: Unsure if this is comprehensive. Generated from a handful of entries
type MiniDestinyGrimoireCardDefinition struct {
	CardID          int
	CardName        string
	CardIntro       string
	CardDescription string
	CardLabel       string
	UnlockHowToText string
	UnlockFlagHash  int64
	Points          int
	// StatisticCollection []struct{}
}
