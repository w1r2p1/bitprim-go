/**
 * Copyright (c) 2017 Bitprim developers (see AUTHORS)
 *
 * This file is part of Bitprim.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

// --------------------------------
// Executor Golang idiomatic Interface
// --------------------------------

package bitprim

// --------------------------------------------------------------------------------

import (
	"syscall" // or "runtime"
	"unsafe"
)

//Executor structure
type Executor struct {
	ptr unsafe.Pointer
}

func NewExecutor(path string) *Executor {
	return NewExecutorWithStd(path, syscall.Stdin, syscall.Stdout, syscall.Stderr)
}

func (x *Executor) Close() {
	ExecutorDestruct(x.ptr)
	x.ptr = nil
}

func (x Executor) Run() int {
	return ExecutorRun(x.ptr)
}

func (x Executor) RunAndWait() int {
	return ExecutorRunAndWait(x.ptr)
}

func (x Executor) Initchain() int {
	return ExecutorInitchain(x.ptr)
}

func (x Executor) FetchLastHeight() int {
	return FetchLastHeight(x.ptr)
}

func (x Executor) FetchBlockHeight(hash HashT) int {
	return FetchBlockHeight(x.ptr, hash)
}

func (x Executor) FetchBlockHeader(height int) *Header {
	res := NewHeader(FetchBlockHeader(x.ptr, height))
	return res
}

func (x Executor) FetchBlockHeaderByHash(hash HashT) *Header {
	res := NewHeader(FetchBlockHeaderByHash(x.ptr, hash))
	return res
}

func (x Executor) FetchBlock(height int) *Block {
	res := NewBlock(FetchBlock(x.ptr, height))
	return res
}

func (x Executor) FetchBlockByHash(hash HashT) *Block {
	res := NewBlock(FetchBlockByHash(x.ptr, hash))
	return res
}

func (x Executor) FetchTransaction(hash HashT, requiredConfirmed bool) *Transaction {
	res := NewTransaction(FetchTransaction(x.ptr, hash, requiredConfirmed))
	return res
}

func (x Executor) FetchOutput(hash HashT, index int, requiredConfirmed bool) *Output {
	res := NewOutput(FetchOutput(x.ptr, hash, index, requiredConfirmed))
	return res
}