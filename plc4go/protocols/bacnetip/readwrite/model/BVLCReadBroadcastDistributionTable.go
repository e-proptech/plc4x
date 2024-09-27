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

package model

import (
	"context"
	"encoding/binary"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BVLCReadBroadcastDistributionTable is the corresponding interface of BVLCReadBroadcastDistributionTable
type BVLCReadBroadcastDistributionTable interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BVLC
	// IsBVLCReadBroadcastDistributionTable is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBVLCReadBroadcastDistributionTable()
	// CreateBuilder creates a BVLCReadBroadcastDistributionTableBuilder
	CreateBVLCReadBroadcastDistributionTableBuilder() BVLCReadBroadcastDistributionTableBuilder
}

// _BVLCReadBroadcastDistributionTable is the data-structure of this message
type _BVLCReadBroadcastDistributionTable struct {
	BVLCContract
}

var _ BVLCReadBroadcastDistributionTable = (*_BVLCReadBroadcastDistributionTable)(nil)
var _ BVLCRequirements = (*_BVLCReadBroadcastDistributionTable)(nil)

// NewBVLCReadBroadcastDistributionTable factory function for _BVLCReadBroadcastDistributionTable
func NewBVLCReadBroadcastDistributionTable() *_BVLCReadBroadcastDistributionTable {
	_result := &_BVLCReadBroadcastDistributionTable{
		BVLCContract: NewBVLC(),
	}
	_result.BVLCContract.(*_BVLC)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BVLCReadBroadcastDistributionTableBuilder is a builder for BVLCReadBroadcastDistributionTable
type BVLCReadBroadcastDistributionTableBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() BVLCReadBroadcastDistributionTableBuilder
	// Build builds the BVLCReadBroadcastDistributionTable or returns an error if something is wrong
	Build() (BVLCReadBroadcastDistributionTable, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BVLCReadBroadcastDistributionTable
}

// NewBVLCReadBroadcastDistributionTableBuilder() creates a BVLCReadBroadcastDistributionTableBuilder
func NewBVLCReadBroadcastDistributionTableBuilder() BVLCReadBroadcastDistributionTableBuilder {
	return &_BVLCReadBroadcastDistributionTableBuilder{_BVLCReadBroadcastDistributionTable: new(_BVLCReadBroadcastDistributionTable)}
}

type _BVLCReadBroadcastDistributionTableBuilder struct {
	*_BVLCReadBroadcastDistributionTable

	parentBuilder *_BVLCBuilder

	err *utils.MultiError
}

var _ (BVLCReadBroadcastDistributionTableBuilder) = (*_BVLCReadBroadcastDistributionTableBuilder)(nil)

func (b *_BVLCReadBroadcastDistributionTableBuilder) setParent(contract BVLCContract) {
	b.BVLCContract = contract
}

func (b *_BVLCReadBroadcastDistributionTableBuilder) WithMandatoryFields() BVLCReadBroadcastDistributionTableBuilder {
	return b
}

func (b *_BVLCReadBroadcastDistributionTableBuilder) Build() (BVLCReadBroadcastDistributionTable, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BVLCReadBroadcastDistributionTable.deepCopy(), nil
}

func (b *_BVLCReadBroadcastDistributionTableBuilder) MustBuild() BVLCReadBroadcastDistributionTable {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BVLCReadBroadcastDistributionTableBuilder) Done() BVLCBuilder {
	return b.parentBuilder
}

func (b *_BVLCReadBroadcastDistributionTableBuilder) buildForBVLC() (BVLC, error) {
	return b.Build()
}

func (b *_BVLCReadBroadcastDistributionTableBuilder) DeepCopy() any {
	_copy := b.CreateBVLCReadBroadcastDistributionTableBuilder().(*_BVLCReadBroadcastDistributionTableBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBVLCReadBroadcastDistributionTableBuilder creates a BVLCReadBroadcastDistributionTableBuilder
func (b *_BVLCReadBroadcastDistributionTable) CreateBVLCReadBroadcastDistributionTableBuilder() BVLCReadBroadcastDistributionTableBuilder {
	if b == nil {
		return NewBVLCReadBroadcastDistributionTableBuilder()
	}
	return &_BVLCReadBroadcastDistributionTableBuilder{_BVLCReadBroadcastDistributionTable: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BVLCReadBroadcastDistributionTable) GetBvlcFunction() uint8 {
	return 0x02
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BVLCReadBroadcastDistributionTable) GetParent() BVLCContract {
	return m.BVLCContract
}

// Deprecated: use the interface for direct cast
func CastBVLCReadBroadcastDistributionTable(structType any) BVLCReadBroadcastDistributionTable {
	if casted, ok := structType.(BVLCReadBroadcastDistributionTable); ok {
		return casted
	}
	if casted, ok := structType.(*BVLCReadBroadcastDistributionTable); ok {
		return *casted
	}
	return nil
}

func (m *_BVLCReadBroadcastDistributionTable) GetTypeName() string {
	return "BVLCReadBroadcastDistributionTable"
}

func (m *_BVLCReadBroadcastDistributionTable) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BVLCContract.(*_BVLC).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_BVLCReadBroadcastDistributionTable) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BVLCReadBroadcastDistributionTable) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BVLC) (__bVLCReadBroadcastDistributionTable BVLCReadBroadcastDistributionTable, err error) {
	m.BVLCContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BVLCReadBroadcastDistributionTable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BVLCReadBroadcastDistributionTable")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("BVLCReadBroadcastDistributionTable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BVLCReadBroadcastDistributionTable")
	}

	return m, nil
}

func (m *_BVLCReadBroadcastDistributionTable) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BVLCReadBroadcastDistributionTable) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BVLCReadBroadcastDistributionTable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BVLCReadBroadcastDistributionTable")
		}

		if popErr := writeBuffer.PopContext("BVLCReadBroadcastDistributionTable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BVLCReadBroadcastDistributionTable")
		}
		return nil
	}
	return m.BVLCContract.(*_BVLC).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BVLCReadBroadcastDistributionTable) IsBVLCReadBroadcastDistributionTable() {}

func (m *_BVLCReadBroadcastDistributionTable) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BVLCReadBroadcastDistributionTable) deepCopy() *_BVLCReadBroadcastDistributionTable {
	if m == nil {
		return nil
	}
	_BVLCReadBroadcastDistributionTableCopy := &_BVLCReadBroadcastDistributionTable{
		m.BVLCContract.(*_BVLC).deepCopy(),
	}
	m.BVLCContract.(*_BVLC)._SubType = m
	return _BVLCReadBroadcastDistributionTableCopy
}

func (m *_BVLCReadBroadcastDistributionTable) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
