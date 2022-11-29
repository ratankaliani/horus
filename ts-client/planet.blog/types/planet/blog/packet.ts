/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "planet.blog";

export interface BlogPacketData {
  noData:
    | NoData
    | undefined;
  /** this line is used by starport scaffolding # ibc/packet/proto/field */
  ibcHorusActionPacket:
    | IbcHorusActionPacketData
    | undefined;
  /** this line is used by starport scaffolding # ibc/packet/proto/field/number */
  ibcPostPacket: IbcPostPacketData | undefined;
}

export interface NoData {
}

/** IbcPostPacketData defines a struct for the packet payload */
export interface IbcPostPacketData {
  title: string;
  content: string;
  creator: string;
}

/** IbcPostPacketAck defines a struct for the packet acknowledgment */
export interface IbcPostPacketAck {
  postID: string;
}

/** IbcHorusActionPacketData defines a struct for the packet payload */
export interface IbcHorusActionPacketData {
  title: string;
  action: string;
  creator: string;
}

/** IbcHorusActionPacketAck defines a struct for the packet acknowledgment */
export interface IbcHorusActionPacketAck {
  actionID: string;
}

function createBaseBlogPacketData(): BlogPacketData {
  return { noData: undefined, ibcHorusActionPacket: undefined, ibcPostPacket: undefined };
}

