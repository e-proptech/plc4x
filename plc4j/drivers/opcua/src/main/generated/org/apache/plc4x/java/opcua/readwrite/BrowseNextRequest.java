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
package org.apache.plc4x.java.opcua.readwrite;

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

public class BrowseNextRequest extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public String getIdentifier() {
    return (String) "533";
  }

  // Properties.
  protected final ExtensionObjectDefinition requestHeader;
  protected final boolean releaseContinuationPoints;
  protected final int noOfContinuationPoints;
  protected final List<PascalByteString> continuationPoints;

  public BrowseNextRequest(
      ExtensionObjectDefinition requestHeader,
      boolean releaseContinuationPoints,
      int noOfContinuationPoints,
      List<PascalByteString> continuationPoints) {
    super();
    this.requestHeader = requestHeader;
    this.releaseContinuationPoints = releaseContinuationPoints;
    this.noOfContinuationPoints = noOfContinuationPoints;
    this.continuationPoints = continuationPoints;
  }

  public ExtensionObjectDefinition getRequestHeader() {
    return requestHeader;
  }

  public boolean getReleaseContinuationPoints() {
    return releaseContinuationPoints;
  }

  public int getNoOfContinuationPoints() {
    return noOfContinuationPoints;
  }

  public List<PascalByteString> getContinuationPoints() {
    return continuationPoints;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BrowseNextRequest");

    // Simple Field (requestHeader)
    writeSimpleField("requestHeader", requestHeader, writeComplex(writeBuffer));

    // Reserved Field (reserved)
    writeReservedField("reserved", (byte) 0x00, writeUnsignedByte(writeBuffer, 7));

    // Simple Field (releaseContinuationPoints)
    writeSimpleField(
        "releaseContinuationPoints", releaseContinuationPoints, writeBoolean(writeBuffer));

    // Simple Field (noOfContinuationPoints)
    writeSimpleField(
        "noOfContinuationPoints", noOfContinuationPoints, writeSignedInt(writeBuffer, 32));

    // Array Field (continuationPoints)
    writeComplexTypeArrayField("continuationPoints", continuationPoints, writeBuffer);

    writeBuffer.popContext("BrowseNextRequest");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BrowseNextRequest _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (requestHeader)
    lengthInBits += requestHeader.getLengthInBits();

    // Reserved Field (reserved)
    lengthInBits += 7;

    // Simple field (releaseContinuationPoints)
    lengthInBits += 1;

    // Simple field (noOfContinuationPoints)
    lengthInBits += 32;

    // Array field
    if (continuationPoints != null) {
      int i = 0;
      for (PascalByteString element : continuationPoints) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= continuationPoints.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, String identifier) throws ParseException {
    readBuffer.pullContext("BrowseNextRequest");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    ExtensionObjectDefinition requestHeader =
        readSimpleField(
            "requestHeader",
            readComplex(
                () -> ExtensionObjectDefinition.staticParse(readBuffer, (String) ("391")),
                readBuffer));

    Byte reservedField0 =
        readReservedField("reserved", readUnsignedByte(readBuffer, 7), (byte) 0x00);

    boolean releaseContinuationPoints =
        readSimpleField("releaseContinuationPoints", readBoolean(readBuffer));

    int noOfContinuationPoints =
        readSimpleField("noOfContinuationPoints", readSignedInt(readBuffer, 32));

    List<PascalByteString> continuationPoints =
        readCountArrayField(
            "continuationPoints",
            readComplex(() -> PascalByteString.staticParse(readBuffer), readBuffer),
            noOfContinuationPoints);

    readBuffer.closeContext("BrowseNextRequest");
    // Create the instance
    return new BrowseNextRequestBuilderImpl(
        requestHeader, releaseContinuationPoints, noOfContinuationPoints, continuationPoints);
  }

  public static class BrowseNextRequestBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final ExtensionObjectDefinition requestHeader;
    private final boolean releaseContinuationPoints;
    private final int noOfContinuationPoints;
    private final List<PascalByteString> continuationPoints;

    public BrowseNextRequestBuilderImpl(
        ExtensionObjectDefinition requestHeader,
        boolean releaseContinuationPoints,
        int noOfContinuationPoints,
        List<PascalByteString> continuationPoints) {
      this.requestHeader = requestHeader;
      this.releaseContinuationPoints = releaseContinuationPoints;
      this.noOfContinuationPoints = noOfContinuationPoints;
      this.continuationPoints = continuationPoints;
    }

    public BrowseNextRequest build() {
      BrowseNextRequest browseNextRequest =
          new BrowseNextRequest(
              requestHeader, releaseContinuationPoints, noOfContinuationPoints, continuationPoints);
      return browseNextRequest;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BrowseNextRequest)) {
      return false;
    }
    BrowseNextRequest that = (BrowseNextRequest) o;
    return (getRequestHeader() == that.getRequestHeader())
        && (getReleaseContinuationPoints() == that.getReleaseContinuationPoints())
        && (getNoOfContinuationPoints() == that.getNoOfContinuationPoints())
        && (getContinuationPoints() == that.getContinuationPoints())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getRequestHeader(),
        getReleaseContinuationPoints(),
        getNoOfContinuationPoints(),
        getContinuationPoints());
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
