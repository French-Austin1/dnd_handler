package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"github.com/fatih/color"
	//"french-austin.com/dungeon_handler/v2/Handler"
)

type Spell struct {
	Index         string   `json:"index"`
	Name          string   `json:"name"`
	Desc          []string `json:"desc"`
	HigherLevel   []string `json:"higher_level"`
	Range         string   `json:"range"`
	Components    []string `json:"components"`
	Material      string   `json:"material"`
	Ritual        bool     `json:"ritual"`
	Duration      string   `json:"duration"`
	Concentration bool     `json:"concentration"`
	CastingTime   string   `json:"casting_time"`
	Level         int      `json:"level"`
	AttackType    string   `json:"attack_type"`
	Damage        struct {
		DamageType struct {
			Index string `json:"index"`
			Name  string `json:"name"`
			URL   string `json:"url"`
		} `json:"damage_type"`
		DamageAtSlotLevel struct {

			Num2 string `json:"2"`

			Num3 string `json:"3"`
			Num4 string `json:"4"`

			Num5 string `json:"5"`

			Num6 string `json:"6"`

			Num7 string `json:"7"`

			Num8 string `json:"8"`

			Num9 string `json:"9"`

		} `json:"damage_at_slot_level"`
	} `json:"damage"`
	School struct {
		Index string `json:"index"`
		Name  string `json:"name"`

		URL   string `json:"url"`
	} `json:"school"`
	Classes []struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		URL   string `json:"url"`

	} `json:"classes"`
	Subclasses []struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		URL   string `json:"url"`

	} `json:"subclasses"`
	URL string `json:"url"`
}

type GenCategory struct {
	Count   int `json:"count"`
	Results []struct {
		Index string `json:"index"`

		Name  string `json:"name"`
		URL   string `json:"url"`
	} `json:"results"`
}

func GetRequestBody(category string, context string) []byte {
    requestURL := fmt.Sprintf("https://dnd5eapi.co/api/"+category+"/"+context)
    req, err := http.NewRequest(http.MethodGet, requestURL, nil)
    if err != nil {
        fmt.Printf("client: error making http request: %s\n", err)
        os.Exit(1)
    }

    res, err := http.DefaultClient.Do(req)
    if err != nil {
        fmt.Printf("client: could not read respones body: %s\n", err)
        os.Exit(1)
    }

    resBody, err := io.ReadAll(res.Body)
    if err != nil {
        fmt.Printf("client: could not read response body: %s\n", err)
        os.Exit(1)
    }
    return resBody
}

func GetGeneric(category string) {
    resBody := GetRequestBody(category, "")

    blueUnder := color.New(color.FgBlue, color.Underline)
    bMag := color.New(color.FgMagenta, color.Bold)

    var generic_res GenCategory
    json.Unmarshal([]byte(resBody), &generic_res)

    for i := 0; i < len(generic_res.Results); i++ {
        blueUnder.Print("Name:")
        bMag.Println(" "+generic_res.Results[i].Name)
    }
}

func GetSpell(category string, context string) {
    resBody := GetRequestBody(category, context)

    blueUnder := color.New(color.FgBlue, color.Underline)
    bMag := color.New(color.FgMagenta, color.Bold)

    var spell Spell
    json.Unmarshal([]byte(resBody), &spell)

    blueUnder.Print("Name:")
    bMag.Println(" "+spell.Name)
    fmt.Println()
    WriteAtt(
        "------Description------",
        spell.Desc[0],
    )
    WriteAtt(
        "---------Damage--------",
        "Slot Level -- Damage At Slot Level",
        "  2nd ------- "+spell.Damage.DamageAtSlotLevel.Num2,
        "  3rd ------- "+spell.Damage.DamageAtSlotLevel.Num3,
        "  4th ------- "+spell.Damage.DamageAtSlotLevel.Num4,
        "  5th ------- "+spell.Damage.DamageAtSlotLevel.Num5,
        "  6th ------- "+spell.Damage.DamageAtSlotLevel.Num6,
        "  7th ------- "+spell.Damage.DamageAtSlotLevel.Num7,
        "  8th ------- "+spell.Damage.DamageAtSlotLevel.Num8,
        "  9th ------- "+spell.Damage.DamageAtSlotLevel.Num9,
    )
    WriteAtt("--------Classes--------")
    for i := 0; i < len(spell.Classes); i++ {
        bMag.Println(spell.Classes[i].Name)
    }
}

func WriteAtt (args ...string) {
    blueUnder := color.New(color.FgBlue, color.Underline)
    cyn := color.New(color.FgCyan, color.Underline)
    bMag := color.New(color.FgMagenta, color.Bold)
    switch i := len(args); i {
        case 1:
            blueUnder.Println(args[0])
        case 2:
            blueUnder.Println(args[0])
            bMag.Println(args[1])
        default:
            blueUnder.Println(args[0])
            cyn.Println(args[1])
            for ii := 2; ii < len(args); ii++ {
                bMag.Println(args[ii])
         }
    }
}

func main()  {
    const helpMessage = "Welcome the the Dungeon Handler!\n Available Commands:\n- spells\n- ablility-scores\n- alignments\n- backgrounds\n- classes\n- conditions\n- damage-types\n- equipment\n- equipment-categories\n- feats\n- features\n- languages\n- magic-items\n- magic-schools\n- monsters\n- proficiencies\n- races\n- rules\n- skills\n- subclasses\n- subraces\n- traits\n- weapon-properties\n"
    var context string

    help := flag.Bool("h", false, helpMessage)
    flag.StringVar(&context, "c", "", "The index you would like to get more information on")
    flag.Parse()

    if *help {
        fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS]\n", os.Args[0])
        flag.PrintDefaults()
        os.Exit(0)
    }

    if flag.NArg() < 1 {
        fmt.Fprintf(os.Stderr, "Error: Missing argument\n")
        fmt.Fprintf(os.Stderr, "Error: Usage %s [OPTIONS] ARGUMENT\n", os.Args[0])
        flag.PrintDefaults()
        os.Exit(1)
    }

    category := flag.Arg(0)
    fmt.Println(category)
    fmt.Println(context)


    // match the value for context
    switch context {
        case "":
            GetGeneric(category)
        default:
            switch category {
                case "spells":
                    GetSpell(category, context)
            }
    }



}


