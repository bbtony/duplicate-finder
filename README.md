# DUPLICATE-FINDER

**THIS IS MVP**

**NOT PRODUCTION READY**

In developing.

**Duplicate-Finder** is the **simple** tool which **find duplicate** key and value in different dotenv-files.

This util was produced for DevSecOps integration one of company. The main goal provides information about 
duplicates in different dotenv files. 

### USAGE

example of usage:
```bash
    dpf find -o report.json -t value -v testdata/env1 testdata/env2 testdata/env3
```

-v verbose:
```
Duplicate value: val1
File: testdata/env1 Variable: test 
File: testdata/env1 Variable: fish 
File: testdata/env2 Variable: test2
```
Examples of work [by value](examples/report_by_value.json) and [by key](examples/report_by_key.json)

Combine with `jq` util:

```bash
    dpf find -o report.json testdata/env1 testdata/env2 testdata/env3 | jq  '."result"[] | ."matches"'
```

after we get the next result:
```
"test"
"val1"
```

The power of this tool will open if you combine result of work with [trufflehog](https://github.com/trufflesecurity/trufflehog) 
or [gitleaks](https://github.com/gitleaks/gitleaks) and etc.

### DEPENDENCIES

**Duplicate-Finder** is used [godotenv](github.com/joho/godotenv) for reading dotenv-files

### ROADMAP
- use JSON output format like base-format for stdout and jq-combination