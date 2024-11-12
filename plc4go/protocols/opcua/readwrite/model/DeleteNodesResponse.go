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

// DeleteNodesResponse is the corresponding interface of DeleteNodesResponse
type DeleteNodesResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetResponseHeader returns ResponseHeader (property field)
	GetResponseHeader() ResponseHeader
	// GetResults returns Results (property field)
	GetResults() []StatusCode
	// GetDiagnosticInfos returns DiagnosticInfos (property field)
	GetDiagnosticInfos() []DiagnosticInfo
	// IsDeleteNodesResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsDeleteNodesResponse()
	// CreateBuilder creates a DeleteNodesResponseBuilder
	CreateDeleteNodesResponseBuilder() DeleteNodesResponseBuilder
}

// _DeleteNodesResponse is the data-structure of this message
type _DeleteNodesResponse struct {
	ExtensionObjectDefinitionContract
	ResponseHeader  ResponseHeader
	Results         []StatusCode
	DiagnosticInfos []DiagnosticInfo
}

var _ DeleteNodesResponse = (*_DeleteNodesResponse)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_DeleteNodesResponse)(nil)

// NewDeleteNodesResponse factory function for _DeleteNodesResponse
func NewDeleteNodesResponse(responseHeader ResponseHeader, results []StatusCode, diagnosticInfos []DiagnosticInfo) *_DeleteNodesResponse {
	if responseHeader == nil {
		panic("responseHeader of type ResponseHeader for DeleteNodesResponse must not be nil")
	}
	_result := &_DeleteNodesResponse{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		ResponseHeader:                    responseHeader,
		Results:                           results,
		DiagnosticInfos:                   diagnosticInfos,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// DeleteNodesResponseBuilder is a builder for DeleteNodesResponse
type DeleteNodesResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(responseHeader ResponseHeader, results []StatusCode, diagnosticInfos []DiagnosticInfo) DeleteNodesResponseBuilder
	// WithResponseHeader adds ResponseHeader (property field)
	WithResponseHeader(ResponseHeader) DeleteNodesResponseBuilder
	// WithResponseHeaderBuilder adds ResponseHeader (property field) which is build by the builder
	WithResponseHeaderBuilder(func(ResponseHeaderBuilder) ResponseHeaderBuilder) DeleteNodesResponseBuilder
	// WithResults adds Results (property field)
	WithResults(...StatusCode) DeleteNodesResponseBuilder
	// WithDiagnosticInfos adds DiagnosticInfos (property field)
	WithDiagnosticInfos(...DiagnosticInfo) DeleteNodesResponseBuilder
	// Build builds the DeleteNodesResponse or returns an error if something is wrong
	Build() (DeleteNodesResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() DeleteNodesResponse
}

// NewDeleteNodesResponseBuilder() creates a DeleteNodesResponseBuilder
func NewDeleteNodesResponseBuilder() DeleteNodesResponseBuilder {
	return &_DeleteNodesResponseBuilder{_DeleteNodesResponse: new(_DeleteNodesResponse)}
}

type _DeleteNodesResponseBuilder struct {
	*_DeleteNodesResponse

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (DeleteNodesResponseBuilder) = (*_DeleteNodesResponseBuilder)(nil)

func (b *_DeleteNodesResponseBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_DeleteNodesResponseBuilder) WithMandatoryFields(responseHeader ResponseHeader, results []StatusCode, diagnosticInfos []DiagnosticInfo) DeleteNodesResponseBuilder {
	return b.WithResponseHeader(responseHeader).WithResults(results...).WithDiagnosticInfos(diagnosticInfos...)
}

func (b *_DeleteNodesResponseBuilder) WithResponseHeader(responseHeader ResponseHeader) DeleteNodesResponseBuilder {
	b.ResponseHeader = responseHeader
	return b
}

func (b *_DeleteNodesResponseBuilder) WithResponseHeaderBuilder(builderSupplier func(ResponseHeaderBuilder) ResponseHeaderBuilder) DeleteNodesResponseBuilder {
	builder := builderSupplier(b.ResponseHeader.CreateResponseHeaderBuilder())
	var err error
	b.ResponseHeader, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ResponseHeaderBuilder failed"))
	}
	return b
}

func (b *_DeleteNodesResponseBuilder) WithResults(results ...StatusCode) DeleteNodesResponseBuilder {
	b.Results = results
	return b
}

func (b *_DeleteNodesResponseBuilder) WithDiagnosticInfos(diagnosticInfos ...DiagnosticInfo) DeleteNodesResponseBuilder {
	b.DiagnosticInfos = diagnosticInfos
	return b
}

func (b *_DeleteNodesResponseBuilder) Build() (DeleteNodesResponse, error) {
	if b.ResponseHeader == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'responseHeader' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._DeleteNodesResponse.deepCopy(), nil
}

func (b *_DeleteNodesResponseBuilder) MustBuild() DeleteNodesResponse {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_DeleteNodesResponseBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_DeleteNodesResponseBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_DeleteNodesResponseBuilder) DeepCopy() any {
	_copy := b.CreateDeleteNodesResponseBuilder().(*_DeleteNodesResponseBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateDeleteNodesResponseBuilder creates a DeleteNodesResponseBuilder
func (b *_DeleteNodesResponse) CreateDeleteNodesResponseBuilder() DeleteNodesResponseBuilder {
	if b == nil {
		return NewDeleteNodesResponseBuilder()
	}
	return &_DeleteNodesResponseBuilder{_DeleteNodesResponse: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DeleteNodesResponse) GetExtensionId() int32 {
	return int32(503)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DeleteNodesResponse) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DeleteNodesResponse) GetResponseHeader() ResponseHeader {
	return m.ResponseHeader
}

func (m *_DeleteNodesResponse) GetResults() []StatusCode {
	return m.Results
}

func (m *_DeleteNodesResponse) GetDiagnosticInfos() []DiagnosticInfo {
	return m.DiagnosticInfos
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastDeleteNodesResponse(structType any) DeleteNodesResponse {
	if casted, ok := structType.(DeleteNodesResponse); ok {
		return casted
	}
	if casted, ok := structType.(*DeleteNodesResponse); ok {
		return *casted
	}
	return nil
}

func (m *_DeleteNodesResponse) GetTypeName() string {
	return "DeleteNodesResponse"
}

func (m *_DeleteNodesResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (responseHeader)
	lengthInBits += m.ResponseHeader.GetLengthInBits(ctx)

	// Implicit Field (noOfResults)
	lengthInBits += 32

	// Array field
	if len(m.Results) > 0 {
		for _curItem, element := range m.Results {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Results), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Implicit Field (noOfDiagnosticInfos)
	lengthInBits += 32

	// Array field
	if len(m.DiagnosticInfos) > 0 {
		for _curItem, element := range m.DiagnosticInfos {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.DiagnosticInfos), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_DeleteNodesResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_DeleteNodesResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__deleteNodesResponse DeleteNodesResponse, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DeleteNodesResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DeleteNodesResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	responseHeader, err := ReadSimpleField[ResponseHeader](ctx, "responseHeader", ReadComplex[ResponseHeader](ExtensionObjectDefinitionParseWithBufferProducer[ResponseHeader]((int32)(int32(394))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'responseHeader' field"))
	}
	m.ResponseHeader = responseHeader

	noOfResults, err := ReadImplicitField[int32](ctx, "noOfResults", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfResults' field"))
	}
	_ = noOfResults

	results, err := ReadCountArrayField[StatusCode](ctx, "results", ReadComplex[StatusCode](StatusCodeParseWithBuffer, readBuffer), uint64(noOfResults))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'results' field"))
	}
	m.Results = results

	noOfDiagnosticInfos, err := ReadImplicitField[int32](ctx, "noOfDiagnosticInfos", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfDiagnosticInfos' field"))
	}
	_ = noOfDiagnosticInfos

	diagnosticInfos, err := ReadCountArrayField[DiagnosticInfo](ctx, "diagnosticInfos", ReadComplex[DiagnosticInfo](DiagnosticInfoParseWithBuffer, readBuffer), uint64(noOfDiagnosticInfos))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'diagnosticInfos' field"))
	}
	m.DiagnosticInfos = diagnosticInfos

	if closeErr := readBuffer.CloseContext("DeleteNodesResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DeleteNodesResponse")
	}

	return m, nil
}

func (m *_DeleteNodesResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DeleteNodesResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DeleteNodesResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DeleteNodesResponse")
		}

		if err := WriteSimpleField[ResponseHeader](ctx, "responseHeader", m.GetResponseHeader(), WriteComplex[ResponseHeader](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'responseHeader' field")
		}
		noOfResults := int32(utils.InlineIf(bool((m.GetResults()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetResults()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfResults", noOfResults, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfResults' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "results", m.GetResults(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'results' field")
		}
		noOfDiagnosticInfos := int32(utils.InlineIf(bool((m.GetDiagnosticInfos()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetDiagnosticInfos()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfDiagnosticInfos", noOfDiagnosticInfos, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfDiagnosticInfos' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "diagnosticInfos", m.GetDiagnosticInfos(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'diagnosticInfos' field")
		}

		if popErr := writeBuffer.PopContext("DeleteNodesResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DeleteNodesResponse")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DeleteNodesResponse) IsDeleteNodesResponse() {}

func (m *_DeleteNodesResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_DeleteNodesResponse) deepCopy() *_DeleteNodesResponse {
	if m == nil {
		return nil
	}
	_DeleteNodesResponseCopy := &_DeleteNodesResponse{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopy[ResponseHeader](m.ResponseHeader),
		utils.DeepCopySlice[StatusCode, StatusCode](m.Results),
		utils.DeepCopySlice[DiagnosticInfo, DiagnosticInfo](m.DiagnosticInfos),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _DeleteNodesResponseCopy
}

func (m *_DeleteNodesResponse) String() string {
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
