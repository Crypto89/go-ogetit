package ogetit

import (
	"encoding/json"
	"encoding/xml"
	"strconv"
	"time"
)

type ServerData struct {
	XMLName                                         xml.Name `xml:"serverData" json:"-"`
	Name                                            string   `xml:"name"`
	Number                                          int64    `xml:"number"`
	Language                                        string   `xml:"language"`
	Timezone                                        string   `xml:"timezone"`
	TimezoneOffset                                  string   `xml:"timezoneOffset"`
	Domain                                          string   `xml:"domain"`
	Version                                         string   `xml:"version"`
	Speed                                           int64    `xml:"speed"`
	SpeedFleetPeaceful                              int64    `xml:"speedFleetPeaceful"`
	SpeedFleetWar                                   int64    `xml:"speedFleetWar"`
	SpeedFleetHolding                               int64    `xml:"speedFleetHolding"`
	Galaxies                                        int64    `xml:"galaxies"`
	Systems                                         int64    `xml:"systems"`
	ACS                                             bool     `xml:"acs"`
	RapidFire                                       bool     `xml:"rapidfire"`
	DefToTF                                         bool     `xml:"defToTF"`
	DebrisFactor                                    float64  `xml:"debrisFactor"`
	DebrisFactorDef                                 float64  `xml:"debrisFactorDef"`
	RepairFactor                                    float64  `xml:"repairFactor"`
	NewbieProtectionLimit                           int64    `xml:"newbieProtectionLimit"`
	NewbieProtectionHigh                            int64    `xml:"newbieProtectionHigh"`
	TopScore                                        float64  `xml:"topScore"`
	BonusFields                                     int64    `xml:"bonusFields"`
	DonutGalaxy                                     bool     `xml:"donutGalaxy"`
	DonutSystem                                     bool     `xml:"donutSystem"`
	WFEnabled                                       bool     `xml:"wfEnabled"`
	WFMinimumRessLost                               int64    `xml:"wfMinimumRessLost"`
	WFMinimumLossPercentage                         int64    `xml:"wfMinimumLossPercentage"`
	WFBasicPercentageRepairable                     int64    `xml:"wfBasicPercentageRepairable"`
	GlobalDeuteriumSaveFactor                       float64  `xml:"globalDeuteriumSaveFactor"`
	BashLimit                                       int64    `xml:"bashLimit"`
	ProbeCargo                                      bool     `xml:"probeCargo"`
	ResearchDurationDivisor                         int64    `xml:"researchDurationDivisor"`
	DarkMatterNewAcount                             int64    `xml:"darkMatterNewAcount"`
	CargoHyperspaceTechMultiplier                   int64    `xml:"cargoHyperspaceTechMultiplier"`
	CharacterClassesEnabled                         bool     `xml:"characterClassesEnabled"`
	MinerBonusResourceProduction                    float64  `xml:"minerBonusResourceProduction"`
	MinerBonusFasterTradingShips                    bool     `xml:"minerBonusFasterTradingShips"`
	MinerBonusIncreasedCargoCapacityForTradingShips float64  `xml:"minerBonusIncreasedCargoCapacityForTradingShips"`
	MinerBonusAdditionalFleetSlots                  int64    `xml:"minerBonusAdditionalFleetSlots"`
	MinerBonusAdditionalMarketSlots                 int64    `xml:"minerBonusAdditionalMarketSlots"`
	MinerBonusAdditionalCrawler                     float64  `xml:"minerBonusAdditionalCrawler"`
	MinerBonusMaxCrawler                            float64  `xml:"minerBonusMaxCrawler"`
	MinerBonusEnergy                                float64  `xml:"minerBonusEnergy"`
	MinerBonusOverloadCrawler                       bool     `xml:"minerBonusOverloadCrawler"`
	ResourceBuggyProductionBoost                    float64  `xml:"resourceBuggyProductionBoost"`
	ResourceBuggyMaxProductionBoost                 float64  `xml:"resourceBuggyMaxProductionBoost"`
	ResourceBuggyEnergyConsumptionPerUnit           int64    `xml:"resourceBuggyEnergyConsumptionPerUnit"`
	WarriorBonusFasterCombatShips                   bool     `xml:"warriorBonusFasterCombatShips"`
	WarriorBonusFasterRecyclers                     bool     `xml:"warriorBonusFasterRecyclers"`
	WarriorBonusFuelConsumption                     float64  `xml:"warriorBonusFuelConsumption"`
	WarriorBonusRecyclerFuelConsumption             float64  `xml:"warriorBonusRecyclerFuelConsumption"`
	WarriorBonusRecyclerCargoCapacity               float64  `xml:"warriorBonusRecyclerCargoCapacity"`
	WarriorBonusAdditionalFleetSlots                int64    `xml:"warriorBonusAdditionalFleetSlots"`
	WarriorBonusAdditionalMoonFields                int64    `xml:"warriorBonusAdditionalMoonFields"`
	WarriorBonusFleetHalfSpeed                      bool     `xml:"warriorBonusFleetHalfSpeed"`
	WarriorBonusAttackerWreckfield                  bool     `xml:"warriorBonusAttackerWreckfield"`
	CombatDebrisFieldLimit                          float64  `xml:"combatDebrisFieldLimit"`
	ExplorerBonusIncreasedResearchSpeed             float64  `xml:"explorerBonusIncreasedResearchSpeed"`
	ExplorerBonusIncreasedExpeditionOutcome         float64  `xml:"explorerBonusIncreasedExpeditionOutcome"`
	ExplorerBonusLargerPlanets                      float64  `xml:"explorerBonusLargerPlanets"`
	ExplorerUnitItemsPerDay                         int64    `xml:"explorerUnitItemsPerDay"`
	ExplorerBonusPhalanxRange                       float64  `xml:"explorerBonusPhalanxRange"`
	ExplorerBonusPlunderInactive                    bool     `xml:"explorerBonusPlunderInactive"`
	ExplorerBonusExpeditionEnemyReduction           float64  `xml:"explorerBonusExpeditionEnemyReduction"`
	ExplorerBonusAdditionalExpeditionSlots          int64    `xml:"explorerBonusAdditionalExpeditionSlots"`
	ResourceProductionIncreaseCrystalDefault        float64  `xml:"resourceProductionIncreaseCrystalDefault"`
	ResourceProductionIncreaseCrystalPos1           float64  `xml:"resourceProductionIncreaseCrystalPos1"`
	ResourceProductionIncreaseCrystalPos2           float64  `xml:"resourceProductionIncreaseCrystalPos2"`
	ResourceProductionIncreaseCrystalPos3           float64  `xml:"resourceProductionIncreaseCrystalPos3"`
}

