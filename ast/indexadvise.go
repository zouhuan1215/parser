// Copyright 2017 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package ast

import (
	. "github.com/pingcap/parser/format"
)

var (
	_ StmtNode = &IndexAdviseStmt{}
)

// IndexAdviseStmt is a statement to recommend indexes for a given workload
// See ???????
// TODO: input sql file parse, like line-breaker
// rebase just didn't work
type IndexAdviseStmt struct {
	stmtNode

	IsLocal   bool
	Path      string
	MaxTime   uint64
	MaxIdxNum uint64
}

// Restore implements Nodes interface.
func (n *IndexAdviseStmt) Restore(ctx *RestoreCtx) error {
	ctx.WriteKeyWord("Index Advise ")
	if n.IsLocal {
		ctx.WriteKeyWord("LOCAL ")
	}

	ctx.WriteKeyWord("INFILE ")
	ctx.WriteString(n.Path)

	if n.MaxTime != 0 {
		ctx.WriteKeyWord(" MAXTIME ")
		ctx.WritePlainf("%d", n.MaxTime)
	}

	if n.MaxIdxNum != 0 {
		ctx.WriteKeyWord(" MAX RECOMMEND INDEX NUMBER ")
		ctx.WritePlainf("%d", n.MaxIdxNum)
	}

	return nil
}

// Accept implements Node Accept interface.
func (n *IndexAdviseStmt) Accept(v Visitor) (Node, bool) {
	newNode, skipChildren := v.Enter(n)
	if skipChildren {
		return v.Leave(newNode)
	}

	n = newNode.(*IndexAdviseStmt)
	return v.Leave(n)
}
