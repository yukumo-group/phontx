package aqdata

type Aq1Preset struct{ preset }

var aq1Presets = presetList[Aq1Preset]{
	build: func(index int, name string) Aq1Preset {
		return Aq1Preset{preset{index, name}}
	},
}

var (
	Aq1F1   = aq1Presets.define("f1")
	Aq1F2   = aq1Presets.define("f2")
	Aq1F3   = aq1Presets.define("f3")
	Aq1M1   = aq1Presets.define("m1")
	Aq1M2   = aq1Presets.define("m2")
	Aq1R1   = aq1Presets.define("r1")
	Aq1DVD  = aq1Presets.define("dvd")
	Aq1IMD1 = aq1Presets.define("imd1")
	Aq1JGR  = aq1Presets.define("jgr")
)

type Aq2Preset struct{ preset }

var aq2 = presetList[Aq2Preset]{
	build: func(index int, name string) Aq2Preset {
		return Aq2Preset{preset{index, name}}
	},
}

var (
	Aq2AqYukkuri = aq2.define("aq_yukkuri")
	Aq2AqDefo1   = aq2.define("aq_defo1")
	Aq2AqF1c     = aq2.define("aq_f1c")
	Aq2AqF3a     = aq2.define("aq_f3a")
	Aq2AqHuskey  = aq2.define("aq_huskey")
	Aq2AqM4b     = aq2.define("aq_m4b")
	Aq2AqMf1     = aq2.define("aq_mf1")
	Aq2AqRb2     = aq2.define("aq_rb2")
	Aq2AqRb3     = aq2.define("aq_rb3")
	Aq2AqRm      = aq2.define("aq_rm")
	Aq2AqRobo    = aq2.define("aq_robo")
	Aq2AqTeto1   = aq2.define("aq_teto1")
	Aq2ArF4      = aq2.define("ar_f4")
	Aq2ArM5      = aq2.define("ar_m5")
	Aq2ArMF2     = aq2.define("ar_mf2")
	Aq2ArRM3     = aq2.define("ar_rm3")
)

type Aq10Preset struct{ preset }

var aq10 = presetList[Aq10Preset]{
	build: func(index int, name string) Aq10Preset {
		return Aq10Preset{preset{index, name}}
	},
}

var (
	Aq10Custom = aq10.define("custom")
	Aq10F1     = aq10.define("f1")
	Aq10F2     = aq10.define("f2")
	Aq10F3     = aq10.define("f3")
	Aq10M1     = aq10.define("m1")
	Aq10M2     = aq10.define("m2")
	Aq10R1     = aq10.define("r1")
	Aq10R2     = aq10.define("r2")
)
