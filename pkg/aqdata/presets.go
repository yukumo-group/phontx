package aqdata

// Aq1Preset defines the presets for aquestak1
type Aq1Preset struct{ Preset }

var aq1Presets PresetList[Aq1Preset, *Aq1Preset]

var (
	// Aq1F1 -> f1
	Aq1F1 = aq1Presets.define("f1")
	// Aq1F2 -> f2
	Aq1F2 = aq1Presets.define("f2")
	// Aq1F3 -> f3
	Aq1F3 = aq1Presets.define("f3")
	// Aq1M1 -> m1
	Aq1M1 = aq1Presets.define("m1")
	// Aq1M2 -> m2
	Aq1M2 = aq1Presets.define("m2")
	// Aq1R1 -> r1
	Aq1R1 = aq1Presets.define("r1")
	// Aq1DVD -> dvd
	Aq1DVD = aq1Presets.define("dvd")
	// Aq1IMD1 -> imd1
	Aq1IMD1 = aq1Presets.define("imd1")
	// Aq1JGR -> jgr
	Aq1JGR = aq1Presets.define("jgr")
)

// Aq2Preset defines the presets for aquestak2
type Aq2Preset struct{ Preset }

var aq2Presets PresetList[Aq2Preset, *Aq2Preset]

var (
	// Aq2AqYukkuri -> aq_yukkuri
	Aq2AqYukkuri = aq2Presets.define("aq_yukkuri")
	// Aq2AqDefo1 -> aq_defo1
	Aq2AqDefo1 = aq2Presets.define("aq_defo1")
	// Aq2AqF1c -> aq_f1c
	Aq2AqF1c = aq2Presets.define("aq_f1c")
	// Aq2AqF3a -> aq_f3a
	Aq2AqF3a = aq2Presets.define("aq_f3a")
	// Aq2AqHuskey -> aq_huskey
	Aq2AqHuskey = aq2Presets.define("aq_huskey")
	// Aq2AqM4b -> aq_m4b
	Aq2AqM4b = aq2Presets.define("aq_m4b")
	// Aq2AqMf1 -> aq_mf1
	Aq2AqMf1 = aq2Presets.define("aq_mf1")
	// Aq2AqRb2 -> aq_rb2
	Aq2AqRb2 = aq2Presets.define("aq_rb2")
	// Aq2AqRb3 -> aq_rb3
	Aq2AqRb3 = aq2Presets.define("aq_rb3")
	// Aq2AqRm -> aq_rm
	Aq2AqRm = aq2Presets.define("aq_rm")
	// Aq2AqRobo -> aq_robo
	Aq2AqRobo = aq2Presets.define("aq_robo")
	// Aq2AqTeto1 -> aq_teto1
	Aq2AqTeto1 = aq2Presets.define("aq_teto1")
	// Aq2ArF4 -> ar_f4
	Aq2ArF4 = aq2Presets.define("ar_f4")
	// Aq2ArM5 -> ar_m5
	Aq2ArM5 = aq2Presets.define("ar_m5")
	// Aq2ArMF2 -> ar_mf2
	Aq2ArMF2 = aq2Presets.define("ar_mf2")
	// Aq2ArRM3 -> ar_rm3
	Aq2ArRM3 = aq2Presets.define("ar_rm3")
)

// Aq10Preset defines the presets for aquestak10
type Aq10Preset struct{ Preset }

var aq10Presets PresetList[Aq10Preset, *Aq10Preset]

var (
	// Aq10Custom -> custom
	Aq10Custom = aq10Presets.define("custom")
	// Aq10F1 -> f1
	Aq10F1 = aq10Presets.define("f1")
	// Aq10F2 -> f2
	Aq10F2 = aq10Presets.define("f2")
	// Aq10F3 -> f3
	Aq10F3 = aq10Presets.define("f3")
	// Aq10M1 -> m1
	Aq10M1 = aq10Presets.define("m1")
	// Aq10M2 -> m2
	Aq10M2 = aq10Presets.define("m2")
	// Aq10R1 -> r1
	Aq10R1 = aq10Presets.define("r1")
	// Aq10R2 -> r2
	Aq10R2 = aq10Presets.define("r2")
)
