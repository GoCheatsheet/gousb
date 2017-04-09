// Copyright 2013 Google Inc.  All rights reserved.
// Copyright 2016 the gousb Authors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package usb

import (
	"fmt"
)

// #include <libusb.h>
import "C"

// USBError is an error code returned by libusb.
type USBError C.int

// Error implements the error interface.
func (e USBError) Error() string {
	return fmt.Sprintf("libusb: %s [code %d]", USBErrorString[e], e)
}

func fromUSBError(errno C.int) error {
	err := USBError(errno)
	if err == Success {
		return nil
	}
	return err
}

// Error codes defined by libusb.
const (
	Success           USBError = C.LIBUSB_SUCCESS
	ErrorIO           USBError = C.LIBUSB_ERROR_IO
	ErrorInvalidParam USBError = C.LIBUSB_ERROR_INVALID_PARAM
	ErrorAccess       USBError = C.LIBUSB_ERROR_ACCESS
	ErrorNoDevice     USBError = C.LIBUSB_ERROR_NO_DEVICE
	ErrorNotFound     USBError = C.LIBUSB_ERROR_NOT_FOUND
	ErrorBusy         USBError = C.LIBUSB_ERROR_BUSY
	ErrorTimeout      USBError = C.LIBUSB_ERROR_TIMEOUT
	ErrorOverflow     USBError = C.LIBUSB_ERROR_OVERFLOW
	ErrorPipe         USBError = C.LIBUSB_ERROR_PIPE
	ErrorInterrupted  USBError = C.LIBUSB_ERROR_INTERRUPTED
	ErrorNoMem        USBError = C.LIBUSB_ERROR_NO_MEM
	ErrorNotSupported USBError = C.LIBUSB_ERROR_NOT_SUPPORTED
	ErrorOther        USBError = C.LIBUSB_ERROR_OTHER
)

var USBErrorString = map[USBError]string{
	Success:           "success",
	ErrorIO:           "i/o error",
	ErrorInvalidParam: "invalid param",
	ErrorAccess:       "bad access",
	ErrorNoDevice:     "no device",
	ErrorNotFound:     "not found",
	ErrorBusy:         "device or resource busy",
	ErrorTimeout:      "timeout",
	ErrorOverflow:     "overflow",
	ErrorPipe:         "pipe error",
	ErrorInterrupted:  "interrupted",
	ErrorNoMem:        "out of memory",
	ErrorNotSupported: "not supported",
	ErrorOther:        "unknown error",
}

// TransferStatus contains information about the result of a transfer.
type TransferStatus uint8

const (
	TransferCompleted TransferStatus = C.LIBUSB_TRANSFER_COMPLETED
	TransferError     TransferStatus = C.LIBUSB_TRANSFER_ERROR
	TransferTimedOut  TransferStatus = C.LIBUSB_TRANSFER_TIMED_OUT
	TransferCancelled TransferStatus = C.LIBUSB_TRANSFER_CANCELLED
	TransferStall     TransferStatus = C.LIBUSB_TRANSFER_STALL
	TransferNoDevice  TransferStatus = C.LIBUSB_TRANSFER_NO_DEVICE
	TransferOverflow  TransferStatus = C.LIBUSB_TRANSFER_OVERFLOW
)

var transferStatusDescription = map[TransferStatus]string{
	TransferCompleted: "transfer completed without error",
	TransferError:     "transfer failed",
	TransferTimedOut:  "transfer timed out",
	TransferCancelled: "transfer was cancelled",
	TransferStall:     "halt condition detected (endpoint stalled) or control request not supported",
	TransferNoDevice:  "device was disconnected",
	TransferOverflow:  "device sent more data than requested",
}

// String returns a human-readable transfer status.
func (ts TransferStatus) String() string {
	return transferStatusDescription[ts]
}

// Error implements the error interface.
func (ts TransferStatus) Error() string {
	return ts.String()
}
