# Goscrape

## Configuration file

The tool uses a configuration file to define starting URL, rulesets etc.

### Example

```json
{
    "url": "https://ndla.no/",
    "rules": [
        {
            "querySelector": "article[data-ndla-article]",
            "titleSelector": "h1[data-style=h1-resource]",
            "excerptSelector": "div[class*=ingress]",
            "contentSelector": "p"
        }
        ...
    ],
    "depth": 1
}
```