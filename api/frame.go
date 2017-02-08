// Copyright (c) 2016 Intel Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

// Version encodes the proxy protocol version.
//
// List of changes:
//
//   • version 2: initial version released with Clear Containers 3.0
//
//                ⚠⚠⚠ backward incompatible with version 1 ⚠⚠⚠
//
//     List of changes:
//
//       • Changed the frame header to include additional fields: version,
//         header length, type and opcode.
//
//   • version 1: initial version released with Clear Containers 2.1
const Version = 2

// FrameType is the type of frame and is part of the frame header.
type FrameType int

const (
	// TypeCommand is a command from a client to the proxy.
	TypeCommand FrameType = iota
	// TypeResponse is a command response back from the proxy to a client.
	TypeResponse
	// TypeStream is a stream of data from a client to the proxy. Streams
	// are to be forwarded onto the VM agent.
	TypeStream
	// TypeNotification is a notification sent by either the proxy or
	// clients. Notifications are one way only and do not prompt a
	// response.
	TypeNotification
	typeMax
)

// Command is the kind of command being sent. In the frame header, Opcode must
// have one of these values when Type is api.TypeCommand.
type Command int

const (
	// CmdRegisterVM registers a new VM/POD.
	CmdRegisterVM Command = iota
	// CmdUnregisterVM unregisters a VM/POD.
	CmdUnregisterVM
	// CmdAttachVM attaches to a registered VM.
	CmdAttachVM
	// CmdHyper sends a hyperstart command through the proxy.
	CmdHyper
	// CmdConnectShim identifies the client as a shim.
	CmdConnectShim
	// CmdDisconnectShim unregisters a shim.
	CmdDisconnectShim
	// CmdSignal sends a signal to the process inside the VM. A client
	// needs to be connected as a shim before it can issue that command.
	CmdSignal
	cmdMax
)

// Stream is the kind of stream being sent. In the frame header, Opcode must
// have one of the these values when Type is api.TypeStream.
type Stream int

const (
	// StreamStdin is a stream conveying stdin data.
	StreamStdin Stream = iota
	// StreamStdout is a stream conveying stdout data.
	StreamStdout
	// StreamStderr is a stream conveying stderr data.
	StreamStderr
	streamMax
)

// FrameHeader is a structure holding a frame header.
type FrameHeader struct {
	Version       uint16
	HeaderLength  uint8
	pad0          uint8
	pad1          uint16
	pad2          uint8
	Type          uint8
	Opcode        uint16
	PayloadLength uint32
}
