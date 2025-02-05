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

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

// Code generated by code-generation. DO NOT EDIT.

public enum AccessControlCommandTypeContainer {
  AccessControlCommandCloseAccessPoint(
      (short) 0x02,
      (byte) 2,
      AccessControlCommandType.CLOSE_ACCESS_POINT,
      AccessControlCategory.SYSTEM_REQUEST),
  AccessControlCommandLockAccessPoint(
      (short) 0x0A,
      (byte) 2,
      AccessControlCommandType.LOCK_ACCESS_POINT,
      AccessControlCategory.SYSTEM_REQUEST),
  AccessControlCommandAccessPointLeftOpen(
      (short) 0x12,
      (byte) 2,
      AccessControlCommandType.ACCESS_POINT_LEFT_OPEN,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandAccessPointForcedOpen(
      (short) 0x1A,
      (byte) 2,
      AccessControlCommandType.ACCESS_POINT_FORCED_OPEN,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandAccessPointClosed(
      (short) 0x22,
      (byte) 2,
      AccessControlCommandType.ACCESS_POINT_CLOSED,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandRequestToExit(
      (short) 0x32,
      (byte) 2,
      AccessControlCommandType.REQUEST_TO_EXIT,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_0Bytes(
      (short) 0xA0,
      (byte) 0,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_1Bytes(
      (short) 0xA1,
      (byte) 1,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_2Bytes(
      (short) 0xA2,
      (byte) 2,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_3Bytes(
      (short) 0xA3,
      (byte) 3,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_4Bytes(
      (short) 0xA4,
      (byte) 4,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_5Bytes(
      (short) 0xA5,
      (byte) 5,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_6Bytes(
      (short) 0xA6,
      (byte) 6,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_7Bytes(
      (short) 0xA7,
      (byte) 7,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_8Bytes(
      (short) 0xA8,
      (byte) 8,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_9Bytes(
      (short) 0xA9,
      (byte) 9,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_10Bytes(
      (short) 0xAA,
      (byte) 10,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_11Bytes(
      (short) 0xAB,
      (byte) 11,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_12Bytes(
      (short) 0xAC,
      (byte) 12,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_13Bytes(
      (short) 0xAD,
      (byte) 13,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_14Bytes(
      (short) 0xAE,
      (byte) 14,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_15Bytes(
      (short) 0xAF,
      (byte) 15,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_16Bytes(
      (short) 0xB0,
      (byte) 16,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_17Bytes(
      (short) 0xB1,
      (byte) 17,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_18Bytes(
      (short) 0xB2,
      (byte) 18,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_19Bytes(
      (short) 0xB3,
      (byte) 19,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_20Bytes(
      (short) 0xB4,
      (byte) 20,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_21Bytes(
      (short) 0xB5,
      (byte) 21,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_22Bytes(
      (short) 0xB6,
      (byte) 22,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_23Bytes(
      (short) 0xB7,
      (byte) 23,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_24Bytes(
      (short) 0xB8,
      (byte) 24,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_25Bytes(
      (short) 0xB9,
      (byte) 25,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_26Bytes(
      (short) 0xBA,
      (byte) 26,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_27Bytes(
      (short) 0xBB,
      (byte) 27,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_28Bytes(
      (short) 0xBC,
      (byte) 28,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_29Bytes(
      (short) 0xBD,
      (byte) 29,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_30Bytes(
      (short) 0xBE,
      (byte) 30,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandValidAccessRequest_31Bytes(
      (short) 0xBF,
      (byte) 31,
      AccessControlCommandType.VALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_0Bytes(
      (short) 0xC0,
      (byte) 0,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_1Bytes(
      (short) 0xC1,
      (byte) 1,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_2Bytes(
      (short) 0xC2,
      (byte) 2,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_3Bytes(
      (short) 0xC3,
      (byte) 3,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_4Bytes(
      (short) 0xC4,
      (byte) 4,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_5Bytes(
      (short) 0xC5,
      (byte) 5,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_6Bytes(
      (short) 0xC6,
      (byte) 6,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_7Bytes(
      (short) 0xC7,
      (byte) 7,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_8Bytes(
      (short) 0xC8,
      (byte) 8,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_9Bytes(
      (short) 0xC9,
      (byte) 9,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_10Bytes(
      (short) 0xCA,
      (byte) 10,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_11Bytes(
      (short) 0xCB,
      (byte) 11,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_12Bytes(
      (short) 0xCC,
      (byte) 12,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_13Bytes(
      (short) 0xCD,
      (byte) 13,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_14Bytes(
      (short) 0xCE,
      (byte) 14,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_15Bytes(
      (short) 0xCF,
      (byte) 15,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_16Bytes(
      (short) 0xD0,
      (byte) 16,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_17Bytes(
      (short) 0xD1,
      (byte) 17,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_18Bytes(
      (short) 0xD2,
      (byte) 18,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_19Bytes(
      (short) 0xD3,
      (byte) 19,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_20Bytes(
      (short) 0xD4,
      (byte) 20,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_21Bytes(
      (short) 0xD5,
      (byte) 21,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_22Bytes(
      (short) 0xD6,
      (byte) 22,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_23Bytes(
      (short) 0xD7,
      (byte) 23,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_24Bytes(
      (short) 0xD8,
      (byte) 24,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_25Bytes(
      (short) 0xD9,
      (byte) 25,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_26Bytes(
      (short) 0xDA,
      (byte) 26,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_27Bytes(
      (short) 0xDB,
      (byte) 27,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_28Bytes(
      (short) 0xDC,
      (byte) 28,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_29Bytes(
      (short) 0xDD,
      (byte) 29,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_30Bytes(
      (short) 0xDE,
      (byte) 30,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY),
  AccessControlCommandInvalidAccessRequest_31Bytes(
      (short) 0xDF,
      (byte) 31,
      AccessControlCommandType.INVALID_ACCESS,
      AccessControlCategory.SYSTEM_ACTIVITY);
  private static final Map<Short, AccessControlCommandTypeContainer> map;

  static {
    map = new HashMap<>();
    for (AccessControlCommandTypeContainer value : AccessControlCommandTypeContainer.values()) {
      map.put((short) value.getValue(), value);
    }
  }

  private final short value;
  private final byte numBytes;
  private final AccessControlCommandType commandType;
  private final AccessControlCategory category;

  AccessControlCommandTypeContainer(
      short value,
      byte numBytes,
      AccessControlCommandType commandType,
      AccessControlCategory category) {
    this.value = value;
    this.numBytes = numBytes;
    this.commandType = commandType;
    this.category = category;
  }

  public short getValue() {
    return value;
  }

  public byte getNumBytes() {
    return numBytes;
  }

  public static AccessControlCommandTypeContainer firstEnumForFieldNumBytes(byte fieldValue) {
    for (AccessControlCommandTypeContainer _val : AccessControlCommandTypeContainer.values()) {
      if (_val.getNumBytes() == fieldValue) {
        return _val;
      }
    }
    return null;
  }

  public static List<AccessControlCommandTypeContainer> enumsForFieldNumBytes(byte fieldValue) {
    List<AccessControlCommandTypeContainer> _values = new ArrayList<>();
    for (AccessControlCommandTypeContainer _val : AccessControlCommandTypeContainer.values()) {
      if (_val.getNumBytes() == fieldValue) {
        _values.add(_val);
      }
    }
    return _values;
  }

  public AccessControlCommandType getCommandType() {
    return commandType;
  }

  public static AccessControlCommandTypeContainer firstEnumForFieldCommandType(
      AccessControlCommandType fieldValue) {
    for (AccessControlCommandTypeContainer _val : AccessControlCommandTypeContainer.values()) {
      if (_val.getCommandType() == fieldValue) {
        return _val;
      }
    }
    return null;
  }

  public static List<AccessControlCommandTypeContainer> enumsForFieldCommandType(
      AccessControlCommandType fieldValue) {
    List<AccessControlCommandTypeContainer> _values = new ArrayList<>();
    for (AccessControlCommandTypeContainer _val : AccessControlCommandTypeContainer.values()) {
      if (_val.getCommandType() == fieldValue) {
        _values.add(_val);
      }
    }
    return _values;
  }

  public AccessControlCategory getCategory() {
    return category;
  }

  public static AccessControlCommandTypeContainer firstEnumForFieldCategory(
      AccessControlCategory fieldValue) {
    for (AccessControlCommandTypeContainer _val : AccessControlCommandTypeContainer.values()) {
      if (_val.getCategory() == fieldValue) {
        return _val;
      }
    }
    return null;
  }

  public static List<AccessControlCommandTypeContainer> enumsForFieldCategory(
      AccessControlCategory fieldValue) {
    List<AccessControlCommandTypeContainer> _values = new ArrayList<>();
    for (AccessControlCommandTypeContainer _val : AccessControlCommandTypeContainer.values()) {
      if (_val.getCategory() == fieldValue) {
        _values.add(_val);
      }
    }
    return _values;
  }

  public static AccessControlCommandTypeContainer enumForValue(short value) {
    return map.get(value);
  }

  public static Boolean isDefined(short value) {
    return map.containsKey(value);
  }
}
