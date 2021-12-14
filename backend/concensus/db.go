package concensus

import (
	"bytes"
	"compress/gzip"
	"context"
	"database/sql"
	"encoding/binary"
	"fmt"
	"github.com/rqlite/go-sqlite3"
	"io"
	"io/ioutil"
	"math"
	"math/rand"
	"os"
	pb "simple-twitter.com/backend/rpc/proto"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

// DB is the SQL backend.
type DB struct {
	path   string // Path to backend file, if running on-disk.
	memory bool   // In-memory only.

	RwDB *sql.DB // Database connection for backend reads and writes.
	RoDB *sql.DB // Database connection backend reads.

	RwDSN string // DSN used for read-write connection
	RoDSN string // DSN used for read-only connections
}

func CreateTables(db *sql.DB) error {
	conn, err := db.Conn(context.Background())
	defer func(conn *sql.Conn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)

	createUserTableSQL := `CREATE TABLE IF NOT EXISTS user (
		id integer NOT NULL PRIMARY KEY,
		name TEXT,
		password TEXT,
		bio TEXT,
		handle TEXT UNIQUE,
		created_at DATETIME
	);`

	statement, err := conn.PrepareContext(context.Background(), createUserTableSQL) // Prepare SQL Statement
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create table...")
		return err
	}
	defer statement.Close()
	_, err = statement.Exec()
	if err != nil {
		return err
	}

	createFollowersTableSQL := `CREATE TABLE IF NOT EXISTS follow (
		follower_id INTEGER NOT NULL,
		followed_id INTEGER NOT NULL,
		FOREIGN KEY (follower_id) 
		REFERENCES user (id)
			ON UPDATE CASCADE
			ON DELETE CASCADE
		FOREIGN KEY (followed_id) 
		REFERENCES user (id)
			ON UPDATE CASCADE
			ON DELETE CASCADE
	 );`

	statement, err = conn.PrepareContext(context.Background(), createFollowersTableSQL) // Prepare SQL Statement
	if err != nil {
		return err
	}
	defer statement.Close()
	statement.Exec()

	createTweetsTableSQL := `CREATE TABLE IF NOT EXISTS tweet (
		id integer NOT NULL PRIMARY KEY,
		content TEXT,
		author_id TEXT,
		author_name TEXT,
		author_handle TEXT,
		created_at DATETIME,
		FOREIGN KEY (author_iD) 
		REFERENCES user (id)
			ON UPDATE CASCADE
			ON DELETE CASCADE
	 );`

	statement, err = conn.PrepareContext(context.Background(), createTweetsTableSQL) // Prepare SQL Statement
	if err != nil {
		return err
	}
	defer statement.Close()
	statement.Exec()

	return nil
}

