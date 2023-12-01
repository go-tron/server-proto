// source: control.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {missingRequire} reports error on implicit type usages.
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global =
    (typeof globalThis !== 'undefined' && globalThis) ||
    (typeof window !== 'undefined' && window) ||
    (typeof global !== 'undefined' && global) ||
    (typeof self !== 'undefined' && self) ||
    (function () { return this; }).call(null) ||
    Function('return this')();

goog.exportSymbol('proto.AssignToTable', null, global);
goog.exportSymbol('proto.AssignToTableResult', null, global);
goog.exportSymbol('proto.ForceStartTable', null, global);
goog.exportSymbol('proto.TableEnd', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.AssignToTable = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.AssignToTable.repeatedFields_, null);
};
goog.inherits(proto.AssignToTable, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.AssignToTable.displayName = 'proto.AssignToTable';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.AssignToTableResult = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.AssignToTableResult, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.AssignToTableResult.displayName = 'proto.AssignToTableResult';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.ForceStartTable = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.ForceStartTable, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.ForceStartTable.displayName = 'proto.ForceStartTable';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.TableEnd = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.TableEnd, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.TableEnd.displayName = 'proto.TableEnd';
}

/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.AssignToTable.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.AssignToTable.prototype.toObject = function(opt_includeInstance) {
  return proto.AssignToTable.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.AssignToTable} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.AssignToTable.toObject = function(includeInstance, msg) {
  var f, obj = {
    usersList: (f = jspb.Message.getRepeatedField(msg, 1)) == null ? undefined : f,
    gameconfigid: jspb.Message.getFieldWithDefault(msg, 2, 0),
    gameruleid: jspb.Message.getFieldWithDefault(msg, 3, 0),
    gameid: jspb.Message.getFieldWithDefault(msg, 4, 0),
    tableid: jspb.Message.getFieldWithDefault(msg, 5, 0),
    jointimeout: jspb.Message.getFieldWithDefault(msg, 6, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.AssignToTable}
 */
proto.AssignToTable.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.AssignToTable;
  return proto.AssignToTable.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.AssignToTable} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.AssignToTable}
 */
proto.AssignToTable.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedInt64() : [reader.readInt64()]);
      for (var i = 0; i < values.length; i++) {
        msg.addUsers(values[i]);
      }
      break;
    case 2:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setGameconfigid(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setGameruleid(value);
      break;
    case 4:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setGameid(value);
      break;
    case 5:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setTableid(value);
      break;
    case 6:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setJointimeout(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.AssignToTable.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.AssignToTable.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.AssignToTable} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.AssignToTable.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getUsersList();
  if (f.length > 0) {
    writer.writePackedInt64(
      1,
      f
    );
  }
  f = message.getGameconfigid();
  if (f !== 0) {
    writer.writeInt32(
      2,
      f
    );
  }
  f = message.getGameruleid();
  if (f !== 0) {
    writer.writeInt32(
      3,
      f
    );
  }
  f = message.getGameid();
  if (f !== 0) {
    writer.writeInt64(
      4,
      f
    );
  }
  f = message.getTableid();
  if (f !== 0) {
    writer.writeInt64(
      5,
      f
    );
  }
  f = message.getJointimeout();
  if (f !== 0) {
    writer.writeInt32(
      6,
      f
    );
  }
};


/**
 * repeated int64 Users = 1;
 * @return {!Array<number>}
 */
proto.AssignToTable.prototype.getUsersList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 1));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.AssignToTable} returns this
 */
proto.AssignToTable.prototype.setUsersList = function(value) {
  return jspb.Message.setField(this, 1, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.AssignToTable} returns this
 */
proto.AssignToTable.prototype.addUsers = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 1, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.AssignToTable} returns this
 */
proto.AssignToTable.prototype.clearUsersList = function() {
  return this.setUsersList([]);
};


/**
 * optional int32 GameConfigId = 2;
 * @return {number}
 */
proto.AssignToTable.prototype.getGameconfigid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {number} value
 * @return {!proto.AssignToTable} returns this
 */
proto.AssignToTable.prototype.setGameconfigid = function(value) {
  return jspb.Message.setProto3IntField(this, 2, value);
};


/**
 * optional int32 GameRuleId = 3;
 * @return {number}
 */
proto.AssignToTable.prototype.getGameruleid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/**
 * @param {number} value
 * @return {!proto.AssignToTable} returns this
 */
proto.AssignToTable.prototype.setGameruleid = function(value) {
  return jspb.Message.setProto3IntField(this, 3, value);
};


/**
 * optional int64 GameId = 4;
 * @return {number}
 */
proto.AssignToTable.prototype.getGameid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 4, 0));
};


/**
 * @param {number} value
 * @return {!proto.AssignToTable} returns this
 */
proto.AssignToTable.prototype.setGameid = function(value) {
  return jspb.Message.setProto3IntField(this, 4, value);
};


/**
 * optional int64 TableId = 5;
 * @return {number}
 */
proto.AssignToTable.prototype.getTableid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 5, 0));
};


/**
 * @param {number} value
 * @return {!proto.AssignToTable} returns this
 */
proto.AssignToTable.prototype.setTableid = function(value) {
  return jspb.Message.setProto3IntField(this, 5, value);
};


