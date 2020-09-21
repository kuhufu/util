package util

import "github.com/go-xorm/xorm"

type XormEngineWrapper struct {
	*xorm.Engine
}

func (x XormEngineWrapper) WithSession(f func(session *xorm.Session) (err error)) error {
	session := x.Engine.NewSession()
	defer session.Close()
	return f(session)
}

func (x XormEngineWrapper) WithTx(f func(session *xorm.Session) (err error)) error {
	session := x.Engine.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		return err
	}

	err := f(session)
	if err != nil {
		session.Rollback()
		return err
	}

	return session.Commit()
}
