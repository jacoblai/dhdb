// Commands from http://redis.io/commands#list

package Redico

import (
	"github.com/bsm/redeo"
)

// commandsList handles list commands
func commandsList(m *Redico, srv *redeo.Server) {
	srv.HandleFunc("LPOP", m.cmdPop)
	srv.HandleFunc("RPUSH", m.cmdRpush)
}

func (m *Redico) cmdPop(out *redeo.Responder, r *redeo.Request) error {
	if len(r.Args) != 0 {
		setDirty(r.Client())
		return r.WrongNumberOfArgs()
	}
	if !m.handleAuth(r.Client(), out) {
		return nil
	}

	return withTx(m, out, r, func(out *redeo.Responder, ctx *connCtx) {
		db := m.db(ctx.selectedDB)

		elem, err := db.Pop()
		if err != nil {
			out.WriteError(err)
		}
		out.WriteString(elem)
	})
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
