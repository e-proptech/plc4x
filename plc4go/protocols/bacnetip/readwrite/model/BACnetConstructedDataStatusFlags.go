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

// BACnetConstructedDataStatusFlags is the corresponding interface of BACnetConstructedDataStatusFlags
type BACnetConstructedDataStatusFlags interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetStatusFlags returns StatusFlags (property field)
	GetStatusFlags() BACnetStatusFlagsTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetStatusFlagsTagged
	// IsBACnetConstructedDataStatusFlags is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataStatusFlags()
	// CreateBuilder creates a BACnetConstructedDataStatusFlagsBuilder
	CreateBACnetConstructedDataStatusFlagsBuilder() BACnetConstructedDataStatusFlagsBuilder
}

// _BACnetConstructedDataStatusFlags is the data-structure of this message
type _BACnetConstructedDataStatusFlags struct {
	BACnetConstructedDataContract
	StatusFlags BACnetStatusFlagsTagged
}

var _ BACnetConstructedDataStatusFlags = (*_BACnetConstructedDataStatusFlags)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataStatusFlags)(nil)

// NewBACnetConstructedDataStatusFlags factory function for _BACnetConstructedDataStatusFlags
func NewBACnetConstructedDataStatusFlags(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, statusFlags BACnetStatusFlagsTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataStatusFlags {
	if statusFlags == nil {
		panic("statusFlags of type BACnetStatusFlagsTagged for BACnetConstructedDataStatusFlags must not be nil")
	}
	_result := &_BACnetConstructedDataStatusFlags{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		StatusFlags:                   statusFlags,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataStatusFlagsBuilder is a builder for BACnetConstructedDataStatusFlags
type BACnetConstructedDataStatusFlagsBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(statusFlags BACnetStatusFlagsTagged) BACnetConstructedDataStatusFlagsBuilder
	// WithStatusFlags adds StatusFlags (property field)
	WithStatusFlags(BACnetStatusFlagsTagged) BACnetConstructedDataStatusFlagsBuilder
	// WithStatusFlagsBuilder adds StatusFlags (property field) which is build by the builder
	WithStatusFlagsBuilder(func(BACnetStatusFlagsTaggedBuilder) BACnetStatusFlagsTaggedBuilder) BACnetConstructedDataStatusFlagsBuilder
	// Build builds the BACnetConstructedDataStatusFlags or returns an error if something is wrong
	Build() (BACnetConstructedDataStatusFlags, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataStatusFlags
}

// NewBACnetConstructedDataStatusFlagsBuilder() creates a BACnetConstructedDataStatusFlagsBuilder
func NewBACnetConstructedDataStatusFlagsBuilder() BACnetConstructedDataStatusFlagsBuilder {
	return &_BACnetConstructedDataStatusFlagsBuilder{_BACnetConstructedDataStatusFlags: new(_BACnetConstructedDataStatusFlags)}
}

type _BACnetConstructedDataStatusFlagsBuilder struct {
	*_BACnetConstructedDataStatusFlags

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataStatusFlagsBuilder) = (*_BACnetConstructedDataStatusFlagsBuilder)(nil)

func (b *_BACnetConstructedDataStatusFlagsBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataStatusFlagsBuilder) WithMandatoryFields(statusFlags BACnetStatusFlagsTagged) BACnetConstructedDataStatusFlagsBuilder {
	return b.WithStatusFlags(statusFlags)
}

func (b *_BACnetConstructedDataStatusFlagsBuilder) WithStatusFlags(statusFlags BACnetStatusFlagsTagged) BACnetConstructedDataStatusFlagsBuilder {
	b.StatusFlags = statusFlags
	return b
}

func (b *_BACnetConstructedDataStatusFlagsBuilder) WithStatusFlagsBuilder(builderSupplier func(BACnetStatusFlagsTaggedBuilder) BACnetStatusFlagsTaggedBuilder) BACnetConstructedDataStatusFlagsBuilder {
	builder := builderSupplier(b.StatusFlags.CreateBACnetStatusFlagsTaggedBuilder())
	var err error
	b.StatusFlags, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetStatusFlagsTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataStatusFlagsBuilder) Build() (BACnetConstructedDataStatusFlags, error) {
	if b.StatusFlags == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'statusFlags' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataStatusFlags.deepCopy(), nil
}

func (b *_BACnetConstructedDataStatusFlagsBuilder) MustBuild() BACnetConstructedDataStatusFlags {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataStatusFlagsBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataStatusFlagsBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataStatusFlagsBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataStatusFlagsBuilder().(*_BACnetConstructedDataStatusFlagsBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataStatusFlagsBuilder creates a BACnetConstructedDataStatusFlagsBuilder
func (b *_BACnetConstructedDataStatusFlags) CreateBACnetConstructedDataStatusFlagsBuilder() BACnetConstructedDataStatusFlagsBuilder {
	if b == nil {
		return NewBACnetConstructedDataStatusFlagsBuilder()
	}
	return &_BACnetConstructedDataStatusFlagsBuilder{_BACnetConstructedDataStatusFlags: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataStatusFlags) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataStatusFlags) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_STATUS_FLAGS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataStatusFlags) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataStatusFlags) GetStatusFlags() BACnetStatusFlagsTagged {
	return m.StatusFlags
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataStatusFlags) GetActualValue() BACnetStatusFlagsTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetStatusFlagsTagged(m.GetStatusFlags())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataStatusFlags(structType any) BACnetConstructedDataStatusFlags {
	if casted, ok := structType.(BACnetConstructedDataStatusFlags); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataStatusFlags); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataStatusFlags) GetTypeName() string {
	return "BACnetConstructedDataStatusFlags"
}

func (m *_BACnetConstructedDataStatusFlags) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (statusFlags)
	lengthInBits += m.StatusFlags.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataStatusFlags) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataStatusFlags) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataStatusFlags BACnetConstructedDataStatusFlags, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataStatusFlags"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataStatusFlags")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	statusFlags, err := ReadSimpleField[BACnetStatusFlagsTagged](ctx, "statusFlags", ReadComplex[BACnetStatusFlagsTagged](BACnetStatusFlagsTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'statusFlags' field"))
	}
	m.StatusFlags = statusFlags

	actualValue, err := ReadVirtualField[BACnetStatusFlagsTagged](ctx, "actualValue", (*BACnetStatusFlagsTagged)(nil), statusFlags)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataStatusFlags"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataStatusFlags")
	}

	return m, nil
}

func (m *_BACnetConstructedDataStatusFlags) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataStatusFlags) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataStatusFlags"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataStatusFlags")
		}

		if err := WriteSimpleField[BACnetStatusFlagsTagged](ctx, "statusFlags", m.GetStatusFlags(), WriteComplex[BACnetStatusFlagsTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'statusFlags' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataStatusFlags"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataStatusFlags")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataStatusFlags) IsBACnetConstructedDataStatusFlags() {}

func (m *_BACnetConstructedDataStatusFlags) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataStatusFlags) deepCopy() *_BACnetConstructedDataStatusFlags {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataStatusFlagsCopy := &_BACnetConstructedDataStatusFlags{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetStatusFlagsTagged](m.StatusFlags),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataStatusFlagsCopy
}

func (m *_BACnetConstructedDataStatusFlags) String() string {
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
