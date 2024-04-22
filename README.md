## Association As Code
This project is designed for associations in order to run a templated webserver.  
The Goal is to make it easy to run a simple, customizeable website.

## Getting started
You can either run the webserver within docker or directly from source.

### Docker

To run the pre-defined example webserver, run:
``` docker
docker run -d -p 8090:8090 julian47/association-as-code:webserver
```

For customization, you can add you own config- and language-file.  
To do that, you have to create the files config.yaml and en.yaml. See [Custom Config](#custom-config) and [Custom language](#custom-language) for specification and how to use your own language-file.  
After that, mount your config into :/app/config.yaml and your language-file to :/app/langs/en.yaml  

``` docker
-v config.yaml:/app/config.yaml -v en.yaml:/app/langs/en.yaml
```

### From Source
**Requirements**: golang/1.22.2  
First you habe to clone this repository:
``` bash
git clone https://github.com/JuLiaN47V/association-as-code.git && cd association-as-code
```
After that, simply run:
``` bash
go run main.go
```

Either way, you can access your webserver now at http://localhost:8090.

## Custom Config
The config got following top-level-keys:
  - head -> Banner + navbar
  - body -> Different departments of association
  - footer -> Social + Contact
  - fonts -> Custom Fonts
  - calendar -> Navbar dropdown menu
  - files -> Navbar dropdown menu
  - linked_sites -> Navbar element linked to different site
### Head
Under the head-key are following second-level-keys:
  - logo[String] -> Filename of logo-image-file
  - name[String] -> Name of association
  - icon[String] -> Filename of icon-file
  - description[String] -> Description display on google search
  - sloga -> Slogan above and underneath the main headline with custom font
#### Slogan
  - font[String] -> Fontname, Kind
  - top[String] -> Slogan above main headline
  - bottom[String] -> Slogan underneath main headline
### Body
The body is a list, consisting of maps as list-elemnts with the second-level-keys:
   - name[String] -> Headline of section
   - contact -> List of contacts
   - gallery -> List of images
#### Contact
Contact is a list, consisting of maps as list-elemnts with the third-level-keys:
  - name[String] -> Name of Person
  - title[String] -> Title in association
  - description[String] -> Description of work
  - tel[String] -> telephonenumber

#### Gallery


## Custom Language