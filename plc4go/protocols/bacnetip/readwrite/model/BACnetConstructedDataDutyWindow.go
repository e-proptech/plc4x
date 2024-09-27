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

// BACnetConstructedDataDutyWindow is the corresponding interface of BACnetConstructedDataDutyWindow
type BACnetConstructedDataDutyWindow interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetDutyWindow returns DutyWindow (property field)
	GetDutyWindow() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataDutyWindow is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataDutyWindow()
	// CreateBuilder creates a BACnetConstructedDataDutyWindowBuilder
	CreateBACnetConstructedDataDutyWindowBuilder() BACnetConstructedDataDutyWindowBuilder
}

// _BACnetConstructedDataDutyWindow is the data-structure of this message
type _BACnetConstructedDataDutyWindow struct {
	BACnetConstructedDataContract
	DutyWindow BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataDutyWindow = (*_BACnetConstructedDataDutyWindow)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataDutyWindow)(nil)

// NewBACnetConstructedDataDutyWindow factory function for _BACnetConstructedDataDutyWindow
func NewBACnetConstructedDataDutyWindow(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, dutyWindow BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataDutyWindow {
	if dutyWindow == nil {
		panic("dutyWindow of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataDutyWindow must not be nil")
	}
	_result := &_BACnetConstructedDataDutyWindow{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		DutyWindow:                    dutyWindow,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataDutyWindowBuilder is a builder for BACnetConstructedDataDutyWindow
type BACnetConstructedDataDutyWindowBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(dutyWindow BACnetApplicationTagUnsignedInteger) BACnetConstructedDataDutyWindowBuilder
	// WithDutyWindow adds DutyWindow (property field)
	WithDutyWindow(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataDutyWindowBuilder
	// WithDutyWindowBuilder adds DutyWindow (property field) which is build by the builder
	WithDutyWindowBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataDutyWindowBuilder
	// Build builds the BACnetConstructedDataDutyWindow or returns an error if something is wrong
	Build() (BACnetConstructedDataDutyWindow, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataDutyWindow
}

// NewBACnetConstructedDataDutyWindowBuilder() creates a BACnetConstructedDataDutyWindowBuilder
func NewBACnetConstructedDataDutyWindowBuilder() BACnetConstructedDataDutyWindowBuilder {
	return &_BACnetConstructedDataDutyWindowBuilder{_BACnetConstructedDataDutyWindow: new(_BACnetConstructedDataDutyWindow)}
}

type _BACnetConstructedDataDutyWindowBuilder struct {
	*_BACnetConstructedDataDutyWindow

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataDutyWindowBuilder) = (*_BACnetConstructedDataDutyWindowBuilder)(nil)

func (b *_BACnetConstructedDataDutyWindowBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataDutyWindowBuilder) WithMandatoryFields(dutyWindow BACnetApplicationTagUnsignedInteger) BACnetConstructedDataDutyWindowBuilder {
	return b.WithDutyWindow(dutyWindow)
}

func (b *_BACnetConstructedDataDutyWindowBuilder) WithDutyWindow(dutyWindow BACnetApplicationTagUnsignedInteger) BACnetConstructedDataDutyWindowBuilder {
	b.DutyWindow = dutyWindow
	return b
}

func (b *_BACnetConstructedDataDutyWindowBuilder) WithDutyWindowBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataDutyWindowBuilder {
	builder := builderSupplier(b.DutyWindow.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.DutyWindow, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataDutyWindowBuilder) Build() (BACnetConstructedDataDutyWindow, error) {
	if b.DutyWindow == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'dutyWindow' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataDutyWindow.deepCopy(), nil
}

func (b *_BACnetConstructedDataDutyWindowBuilder) MustBuild() BACnetConstructedDataDutyWindow {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataDutyWindowBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataDutyWindowBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataDutyWindowBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataDutyWindowBuilder().(*_BACnetConstructedDataDutyWindowBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataDutyWindowBuilder creates a BACnetConstructedDataDutyWindowBuilder
func (b *_BACnetConstructedDataDutyWindow) CreateBACnetConstructedDataDutyWindowBuilder() BACnetConstructedDataDutyWindowBuilder {
	if b == nil {
		return NewBACnetConstructedDataDutyWindowBuilder()
	}
	return &_BACnetConstructedDataDutyWindowBuilder{_BACnetConstructedDataDutyWindow: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDutyWindow) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataDutyWindow) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_DUTY_WINDOW
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDutyWindow) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataDutyWindow) GetDutyWindow() BACnetApplicationTagUnsignedInteger {
	return m.DutyWindow
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataDutyWindow) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetDutyWindow())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDutyWindow(structType any) BACnetConstructedDataDutyWindow {
	if casted, ok := structType.(BACnetConstructedDataDutyWindow); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDutyWindow); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDutyWindow) GetTypeName() string {
	return "BACnetConstructedDataDutyWindow"
}

func (m *_BACnetConstructedDataDutyWindow) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (dutyWindow)
	lengthInBits += m.DutyWindow.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataDutyWindow) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataDutyWindow) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataDutyWindow BACnetConstructedDataDutyWindow, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDutyWindow"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDutyWindow")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	dutyWindow, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "dutyWindow", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dutyWindow' field"))
	}
	m.DutyWindow = dutyWindow

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), dutyWindow)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDutyWindow"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDutyWindow")
	}

	return m, nil
}

func (m *_BACnetConstructedDataDutyWindow) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataDutyWindow) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDutyWindow"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDutyWindow")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "dutyWindow", m.GetDutyWindow(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'dutyWindow' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDutyWindow"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDutyWindow")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataDutyWindow) IsBACnetConstructedDataDutyWindow() {}

func (m *_BACnetConstructedDataDutyWindow) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataDutyWindow) deepCopy() *_BACnetConstructedDataDutyWindow {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataDutyWindowCopy := &_BACnetConstructedDataDutyWindow{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.DutyWindow.DeepCopy().(BACnetApplicationTagUnsignedInteger),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataDutyWindowCopy
}

func (m *_BACnetConstructedDataDutyWindow) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