type SpyReport struct {
	ResultCode int                  `json:"RESULT_CODE"`
	Data       SpyReportDataWrapper `json:"RESULT_DATA"`
}

type SpyReportData struct {
	Generic struct {
		ID                        string    `json:"sr_id"`
		EventTime                 time.Time `json:"event_time"`
		EventTimestamp            int64     `json:"event_timestamp"`
		TotalShipCount            int64     `json:"total_ship_count"`
		TotalDefenceCount         int64     `json:"total_defence_count"`
		LootPercentage            int64     `json:"loot_percentage"`
		SpyChanceFail             int64     `json:"spy_chance_fail"`
		Activity                  int64     `json:"activity"`
		AttackerName              string    `json:"attacker_name"`
		AttackerUserID            int64     `json:"attacker_user_id"`
		AttackerCharacterClassID  int64     `json:"attacker_character_class_id"`
		AttackerPlanetName        string    `json:"attacker_planet_name"`
		AttackerPlanetCoordinates string    `json:"attacker_planet_coordinates"`
		AttackerPlanetType        int64     `json:"attacker_planet_type"`
		AttackerAllianceName      *string   `json:"attacker_alliance_name"`
		AttackerAllianceTag       *string   `json:"attacker_alliance_tag"`
		DefenderName              string    `json:"defender_name"`
		DefenderUserID            int64     `json:"defender_user_id"`
		DefenderCharacterClassID  int64     `json:"defender_character_class_id"`
		DefenderPlanetName        string    `json:"defender_planet_name"`
		DefenderPlanetCoordinates string    `json:"defender_planet_coordinates"`
		DefenderPlanetType        int64     `json:"defender_planet_type"`
		DefenderAllianceName      *string   `json:"defender_alliance_name"`
		DefenderAllianceTag       *string   `json:"defender_alliance_tag"`
		FailedShips               bool      `json:"failed_ships"`
		FailedDefense             bool      `json:"failed_defense"`
		FailedBuildings           bool      `json:"failed_buildings"`
		FailedResearch            bool      `json:"failed_research"`
		DebrisMetal               int64     `json:"debris_metal"`
		DebrisCrystal             int64     `json:"debris_crystal"`
	} `json:"generic"`
	Details struct {
		Resources struct {
			Metal     int64 `json:"metal"`
			Crystal   int64 `json:"crystal"`
			Deuterium int64 `json:"deuterium"`
		} `json:"resources"`
		Buildings []Building `json:"buildings"`
		Ships     []Ship     `json:"ships"`
		Defence   []Defence  `json:"defence"`
		Research  []Research `json:"research"`
	} `json:"details"`
}

