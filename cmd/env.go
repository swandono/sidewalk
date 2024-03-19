/*
Copyright © 2024 GUNAWAN SWANDONO <me@swandono.com>
*/
package cmd

import (
	"os/exec"
	"strings"
)

type env interface {
	check(name string) ([]string, error)
	install(name string) error
	uninstall(name string) error
	update(name string) error
}

type macos struct {
}

func (m *macos) check(name string) ([]string, error) {
	list, err := exec.Command("brew", "list", name).Output()
	if err != nil {
		return nil, err
	}
	return strings.Split(string(list), "\n"), nil
}

func (m *macos) install(name string) error {
	return nil
}

func (m *macos) uninstall(name string) error {
	return nil
}

func (m *macos) update(name string) error {
	return nil
}

type linux struct {
}

func (l *linux) check(name string) ([]string, error) {
	return []string{}, nil
}

func (l *linux) install(name string) error {
	return nil
}

func (l *linux) uninstall(name string) error {
	return nil
}

func (l *linux) update(name string) error {
	return nil
}
