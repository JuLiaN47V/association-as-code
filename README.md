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

For customization, you can add you own config-file, language-file and customized bootstap.  
To do that, you have to create the files config.yaml and en.yaml. See [Custom Config](https://github.com/JuLiaN47V/association-as-code/wiki/config.yaml) and [Custom language](#custom-language) for specification and how to use your own language-file.  
After that, mount your config into :/app/config.yaml and your language-file to :/app/langs/en.yaml  
For bootstrap customization like color change, follow [Custom-CSS](https://github.com/JuLiaN47V/association-as-code/wiki/Custom-CSS) wiki.

``` docker
-v config.yaml:/app/config.yaml -v en.yaml:/app/langs/en.yaml -v static/css/bootstrap.css:/app/static/bootstrap/bootstrap.css
```

### From Source
**Requirements**: golang/1.22.2  
First you habe to clone this repository:
``` bash
git clone https://github.com/JuLiaN47V/association-as-code.git && cd association-as-code
```
After that, simply run for a example webserver:
``` bash
go run main.go -d
```

#### main.go Flags
-d Debug Mode  
-w Watch config.yaml for changes

Either way, you can access your webserver now at http://localhost:8090.

## Custom Config
See [config.yaml](https://github.com/JuLiaN47V/association-as-code/wiki/config.yaml) for detailed informations.
## Custom Language
See [Custom Language](https://github.com/JuLiaN47V/association-as-code/wiki/Custom-Language) for detailed Informations
## Custom CSS
If you want to make changes to default bootstrap configuration, follow [Custom-CSS](https://github.com/JuLiaN47V/association-as-code/wiki/Custom-CSS).
