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

// BACnetConstructedDataBias is the corresponding interface of BACnetConstructedDataBias
type BACnetConstructedDataBias interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetBias returns Bias (property field)
	GetBias() BACnetApplicationTagReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagReal
	// IsBACnetConstructedDataBias is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataBias()
	// CreateBuilder creates a BACnetConstructedDataBiasBuilder
	CreateBACnetConstructedDataBiasBuilder() BACnetConstructedDataBiasBuilder
}

// _BACnetConstructedDataBias is the data-structure of this message
type _BACnetConstructedDataBias struct {
	BACnetConstructedDataContract
	Bias BACnetApplicationTagReal
}

var _ BACnetConstructedDataBias = (*_BACnetConstructedDataBias)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataBias)(nil)

// NewBACnetConstructedDataBias factory function for _BACnetConstructedDataBias
func NewBACnetConstructedDataBias(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, bias BACnetApplicationTagReal, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataBias {
	if bias == nil {
		panic("bias of type BACnetApplicationTagReal for BACnetConstructedDataBias must not be nil")
	}
	_result := &_BACnetConstructedDataBias{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		Bias:                          bias,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataBiasBuilder is a builder for BACnetConstructedDataBias
type BACnetConstructedDataBiasBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(bias BACnetApplicationTagReal) BACnetConstructedDataBiasBuilder
	// WithBias adds Bias (property field)
	WithBias(BACnetApplicationTagReal) BACnetConstructedDataBiasBuilder
	// WithBiasBuilder adds Bias (property field) which is build by the builder
	WithBiasBuilder(func(BACnetApplicationTagRealBuilder) BACnetApplicationTagRealBuilder) BACnetConstructedDataBiasBuilder
	// Build builds the BACnetConstructedDataBias or returns an error if something is wrong
	Build() (BACnetConstructedDataBias, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataBias
}

// NewBACnetConstructedDataBiasBuilder() creates a BACnetConstructedDataBiasBuilder
func NewBACnetConstructedDataBiasBuilder() BACnetConstructedDataBiasBuilder {
	return &_BACnetConstructedDataBiasBuilder{_BACnetConstructedDataBias: new(_BACnetConstructedDataBias)}
}

type _BACnetConstructedDataBiasBuilder struct {
	*_BACnetConstructedDataBias

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataBiasBuilder) = (*_BACnetConstructedDataBiasBuilder)(nil)

func (b *_BACnetConstructedDataBiasBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataBiasBuilder) WithMandatoryFields(bias BACnetApplicationTagReal) BACnetConstructedDataBiasBuilder {
	return b.WithBias(bias)
}

func (b *_BACnetConstructedDataBiasBuilder) WithBias(bias BACnetApplicationTagReal) BACnetConstructedDataBiasBuilder {
	b.Bias = bias
	return b
}

func (b *_BACnetConstructedDataBiasBuilder) WithBiasBuilder(builderSupplier func(BACnetApplicationTagRealBuilder) BACnetApplicationTagRealBuilder) BACnetConstructedDataBiasBuilder {
	builder := builderSupplier(b.Bias.CreateBACnetApplicationTagRealBuilder())
	var err error
	b.Bias, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagRealBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataBiasBuilder) Build() (BACnetConstructedDataBias, error) {
	if b.Bias == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'bias' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataBias.deepCopy(), nil
}

func (b *_BACnetConstructedDataBiasBuilder) MustBuild() BACnetConstructedDataBias {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataBiasBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataBiasBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataBiasBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataBiasBuilder().(*_BACnetConstructedDataBiasBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataBiasBuilder creates a BACnetConstructedDataBiasBuilder
func (b *_BACnetConstructedDataBias) CreateBACnetConstructedDataBiasBuilder() BACnetConstructedDataBiasBuilder {
	if b == nil {
		return NewBACnetConstructedDataBiasBuilder()
	}
	return &_BACnetConstructedDataBiasBuilder{_BACnetConstructedDataBias: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataBias) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataBias) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_BIAS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataBias) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataBias) GetBias() BACnetApplicationTagReal {
	return m.Bias
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataBias) GetActualValue() BACnetApplicationTagReal {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagReal(m.GetBias())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataBias(structType any) BACnetConstructedDataBias {
	if casted, ok := structType.(BACnetConstructedDataBias); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBias); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataBias) GetTypeName() string {
	return "BACnetConstructedDataBias"
}

func (m *_BACnetConstructedDataBias) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (bias)
	lengthInBits += m.Bias.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataBias) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataBias) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataBias BACnetConstructedDataBias, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBias"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBias")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	bias, err := ReadSimpleField[BACnetApplicationTagReal](ctx, "bias", ReadComplex[BACnetApplicationTagReal](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagReal](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'bias' field"))
	}
	m.Bias = bias

	actualValue, err := ReadVirtualField[BACnetApplicationTagReal](ctx, "actualValue", (*BACnetApplicationTagReal)(nil), bias)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBias"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBias")
	}

	return m, nil
}

func (m *_BACnetConstructedDataBias) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataBias) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBias"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBias")
		}

		if err := WriteSimpleField[BACnetApplicationTagReal](ctx, "bias", m.GetBias(), WriteComplex[BACnetApplicationTagReal](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'bias' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBias"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBias")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataBias) IsBACnetConstructedDataBias() {}

func (m *_BACnetConstructedDataBias) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataBias) deepCopy() *_BACnetConstructedDataBias {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataBiasCopy := &_BACnetConstructedDataBias{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagReal](m.Bias),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataBiasCopy
}

func (m *_BACnetConstructedDataBias) String() string {
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
