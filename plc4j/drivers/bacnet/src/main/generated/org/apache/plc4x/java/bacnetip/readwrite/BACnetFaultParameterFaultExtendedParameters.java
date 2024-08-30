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
package org.apache.plc4x.java.bacnetip.readwrite;

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class BACnetFaultParameterFaultExtendedParameters implements Message {

  // Properties.
  protected final BACnetOpeningTag openingTag;
  protected final List<BACnetFaultParameterFaultExtendedParametersEntry> parameters;
  protected final BACnetClosingTag closingTag;

  // Arguments.
  protected final Short tagNumber;

  public BACnetFaultParameterFaultExtendedParameters(
      BACnetOpeningTag openingTag,
      List<BACnetFaultParameterFaultExtendedParametersEntry> parameters,
      BACnetClosingTag closingTag,
      Short tagNumber) {
    super();
    this.openingTag = openingTag;
    this.parameters = parameters;
    this.closingTag = closingTag;
    this.tagNumber = tagNumber;
  }

  public BACnetOpeningTag getOpeningTag() {
    return openingTag;
  }

  public List<BACnetFaultParameterFaultExtendedParametersEntry> getParameters() {
    return parameters;
  }

  public BACnetClosingTag getClosingTag() {
    return closingTag;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetFaultParameterFaultExtendedParameters");

    // Simple Field (openingTag)
    writeSimpleField("openingTag", openingTag, writeComplex(writeBuffer));

    // Array Field (parameters)
    writeComplexTypeArrayField("parameters", parameters, writeBuffer);

    // Simple Field (closingTag)
    writeSimpleField("closingTag", closingTag, writeComplex(writeBuffer));

    writeBuffer.popContext("BACnetFaultParameterFaultExtendedParameters");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    BACnetFaultParameterFaultExtendedParameters _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (openingTag)
    lengthInBits += openingTag.getLengthInBits();

    // Array field
    if (parameters != null) {
      for (Message element : parameters) {
        lengthInBits += element.getLengthInBits();
      }
    }

    // Simple field (closingTag)
    lengthInBits += closingTag.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetFaultParameterFaultExtendedParameters staticParse(
      ReadBuffer readBuffer, Short tagNumber) throws ParseException {
    readBuffer.pullContext("BACnetFaultParameterFaultExtendedParameters");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetOpeningTag openingTag =
        readSimpleField(
            "openingTag",
            readComplex(
                () -> BACnetOpeningTag.staticParse(readBuffer, (short) (tagNumber)), readBuffer));

    List<BACnetFaultParameterFaultExtendedParametersEntry> parameters =
        readTerminatedArrayField(
            "parameters",
            readComplex(
                () -> BACnetFaultParameterFaultExtendedParametersEntry.staticParse(readBuffer),
                readBuffer),
            () ->
                ((boolean)
                    (org.apache.plc4x.java.bacnetip.readwrite.utils.StaticHelper
                        .isBACnetConstructedDataClosingTag(readBuffer, false, tagNumber))));

    BACnetClosingTag closingTag =
        readSimpleField(
            "closingTag",
            readComplex(
                () -> BACnetClosingTag.staticParse(readBuffer, (short) (tagNumber)), readBuffer));

    readBuffer.closeContext("BACnetFaultParameterFaultExtendedParameters");
    // Create the instance
    BACnetFaultParameterFaultExtendedParameters _bACnetFaultParameterFaultExtendedParameters;
    _bACnetFaultParameterFaultExtendedParameters =
        new BACnetFaultParameterFaultExtendedParameters(
            openingTag, parameters, closingTag, tagNumber);
    return _bACnetFaultParameterFaultExtendedParameters;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetFaultParameterFaultExtendedParameters)) {
      return false;
    }
    BACnetFaultParameterFaultExtendedParameters that =
        (BACnetFaultParameterFaultExtendedParameters) o;
    return (getOpeningTag() == that.getOpeningTag())
        && (getParameters() == that.getParameters())
        && (getClosingTag() == that.getClosingTag())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getOpeningTag(), getParameters(), getClosingTag());
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}