export const BlogPacketData = {
  encode(message: BlogPacketData, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.noData !== undefined) {
      NoData.encode(message.noData, writer.uint32(10).fork()).ldelim();
    }
    if (message.ibcHorusActionPacket !== undefined) {
      IbcHorusActionPacketData.encode(message.ibcHorusActionPacket, writer.uint32(26).fork()).ldelim();
    }
    if (message.ibcPostPacket !== undefined) {
      IbcPostPacketData.encode(message.ibcPostPacket, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): BlogPacketData {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseBlogPacketData();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.noData = NoData.decode(reader, reader.uint32());
          break;
        case 3:
          message.ibcHorusActionPacket = IbcHorusActionPacketData.decode(reader, reader.uint32());
          break;
        case 2:
          message.ibcPostPacket = IbcPostPacketData.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): BlogPacketData {
    return {
      noData: isSet(object.noData) ? NoData.fromJSON(object.noData) : undefined,
      ibcHorusActionPacket: isSet(object.ibcHorusActionPacket)
        ? IbcHorusActionPacketData.fromJSON(object.ibcHorusActionPacket)
        : undefined,
      ibcPostPacket: isSet(object.ibcPostPacket) ? IbcPostPacketData.fromJSON(object.ibcPostPacket) : undefined,
    };
  },

  toJSON(message: BlogPacketData): unknown {
    const obj: any = {};
    message.noData !== undefined && (obj.noData = message.noData ? NoData.toJSON(message.noData) : undefined);
    message.ibcHorusActionPacket !== undefined && (obj.ibcHorusActionPacket = message.ibcHorusActionPacket
      ? IbcHorusActionPacketData.toJSON(message.ibcHorusActionPacket)
      : undefined);
    message.ibcPostPacket !== undefined
      && (obj.ibcPostPacket = message.ibcPostPacket ? IbcPostPacketData.toJSON(message.ibcPostPacket) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<BlogPacketData>, I>>(object: I): BlogPacketData {
    const message = createBaseBlogPacketData();
    message.noData = (object.noData !== undefined && object.noData !== null)
      ? NoData.fromPartial(object.noData)
      : undefined;
    message.ibcHorusActionPacket = (object.ibcHorusActionPacket !== undefined && object.ibcHorusActionPacket !== null)
      ? IbcHorusActionPacketData.fromPartial(object.ibcHorusActionPacket)
      : undefined;
    message.ibcPostPacket = (object.ibcPostPacket !== undefined && object.ibcPostPacket !== null)
      ? IbcPostPacketData.fromPartial(object.ibcPostPacket)
      : undefined;
    return message;
  },
};

function createBaseNoData(): NoData {
  return {};
}

export const NoData = {
  encode(_: NoData, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): NoData {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseNoData();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): NoData {
    return {};
  },

  toJSON(_: NoData): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<NoData>, I>>(_: I): NoData {
    const message = createBaseNoData();
    return message;
  },
};

function createBaseIbcPostPacketData(): IbcPostPacketData {
  return { title: "", content: "", creator: "" };
}

export const IbcPostPacketData = {
  encode(message: IbcPostPacketData, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.title !== "") {
      writer.uint32(10).string(message.title);
    }
    if (message.content !== "") {
      writer.uint32(18).string(message.content);
    }
    if (message.creator !== "") {
      writer.uint32(26).string(message.creator);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): IbcPostPacketData {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseIbcPostPacketData();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.title = reader.string();
          break;
        case 2:
          message.content = reader.string();
          break;
        case 3:
          message.creator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): IbcPostPacketData {
    return {
      title: isSet(object.title) ? String(object.title) : "",
      content: isSet(object.content) ? String(object.content) : "",
      creator: isSet(object.creator) ? String(object.creator) : "",
    };
  },

  toJSON(message: IbcPostPacketData): unknown {
    const obj: any = {};
    message.title !== undefined && (obj.title = message.title);
    message.content !== undefined && (obj.content = message.content);
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<IbcPostPacketData>, I>>(object: I): IbcPostPacketData {
    const message = createBaseIbcPostPacketData();
    message.title = object.title ?? "";
    message.content = object.content ?? "";
    message.creator = object.creator ?? "";
    return message;
  },
};

function createBaseIbcPostPacketAck(): IbcPostPacketAck {
  return { postID: "" };
}

export const IbcPostPacketAck = {
  encode(message: IbcPostPacketAck, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.postID !== "") {
      writer.uint32(10).string(message.postID);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): IbcPostPacketAck {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseIbcPostPacketAck();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.postID = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): IbcPostPacketAck {
    return { postID: isSet(object.postID) ? String(object.postID) : "" };
  },

  toJSON(message: IbcPostPacketAck): unknown {
    const obj: any = {};
    message.postID !== undefined && (obj.postID = message.postID);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<IbcPostPacketAck>, I>>(object: I): IbcPostPacketAck {
    const message = createBaseIbcPostPacketAck();
    message.postID = object.postID ?? "";
    return message;
  },
};

function createBaseIbcHorusActionPacketData(): IbcHorusActionPacketData {
  return { title: "", action: "", creator: "" };
}

export const IbcHorusActionPacketData = {
  encode(message: IbcHorusActionPacketData, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.title !== "") {
      writer.uint32(10).string(message.title);
    }
    if (message.action !== "") {
      writer.uint32(18).string(message.action);
    }
    if (message.creator !== "") {
      writer.uint32(26).string(message.creator);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): IbcHorusActionPacketData {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseIbcHorusActionPacketData();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.title = reader.string();
          break;
        case 2:
          message.action = reader.string();
          break;
        case 3:
          message.creator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): IbcHorusActionPacketData {
    return {
      title: isSet(object.title) ? String(object.title) : "",
      action: isSet(object.action) ? String(object.action) : "",
      creator: isSet(object.creator) ? String(object.creator) : "",
    };
  },

  toJSON(message: IbcHorusActionPacketData): unknown {
    const obj: any = {};
    message.title !== undefined && (obj.title = message.title);
    message.action !== undefined && (obj.action = message.action);
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<IbcHorusActionPacketData>, I>>(object: I): IbcHorusActionPacketData {
    const message = createBaseIbcHorusActionPacketData();
    message.title = object.title ?? "";
    message.action = object.action ?? "";
    message.creator = object.creator ?? "";
    return message;
  },
};

function createBaseIbcHorusActionPacketAck(): IbcHorusActionPacketAck {
  return { actionID: "" };
}

export const IbcHorusActionPacketAck = {
  encode(message: IbcHorusActionPacketAck, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.actionID !== "") {
      writer.uint32(10).string(message.actionID);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): IbcHorusActionPacketAck {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseIbcHorusActionPacketAck();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.actionID = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): IbcHorusActionPacketAck {
    return { actionID: isSet(object.actionID) ? String(object.actionID) : "" };
  },

  toJSON(message: IbcHorusActionPacketAck): unknown {
    const obj: any = {};
    message.actionID !== undefined && (obj.actionID = message.actionID);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<IbcHorusActionPacketAck>, I>>(object: I): IbcHorusActionPacketAck {
    const message = createBaseIbcHorusActionPacketAck();
    message.actionID = object.actionID ?? "";
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
