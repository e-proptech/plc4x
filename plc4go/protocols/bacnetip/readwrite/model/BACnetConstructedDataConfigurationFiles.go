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

// BACnetConstructedDataConfigurationFiles is the corresponding interface of BACnetConstructedDataConfigurationFiles
type BACnetConstructedDataConfigurationFiles interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetConfigurationFiles returns ConfigurationFiles (property field)
	GetConfigurationFiles() []BACnetApplicationTagObjectIdentifier
	// GetZero returns Zero (virtual field)
	GetZero() uint64
	// IsBACnetConstructedDataConfigurationFiles is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataConfigurationFiles()
	// CreateBuilder creates a BACnetConstructedDataConfigurationFilesBuilder
	CreateBACnetConstructedDataConfigurationFilesBuilder() BACnetConstructedDataConfigurationFilesBuilder
}

// _BACnetConstructedDataConfigurationFiles is the data-structure of this message
type _BACnetConstructedDataConfigurationFiles struct {
	BACnetConstructedDataContract
	NumberOfDataElements BACnetApplicationTagUnsignedInteger
	ConfigurationFiles   []BACnetApplicationTagObjectIdentifier
}

var _ BACnetConstructedDataConfigurationFiles = (*_BACnetConstructedDataConfigurationFiles)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataConfigurationFiles)(nil)

// NewBACnetConstructedDataConfigurationFiles factory function for _BACnetConstructedDataConfigurationFiles
func NewBACnetConstructedDataConfigurationFiles(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, numberOfDataElements BACnetApplicationTagUnsignedInteger, configurationFiles []BACnetApplicationTagObjectIdentifier, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataConfigurationFiles {
	_result := &_BACnetConstructedDataConfigurationFiles{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		NumberOfDataElements:          numberOfDataElements,
		ConfigurationFiles:            configurationFiles,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataConfigurationFilesBuilder is a builder for BACnetConstructedDataConfigurationFiles
type BACnetConstructedDataConfigurationFilesBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(configurationFiles []BACnetApplicationTagObjectIdentifier) BACnetConstructedDataConfigurationFilesBuilder
	// WithNumberOfDataElements adds NumberOfDataElements (property field)
	WithOptionalNumberOfDataElements(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataConfigurationFilesBuilder
	// WithOptionalNumberOfDataElementsBuilder adds NumberOfDataElements (property field) which is build by the builder
	WithOptionalNumberOfDataElementsBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataConfigurationFilesBuilder
	// WithConfigurationFiles adds ConfigurationFiles (property field)
	WithConfigurationFiles(...BACnetApplicationTagObjectIdentifier) BACnetConstructedDataConfigurationFilesBuilder
	// Build builds the BACnetConstructedDataConfigurationFiles or returns an error if something is wrong
	Build() (BACnetConstructedDataConfigurationFiles, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataConfigurationFiles
}

// NewBACnetConstructedDataConfigurationFilesBuilder() creates a BACnetConstructedDataConfigurationFilesBuilder
func NewBACnetConstructedDataConfigurationFilesBuilder() BACnetConstructedDataConfigurationFilesBuilder {
	return &_BACnetConstructedDataConfigurationFilesBuilder{_BACnetConstructedDataConfigurationFiles: new(_BACnetConstructedDataConfigurationFiles)}
}

type _BACnetConstructedDataConfigurationFilesBuilder struct {
	*_BACnetConstructedDataConfigurationFiles

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataConfigurationFilesBuilder) = (*_BACnetConstructedDataConfigurationFilesBuilder)(nil)

func (b *_BACnetConstructedDataConfigurationFilesBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataConfigurationFilesBuilder) WithMandatoryFields(configurationFiles []BACnetApplicationTagObjectIdentifier) BACnetConstructedDataConfigurationFilesBuilder {
	return b.WithConfigurationFiles(configurationFiles...)
}

func (b *_BACnetConstructedDataConfigurationFilesBuilder) WithOptionalNumberOfDataElements(numberOfDataElements BACnetApplicationTagUnsignedInteger) BACnetConstructedDataConfigurationFilesBuilder {
	b.NumberOfDataElements = numberOfDataElements
	return b
}

func (b *_BACnetConstructedDataConfigurationFilesBuilder) WithOptionalNumberOfDataElementsBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataConfigurationFilesBuilder {
	builder := builderSupplier(b.NumberOfDataElements.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.NumberOfDataElements, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataConfigurationFilesBuilder) WithConfigurationFiles(configurationFiles ...BACnetApplicationTagObjectIdentifier) BACnetConstructedDataConfigurationFilesBuilder {
	b.ConfigurationFiles = configurationFiles
	return b
}

func (b *_BACnetConstructedDataConfigurationFilesBuilder) Build() (BACnetConstructedDataConfigurationFiles, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataConfigurationFiles.deepCopy(), nil
}

func (b *_BACnetConstructedDataConfigurationFilesBuilder) MustBuild() BACnetConstructedDataConfigurationFiles {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataConfigurationFilesBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataConfigurationFilesBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataConfigurationFilesBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataConfigurationFilesBuilder().(*_BACnetConstructedDataConfigurationFilesBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataConfigurationFilesBuilder creates a BACnetConstructedDataConfigurationFilesBuilder
func (b *_BACnetConstructedDataConfigurationFiles) CreateBACnetConstructedDataConfigurationFilesBuilder() BACnetConstructedDataConfigurationFilesBuilder {
	if b == nil {
		return NewBACnetConstructedDataConfigurationFilesBuilder()
	}
	return &_BACnetConstructedDataConfigurationFilesBuilder{_BACnetConstructedDataConfigurationFiles: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataConfigurationFiles) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataConfigurationFiles) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_CONFIGURATION_FILES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataConfigurationFiles) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataConfigurationFiles) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetConstructedDataConfigurationFiles) GetConfigurationFiles() []BACnetApplicationTagObjectIdentifier {
	return m.ConfigurationFiles
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataConfigurationFiles) GetZero() uint64 {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.GetNumberOfDataElements()
	_ = numberOfDataElements
	return uint64(uint64(0))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataConfigurationFiles(structType any) BACnetConstructedDataConfigurationFiles {
	if casted, ok := structType.(BACnetConstructedDataConfigurationFiles); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataConfigurationFiles); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataConfigurationFiles) GetTypeName() string {
	return "BACnetConstructedDataConfigurationFiles"
}

func (m *_BACnetConstructedDataConfigurationFiles) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits(ctx)
	}

	// Array field
	if len(m.ConfigurationFiles) > 0 {
		for _, element := range m.ConfigurationFiles {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataConfigurationFiles) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataConfigurationFiles) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataConfigurationFiles BACnetConstructedDataConfigurationFiles, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataConfigurationFiles"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataConfigurationFiles")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	zero, err := ReadVirtualField[uint64](ctx, "zero", (*uint64)(nil), uint64(0))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zero' field"))
	}
	_ = zero

	var numberOfDataElements BACnetApplicationTagUnsignedInteger
	_numberOfDataElements, err := ReadOptionalField[BACnetApplicationTagUnsignedInteger](ctx, "numberOfDataElements", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer), bool(bool((arrayIndexArgument) != (nil))) && bool(bool((arrayIndexArgument.GetActualValue()) == (zero))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numberOfDataElements' field"))
	}
	if _numberOfDataElements != nil {
		numberOfDataElements = *_numberOfDataElements
		m.NumberOfDataElements = numberOfDataElements
	}

	configurationFiles, err := ReadTerminatedArrayField[BACnetApplicationTagObjectIdentifier](ctx, "configurationFiles", ReadComplex[BACnetApplicationTagObjectIdentifier](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagObjectIdentifier](), readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'configurationFiles' field"))
	}
	m.ConfigurationFiles = configurationFiles

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataConfigurationFiles"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataConfigurationFiles")
	}

	return m, nil
}

