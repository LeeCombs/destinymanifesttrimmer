# Destiny Manifest Trimmer

Tool that will trim and transform the full DestinyManifest.json into a different format. It will read the DestinyManifest.json within the same directory, transform the data, and output it as MiniMani.json.

### **Note: The supplied Manifest was generated years ago and is to be updated**

## Element Access
Original Access
```go
manifest.Manifest[0].DestinyActivityDefinition[0].ActivityName
manifest.Manifest[0].DestinyActivityDefinition[1].ActivityName
manifest.Manifest[1].DestinyActivityTypeDefinition[0].ActivityTypeName
```
Converted Access By Hash
```go
manifest.DestinyActivityDefinition[1005705920].ActivityName
manifest.DestinyActivityDefinition[2001232120].ActivityName
manifest.DestinyActivityTypeDefinition[2751293122].ActivityTypeName
```

## Manifest Layout
Original
```json
{
  "manifest":[{
      "DestinyActivityDefinition": [{
        "hash": 1005705920
        "name": ""
        "icon": ""
        ...
      }]
    },{
      "DestinyActivityTypeDefinition": [{}]
    }
    ...
```

Converted:
```json
{
  "DestinyActivityDefinition": {
    "1005705920": {
      "name": ""
      "icon": ""
    }
  },
  "DestinyActivityTypeDefinition": {
  },
  ...
}

