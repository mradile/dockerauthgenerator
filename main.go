package main

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"

	"github.com/urfave/cli/v2"
	"golang.org/x/crypto/ssh/terminal"
)

var version = "0.0.1"

func main() {
	app := &cli.App{
		Name:  "docker auth generator",
		Usage: "create a docker auth json",
		Authors: []*cli.Author{{
			Name:  "Martin Radile",
			Email: "martin.radile@gmail.com",
		}},
		Description: "the tool creates a docker auth json with the login and password base64 encoded",
		Version:     version,
		Copyright:   "Copyright (c) 2020 Martin Radile",
		Action: func(c *cli.Context) error {
			return Run(
				c.String("registry"),
				c.String("login"),
				c.String("password"),
				c.Bool("password-stdin"),
			)
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "registry",
				Aliases:  []string{"r"},
				Usage:    "registry host with port",
				EnvVars:  []string{"REGISTRY"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     "login",
				Aliases:  []string{"l"},
				Usage:    "login",
				EnvVars:  []string{"LOGIN"},
				Required: true,
			},
			&cli.StringFlag{
				Name:    "password",
				Aliases: []string{"p"},
				Usage:   "provide a password as a flag",
				EnvVars: []string{"PASSWORD"},
			},
			&cli.BoolFlag{
				Name:    "password-stdin",
				Aliases: []string{"s"},
				Usage:   "read password from stdin",
				EnvVars: []string{"PASSWORD_STDIN"},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

type DockerAuth struct {
	Auths map[string]map[string]string `json:"auths"`
}

func Run(registry, login, password string, pwStdin bool) error {
	if pwStdin {

		pw, err := ReadPasswordFromStdin()
		if err != nil {
			return err
		}
		password = pw
	} else if password == "" {
		pw, err := ReadPasswordFromTerminal()
		if err != nil {
			return err
		}
		password = pw
	}
	if password == "" {
		return errors.New("password must not be empty")
	}

	auths := &DockerAuth{
		Auths: make(map[string]map[string]string),
	}

	lp := fmt.Sprintf("%s:%s", login, password)
	lp64 := base64.StdEncoding.EncodeToString([]byte(lp))
	auths.Auths[registry] = map[string]string{"auth": lp64}

	data, err := json.MarshalIndent(auths, "", "  ")
	if err != nil {
		return fmt.Errorf("could not marshal to json: %w", err)
	}

	fmt.Println(string(data))

	return nil
}

func ReadPasswordFromStdin() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	bytePassword, _, err := reader.ReadLine()
	if err != nil {
		return "", fmt.Errorf("could not read password from stding: %w", err)
	}
	password := string(bytePassword)
	password = strings.TrimSpace(password)
	return password, nil
}

func ReadPasswordFromTerminal() (string, error) {
	fmt.Println("Enter password:")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", fmt.Errorf("could not read password from terminal: %w", err)
	}
	password := string(bytePassword)
	password = strings.TrimSpace(password)

	return password, nil
}
