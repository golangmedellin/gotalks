// +build ignore

package main

// START TYPES OMIT
type signal struct{}

// Based on from github.com/astaxie/beego/session
type SessionStore interface {
	Set(key string, value interface{}) //set session value
	Get(key string) interface{}        //get session value
	Delete(key string)                 //delete session value
	SessionID() string                 //back current sessionID
	Flush()                            //delete all data
}

// memory session store.
// it saved sessions in a map in memory.
// based on from github.com/astaxie/beego/session/sess_mem.go
type MemorySessionStore struct {
	SessionStore                        //implements SessionStore interface
	sessionId    string                 //session id
	session      map[string]interface{} //session store
	semaphore    chan signal            //semaphore
}

// END TYPES OMIT

// START SEM OMIT
func NewMemorySessionStore(sessionId string) *MemorySessionStore {
	mStore := &MemorySessionStore{
		sessionId: sessionId,
		session:   make(map[string]interface{}),
		semaphore: make(chan signal, 1),
	}
	return mStore
}

// set user to memory session
func (st *MemorySessionStore) Set(key string, value interface{}) {
	st.semaphore <- signal{}
	st.session[key] = value
	<-st.semaphore
}

// END SEM OMIT

// START SEM2 OMIT

// get user from memory session by key
func (st *MemorySessionStore) Get(key string) interface{} {
	st.semaphore <- signal{}
	user, ok := st.session[key]
	<-st.semaphore
	if ok {
		return user
	} else {
		return nil
	}
}

// delete in memory session by key
func (st *MemorySessionStore) Delete(key string) {
	st.semaphore <- signal{}
	delete(st.session, key)
	<-st.semaphore
}

// END SEM2 OMIT
// clear all users in memory session
func (st *MemorySessionStore) Flush() {
	st.semaphore <- signal{}
	st.session = make(map[string]interface{})
	<-st.semaphore
}
