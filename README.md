## Association As Code
This project is designed for associations in order to run a templated webserver.  
The Goal is to make it easy to run a simple, customizeable website.  
Licensed under [Apache 2.0](https://github.com/JuLiaN47V/association-as-code/blob/main/LICENSE.md)

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
See [config.yaml](https://github.com/JuLiaN47V/association-as-code/wiki/config.yaml) for detailed informations.
## Custom Language
