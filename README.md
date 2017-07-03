# Destiny Manifest Trimmer

Tool that will trim and transform the full DestinyManifest.json into a different format. 

It will read the DestinyManifest.json within the same directory, transform the data, and output it as MiniMani.json.

### **Note: The supplied Manifest was generated years ago and is to be updated**

## Manifest Layout
Original
```
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
```
{
  "DestinyActivityDefinition": {
    "1005705920": {
      "name": ""
      "icon": ""
      ...
    }
  },
  "DestinyActivityTypeDefinition": {
  },
  ...
}

