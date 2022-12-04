/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "planet.blog";

export interface BridgeStatus {
  isShutdown: string;
}

function createBaseBridgeStatus(): BridgeStatus {
  return { isShutdown: "" };
}

export const BridgeStatus = {
  encode(message: BridgeStatus, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.isShutdown !== "") {
      writer.uint32(10).string(message.isShutdown);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): BridgeStatus {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseBridgeStatus();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.isShutdown = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): BridgeStatus {
    return { isShutdown: isSet(object.isShutdown) ? String(object.isShutdown) : "" };
  },

  toJSON(message: BridgeStatus): unknown {
    const obj: any = {};
    message.isShutdown !== undefined && (obj.isShutdown = message.isShutdown);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<BridgeStatus>, I>>(object: I): BridgeStatus {
    const message = createBaseBridgeStatus();
    message.isShutdown = object.isShutdown ?? "";
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
