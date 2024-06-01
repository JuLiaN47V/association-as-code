package main

import (
	"aas/config"
	"flag"
	"html/template"
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
)

func toLower(input string) string {
    return strings.ToLower(input)
}



func main() {
	
	configFilePath := "config.yaml"
	configStruct := config.ReadConfig(configFilePath)
	println(&configStruct)	
	languageStruct := config.ReadLanguage("langs/"+configStruct.LangFile)
	println(&languageStruct.Calendars)	
	println(configStruct.CustomPages[0].Name)


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
