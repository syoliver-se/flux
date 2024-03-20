package table

import "github.com/syoliver-se/flux"

type Iterator []flux.Table

func (t Iterator) Do(f func(flux.Table) error) error {
	for _, tbl := range t {
		if err := f(tbl); err != nil {
			return err
		}
	}
	return nil
}
