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

public class BACnetConfirmedServiceRequestGetEventInformation extends BACnetConfirmedServiceRequest
    implements Message {

  // Accessors for discriminator values.
  public BACnetConfirmedServiceChoice getServiceChoice() {
    return BACnetConfirmedServiceChoice.GET_EVENT_INFORMATION;
  }

  // Properties.
  protected final BACnetContextTagObjectIdentifier lastReceivedObjectIdentifier;

  // Arguments.
  protected final Long serviceRequestLength;

  public BACnetConfirmedServiceRequestGetEventInformation(
      BACnetContextTagObjectIdentifier lastReceivedObjectIdentifier, Long serviceRequestLength) {
    super(serviceRequestLength);
    this.lastReceivedObjectIdentifier = lastReceivedObjectIdentifier;
    this.serviceRequestLength = serviceRequestLength;
  }

  public BACnetContextTagObjectIdentifier getLastReceivedObjectIdentifier() {
    return lastReceivedObjectIdentifier;
  }

  @Override
  protected void serializeBACnetConfirmedServiceRequestChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetConfirmedServiceRequestGetEventInformation");

    // Optional Field (lastReceivedObjectIdentifier) (Can be skipped, if the value is null)
    writeOptionalField(
        "lastReceivedObjectIdentifier", lastReceivedObjectIdentifier, writeComplex(writeBuffer));

    writeBuffer.popContext("BACnetConfirmedServiceRequestGetEventInformation");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetConfirmedServiceRequestGetEventInformation _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Optional Field (lastReceivedObjectIdentifier)
    if (lastReceivedObjectIdentifier != null) {
      lengthInBits += lastReceivedObjectIdentifier.getLengthInBits();
    }

    return lengthInBits;
  }

  public static BACnetConfirmedServiceRequestBuilder
      staticParseBACnetConfirmedServiceRequestBuilder(
          ReadBuffer readBuffer, Long serviceRequestLength) throws ParseException {
    readBuffer.pullContext("BACnetConfirmedServiceRequestGetEventInformation");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetContextTagObjectIdentifier lastReceivedObjectIdentifier =
        readOptionalField(
            "lastReceivedObjectIdentifier",
            readComplex(
                () ->
                    (BACnetContextTagObjectIdentifier)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (0),
                            (BACnetDataType) (BACnetDataType.BACNET_OBJECT_IDENTIFIER)),
                readBuffer));

    readBuffer.closeContext("BACnetConfirmedServiceRequestGetEventInformation");
    // Create the instance
    return new BACnetConfirmedServiceRequestGetEventInformationBuilderImpl(
        lastReceivedObjectIdentifier, serviceRequestLength);
  }

  public static class BACnetConfirmedServiceRequestGetEventInformationBuilderImpl
      implements BACnetConfirmedServiceRequest.BACnetConfirmedServiceRequestBuilder {
    private final BACnetContextTagObjectIdentifier lastReceivedObjectIdentifier;
    private final Long serviceRequestLength;

    public BACnetConfirmedServiceRequestGetEventInformationBuilderImpl(
        BACnetContextTagObjectIdentifier lastReceivedObjectIdentifier, Long serviceRequestLength) {
      this.lastReceivedObjectIdentifier = lastReceivedObjectIdentifier;
      this.serviceRequestLength = serviceRequestLength;
    }

    public BACnetConfirmedServiceRequestGetEventInformation build(Long serviceRequestLength) {

      BACnetConfirmedServiceRequestGetEventInformation
          bACnetConfirmedServiceRequestGetEventInformation =
              new BACnetConfirmedServiceRequestGetEventInformation(
                  lastReceivedObjectIdentifier, serviceRequestLength);
      return bACnetConfirmedServiceRequestGetEventInformation;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetConfirmedServiceRequestGetEventInformation)) {
      return false;
    }
    BACnetConfirmedServiceRequestGetEventInformation that =
        (BACnetConfirmedServiceRequestGetEventInformation) o;
    return (getLastReceivedObjectIdentifier() == that.getLastReceivedObjectIdentifier())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getLastReceivedObjectIdentifier());
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
