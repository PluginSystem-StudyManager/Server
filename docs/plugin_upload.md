# Plugin Upload

Plugins can be uploaded via a zip file. This file contains:
- plugin_info file
- Files needed when installed
    - jar file
- Files needed for marketplace
    - README
    - files referenced in README


## Plugin info file

description: File contains key value data about the plugin
file name: `plugin_info.yaml`
content:

```yaml
# [metadata]
name: TAB_NAME
shortDescription: SHORT_DESCRIPTION

# [metadata] optional
authors:
  - AUTHOR_1
version: '1.0'
repository: 'REPOSITORY_URL'
tags:
  - TAG_1
  - TAG_2
languages:
  - de
  - en

# [options] optional
allowMultipleInstallations: false

```

## Jar File

description: JavaFX Application build according to this guide: [TODO: link to development guide]
file name: `plugin_{PLUGIN_NAME}.jar`

## README

description: Description about the plugin. Can be shown on the marketplace or website.
file_name: `README.md`

## Files referenced in README

description: Images, ... . The path of the files must be correct.
file_name: same as in README


# Upload Process

Post Request to the server with the following data:

- Zip file
- authentication
    - Cookie: username and token

Response:

JSON with result