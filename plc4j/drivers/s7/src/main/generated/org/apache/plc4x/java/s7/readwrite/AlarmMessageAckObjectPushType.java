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
package org.apache.plc4x.java.s7.readwrite;

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

public class AlarmMessageAckObjectPushType implements Message {

  // Constant values.
  public static final Short VARIABLESPEC = 0x12;

  // Properties.
  protected final short lengthSpec;
  protected final SyntaxIdType syntaxId;
  protected final short numberOfValues;
  protected final long eventId;
  protected final State ackStateGoing;
  protected final State ackStateComing;

  public AlarmMessageAckObjectPushType(
      short lengthSpec,
      SyntaxIdType syntaxId,
      short numberOfValues,
      long eventId,
      State ackStateGoing,
      State ackStateComing) {
    super();
    this.lengthSpec = lengthSpec;
    this.syntaxId = syntaxId;
    this.numberOfValues = numberOfValues;
    this.eventId = eventId;
    this.ackStateGoing = ackStateGoing;
    this.ackStateComing = ackStateComing;
  }

  public short getLengthSpec() {
    return lengthSpec;
  }

  public SyntaxIdType getSyntaxId() {
    return syntaxId;
  }

  public short getNumberOfValues() {
    return numberOfValues;
  }

  public long getEventId() {
    return eventId;
  }

  public State getAckStateGoing() {
    return ackStateGoing;
  }

  public State getAckStateComing() {
    return ackStateComing;
  }

  public short getVariableSpec() {
    return VARIABLESPEC;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("AlarmMessageAckObjectPushType");

    // Const Field (variableSpec)
    writeConstField("variableSpec", VARIABLESPEC, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (lengthSpec)
    writeSimpleField("lengthSpec", lengthSpec, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (syntaxId)
    writeSimpleEnumField(
        "syntaxId",
        "SyntaxIdType",
        syntaxId,
        new DataWriterEnumDefault<>(
            SyntaxIdType::getValue, SyntaxIdType::name, writeUnsignedShort(writeBuffer, 8)));

    // Simple Field (numberOfValues)
    writeSimpleField("numberOfValues", numberOfValues, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (eventId)
    writeSimpleField("eventId", eventId, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (ackStateGoing)
    writeSimpleField("ackStateGoing", ackStateGoing, writeComplex(writeBuffer));

    // Simple Field (ackStateComing)
    writeSimpleField("ackStateComing", ackStateComing, writeComplex(writeBuffer));

    writeBuffer.popContext("AlarmMessageAckObjectPushType");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    AlarmMessageAckObjectPushType _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Const Field (variableSpec)
    lengthInBits += 8;

    // Simple field (lengthSpec)
    lengthInBits += 8;

    // Simple field (syntaxId)
    lengthInBits += 8;

    // Simple field (numberOfValues)
    lengthInBits += 8;

    // Simple field (eventId)
    lengthInBits += 32;

    // Simple field (ackStateGoing)
    lengthInBits += ackStateGoing.getLengthInBits();

    // Simple field (ackStateComing)
    lengthInBits += ackStateComing.getLengthInBits();

    return lengthInBits;
  }

  public static AlarmMessageAckObjectPushType staticParse(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("AlarmMessageAckObjectPushType");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    short variableSpec =
        readConstField(
            "variableSpec",
            readUnsignedShort(readBuffer, 8),
            AlarmMessageAckObjectPushType.VARIABLESPEC);

    short lengthSpec = readSimpleField("lengthSpec", readUnsignedShort(readBuffer, 8));

    SyntaxIdType syntaxId =
        readEnumField(
            "syntaxId",
            "SyntaxIdType",
            new DataReaderEnumDefault<>(
                SyntaxIdType::enumForValue, readUnsignedShort(readBuffer, 8)));

    short numberOfValues = readSimpleField("numberOfValues", readUnsignedShort(readBuffer, 8));

    long eventId = readSimpleField("eventId", readUnsignedLong(readBuffer, 32));

    State ackStateGoing =
        readSimpleField(
            "ackStateGoing", readComplex(() -> State.staticParse(readBuffer), readBuffer));

    State ackStateComing =
        readSimpleField(
            "ackStateComing", readComplex(() -> State.staticParse(readBuffer), readBuffer));

    readBuffer.closeContext("AlarmMessageAckObjectPushType");
    // Create the instance
    AlarmMessageAckObjectPushType _alarmMessageAckObjectPushType;
    _alarmMessageAckObjectPushType =
        new AlarmMessageAckObjectPushType(
            lengthSpec, syntaxId, numberOfValues, eventId, ackStateGoing, ackStateComing);
    return _alarmMessageAckObjectPushType;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof AlarmMessageAckObjectPushType)) {
      return false;
    }
    AlarmMessageAckObjectPushType that = (AlarmMessageAckObjectPushType) o;
    return (getLengthSpec() == that.getLengthSpec())
        && (getSyntaxId() == that.getSyntaxId())
        && (getNumberOfValues() == that.getNumberOfValues())
        && (getEventId() == that.getEventId())
        && (getAckStateGoing() == that.getAckStateGoing())
        && (getAckStateComing() == that.getAckStateComing())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        getLengthSpec(),
        getSyntaxId(),
        getNumberOfValues(),
        getEventId(),
        getAckStateGoing(),
        getAckStateComing());
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
