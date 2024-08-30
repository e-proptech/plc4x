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
package org.apache.plc4x.java.cbus.readwrite;

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

public class CBusPointToMultiPointCommandNormal extends CBusPointToMultiPointCommand
    implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final ApplicationIdContainer application;
  protected final SALData salData;

  // Arguments.
  protected final CBusOptions cBusOptions;
  // Reserved Fields
  private Byte reservedField0;

  public CBusPointToMultiPointCommandNormal(
      byte peekedApplication,
      ApplicationIdContainer application,
      SALData salData,
      CBusOptions cBusOptions) {
    super(peekedApplication, cBusOptions);
    this.application = application;
    this.salData = salData;
    this.cBusOptions = cBusOptions;
  }

  public ApplicationIdContainer getApplication() {
    return application;
  }

  public SALData getSalData() {
    return salData;
  }

  @Override
  protected void serializeCBusPointToMultiPointCommandChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("CBusPointToMultiPointCommandNormal");

    // Simple Field (application)
    writeSimpleEnumField(
        "application",
        "ApplicationIdContainer",
        application,
        new DataWriterEnumDefault<>(
            ApplicationIdContainer::getValue,
            ApplicationIdContainer::name,
            writeUnsignedShort(writeBuffer, 8)));

    // Reserved Field (reserved)
    writeReservedField(
        "reserved",
        reservedField0 != null ? reservedField0 : (byte) 0x00,
        writeByte(writeBuffer, 8));

    // Simple Field (salData)
    writeSimpleField("salData", salData, writeComplex(writeBuffer));

    writeBuffer.popContext("CBusPointToMultiPointCommandNormal");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    CBusPointToMultiPointCommandNormal _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (application)
    lengthInBits += 8;

    // Reserved Field (reserved)
    lengthInBits += 8;

    // Simple field (salData)
    lengthInBits += salData.getLengthInBits();

    return lengthInBits;
  }

  public static CBusPointToMultiPointCommandBuilder staticParseCBusPointToMultiPointCommandBuilder(
      ReadBuffer readBuffer, CBusOptions cBusOptions) throws ParseException {
    readBuffer.pullContext("CBusPointToMultiPointCommandNormal");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    ApplicationIdContainer application =
        readEnumField(
            "application",
            "ApplicationIdContainer",
            new DataReaderEnumDefault<>(
                ApplicationIdContainer::enumForValue, readUnsignedShort(readBuffer, 8)));

    Byte reservedField0 = readReservedField("reserved", readByte(readBuffer, 8), (byte) 0x00);

    SALData salData =
        readSimpleField(
            "salData",
            readComplex(
                () ->
                    SALData.staticParse(
                        readBuffer, (ApplicationId) (application.getApplicationId())),
                readBuffer));

    readBuffer.closeContext("CBusPointToMultiPointCommandNormal");
    // Create the instance
    return new CBusPointToMultiPointCommandNormalBuilderImpl(
        application, salData, cBusOptions, reservedField0);
  }

  public static class CBusPointToMultiPointCommandNormalBuilderImpl
      implements CBusPointToMultiPointCommand.CBusPointToMultiPointCommandBuilder {
    private final ApplicationIdContainer application;
    private final SALData salData;
    private final CBusOptions cBusOptions;
    private final Byte reservedField0;

    public CBusPointToMultiPointCommandNormalBuilderImpl(
        ApplicationIdContainer application,
        SALData salData,
        CBusOptions cBusOptions,
        Byte reservedField0) {
      this.application = application;
      this.salData = salData;
      this.cBusOptions = cBusOptions;
      this.reservedField0 = reservedField0;
    }

    public CBusPointToMultiPointCommandNormal build(
        byte peekedApplication, CBusOptions cBusOptions) {
      CBusPointToMultiPointCommandNormal cBusPointToMultiPointCommandNormal =
          new CBusPointToMultiPointCommandNormal(
              peekedApplication, application, salData, cBusOptions);
      cBusPointToMultiPointCommandNormal.reservedField0 = reservedField0;
      return cBusPointToMultiPointCommandNormal;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof CBusPointToMultiPointCommandNormal)) {
      return false;
    }
    CBusPointToMultiPointCommandNormal that = (CBusPointToMultiPointCommandNormal) o;
    return (getApplication() == that.getApplication())
        && (getSalData() == that.getSalData())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getApplication(), getSalData());
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