func (m *_BACnetConstructedDataConfigurationFiles) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataConfigurationFiles) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataConfigurationFiles"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataConfigurationFiles")
		}
		// Virtual field
		zero := m.GetZero()
		_ = zero
		if _zeroErr := writeBuffer.WriteVirtual(ctx, "zero", m.GetZero()); _zeroErr != nil {
			return errors.Wrap(_zeroErr, "Error serializing 'zero' field")
		}

		if err := WriteOptionalField[BACnetApplicationTagUnsignedInteger](ctx, "numberOfDataElements", GetRef(m.GetNumberOfDataElements()), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'numberOfDataElements' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "configurationFiles", m.GetConfigurationFiles(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'configurationFiles' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataConfigurationFiles"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataConfigurationFiles")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataConfigurationFiles) IsBACnetConstructedDataConfigurationFiles() {}

func (m *_BACnetConstructedDataConfigurationFiles) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataConfigurationFiles) deepCopy() *_BACnetConstructedDataConfigurationFiles {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataConfigurationFilesCopy := &_BACnetConstructedDataConfigurationFiles{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.NumberOfDataElements.DeepCopy().(BACnetApplicationTagUnsignedInteger),
		utils.DeepCopySlice[BACnetApplicationTagObjectIdentifier, BACnetApplicationTagObjectIdentifier](m.ConfigurationFiles),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataConfigurationFilesCopy
}

func (m *_BACnetConstructedDataConfigurationFiles) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
