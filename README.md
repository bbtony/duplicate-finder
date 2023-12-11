# DUPLICATE-FINDER

**Duplicate-Finder** is the **simple** tool which **find duplicate** value in different dotenv-files.

I developed this tool for myself and for analyze different environments (Docker/K8S etc.) at work.

### USAGE

example of usage:
```bash
    duplicate-finder </path/envfile> </path/envfile> </path/envfile>
```

### DEPENDENCIES

**Duplicate-Finder** is used [godotenv](github.com/joho/godotenv) for reading dotenv-files

### ROADMAP
- use JSON output format like base-format for stdout and jq-combination
- zero dependencies - implement own solution for reading dotenv-files it will allow use only one repository