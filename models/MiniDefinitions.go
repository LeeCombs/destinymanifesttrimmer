package models

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
