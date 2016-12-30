// Commands from http://redis.io/commands#list

package Redico

import (
	"strconv"
	"time"

	"github.com/bsm/redeo"
)

// commandsList handles list commands
func commandsList(m *Redico, srv *redeo.Server) {
	srv.HandleFunc("LPOP", m.cmdPop)
	srv.HandleFunc("RPUSH", m.cmdRpush)
}

func (m *Redico) cmdPop(out *redeo.Responder, r *redeo.Request) error {
	if len(r.Args) < 1 {
		setDirty(r.Client())
		return r.WrongNumberOfArgs()
	}
	if !m.handleAuth(r.Client(), out) {
		return nil
	}
	args := r.Args
	timeoutS := args[0]

	timeout, err := strconv.Atoi(timeoutS)
	if err != nil {
		setDirty(r.Client())
		out.WriteErrorString(msgInvalidTimeout)
		return nil
	}
	if timeout < 0 {
		setDirty(r.Client())
		out.WriteErrorString(msgNegTimeout)
		return nil
	}

	blocking(
		m,
		out,
		r,
		time.Duration(timeout) * time.Second,
		func(out *redeo.Responder, ctx *connCtx) bool {
			db := m.db(ctx.selectedDB)
			v, err := db.Pop()
			if err != nil {
				return false
			}
			out.WriteString(v)
			return true
		},
		func(out *redeo.Responder) {
			// timeout
			out.WriteNil()
		},
	)
	return nil
}

func (m *Redico) cmdRpush(out *redeo.Responder, r *redeo.Request) error {
	if len(r.Args) < 1 {
		setDirty(r.Client())
		return r.WrongNumberOfArgs()
	}
	if !m.handleAuth(r.Client(), out) {
		return nil
	}
	args := r.Args[0:]

	return withTx(m, out, r, func(out *redeo.Responder, ctx *connCtx) {
		db := m.db(ctx.selectedDB)

		db.Push(args)
		out.WriteInt(len(r.Args))
	})
}
