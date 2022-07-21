/*
 * Copyright (c) 2022, AcmeStack
 * All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package stringutils

import (
	"testing"
)

func TestSnake2Camel(t *testing.T) {
	t.Run("compute_node", func(t *testing.T) {
		s := Snake2camel("compute_node")
		if s != "ComputeNode" {
			t.Fatal("expect ComputeNode got: ", s)
		}
	})
	t.Run("COMPUTE_NODE", func(t *testing.T) {
		s := Snake2camel("COMPUTE_NODE")
		if s != "ComputeNode" {
			t.Fatal("expect ComputeNode got: ", s)
		}
	})

	t.Run("COmpute_NODE", func(t *testing.T) {
		s := Snake2camel("COmpute_NODE")
		if s != "ComputeNode" {
			t.Fatal("expect ComputeNode got: ", s)
		}
	})

	t.Run("COmuteNODE", func(t *testing.T) {
		s := Snake2camel("COmuteNODE")
		if s != "COmuteNODE" {
			t.Fatal("expect COmuteNODE got: ", s)
		}
	})

	t.Run("__compute_nOde", func(t *testing.T) {
		s := Snake2camel("__compute_nOde")
		if s != "ComputeNode" {
			t.Fatal("expect ComputeNode got: ", s)
		}
	})

	t.Run("__compute__nOde", func(t *testing.T) {
		s := Snake2camel("__compute__nOde")
		if s != "ComputeNode" {
			t.Fatal("expect ComputeNode got: ", s)
		}
	})

	t.Run("computeNode", func(t *testing.T) {
		s := Snake2camel("computeNode")
		if s != "ComputeNode" {
			t.Fatal("expect ComputeNode got: ", s)
		}
	})

	t.Run("_computeNode_", func(t *testing.T) {
		s := Snake2camel("_computeNode_")
		if s != "ComputeNode" {
			t.Fatal("expect ComputeNode got: ", s)
		}
	})

	t.Run("ComputeNode", func(t *testing.T) {
		s := Snake2camel("ComputeNode")
		if s != "ComputeNode" {
			t.Fatal("expect ComputeNode got: ", s)
		}
	})

	t.Run("ComputeNode_detail", func(t *testing.T) {
		s := Snake2camel("ComputeNode_detail")
		if s != "ComputeNodeDetail" {
			t.Fatal("expect ComputeNode got: ", s)
		}
	})

	t.Run("_computeNode_detail_", func(t *testing.T) {
		s := Snake2camel("_computeNode_detail_")
		if s != "ComputeNodeDetail" {
			t.Fatal("expect ComputeNode got: ", s)
		}
	})

	t.Run("_compute_Node_detail_", func(t *testing.T) {
		s := Snake2camel("_compute_Node_detail_")
		if s != "ComputeNodeDetail" {
			t.Fatal("expect ComputeNode got: ", s)
		}
	})

	t.Run("_compute_node_detail_", func(t *testing.T) {
		s := Snake2camel("_compute_node_detail_")
		if s != "ComputeNodeDetail" {
			t.Fatal("expect ComputeNode got: ", s)
		}
	})
}
