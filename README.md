# dockerauthgenerator
S simple cli tool to generate a docker auth json 

## Usage
Dockerauthgenerator creates a json file which can be uses to authenticate againt a docker registry.

The out will be similar to this and will be printed on stdout:

```json
{
  "auths": {
    "some.registry.com:5000": {
      "auth": "bG9naW5OYW1lOnNvbWVQYXNzd29yZA=="
    }
  }
}
```

### Password as Parameter

```bash
dockerauthgenerator -r some.registry.com:5000 -l loginName -p somePassword
```

### Password read from Terminal

If the password parameter is omitted the password is read from terminal
```bash
dockerauthgenerator -r some.registry.com:5000 -l loginName
```


### Password read from stdin
```bash
echo somePassword | dockerauthgenerator -r some.registry.com:5000 -l loginName --s
```

