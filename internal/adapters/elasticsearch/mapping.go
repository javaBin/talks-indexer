package elasticsearch

// TalkIndexMapping defines the Elasticsearch mapping for the talks index.
// This mapping optimizes for search and filtering on talk metadata.
const TalkIndexMapping = `{
  "settings": {
    "number_of_shards": 1,
    "number_of_replicas": 1,
    "analysis": {
      "analyzer": {
        "default": {
          "type": "standard"
        }
      }
    }
  },
  "mappings": {
    "properties": {
      "id": {
        "type": "keyword"
      },
      "conferenceId": {
        "type": "keyword"
      },
      "conferenceSlug": {
        "type": "keyword"
      },
      "title": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      },
      "abstract": {
        "type": "text"
      },
      "intendedAudience": {
        "type": "text"
      },
      "language": {
        "type": "keyword"
      },
      "format": {
        "type": "keyword"
      },
      "level": {
        "type": "keyword"
      },
      "keywords": {
        "type": "keyword"
      },
      "status": {
        "type": "keyword"
      },
      "room": {
        "type": "keyword"
      },
      "startTime": {
        "type": "date",
        "format": "strict_date_optional_time||epoch_millis"
      },
      "endTime": {
        "type": "date",
        "format": "strict_date_optional_time||epoch_millis"
      },
      "speakers": {
        "type": "nested",
        "properties": {
          "id": {
            "type": "keyword"
          },
          "name": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "bio": {
            "type": "text"
          },
          "twitter": {
            "type": "keyword"
          },
          "pictureUrl": {
            "type": "keyword",
            "index": false
          }
        }
      },
      "submitterEmail": {
        "type": "keyword"
      },
      "created": {
        "type": "date",
        "format": "strict_date_optional_time||epoch_millis"
      },
      "lastUpdated": {
        "type": "date",
        "format": "strict_date_optional_time||epoch_millis"
      }
    }
  }
}`
