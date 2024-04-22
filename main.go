package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
	"github.com/fsnotify/fsnotify"
	"log"
)


func watchConfig(configYAML *map[string]interface{}, configFilePath *string){
	watcher, err := fsnotify.NewWatcher()
    if err != nil {
        log.Fatal(err)
    }
    defer watcher.Close()
	go func() {
        for {
            select {
            case event, ok := <-watcher.Events:
                if !ok {
                    return
                }
                if event.Has(fsnotify.Write) {
					yamlFile, err := os.ReadFile(*configFilePath)
					if err != nil {
						fmt.Printf("yamlFile.Get err #%v ", err)
					}
					err = yaml.Unmarshal(yamlFile, &configYAML)
					if err != nil {
						fmt.Printf("Unmarshal: %v", err)
					}
                }
            case err, ok := <-watcher.Errors:
                if !ok {
                    return
                }
                log.Println("error:", err)
            }
        }
    }()
	    // Add a path.
		err = watcher.Add("config.yaml")
		if err != nil {
			log.Fatal(err)
		}
	
		// Block main goroutine forever.
		<-make(chan struct{})
}

func getConfig(path string) (map[string]interface{}){


	configYAML := make(map[string]interface{})

	yamlFile, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("yamlFile.Get err #%v ", err)
		os.Exit(1)
	}
	err = yaml.Unmarshal(yamlFile, configYAML)
	if err != nil {
		fmt.Printf("Unmarshal: %v", err)
		os.Exit(1)
	}

	
	
	return configYAML
}


func main() {
	configFilePath := "config.yaml"
	configYAML := getConfig(configFilePath)
	langFile, _ := configYAML["lang_file"].(string)
	var langYAML map[string]interface{}
	if _, ok := configYAML["lang_file"].(string); ok {
		langFilePath := "langs/"+langFile
		langYAML = getConfig(langFilePath)
	}



	var watchFlag bool
	var debugFlag bool
	flag.BoolVar(&watchFlag, "w", false, "Watch config file for changes")
	flag.BoolVar(&debugFlag, "d", false, "Debug mode")
	flag.Parse()

	if watchFlag{
		go watchConfig(&configYAML, &configFilePath)
	}

	if !debugFlag {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	router.Static("/static", "static")

	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"uri":    "http://" + c.Request.Host,
			"config": configYAML,
			"lang": langYAML,
		})
	})
	router.GET("/dataprotection", func(c *gin.Context) {
		c.HTML(http.StatusOK, "dataprotection.html", gin.H{
			"uri":    "http://" + c.Request.Host,
			"config": configYAML,
		})
	})

	if debugFlag {
		router.Run(":8090")
	} else{
		router.Run("0.0.0.0:80")
	}
	
}
