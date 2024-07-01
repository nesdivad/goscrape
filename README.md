# Goscrape

## Configuration file

The tool uses a configuration file to define starting URL, rulesets etc.

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
        "fileName": "goscrape"
    }
}
```
