package config

import (
	"fmt"
	"os"
	"gopkg.in/yaml.v2"
	"log"
	"github.com/fsnotify/fsnotify"
)

type Config struct {
	TLS bool `yaml:"tls"`
	LogToFile string `yaml:"log_to_file"`
	Theme string `yaml:"theme"`
	Head struct {
		Logo        string `yaml:"logo"`
		Name        string `yaml:"name"`
		Icon        string `yaml:"icon"`
		Description string `yaml:"description"`
		BackgroundColor string `yaml:"backgroundcolor"`
		TextColor string `yaml:"textcolor"`
		Slogan      struct {
			Font   string `yaml:"font"`
			Top    string `yaml:"top"`
			Bottom string `yaml:"bottom"`
		} `yaml:"slogan"`
	} `yaml:"head"`
	Body struct {
		Strapi struct {
			URL     string `yaml:"url"`
			Name    string `yaml:"name"`
		} `yaml:"strapi"`
		Departments []struct {
			Name     string `yaml:"name"`
			Contacts []struct {
				Name        string `yaml:"name"`
				Title       string `yaml:"title"`
				Description string `yaml:"description"`
				Email		string `yaml:"email"`
				Tel         string `yaml:"tel"`
			} `yaml:"contacts"`
			Gallery []struct {
				Title string `yaml:"title"`
				Src   string `yaml:"src"`
				Alt   string `yaml:"alt"`
			} `yaml:"gallery"`
			BFVWidgets []struct {
				TeamID string `yaml:"teamid"`
				ClubID string `yaml:"clubid"`
				CompoundID string `yaml:"compoundid"`
				Type string `yaml:"type"`
			} `yaml:"bfvwidgets"`
		} `yaml:"departments"`
	} `yaml:"body"`
	Footer struct {
		BackgroundImage string `yaml:"background_image"`
		Socials         []struct {
			Account string `yaml:"account"`
			Link    string `yaml:"link"`
			Image   string `yaml:"image"`
		} `yaml:"socials"`
		Contacts []struct {
			Title       string `yaml:"title"`
			Name        string `yaml:"name"`
			Address     string `yaml:"address"`
			Email       string `yaml:"email"`
			Tel         string `yaml:"tel"`
			Responsible bool   `yaml:"responsible"`
		} `yaml:"contacts"`
	} `yaml:"footer"`
	Fonts []struct {
		Name       string `yaml:"name"`
		Src        string `yaml:"src"`
		FontWeight string `yaml:"font_weight"`
		FontStyle  string `yaml:"font_style"`
	} `yaml:"fonts"`
	LangFile    string `yaml:"lang_file"`
	LinkedSites []struct {
		Name  string `yaml:"name"`
		URL   string `yaml:"url,omitempty"`
		Type  string `yaml:"type"`
		Links []struct {
			Name string `yaml:"name"`
			URL  string `yaml:"url"`
		} `yaml:"links,omitempty"`
	} `yaml:"linked_sites"`
	CustomPages []struct {
		Name string `yaml:"name"`
		File string `yaml:"file"`
	} `yaml:"custom_pages"`
}

type Language struct {
	Contacts       string `yaml:"contacts"`
	Gallery        string `yaml:"gallery"`
	Files          string `yaml:"files"`
	Calendars      string `yaml:"calendars"`
	ShowMore       string `yaml:"show_more"`
	ShowLess       string `yaml:"show_less"`
}


func ReadConfig(path string) (*Config){

	var configYAML Config
	yamlFile, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("yamlFile.Get err #%v ", err)
		os.Exit(1)
	}
	err = yaml.Unmarshal(yamlFile, &configYAML)
	if err != nil {
		fmt.Printf("Unmarshal: %v", err)
		os.Exit(1)
	}

	return &configYAML
}

func ReadLanguage(path string) (*Language){

	var languageYAML Language
	yamlFile, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("yamlFile.Get err #%v ", err)
		os.Exit(1)
	}
	err = yaml.Unmarshal(yamlFile, &languageYAML)
	if err != nil {
		fmt.Printf("Unmarshal: %v", err)
		os.Exit(1)
	}

	return &languageYAML
}

func Watch(configStruct *Config, configFilePath *string){
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
					err = yaml.Unmarshal(yamlFile, &configStruct)
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
		err = watcher.Add(*configFilePath)
		if err != nil {
			log.Fatal(err)
		}
	
		// Block main goroutine forever.
		<-make(chan struct{})
}
