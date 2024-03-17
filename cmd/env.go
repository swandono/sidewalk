/*
Copyright © 2024 GUNAWAN SWANDONO <me@swandono.com>
*/
package cmd

type env interface {
    check() ([]string, error)
    install() error
    uninstall() error
    update() error
}

type macos struct {
}

func (m *macos) check() ([]string, error) {
    return []string{}, nil
}

func (m *macos) install() error {
    return nil
}

func (m *macos) uninstall() error {
    return nil
}

func (m *macos) update() error {
    return nil
}

type linux struct {
}

func (l *linux) check() ([]string, error) {
    return []string{}, nil
}

func (l *linux) install() error {
    return nil
}

func (l *linux) uninstall() error {
    return nil
}

func (l *linux) update() error {
    return nil
}
