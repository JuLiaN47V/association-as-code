package main

import (
	"aas/config"
	"flag"
	"html/template"
	"net/http"
	"strings"
	"os"
	"github.com/gin-gonic/gin"
	"regexp"
)

func toLower(input string) string {
    return strings.ToLower(input)
}



func main() {
	
	configFilePath := "config.yaml"
	configStruct := config.ReadConfig(configFilePath)
	languageStruct := config.ReadLanguage("langs/"+configStruct.LangFile)


	// Flags
	var watchFlag bool
	var debugFlag bool
	flag.BoolVar(&watchFlag, "w", false, "Watch config file for changes")
	flag.BoolVar(&debugFlag, "d", false, "Debug mode")
	flag.Parse()

	if watchFlag{
		go config.Watch(configStruct, &configFilePath)
	}

	if !debugFlag {
		gin.SetMode(gin.ReleaseMode)
	}
	// End Flags


	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*")


	// Start Config Parsing
	for department_iteration, department := range configStruct.Body.Departments{
		for widget_iteration, widget := range department.BFVWidgets{
			switch toLower(widget.Type) {
			case "compoundtable":
				continue
			case "teamfull":
				continue
			case "teamsmall":
				continue
			case "clubgames":
				continue
			case "clubvideo":
				continue
			case "compoundvideo":
				continue
			case "teamliveticker":
				continue
			case "compound":
				continue
			case "cup":
				continue
			case "gamereport":
				continue
			default:
				configStruct.Body.Departments[department_iteration].BFVWidgets[widget_iteration].Type = "invalid"
			}
		}
	}

	// End Config Parsing


	// Start theme-colors modification
	// Read bootstrap css file
	bootstrapPath := "static/css/bootstrap.css"
	b, err := os.ReadFile(bootstrapPath)
    if err != nil {
        os.Exit(1)
    }
	customSCSScontent := string(b)
	// End Read bootstrap css file
	// Find Primary Color 
	primaryRegex, _ := regexp.Compile("--bs-primary-rgb:\\s?\\w{1,},\\s?\\w{1,},\\s?\\w{1,};")
	primary := primaryRegex.FindString(customSCSScontent)

	primaryRGBRegex, _ := regexp.Compile("\\s?\\w{1,},\\s?\\w{1,},\\s?\\w{1,}")
	primaryRGB := primaryRGBRegex.FindString(primary)
	// End Find Primary Color
	// Replace Primary Color
	if primaryRGB != configStruct.Head.BackgroundColor {
		newContents := strings.Replace(string(customSCSScontent), primaryRGB, configStruct.Head.BackgroundColor, -1)
		err = os.WriteFile(bootstrapPath, []byte(newContents), 0)
		if err != nil {
			panic(err)
		}
	}
	// End Replace Primary Color
	// Reread bootstrap css file
	b, err = os.ReadFile(bootstrapPath)
    if err != nil {
        os.Exit(1)
    }
	customSCSScontent = string(b)
    // End Reread bootstrap css file
	// Find Primary Color 
	secondaryRegex, _ := regexp.Compile("--bs-secondary-rgb:\\s?\\w{1,},\\s?\\w{1,},\\s?\\w{1,};")
	secondary := secondaryRegex.FindString(customSCSScontent)

	secondaryRGBRegex, _ := regexp.Compile("\\s?\\w{1,},\\s?\\w{1,},\\s?\\w{1,}")
	secondaryRGB := secondaryRGBRegex.FindString(secondary)
	// End Find Primary Color
	// Replace Primary Color
	if secondaryRGB != configStruct.Head.TextColor {
		newContents := strings.Replace(string(customSCSScontent), secondaryRGB, configStruct.Head.TextColor, -1)
		err = os.WriteFile(bootstrapPath, []byte(newContents), 0)
		if err != nil {
			panic(err)
		}
	}
	// End Replace Primary Color

	// End theme-colors modification

	// Static Router
	router.Static("/static", "static")
	// End Static Router

	// Custom Pages Router
	for _, customPage := range configStruct.CustomPages{
		router.GET(customPage.Name, func(c *gin.Context) {
			c.HTML(http.StatusOK, customPage.File, gin.H{
				"uri":    "http://" + c.Request.Host,
				"config": configStruct,
				"lang": languageStruct,
				"customPages" : configStruct.CustomPages,
			})
		})
	}
	// End Custom Pages Router

	// Index Router
	router.SetFuncMap(template.FuncMap{
        "toLower": toLower,
    })
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"uri":    "http://" + c.Request.Host,
			"config": configStruct,
			"lang": languageStruct,
			"customPages" : configStruct.CustomPages,
		})
	})
   // End Index Router


	if debugFlag {
		router.Run(":8090")
	} else{
		router.Run("0.0.0.0:80")
	}
	
}
