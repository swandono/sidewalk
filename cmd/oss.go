/*
Copyright © 2024 GUNAWAN SWANDONO <me@swandono.com>
*/
package cmd

import (
	"os/exec"
	"strings"
)

type oss interface {
	check(name string) ([]string, error)
	install(name string) error
	uninstall(name string) error
	update(name string) error
	init(target string, source string) error
}

type macos struct {
	name string
}

func (m *macos) check(name string) ([]string, error) {
	list, err := exec.Command("brew", "list", name).Output()
	if err != nil {
		return nil, err
	}
	return strings.Split(string(list), "\n"), nil
}

func (m *macos) install(name string) error {
	_, err := exec.Command("brew", "install", name).Output()
	if err != nil {
		return err
	}
	return nil
}

func (m *macos) uninstall(name string) error {
	_, err := exec.Command("brew", "uninstall", name).Output()
	if err != nil {
		return err
	}
	return nil
}

func (m *macos) update(name string) error {
	_, err := exec.Command("brew", "upgrade", name).Output()
	if err != nil {
		return err
	}
	return nil
}

func (m *macos) init(target string, source string) error {
	_, err := exec.Command("cp", "-r", source, target).Output()
	if err != nil {
		return err
	}
	return nil
}

type linux struct {
	name string
}

// Ubuntu
func (l *linux) check(name string) ([]string, error) {
	list, err := exec.Command("dpkg", "-l", name).Output()
	if err != nil {
		return nil, err
	}
	return strings.Split(string(list), "\n"), nil
}

func (l *linux) install(name string) error {
	_, err := exec.Command("apt-get", "install", name).Output()
	if err != nil {
		return err
	}
	return nil
}

func (l *linux) uninstall(name string) error {
	_, err := exec.Command("apt-get", "remove", name).Output()
	if err != nil {
		return err
	}
	return nil
}

func (l *linux) update(name string) error {
	_, err := exec.Command("apt-get", "update", name).Output()
	if err != nil {
		return err
	}
	return nil
}

func (l *linux) init(target string, source string) error {
	_, err := exec.Command("cp", "-r", source, target).Output()
	if err != nil {
		return err
	}
	return nil
}
