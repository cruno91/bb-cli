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
bb get project -w <workspace> -n <project>

# Create a project
bb create project -w <workspace> -n "<name>" -k <key> -p -d "<description>"

# List repositories
bb list repositories -w <workspace> -n <project>

# Create repository
bb create repository -w <workspace> -n <project> -r <repository slug> -p

# Add access key
bb add access-key -w <workspace> -r <repository slug> -k <access key> -l <access key label>

# Add a webhook
bb add webhook -w <workspace> -r <repositor slugy> -u <webhook url> -d <webhook label> --events repo:push repo:update

# Add pipeline variable
bb add variable -w <workspace> -r <repository slug> -v <value> -l <label> -s
```