type SpyReportDataWrapper struct {
	SpyReportData
}

func (srd *SpyReportDataWrapper) UnmarshalJSON(b []byte) error {
	if string(b) != "[]" {
		return json.Unmarshal(b, &srd.SpyReportData)
	}
	return nil
}

type Building struct {
	BuildingType int64 `json:"building_type"`
	Level        int64 `json:"level"`
}

type Defence struct {
	DefenceType int64 `json:"defence_type"`
	Count       int64 `json:"count"`
}

type Ship struct {
	Owner    int64     `json:"owner,omitempty"`
	ShipType int64     `json:"ship_type"`
	Armor    int64     `json:"armor,omitempty"`
	Shield   int64     `json:"shield,omitempty"`
	Weapon   int64     `json:"weapon,omitempty"`
	Count    StringInt `json:"count"`
}

type Research struct {
	ResearchType int64 `json:"research_type"`
	Level        int64 `json:"level"`
}

type CombatReport struct {
	ResultCode int                     `json:"RESULT_CODE"`
	Data       CombatReportDataWrapper `json:"RESULT_DATA"`
}

type CombatReportData struct {
	RepairedDeferenses []interface{} `json:"repaired_defenses"`
	Generic            struct {
		ID                               string        `json:"cr_id"`
		EventTime                        time.Time     `json:"event_time"`
		EventTimestamp                   int64         `json:"event_timestamp"`
		CombatCoordinates                string        `json:"combat_coordinates"`
		CombatPlanetType                 int64         `json:"combat_planet_type"`
		CombatRounds                     int64         `json:"combat_rounds"`
		LootPercentage                   int64         `json:"loot_percentage"`
		Winner                           string        `json:"winner"`
		UnitsLostAttacker                int64         `json:"units_lost_attacker"`
		UnitsLostDefender                int64         `json:"units_lost_defender"`
		AttackerCount                    int64         `json:"attacker_count"`
		DefenderCount                    int64         `json:"defender_count"`
		LootMetal                        int64         `json:"loot_metal"`
		LootCrystal                      int64         `json:"loot_crystal"`
		LootDeuterium                    int64         `json:"loot_deuterium"`
		CombatHonorable                  bool          `json:"combat_honorable"`
		AttackerHonorable                bool          `json:"attacker_honorable"`
		AttackerHonorpoints              int64         `json:"attacker_honorpoints"`
		DefenderHonorable                bool          `json:"defender_honorable"`
		DefenderHonorpoints              int64         `json:"defender_honorpoints"`
		MoonCreated                      bool          `json:"moon_created"`
		MoonChance                       int64         `json:"moon_chance"`
		MoonSize                         int64         `json:"moon_size"`
		MoonExists                       bool          `json:"moon_exists"`
		DebrisMetalTotal                 int64         `json:"debris_metal_total"`
		DebrisCrystalTotal               int64         `json:"debris_crystal_total"`
		DebrisMetal                      int64         `json:"debris_metal"`
		DebrisCrystal                    int64         `json:"debris_crystal"`
		DebrisReaperMetalRetrieved       int64         `json:"debris_reaper_metal_retrieved"`
		DebrisReaperCrystalRetrieved     int64         `json:"debris_reaper_crystal_retrieved"`
		DebrisPathfinderMetalRetrieved   int64         `json:"debris_pathfinder_metal_retrieved"`
		DebrisPathfinderCrystalRetrieved int64         `json:"debris_pathfinder_crystal_retrieved"`
		Wreckfield                       []interface{} `json:"wreckfield"`
	} `json:"generic"`
	Attackers []CombatReportPlayer `json:"attackers"`
	Defenders []CombatReportPlayer `json:"defenders"`
	Rounds    []CombatReportRound  `json:"rounds"`
}

type CombatReportDataWrapper struct {
	CombatReportData
}

