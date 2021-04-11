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
	//fmt.Println(chr.Sheet())
	app := cli.NewApp()
	app.Version = "v 0.0.3"
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
			Name:        "create",
			Usage:       "запускает процесс создания PC/NPC с учетом всех флагов и переменных",
			UsageText:   "ТУДУ: сделать очень подробное описание команды create",
			Description: "ТУДУ: сделать ltcrhbgity команды create",
			Action: func(c *cli.Context) error {
				fmt.Println("Start Action")
				fmt.Println(c.GlobalBool("autoroll"))
				fmt.Println("Step 1 & 2")
				chr := character.New(c.GlobalBool("autoroll"))
				chr.SetAttributes()
				fmt.Println(chr.Sheet())
				fmt.Println("Step 3")
				chr.SetBackground()
				fmt.Println(chr.Sheet())
				fmt.Println("Step 4 & 5")
				chr.SetSkills()
				fmt.Println(chr.Sheet())
				fmt.Println("Step 6")
				chr.SetClass()
				fmt.Println(chr.Sheet())
				fmt.Println("Step 7")
				chr.SetFoci()
				fmt.Println(chr.Sheet())
				fmt.Println("Step 9")
				fmt.Println("Pick one skill of your choice to reflect your\nhero’s outside interests, natural talents, hobby expertise,\nor other personal focus.")
				chr.Train("Any Skill")
				fmt.Println(chr.Sheet())

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
