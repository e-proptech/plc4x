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
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetPropertyStatesWriteStatus is the corresponding interface of BACnetPropertyStatesWriteStatus
type BACnetPropertyStatesWriteStatus interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetPropertyStates
	// GetWriteStatus returns WriteStatus (property field)
	GetWriteStatus() BACnetWriteStatusTagged
	// IsBACnetPropertyStatesWriteStatus is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyStatesWriteStatus()
	// CreateBuilder creates a BACnetPropertyStatesWriteStatusBuilder
	CreateBACnetPropertyStatesWriteStatusBuilder() BACnetPropertyStatesWriteStatusBuilder
}

// _BACnetPropertyStatesWriteStatus is the data-structure of this message
type _BACnetPropertyStatesWriteStatus struct {
	BACnetPropertyStatesContract
	WriteStatus BACnetWriteStatusTagged
}

var _ BACnetPropertyStatesWriteStatus = (*_BACnetPropertyStatesWriteStatus)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesWriteStatus)(nil)

// NewBACnetPropertyStatesWriteStatus factory function for _BACnetPropertyStatesWriteStatus
func NewBACnetPropertyStatesWriteStatus(peekedTagHeader BACnetTagHeader, writeStatus BACnetWriteStatusTagged) *_BACnetPropertyStatesWriteStatus {
	if writeStatus == nil {
		panic("writeStatus of type BACnetWriteStatusTagged for BACnetPropertyStatesWriteStatus must not be nil")
	}
	_result := &_BACnetPropertyStatesWriteStatus{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		WriteStatus:                  writeStatus,
	}
	_result.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetPropertyStatesWriteStatusBuilder is a builder for BACnetPropertyStatesWriteStatus
type BACnetPropertyStatesWriteStatusBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(writeStatus BACnetWriteStatusTagged) BACnetPropertyStatesWriteStatusBuilder
	// WithWriteStatus adds WriteStatus (property field)
	WithWriteStatus(BACnetWriteStatusTagged) BACnetPropertyStatesWriteStatusBuilder
	// WithWriteStatusBuilder adds WriteStatus (property field) which is build by the builder
	WithWriteStatusBuilder(func(BACnetWriteStatusTaggedBuilder) BACnetWriteStatusTaggedBuilder) BACnetPropertyStatesWriteStatusBuilder
	// Build builds the BACnetPropertyStatesWriteStatus or returns an error if something is wrong
	Build() (BACnetPropertyStatesWriteStatus, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetPropertyStatesWriteStatus
}

// NewBACnetPropertyStatesWriteStatusBuilder() creates a BACnetPropertyStatesWriteStatusBuilder
func NewBACnetPropertyStatesWriteStatusBuilder() BACnetPropertyStatesWriteStatusBuilder {
	return &_BACnetPropertyStatesWriteStatusBuilder{_BACnetPropertyStatesWriteStatus: new(_BACnetPropertyStatesWriteStatus)}
}

type _BACnetPropertyStatesWriteStatusBuilder struct {
	*_BACnetPropertyStatesWriteStatus

	parentBuilder *_BACnetPropertyStatesBuilder

	err *utils.MultiError
}

var _ (BACnetPropertyStatesWriteStatusBuilder) = (*_BACnetPropertyStatesWriteStatusBuilder)(nil)

func (b *_BACnetPropertyStatesWriteStatusBuilder) setParent(contract BACnetPropertyStatesContract) {
	b.BACnetPropertyStatesContract = contract
}

func (b *_BACnetPropertyStatesWriteStatusBuilder) WithMandatoryFields(writeStatus BACnetWriteStatusTagged) BACnetPropertyStatesWriteStatusBuilder {
	return b.WithWriteStatus(writeStatus)
}

func (b *_BACnetPropertyStatesWriteStatusBuilder) WithWriteStatus(writeStatus BACnetWriteStatusTagged) BACnetPropertyStatesWriteStatusBuilder {
	b.WriteStatus = writeStatus
	return b
}

func (b *_BACnetPropertyStatesWriteStatusBuilder) WithWriteStatusBuilder(builderSupplier func(BACnetWriteStatusTaggedBuilder) BACnetWriteStatusTaggedBuilder) BACnetPropertyStatesWriteStatusBuilder {
	builder := builderSupplier(b.WriteStatus.CreateBACnetWriteStatusTaggedBuilder())
	var err error
	b.WriteStatus, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetWriteStatusTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetPropertyStatesWriteStatusBuilder) Build() (BACnetPropertyStatesWriteStatus, error) {
	if b.WriteStatus == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'writeStatus' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetPropertyStatesWriteStatus.deepCopy(), nil
}

func (b *_BACnetPropertyStatesWriteStatusBuilder) MustBuild() BACnetPropertyStatesWriteStatus {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetPropertyStatesWriteStatusBuilder) Done() BACnetPropertyStatesBuilder {
	return b.parentBuilder
}

func (b *_BACnetPropertyStatesWriteStatusBuilder) buildForBACnetPropertyStates() (BACnetPropertyStates, error) {
	return b.Build()
}

func (b *_BACnetPropertyStatesWriteStatusBuilder) DeepCopy() any {
	_copy := b.CreateBACnetPropertyStatesWriteStatusBuilder().(*_BACnetPropertyStatesWriteStatusBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetPropertyStatesWriteStatusBuilder creates a BACnetPropertyStatesWriteStatusBuilder
func (b *_BACnetPropertyStatesWriteStatus) CreateBACnetPropertyStatesWriteStatusBuilder() BACnetPropertyStatesWriteStatusBuilder {
	if b == nil {
		return NewBACnetPropertyStatesWriteStatusBuilder()
	}
	return &_BACnetPropertyStatesWriteStatusBuilder{_BACnetPropertyStatesWriteStatus: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesWriteStatus) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesWriteStatus) GetWriteStatus() BACnetWriteStatusTagged {
	return m.WriteStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesWriteStatus(structType any) BACnetPropertyStatesWriteStatus {
	if casted, ok := structType.(BACnetPropertyStatesWriteStatus); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesWriteStatus); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesWriteStatus) GetTypeName() string {
	return "BACnetPropertyStatesWriteStatus"
}

func (m *_BACnetPropertyStatesWriteStatus) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (writeStatus)
	lengthInBits += m.WriteStatus.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesWriteStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesWriteStatus) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesWriteStatus BACnetPropertyStatesWriteStatus, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesWriteStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesWriteStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	writeStatus, err := ReadSimpleField[BACnetWriteStatusTagged](ctx, "writeStatus", ReadComplex[BACnetWriteStatusTagged](BACnetWriteStatusTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'writeStatus' field"))
	}
	m.WriteStatus = writeStatus

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesWriteStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesWriteStatus")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesWriteStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesWriteStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesWriteStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesWriteStatus")
		}

		if err := WriteSimpleField[BACnetWriteStatusTagged](ctx, "writeStatus", m.GetWriteStatus(), WriteComplex[BACnetWriteStatusTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'writeStatus' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesWriteStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesWriteStatus")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesWriteStatus) IsBACnetPropertyStatesWriteStatus() {}

func (m *_BACnetPropertyStatesWriteStatus) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetPropertyStatesWriteStatus) deepCopy() *_BACnetPropertyStatesWriteStatus {
	if m == nil {
		return nil
	}
	_BACnetPropertyStatesWriteStatusCopy := &_BACnetPropertyStatesWriteStatus{
		m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).deepCopy(),
		utils.DeepCopy[BACnetWriteStatusTagged](m.WriteStatus),
	}
	m.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = m
	return _BACnetPropertyStatesWriteStatusCopy
}

func (m *_BACnetPropertyStatesWriteStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}
