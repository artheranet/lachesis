package leveldb

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/artheranet/lachesis/kvdb"
)

type Producer struct {
	datadir         string
	getCacheFdLimit func(string) (int, int)
}

// NewProducer of level db.
func NewProducer(datadir string, getCacheFdLimit func(string) (int, int)) kvdb.IterableDBProducer {
	return &Producer{
		datadir:         datadir,
		getCacheFdLimit: getCacheFdLimit,
	}
}

// Names of existing databases.
func (p *Producer) Names() []string {
	var names []string

	files, err := ioutil.ReadDir(p.datadir)
	if err != nil {
		return []string{}
	}

	for _, f := range files {
		if !f.IsDir() {
			continue
		}
		names = append(names, f.Name())
	}
	return names
}

// OpenDB or create db with name.
func (p *Producer) OpenDB(name string) (kvdb.Store, error) {
	path := p.resolvePath(name)

	err := os.MkdirAll(path, 0700)
	if err != nil {
		return nil, err
	}

	onDrop := func() {
		_ = os.RemoveAll(path)
	}

	cache, fdlimit := p.getCacheFdLimit(name)
	db, err := New(path, cache, fdlimit, nil, onDrop)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (p *Producer) resolvePath(name string) string {
	return filepath.Join(p.datadir, name)
}
