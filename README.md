# bb-cli

## Description

This is a specialized command line tool to interact with the Bitbucket API
with a focus on scaffolding repositories and projects.

## Usage

```bash
# Show help
bb -h

# Overide config file
bb --config=config.json

# Authenticate with Bitbucket
bb auth -t <token>
```

## Todo

### 1.0.0

- [x] Add viper for configuration
- [x] Authenticate with Bitbucket via OAuth
- [x] List projects
- [] Create a project
- [] List repositories in a project
- [] Create a repository in a project
- [] Create an authentication key in a repository
- [] Create a pipeline variable in a repository
- [] Create branch restrictions in a repository
