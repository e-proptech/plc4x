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
package org.apache.plc4x.java.canopen.readwrite;

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

public class CANOpenSDORequest extends CANOpenPayload implements Message {

  // Accessors for discriminator values.
  public CANOpenService getService() {
    return CANOpenService.RECEIVE_SDO;
  }

  // Properties.
  protected final SDORequestCommand command;
  protected final SDORequest request;

  public CANOpenSDORequest(SDORequestCommand command, SDORequest request) {
    super();
    this.command = command;
    this.request = request;
  }

  public SDORequestCommand getCommand() {
    return command;
  }

  public SDORequest getRequest() {
    return request;
  }

  @Override
  protected void serializeCANOpenPayloadChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("CANOpenSDORequest");

    // Simple Field (command)
    writeSimpleEnumField(
        "command",
        "SDORequestCommand",
        command,
        new DataWriterEnumDefault<>(
            SDORequestCommand::getValue,
            SDORequestCommand::name,
            writeUnsignedByte(writeBuffer, 3)));

    // Simple Field (request)
    writeSimpleField("request", request, writeComplex(writeBuffer));

    writeBuffer.popContext("CANOpenSDORequest");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    CANOpenSDORequest _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (command)
    lengthInBits += 3;

    // Simple field (request)
    lengthInBits += request.getLengthInBits();

    return lengthInBits;
  }

  public static CANOpenPayloadBuilder staticParseCANOpenPayloadBuilder(
      ReadBuffer readBuffer, CANOpenService service) throws ParseException {
    readBuffer.pullContext("CANOpenSDORequest");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    SDORequestCommand command =
        readEnumField(
            "command",
            "SDORequestCommand",
            new DataReaderEnumDefault<>(
                SDORequestCommand::enumForValue, readUnsignedByte(readBuffer, 3)));

    SDORequest request =
        readSimpleField(
            "request",
            readComplex(
                () -> SDORequest.staticParse(readBuffer, (SDORequestCommand) (command)),
                readBuffer));

    readBuffer.closeContext("CANOpenSDORequest");
    // Create the instance
    return new CANOpenSDORequestBuilderImpl(command, request);
  }

  public static class CANOpenSDORequestBuilderImpl implements CANOpenPayload.CANOpenPayloadBuilder {
    private final SDORequestCommand command;
    private final SDORequest request;

    public CANOpenSDORequestBuilderImpl(SDORequestCommand command, SDORequest request) {
      this.command = command;
      this.request = request;
    }

    public CANOpenSDORequest build() {
      CANOpenSDORequest cANOpenSDORequest = new CANOpenSDORequest(command, request);
      return cANOpenSDORequest;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof CANOpenSDORequest)) {
      return false;
    }
    CANOpenSDORequest that = (CANOpenSDORequest) o;
    return (getCommand() == that.getCommand())
        && (getRequest() == that.getRequest())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getCommand(), getRequest());
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
