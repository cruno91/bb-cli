# bb-cli

## Description

This is an example of a specialized command line tool to interact with the 
Bitbucket API with a focus on scaffolding repositories and projects.

## Usage

```bash
# Show help
bb -h

# Override config file
bb --config=config.json

# Authenticate with Bitbucket
bb auth -t <token>

# List workspaces
bb list workspaces

# List projects
bb list projects -w <workspace>

# Get a project
bb get project -w <workspace> -p <project>

# Create a project
bb create project -w <workspace> -n "<name>" -k <key> -p -d "<description>"
```
