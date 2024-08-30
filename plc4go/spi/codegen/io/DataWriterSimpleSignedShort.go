/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package io

import (
	"context"

	"github.com/apache/plc4x/plc4go/spi/utils"
)

type DataWriterSimpleSignedShort struct {
	*DataWriterSimpleBase[int16]
}

var _ DataWriter[int16] = (*DataWriterSimpleSignedShort)(nil)

func NewDataWriterSimpleSignedShort(WritBuffer utils.WriteBuffer, bitLength uint8) *DataWriterSimpleSignedShort {
	return &DataWriterSimpleSignedShort{
		DataWriterSimpleBase: NewDataWriterSimpleBase[int16](WritBuffer, uint(bitLength)),
	}
}

func (d *DataWriterSimpleSignedShort) Write(ctx context.Context, logicalName string, value int16, args ...utils.WithWriterArgs) error {
	return d.WriteBuffer.WriteInt16(logicalName, uint8(d.bitLength), value, args...)
}
