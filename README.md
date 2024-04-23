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
go run main.go -d
```

Either way, you can access your webserver now at http://localhost:8090.

## Custom Config
The config got following top-level-keys:
  - head[Map] -> Banner + navbar
  - body[List] -> Different departments of association
  - footer[MAP] -> Social + Contact
  - fonts[List] -> Custom Fonts
  - calendars[List] -> Navbar dropdown menu
  - files[List] -> Navbar dropdown menu
  - linked_sites[List] -> Navbar element linked to different site
### Head[MAP]
The head-key is a map with following second-level-keys:
  - logo[String] -> Filename of logo-image-file
  - name[String] -> Name of association
  - icon[String] -> Filename of icon-file
  - description[String] -> Description display on google search
  - slogan[MAP] -> Slogan above and underneath the main headline with custom font
#### Slogan[MAP]
  - font[String] -> Fontname, Kind
  - top[String] -> Slogan above main headline
  - bottom[String] -> Slogan underneath main headline
### Body[List]
The body is a list, consisting of maps as list-elemnts with the second-level-keys:
   - name[String] -> Headline of section
   - contact[List] -> List of contact maps
   - gallery[List] -> List of image maps
#### Contacts[List]
Contact is a list, consisting of maps as list-elemnts with the third-level-keys:
  - name[String] -> Name of Person
  - title[String] -> Title in association
  - description[String] -> Description of work
  - tel[String] -> telephonenumber
#### Gallery[List]
Gallery is a list, consisting of maps as list-elemnts with the third-level-keys:
  - title[String] -> Title of image
  - src[String] -> Filename of image-file
### Footer[Map]
The footer-key is a map with following second-level-keys:
   - background_image[String] -> Filename of image-file for background with low opacity
   - socials[List] -> List of social platform maps
   - contact[List] -> List of contact maps for association
#### Socials[List]
Socials is a list, consisting of maps as list-elemnts with the third-level-keys:
  - account[String] -> Name of Account
  - link[String] -> Link to account of social platform
  - image[String] -> Filename of image-file for social platform
#### Contact[List]
  - title[String] -> Title of Contact
  - name[String] -> Name of Contact
  - address[String] -> Address of Contact
  - email[String] -> eMail of Contact
  - tel[String] -> telephonenumber of contact
  - responsible[Bool] -> If Contact is responsible for the association
### Fonts[List]
  - name[String] -> Name of Font
  - src[String] -> Filename of font-file in side fonts directory
  - font_weight[String] -> Name of font weight
  - font_style[String] -> Name of font style
### Calendars[List]
Calendars is a list, consisting of maps as list-elemnts with the second-level-keys:
  - name[String] -> Name of calendar
  - url[String] -> URL of calendar
### Files[List]
Files is a list, consisting of maps as list-elemnts with the second-level-keys:
  - name[String] -> Name of File
  - url[String] -> URL of File
### linked_sites[List]
Linked_sites is a list, consisting of maps as list-elemnts with the second-level-keys:
  - name[String] -> Name of Site
  - url[String] -> URL of Site
### lang_file[String]
Filename of language-file to use inside "lang" directory

## Custom Language