// CreateUser creates a new user profile
func (db *DB) CreateUser(conn *sql.Conn, newUser *pb.User) (res interface{}, err error) {
	sqlStatement := `
		INSERT INTO user (
			name,
			password,
			bio,
			handle,
			created_at
		) VALUES (?, ?, ?, ?, CURRENT_TIMESTAMP);`

	statement, err := conn.PrepareContext(context.Background(), sqlStatement) // Prepare SQL Statement
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	res, err = statement.Exec(newUser.Fullname, newUser.Password, newUser.Bio, newUser.Handle)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (db *DB) LoginUser(conn *sql.Conn, user *pb.User) (res []*pb.User, err error) {
	sqlStatement := `
		SELECT * FROM user
		WHERE name = ? AND password = ?
		;`

	statement, err := conn.PrepareContext(context.Background(), sqlStatement) // Prepare SQL Statement
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	rows, err := statement.Query(user.Fullname, user.Password)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user pb.User
		err = rows.Scan(&user.Id, &user.Fullname, &user.Password, &user.Bio, &user.Handle, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		res = append(res, &user)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return res, nil
}

// CreateTweet creates a new tweet for the user
func (db *DB) CreateTweet(conn *sql.Conn, tweet *pb.Tweet) (res interface{}, err error) {
	sqlStatement := `
		INSERT OR REPLACE INTO tweet(
			content,
			author_id,
			author_name,
			author_handle,
			created_at
		) VALUES (?, ?, ?, ?, CURRENT_TIMESTAMP);`

	statement, err := conn.PrepareContext(context.Background(), sqlStatement) // Prepare SQL Statement
	if err != nil {
		return nil, err
	}
	defer statement.Close()
	res, err = statement.Exec(tweet.Content, tweet.AuthorId, tweet.AuthorName, tweet.AuthorHandle)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// FollowUser creates a follower-followed relationship
func (db *DB) FollowUser(conn *sql.Conn, newFollow *pb.Follow) (res interface{}, err error) {

	sqlStatement := `
		INSERT OR REPLACE INTO follow(
			follower_id,
			followed_id
		) VALUES (?, ?);`

	statement, err := conn.PrepareContext(context.Background(), sqlStatement) // Prepare SQL Statement
	if err != nil {
		return nil, err
	}
	defer statement.Close()
	res, err = statement.Exec(newFollow.FollowerId, newFollow.FollowedId)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// UnfollowUser removes a follower-followed relationship
func (db *DB) UnfollowUser(conn *sql.Conn, newUnfollow *pb.Follow) (res interface{}, err error) {

	sqlStatement := `
		DELETE FROM follow
      	WHERE follower_id=? AND followed_id=?
	;`

	statement, err := conn.PrepareContext(context.Background(), sqlStatement) // Prepare SQL Statement
	if err != nil {
		return nil, err
	}
	defer statement.Close()
	res, err = statement.Exec(newUnfollow.FollowerId, newUnfollow.FollowedId)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (db *DB) GetUser(conn *sql.Conn, user *pb.User) (res []*pb.User, err error) {
	sqlStatement := `
		SELECT * FROM user
		WHERE id = ?
		;`

	statement, err := conn.PrepareContext(context.Background(), sqlStatement) // Prepare SQL Statement
	if err != nil {
		return nil, err
	}
	defer statement.Close()
	rows, err := statement.Query(user.Id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var user pb.User
		err = rows.Scan(&user.Id, &user.Fullname, &user.Password, &user.Bio, &user.Handle, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		res = append(res, &user)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (db *DB) GetUsers(conn *sql.Conn) (res []*pb.User, err error) {
	sqlStatement := `
		SELECT * FROM user
		;`

	statement, err := conn.PrepareContext(context.Background(), sqlStatement) // Prepare SQL Statement
	if err != nil {
		return nil, err
	}
	defer statement.Close()
	rows, err := statement.Query()
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var user pb.User
		err = rows.Scan(&user.Id, &user.Fullname, &user.Password, &user.Bio, &user.Handle, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		res = append(res, &user)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (db *DB) GetTweetsByUser(conn *sql.Conn, user *pb.User) (res []*pb.Tweet, err error) {

	sqlStatement := `
		SELECT * FROM tweet
		WHERE author_id = ?;`

	statement, err := conn.PrepareContext(context.Background(), sqlStatement) // Prepare SQL Statement
	if err != nil {
		return nil, err
	}
	defer func(statement *sql.Stmt) {
		err := statement.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed closing statement: %v \n", err)
		}
	}(statement)
	rows, err := statement.Query(user.Id)
	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)
	res = []*pb.Tweet{}
	for rows.Next() {
		var tweet pb.Tweet
		err = rows.Scan(&tweet.Id, &tweet.Content, &tweet.AuthorId, &tweet.AuthorName, &tweet.AuthorHandle, &tweet.CreatedAt)
		if err != nil {
			return nil, err
		}
		res = append(res, &tweet)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (db *DB) GetFeedByUser(conn *sql.Conn, user *pb.User) (res []*pb.Tweet, err error) {
	sqlStatement := `
		SELECT * FROM tweet
		WHERE author_id IN (SELECT followed_id FROM follow WHERE follower_id = ?)
		ORDER BY datetime(created_at) DESC
		;`

	statement, err := conn.PrepareContext(context.Background(), sqlStatement) // Prepare SQL Statement
	if err != nil {
		return nil, err
	}
	defer statement.Close()
	rows, err := statement.Query(user.Id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var tweet pb.Tweet
		err = rows.Scan(&tweet.Id, &tweet.Content, &tweet.AuthorId, &tweet.AuthorName, &tweet.AuthorHandle, &tweet.CreatedAt)
		if err != nil {
			return nil, err
		}
		res = append(res, &tweet)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (db *DB) GetFollowedByUser(conn *sql.Conn, user *pb.User) (res []*pb.User, err error) {

	sqlStatement := `
		SELECT * FROM user
		WHERE id IN (SELECT follower_id FROM follow WHERE followed_id = ?)
		;`

	statement, err := conn.PrepareContext(context.Background(), sqlStatement) // Prepare SQL Statement
	if err != nil {
		return nil, err
	}
	defer statement.Close()
	rows, err := statement.Query(user.Id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var user pb.User
		err = rows.Scan(&user.Id, &user.Password, &user.Fullname, &user.Bio, &user.Handle, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		res = append(res, &user)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (db *DB) GetFollowingByUser(conn *sql.Conn, user *pb.User) (res []*pb.User, err error) {
	sqlStatement := `
		SELECT * FROM user
		WHERE id IN (SELECT followed_id FROM follow WHERE follower_id = ?)
		;`

	statement, err := conn.PrepareContext(context.Background(), sqlStatement) // Prepare SQL Statement
	if err != nil {
		return nil, err
	}
	defer statement.Close()
	rows, err := statement.Query(user.Id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var user pb.User
		err = rows.Scan(&user.Id, &user.Password, &user.Fullname, &user.Bio, &user.Handle, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		res = append(res, &user)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (db *DB) GetTweet(conn *sql.Conn, tweet *pb.Tweet) (res *pb.Tweet, err error) {
	sqlStatement := `
		SELECT * FROM tweet
		WHERE id = ?
		;`

	statement, err := conn.PrepareContext(context.Background(), sqlStatement) // Prepare SQL Statement
	if err != nil {
		return nil, err
	}
	defer statement.Close()
	err = statement.QueryRow(tweet.Id).Scan(&res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// createOnDisk opens an on-disk backend file at the fsm's configured path. If
// b is non-nil, any preexisting file will first be overwritten with those contents.
// Otherwise any pre-existing file will be removed before the backend is opened.
func CreateOnDisk(dbPath string, b []byte) (*DB, error) {
	if err := os.Remove(dbPath); err != nil && !os.IsNotExist(err) {
		return nil, err
	}
	if b != nil {
		if err := ioutil.WriteFile(dbPath, b, 0660); err != nil {
			return nil, err
		}
	}
	return open(dbPath, true)
}

// Open opens a file-based backend, creating it if it does not exist. After this
// function returns, an actual SQLite file will always exist.
func open(dbPath string, fkEnabled bool) (*DB, error) {
	rwDSN := fmt.Sprintf("file:%s?_fk=%s", dbPath, strconv.FormatBool(fkEnabled))
	rwDB, err := sql.Open("sqlite3", rwDSN)
	if err != nil {
		return nil, err
	}

	roOpts := []string{
		"mode=ro",
		fmt.Sprintf("_fk=%s", strconv.FormatBool(fkEnabled)),
	}

	roDSN := fmt.Sprintf("file:%s?%s", dbPath, strings.Join(roOpts, "&"))
	roDB, err := sql.Open("sqlite3", roDSN)
	if err != nil {
		return nil, err
	}

	// Force creation of on-disk backend file.
	if err := rwDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping on-disk backend: %s", err.Error())
	}

	// Set some reasonable connection pool behaviour.
	rwDB.SetConnMaxIdleTime(30 * time.Second)
	rwDB.SetConnMaxLifetime(0)
	roDB.SetConnMaxIdleTime(30 * time.Second)
	roDB.SetConnMaxLifetime(0)

	return &DB{
		path:   dbPath,
		memory: true,
		RwDB:   rwDB,
		RoDB:   roDB,
		RwDSN:  rwDSN,
		RoDSN:  roDSN,
	}, nil
}

// OpenInMemory returns a new in-memory backend.
func OpenInMemory(fkEnabled bool, serverID string) (*DB, error) {

	inMemPath := fmt.Sprintf("file:/%s/%s", serverID, randomString(len(serverID)))

	rwOpts := []string{
		"mode=rw",
		"vfs=memdb",
		"_txlock=immediate",
		fmt.Sprintf("_fk=%s", strconv.FormatBool(fkEnabled)),
	}

	rwDSN := fmt.Sprintf("%s?%s", inMemPath, strings.Join(rwOpts, "&"))
	rwDB, err := sql.Open("sqlite3", rwDSN)
	if err != nil {
		return nil, err
	}

	// Ensure there is only one connection and it never closes.
	// If it closed, in-memory backend could be lost.
	rwDB.SetConnMaxIdleTime(0)
	rwDB.SetConnMaxLifetime(0)
	rwDB.SetMaxIdleConns(1)
	rwDB.SetMaxOpenConns(1)

	roOpts := []string{
		"mode=ro",
		"vfs=memdb",
		"_txlock=deferred",
		fmt.Sprintf("_fk=%s", strconv.FormatBool(fkEnabled)),
	}

	roDSN := fmt.Sprintf("%s?%s", inMemPath, strings.Join(roOpts, "&"))
	roDB, err := sql.Open("sqlite3", roDSN)
	if err != nil {
		return nil, err
	}

	// Ensure backend is basically healthy.
	if err := rwDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping in-memory backend: %s", err.Error())
	}

	return &DB{
		memory: true,
		RwDB:   rwDB,
		RoDB:   roDB,
		RwDSN:  rwDSN,
		RoDSN:  roDSN,
	}, nil
}

func (db *DB) Serialize() ([]byte, error) {
	if !db.memory {
		// Simply read and return the SQLite file.
		return os.ReadFile(db.path)
	}
	conn, err := db.RoDB.Conn(context.Background())
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var b []byte
	f := func(driverConn interface{}) error {
		c := driverConn.(*sqlite3.SQLiteConn)
		b = c.Serialize("")
		if b == nil {
			return fmt.Errorf("failed to serialize backend")
		}
		return nil
	}

	if err := conn.Raw(f); err != nil {
		return nil, err
	}
	return b, nil
}

func dbBytesFromSnapshot(rc io.ReadCloser) ([]byte, error) {
	var uint64Size uint64
	inc := int64(unsafe.Sizeof(uint64Size))

	// Read all the data into RAM, since we have to decode known-length
	// chunks of various forms.
	var offset int64
	b, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, fmt.Errorf("readall: %s", err)
	}

	// Get size of backend, checking for compression.
	compressed := false
	sz, err := readUint64(b[offset : offset+inc])
	if err != nil {
		return nil, fmt.Errorf("read compression check: %s", err)
	}
	offset = offset + inc

	if sz == math.MaxUint64 {
		compressed = true
		// Database is actually compressed, read actual size next.
		sz, err = readUint64(b[offset : offset+inc])
		if err != nil {
			return nil, fmt.Errorf("read compressed size: %s", err)
		}
		offset = offset + inc
	}

	// Now read in the backend file data, decompress if necessary, and restore.
	var database []byte
	if sz > 0 {
		if compressed {
			buf := new(bytes.Buffer)
			gz, err := gzip.NewReader(bytes.NewReader(b[offset : offset+int64(sz)]))
			if err != nil {
				return nil, err
			}

			if _, err := io.Copy(buf, gz); err != nil {
				return nil, fmt.Errorf("SQLite backend decompress: %s", err)
			}

			if err := gz.Close(); err != nil {
				return nil, err
			}
			database = buf.Bytes()
		} else {
			database = b[offset : offset+int64(sz)]
		}
	} else {
		database = nil
	}
	return database, nil
}

func readUint64(b []byte) (uint64, error) {
	var sz uint64
	if err := binary.Read(bytes.NewReader(b), binary.LittleEndian, &sz); err != nil {
		return 0, err
	}
	return sz, nil
}
func randomString(seed int) string {
	var output strings.Builder
	chars := "abcdedfghijklmnopqrstABCDEFGHIJKLMNOP"
	for i := 0; i < 20; i++ {
		random := rand.Intn(seed)
		randomChar := chars[random]
		output.WriteString(string(randomChar))
	}
	return output.String()
}
