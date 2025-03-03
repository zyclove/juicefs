/*
 * JuiceFS, Copyright (C) 2020 Juicedata, Inc.
 *
 * This program is free software: you can use, redistribute, and/or modify
 * it under the terms of the GNU Affero General Public License, version 3
 * or later ("AGPL"), as published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful, but WITHOUT
 * ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
 * FITNESS FOR A PARTICULAR PURPOSE.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program. If not, see <http://www.gnu.org/licenses/>.
 */

package utils

import "testing"

func TestAlloc(t *testing.T) {
	old := AllocMemory()
	b := Alloc(10)
	if AllocMemory()-old != 16 {
		t.Fatalf("alloc 16 bytes, but got %d", AllocMemory()-old)
	}
	Free(b)
	if AllocMemory()-old != 0 {
		t.Fatalf("free all allocated memory, but got %d", AllocMemory()-old)
	}
}
