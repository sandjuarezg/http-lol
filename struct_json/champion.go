package struct_json

import (
	"encoding/json"
	"errors"
	"fmt"
)

func GetChampion(body []byte, nameChampion string) (champion Champion, err error) {
	var aux interface{}
	var jsonbody []byte

	err = json.Unmarshal(body, &aux)
	if err != nil {
		fmt.Println(err)
		return
	}

	var champ map[string]interface{} = aux.(map[string]interface{})
	s, ok := champ["data"].(map[string]interface{})

	if ok {
		jsonbody, err = json.Marshal(s[nameChampion])
		if err != nil {
			fmt.Println(err)
			return
		}

		err = json.Unmarshal(jsonbody, &champion)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		err = errors.New("interface not match")
	}

	return
}

type Champion struct {
	Id    string `json:"id"`
	Key   string `json:"key"`
	Name  string `json:"name"`
	Title string `json:"title"`
	Image struct {
		Full   string `json:"full"`
		Sprite string `json:"sprite"`
		Group  string `json:"group"`
		X      int    `json:"x"`
		Y      int    `json:"y"`
		W      int    `json:"w"`
		H      int    `json:"h"`
	} `json:"image"`
	Skins []struct {
		Id      string `json:"id"`
		Num     int    `json:"num"`
		Name    string `json:"name"`
		Chromas bool   `json:"chromas"`
	} `json:"skins"`
	Lore      string   `json:"lore"`
	Blurb     string   `json:"blurb"`
	Allytips  []string `json:"allytips"`
	Enemytips []string `json:"enemytips"`
	Tags      []string `json:"tags"`
	Partype   string   `json:"partype"`
	Info      struct {
		Attack     int `json:"attack"`
		Defense    int `json:"defense"`
		Magic      int `json:"magic"`
		Difficulty int `json:"difficulty"`
	} `json:"info"`
	Stats struct {
		Hp                   float32 `json:"hp"`
		Hpperlevel           float32 `json:"hpperlevel"`
		Mp                   float32 `json:"mp"`
		Mpperlevel           float32 `json:"mpperlevel"`
		Movespeed            float32 `json:"movespeed"`
		Armor                float32 `json:"armor"`
		Armorperlevel        float32 `json:"armorperlevel"`
		Spellblock           float32 `json:"spellblock"`
		Spellblockperlevel   float32 `json:"spellblockperlevel"`
		Attackrange          float32 `json:"attackrange"`
		Hpregen              float32 `json:"hpregen"`
		Hpregenperlevel      float32 `json:"hpregenperlevel"`
		Mpregen              float32 `json:"mpregen"`
		Mpregenperlevel      float32 `json:"mpregenperlevel"`
		Crit                 float32 `json:"crit"`
		Critperlevel         float32 `json:"critperlevel"`
		Attackdamage         float32 `json:"attackdamage"`
		Attackdamageperlevel float32 `json:"attackdamageperlevel"`
		Attackspeedperlevel  float32 `json:"attackspeedperlevel"`
		Attackspeed          float32 `json:"attackspeed"`
	} `json:"stats"`
	Spells []struct {
		Id          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Tooltip     string `json:"tooltip"`
		Leveltip    struct {
			Label  []string `json:"label"`
			Effect []string `json:"effect"`
		} `json:"leveltip"`
		Maxrank      int           `json:"maxrank"`
		Cooldown     []float32     `json:"cooldown"`
		CooldownBurn string        `json:"cooldown_burn"`
		Cost         []int         `json:"cost"`
		CostBurn     int           `json:"cost_burn"`
		Datavalues   struct{}      `json:"datavalues"`
		Effect       []interface{} `json:"effect"`
		EffectBurn   []string      `json:"effect_burn"`
		Vars         []string      `json:"vars"`
		CostType     string        `json:"cost_type"`
		Maxammo      string        `json:"maxammo"`
		Range        []int         `json:"range"`
		RangeBurn    string        `json:"range_burn"`
		Image        struct {
			Full   string `json:"full"`
			Sprite string `json:"sprite"`
			Group  string `json:"group"`
			X      int    `json:"x"`
			Y      int    `json:"y"`
			W      int    `json:"w"`
			H      int    `json:"h"`
		} `json:"image"`
		Resource string `json:"resource"`
	} `json:"spells"`
	Passive struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Image       struct {
			Full   string `json:"full"`
			Sprite string `json:"sprite"`
			Group  string `json:"group"`
			X      int    `json:"x"`
			Y      int    `json:"y"`
			W      int    `json:"w"`
			H      int    `json:"h"`
		} `json:"image"`
	} `json:"passive"`
	Recommended []string `json:"recommended"`
}
