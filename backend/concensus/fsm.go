package concensus

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/raft"
	"io"
	"sync"
)

// FSM holds our backend and lock
type FSM struct {
	Mutex sync.RWMutex
	DB    *DB
}

// Apply applies a Raft log entry to the store
func (fsm *FSM) Apply(logEntry *raft.Log) (res interface{}) {
	var e Event
	if err := json.Unmarshal(logEntry.Data, &e); err != nil {
		panic("Failed unmarshaling Raft log entry.")
	}
	fsm.Mutex.Lock()
	defer fsm.Mutex.Unlock()
	conn, err := fsm.DB.RwDB.Conn(context.Background())
	if err != nil {
		return err
	}
	defer func(conn *sql.Conn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)

	switch e.Type {
	case "CreateUser":
		res, err = fsm.DB.CreateUser(conn, e.NewUserRequest)
	case "CreateTweet":
		res, err = fsm.DB.CreateTweet(conn, e.NewTweetRequest)
	case "FollowUser":
		res, err = fsm.DB.FollowUser(conn, e.NewFollowRequest)
	case "UnfollowUser":
		res, err = fsm.DB.UnfollowUser(conn, e.NewFollowRequest)
	default:
		panic(fmt.Sprintf("Unrecognized Event type in Raft log entry: %s.", e.Type))
	}

	return nil
}

func (fsm *FSM) Snapshot() (raft.FSMSnapshot, error) {
	fsm.Mutex.Lock()
	defer fsm.Mutex.Unlock()

	snapfsm := newFSMSnapshot(fsm.DB)

	return snapfsm, nil
}

// Restore stores the store to a previous state.
func (fsm *FSM) Restore(rc io.ReadCloser) error {

	b, err := dbBytesFromSnapshot(rc)
	if err != nil {
		return fmt.Errorf("restore failed: %s", err.Error())
	}
	if b == nil {
		fmt.Errorf("no backend data present in restored snapshot")
	}

	if err := fsm.DB.RwDB.Close(); err != nil {
		fmt.Errorf("failed to close pre-restore backend: %s")
	}

	var db *DB
	db, err = CreateOnDisk(fsm.DB.path, b)
	if err != nil {
		return fmt.Errorf("open on-disk file during restore: %s", err)
	}
	fmt.Printf("successfully switched to on-disk backend due to restore")

	fsm.DB = db

	fmt.Printf("Node restored")
	return nil
}