/**
 * optional int32 JoinTimeout = 6;
 * @return {number}
 */
proto.AssignToTable.prototype.getJointimeout = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 6, 0));
};


/**
 * @param {number} value
 * @return {!proto.AssignToTable} returns this
 */
proto.AssignToTable.prototype.setJointimeout = function(value) {
  return jspb.Message.setProto3IntField(this, 6, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.AssignToTableResult.prototype.toObject = function(opt_includeInstance) {
  return proto.AssignToTableResult.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.AssignToTableResult} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.AssignToTableResult.toObject = function(includeInstance, msg) {
  var f, obj = {
    gameid: jspb.Message.getFieldWithDefault(msg, 1, 0),
    tableid: jspb.Message.getFieldWithDefault(msg, 2, 0),
    succeed: jspb.Message.getBooleanFieldWithDefault(msg, 3, false)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.AssignToTableResult}
 */
proto.AssignToTableResult.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.AssignToTableResult;
  return proto.AssignToTableResult.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.AssignToTableResult} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.AssignToTableResult}
 */
proto.AssignToTableResult.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setGameid(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setTableid(value);
      break;
    case 3:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setSucceed(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.AssignToTableResult.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.AssignToTableResult.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.AssignToTableResult} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.AssignToTableResult.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getGameid();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = message.getTableid();
  if (f !== 0) {
    writer.writeInt64(
      2,
      f
    );
  }
  f = message.getSucceed();
  if (f) {
    writer.writeBool(
      3,
      f
    );
  }
};


/**
 * optional int64 GameId = 1;
 * @return {number}
 */
proto.AssignToTableResult.prototype.getGameid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {number} value
 * @return {!proto.AssignToTableResult} returns this
 */
proto.AssignToTableResult.prototype.setGameid = function(value) {
  return jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional int64 TableId = 2;
 * @return {number}
 */
proto.AssignToTableResult.prototype.getTableid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {number} value
 * @return {!proto.AssignToTableResult} returns this
 */
proto.AssignToTableResult.prototype.setTableid = function(value) {
  return jspb.Message.setProto3IntField(this, 2, value);
};


/**
 * optional bool Succeed = 3;
 * @return {boolean}
 */
proto.AssignToTableResult.prototype.getSucceed = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 3, false));
};


/**
 * @param {boolean} value
 * @return {!proto.AssignToTableResult} returns this
 */
proto.AssignToTableResult.prototype.setSucceed = function(value) {
  return jspb.Message.setProto3BooleanField(this, 3, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.ForceStartTable.prototype.toObject = function(opt_includeInstance) {
  return proto.ForceStartTable.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.ForceStartTable} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ForceStartTable.toObject = function(includeInstance, msg) {
  var f, obj = {
    gameid: jspb.Message.getFieldWithDefault(msg, 1, 0),
    tableid: jspb.Message.getFieldWithDefault(msg, 2, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.ForceStartTable}
 */
proto.ForceStartTable.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.ForceStartTable;
  return proto.ForceStartTable.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.ForceStartTable} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.ForceStartTable}
 */
proto.ForceStartTable.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setGameid(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setTableid(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.ForceStartTable.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.ForceStartTable.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.ForceStartTable} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ForceStartTable.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getGameid();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = message.getTableid();
  if (f !== 0) {
    writer.writeInt64(
      2,
      f
    );
  }
};


/**
 * optional int64 GameId = 1;
 * @return {number}
 */
proto.ForceStartTable.prototype.getGameid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {number} value
 * @return {!proto.ForceStartTable} returns this
 */
proto.ForceStartTable.prototype.setGameid = function(value) {
  return jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional int64 TableId = 2;
 * @return {number}
 */
proto.ForceStartTable.prototype.getTableid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {number} value
 * @return {!proto.ForceStartTable} returns this
 */
proto.ForceStartTable.prototype.setTableid = function(value) {
  return jspb.Message.setProto3IntField(this, 2, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.TableEnd.prototype.toObject = function(opt_includeInstance) {
  return proto.TableEnd.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.TableEnd} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.TableEnd.toObject = function(includeInstance, msg) {
  var f, obj = {
    gameid: jspb.Message.getFieldWithDefault(msg, 1, 0),
    tableid: jspb.Message.getFieldWithDefault(msg, 2, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.TableEnd}
 */
proto.TableEnd.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.TableEnd;
  return proto.TableEnd.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.TableEnd} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.TableEnd}
 */
proto.TableEnd.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setGameid(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setTableid(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.TableEnd.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.TableEnd.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.TableEnd} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.TableEnd.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getGameid();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = message.getTableid();
  if (f !== 0) {
    writer.writeInt64(
      2,
      f
    );
  }
};


/**
 * optional int64 GameId = 1;
 * @return {number}
 */
proto.TableEnd.prototype.getGameid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {number} value
 * @return {!proto.TableEnd} returns this
 */
proto.TableEnd.prototype.setGameid = function(value) {
  return jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional int64 TableId = 2;
 * @return {number}
 */
proto.TableEnd.prototype.getTableid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {number} value
 * @return {!proto.TableEnd} returns this
 */
proto.TableEnd.prototype.setTableid = function(value) {
  return jspb.Message.setProto3IntField(this, 2, value);
};


goog.object.extend(exports, proto);
