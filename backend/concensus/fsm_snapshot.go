package concensus

import (
	"bytes"
	"compress/gzip"
	"encoding/binary"
	"github.com/hashicorp/raft"
	"io"
	"math"
)

type fsmSnapshot struct {
	database []byte
}

func newFSMSnapshot(db *DB) *fsmSnapshot {
	fsm := &fsmSnapshot{}

	fsm.database, _ = db.Serialize()
	return fsm
}

// Persist writes the snapshot to the given sink.
func (f *fsmSnapshot) Persist(sink raft.SnapshotSink) error {

	err := func() error {
		b := new(bytes.Buffer)

		// Flag compressed backend by writing max uint64 value first.
		// No SQLite backend written by earlier versions will have this
		// as a size. *Surely*.
		err := writeUint64(b, math.MaxUint64)
		if err != nil {
			return err
		}
		if _, err := sink.Write(b.Bytes()); err != nil {
			return err
		}
		b.Reset() // Clear state of buffer for future use.

		// Get compressed copy of backend.
		cdb, err := f.compressedDatabase()
		if err != nil {
			return err
		}

		if cdb != nil {
			// Write size of compressed backend.
			err = writeUint64(b, uint64(len(cdb)))
			if err != nil {
				return err
			}
			if _, err := sink.Write(b.Bytes()); err != nil {
				return err
			}

			// Write compressed backend to sink.
			if _, err := sink.Write(cdb); err != nil {
				return err
			}
		} else {
			err = writeUint64(b, uint64(0))
			if err != nil {
				return err
			}
			if _, err := sink.Write(b.Bytes()); err != nil {
				return err
			}
		}

		// Close the sink.
		return sink.Close()
	}()

	if err != nil {
		sink.Cancel()
		return err
	}

	return nil
}

func (f *fsmSnapshot) Release() {}

func writeUint64(w io.Writer, v uint64) error {
	return binary.Write(w, binary.LittleEndian, v)
}

func (f *fsmSnapshot) compressedDatabase() ([]byte, error) {
	if f.database == nil {
		return nil, nil
	}

	var buf bytes.Buffer
	gz, err := gzip.NewWriterLevel(&buf, gzip.BestCompression)
	if err != nil {
		return nil, err
	}

	if _, err := gz.Write(f.database); err != nil {
		return nil, err
	}
	if err := gz.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