func (crd *CombatReportDataWrapper) UnmarshalJSON(b []byte) error {
	if string(b) != "[]" {
		return json.Unmarshal(b, &crd.CombatReportData)
	}
	return nil
}

type CombatReportPlayer struct {
	FleetOwner                 string  `json:"fleet_owner"`
	FleetOwnerID               int64   `json:"fleet_owner_id"`
	FleetOwnerCharacterClassID int64   `json:"fleet_owner_character_class_id"`
	FleetOwnerCoordinates      string  `json:"fleet_owner_coordinates"`
	FleetOwnerPlanetType       int64   `json:"fleet_owner_planet_type"`
	FleetOwnerPlanetName       string  `json:"fleet_owner_planet_name"`
	FleetOwnerAlliance         *string `json:"fleet_owner_alliance"`
	FleetOwnerAllianceTag      *string `json:"fleet_owner_alliance_tag"`
	FleetArmorPercentage       int64   `json:"fleet_armor_percentage"`
	FleetShieldPercentage      int64   `json:"fleet_shield_percentage"`
	FleetWeaponPercentage      int64   `json:"fleet_weapon_percentage"`
	FleetComposition           []Ship  `json:"fleet_composition"`
}

type CombatReportRound struct {
	RoundNumber int64 `json:"round_number"`
	Statistics  struct {
		AttackerHits         string `json:"attacker_hits"`
		AttackerAbsorbed     string `json:"attacker_absorbed"`
		AttackerFullstrength string `json:"attacker_fullstrength"`
		DefenderHits         string `json:"defender_hits"`
		DefenderAbsorbed     string `json:"defender_absorbed"`
		DefenderFullstrength string `json:"defender_fullstrength"`
	} `json:"statistics"`
	AttackerShips      []Ship `json:"attacker_ships"`
	AttackerShipLosses []Ship `json:"attacker_ship_losses"`
	// DefenderShips      []Ship `json:"defender_ships"`
	// DefenderShipLosses []Ship `json:"defender_ship_losses"`
}

type StringInt int64

func (si *StringInt) UnmarshalJSON(b []byte) error {
	if i, err := strconv.Atoi(string(b)); err == nil {
		(*si) = StringInt(i)
	}
	return nil
}

type RecycleReport struct {
	ResultCode int                      `json:"RESULT_CODE"`
	Data       RecycleReportDataWrapper `json:"RESULT_DATA"`
}

type RecycleReportData struct {
	Generic struct {
		ID                         string    `json:"rr_id"`
		EventTime                  time.Time `json:"event_time"`
		EventTimestamp             int64     `json:"event_timestamp"`
		Coordinates                string    `json:"coordinates"`
		MetalInDebrisField         int64     `json:"metal_in_debris_field"`
		CrystalInDebrisField       int64     `json:"crystal_in_debris_field"`
		RecyclerCount              int64     `json:"recycler_count"`
		RecyclerCapacity           StringInt `json:"recycler_capacity"`
		RecyclerMetalRetrieved     int64     `json:"recycler_metal_retrieved"`
		RecyclerCrystalRetrieved   int64     `json:"recycler_crystal_retrieved"`
		ReaperCount                int64     `json:"reaper_count"`
		ReaperCapacity             StringInt `json:"reaper_capacity"`
		ReaperMetalRetrieved       int64     `json:"reaper_metal_retrieved"`
		ReaperCrystalRetrieved     int64     `json:"reaper_crystal_retrieved"`
		PathfinderCount            int64     `json:"pathfinder_count"`
		PathfinderCapacity         StringInt `json:"pathfinder_capacity"`
		PathfinderMetalRetrieved   int64     `json:"pathfinder_metal_retrieved"`
		PathfinderCrystalRetrieved int64     `json:"pathfinder_crystal_retrieved"`
		MetalRetrieved             int64     `json:"metal_retrieved"`
		CrystalRetrieved           int64     `json:"crystal_retrieved"`
		OwnerName                  string    `json:"owner_name"`
		OwnerID                    int64     `json:"owner_id"`
		OwnerAllianceName          string    `json:"owner_alliance_name"`
		OwnerAllianceTag           string    `json:"owner_alliance_tag"`
	} `json:"generic"`
}

type RecycleReportDataWrapper struct {
	RecycleReportData
}

func (crd *RecycleReportDataWrapper) UnmarshalJSON(b []byte) error {
	if string(b) != "[]" {
		return json.Unmarshal(b, &crd.RecycleReportData)
	}
	return nil
}
