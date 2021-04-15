package main

import (
	"fmt"
	"os"

	"github.com/Galdoba/WWN_tools/pkg/character"
	"github.com/urfave/cli"
)

func main() {
	fmt.Println("Start")
	//chr := character.New(true)
	//fmt.Print(chr.Sheet())
	app := cli.NewApp()
	app.Version = "v 0.1.0"
	app.Name = "wwncharmaker"
	app.Usage = "Создает персонажей для World Without Number в ручном и автоматическом режимах"
	app.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name:        "autoroll",
			Usage:       "Принудительно бросает дайсы во время всех этапов создания PC",
			EnvVar:      "",
			FilePath:    "",
			Required:    false,
			Hidden:      false,
			Destination: new(bool),
		},
	}
	app.Commands = []cli.Command{
		{
			Name:        "new",
			Usage:       "запускает процесс создания PC/NPC с учетом всех флагов и переменных",
			UsageText:   "ТУДУ: сделать очень подробное описание команды new",
			Description: "ТУДУ: сделать ltcrhbgity команды create",
			Action: func(c *cli.Context) error {
				chr := character.New(c.GlobalBool("autoroll"))
				fmt.Print(chr.Sheet())
				fmt.Println("Step 1 & 2")
				fmt.Println("Attributes and modifiers")

				chr.SetAttributes()
				fmt.Print(chr.Sheet())
				fmt.Println("Step 3")
				chr.SetBackground()
				fmt.Print(chr.Sheet())
				fmt.Println("Step 4 & 5")
				chr.SetSkills()
				fmt.Print(chr.Sheet())
				fmt.Println("Step 6")
				chr.SetClass()
				fmt.Print(chr.Sheet())
				fmt.Println("Step 7")
				chr.SetFoci()
				fmt.Print(chr.Sheet())
				fmt.Println("Step 8")
				fmt.Println("Special origins not implemented")
				fmt.Println("SKIP")
				fmt.Print(chr.Sheet())
				fmt.Println("Step 9")
				fmt.Println("Pick one skill of your choice to reflect your\nhero’s outside interests, natural talents, hobby expertise,\nor other personal focus.")
				chr.Train("Any Skill")
				fmt.Print(chr.Sheet())
				fmt.Println("Step 10 & 11")
				fmt.Println("Learning Arts and Spells")
				chr.SetMagicTraditions()
				fmt.Print(chr.Sheet())
				fmt.Println("Final Touches:")
				chr.SetHitPoints()
				fmt.Print(chr.Sheet())

				return nil
			},
		},
	}
	args := os.Args

	if err := app.Run(args); err != nil {
		fmt.Println(err.Error())
	}
}

//wwncharmaker create
