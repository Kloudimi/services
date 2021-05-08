---
title: sentiment
servicename: sentiment
labels: 
- Readme
---
Real time sentiment analysis

# Sentiment Service

The sentiment service provides rudimentary sentiment analysis on text


## cURL


### Sentiment Analyze
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->

```shell
> curl 'https://api.m3o.com/sentiment/Sentiment/Analyze' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "lang": "string",
  "text": "string"
};
# Response
{
  "score": 1
}
```


