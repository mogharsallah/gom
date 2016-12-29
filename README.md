[![CircleCI](https://circleci.com/gh/medhoover/gom.svg?style=svg&circle-token=45019dc7f97b86994b79a44e66305018efd9a22f)](https://circleci.com/gh/medhoover/gom) **_Beta_**

**Got a question? / Want to collaborate?** Feel free to open an [issue](https://github.com/medhoover/gom/issues) or reach me at <gmedhoover@gmail.com>

# Introduction

__gom__ is a Powerful commands manager that simplifies complex scripts execution by defining aliases and execution policies. gom reads your commands from the _YAML_ [config file](#config-file), and execute the passed command alias.

# Usage

[![asciicast](https://asciinema.org/a/1ulj96gi7v6lyk2tqt2mn9oj3.png)](https://asciinema.org/a/1ulj96gi7v6lyk2tqt2mn9oj3)

### Create a config file
Create a configuration file named `gom.yaml` and add your commands. More information in [config file](#config-file)

### Launch your commands!
Just type :  ```$ gom [option] command [command_option...]```

For more details check the help manual : `$ gom -h`

# Install

If you have [**_GO_**](https://golang.org) installed you just need  to type:  
` $ go get github.com/medhoover/gom `
___

For others, since this project is still in beta, you would have to do a manual install. Just download the binary file from [release page](https://github.com/medhoover/gom/releases),  
 rename it, and place it in one of your path folders.

# Config file

By default, gom looks for `gom.yaml` file in your current directory. But you can manually set the file location by (-f, --file) flags:

`$ gom -f path/to/file.yaml ...`

The config file follows the YAML syntax, and contains the name and the commands properties:
```yaml
name: gom
commands:
# To execute a command, use its path name
# Example: $ gom install
  install:
    - go fmt github.com/medhoover/gom
    - go install github.com/medhoover/gom
  greet:
    morning: echo $GREET_M $USER !
    evening: echo $GREET_E $USER !
# Use -e flag to set an environment
# Example: $ gom -e fr greet morning
env:
  fr:
    GREET_M: Bonjour
    GREET_E: Bonne nuit
  en:
    GREET_M: Good morning
    GREET_E: Good Evening

```

A command have a string value, which gets executed using its YAML path: `$ gom greet evening`

An array of strings is considered as a command list, which gets executed sequently: `$ gom install`

Define environment variables within the 'env' property. Set the defined variables using the (-e , --environment) flags: `$ gom -e en greet morning`


# TODO

- Add unit tests
- Add config file initialisation
- Add support for parallel execution
- Simplify installation

# Questions and issues

Please feel free to share your questions and issues in [here](https://github.com/medhoover/gom/issues), I will be happy to help.

# Collaborate

Please do :smiley:.
