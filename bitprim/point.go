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
// Point Golang idiomatic Interface
// --------------------------------

package bitprim

import (
	"unsafe"
)

type Point struct {
	ptr unsafe.Pointer
}

func NewPoint(ptr unsafe.Pointer) *Point {
	x := new(Point)
	x.ptr = ptr
	return x
}

func (x Point) Hash() HashT {
	return pointHash(x.ptr)
}

func (x *Point) IsValid() bool {
	return pointIsValid(x.ptr)
}

func (x *Point) Index() uint32 {
	return pointGetIndex(x.ptr)
}

func (x *Point) Checksum() uint64 {
	return pointGetChecksum(x.ptr)
}
