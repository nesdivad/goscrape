# Goscrape

## Get started

1. To run the program locally, you need [Go](https://go.dev/dl/) installed on your machine.
2. Install it on your local path by running `install.sh` on Linux/Mac, and running `goscrape [arguments]`.
   Alternatively, run the program with the Go CLI, using `go run main.go [arguments]`.

## Help

After installation, run the program with the `--help` argument for a description of the arguments.

## Configuration file

The tool uses a configuration file to define starting URL, rulesets etc.
There are two ways to feed the configuration into the program:
1. Using `--config` flag with the path to the config file.
2. Using `--configjson` flag containing the config as a json-string, preferably compacted and without whitespace, newlines etc.

### Example

```json
{
    "url": "https://ndla.no/subject:1:d92be649-8bda-4514-b04d-2d3c5251aa79/topic:e894b2c5-f7f7-4598-ada8-221d18fba875/topic:0c514c1d-0207-4ac4-b042-126fa5a9acee/resource:0fce6cd6-0db5-47d5-b6ef-65021dbf2497",
    "rules": [
        {
            "querySelector": "article[data-ndla-article]",
            "titleSelector": "h1[data-style=h1-resource]",
            "excerptSelector": "div[class*=ingress]",
            "contentSelector": "p"
        },
        ...
    ],
    "urlFilters": [
        "https://ndla.no/article/erklaering-for-informasjonskapsler",
        ...
    ],
    "settings": {
        "depth": 1,
        "limitRules": [
            {
                "domainGlob":"*ndla.*",
                "parallelism": 2,
                "randomDelay": 3
            }
        ]
    },
    "output": {
        "path": "output",
        "fileType": "json | jsonl",
        "fileName": "goscrape",
        "chunk": 5 //Only needed if fileType == "jsonl". Defaults to 5.
    }
}
```
