# opai-cli

A CLI for running [OpenAI](https://openai.com/) completions.

## Build
```shell
go build
```

## Usage
```shell
❯ ./opai -h
opai - CLI for OpenAI completions

Usage:
  opai [flags]

Flags:
      --config string         config file (default is $HOME/.opai/config.yaml)
  -h, --help                  help for opai
      --max-tokens int        max tokens (default 256)
  -m, --model string          model name (default "text-davinci-003")
  -t, --temperature float32   temperature (default 0.1)
      --token string          OpenAI API token (default from config file      
```

## Sample Config
```yaml
# ~/.opai/config.yaml

token: <OpenAI API token>
model: "text-davinci-003"
temperature: 0.1
```

## Sample invocation
```shell
❯ ./opai "write a linux cmd to list all files in a directory larger than 1GB"
find /path/to/directory -type f -size +1G -exec ls -lh {} \;
```
