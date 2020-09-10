# Dockerauthgenerator

Dockerauthgenerator is simple cli tool to generate a docker auth json file. 

## Usage

Dockerauthgenerator creates a json file which can be used to authenticate against a docker registry.

The output will be similar to this and will be printed on stdout:

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

Most simple way is to provide the password as flag argument:

```bash
dockerauthgenerator -r some.registry.com:5000 -l loginName -p somePassword
```

### Password read from Terminal

If the password parameter is omitted,  the password is read from terminal:

```bash
dockerauthgenerator -r some.registry.com:5000 -l loginName
```

### Password read from stdin

You can also pipe the password from stdin:
  
```bash
echo somePassword | dockerauthgenerator -r some.registry.com:5000 -l loginName --s
```
