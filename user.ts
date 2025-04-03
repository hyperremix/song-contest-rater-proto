// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.7.0
//   protoc               v5.29.3
// source: user.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";
import { Timestamp } from "./google/protobuf/timestamp";

export const protobufPackage = "songcontestrater";

export interface ListCompetitionUsersRequest {
  competition_id: string;
}

export interface GetUserRequest {
  id: string;
}

export interface CreateUserRequest {
  email: string;
  firstname: string;
  lastname: string;
  image_url: string;
}

export interface UpdateUserRequest {
  id: string;
  firstname: string;
  lastname: string;
  image_url: string;
}

export interface DeleteUserRequest {
  id: string;
}

export interface UserResponse {
  id: string;
  email: string;
  firstname: string;
  lastname: string;
  image_url: string;
  created_at: Timestamp | undefined;
  updated_at: Timestamp | undefined;
}

export interface ListUsersResponse {
  users: UserResponse[];
}

export interface GetPresignedURLRequest {
  file_name: string;
  content_type: string;
}

export interface GetPresignedURLResponse {
  presigned_url: string;
  image_url: string;
}

function createBaseListCompetitionUsersRequest(): ListCompetitionUsersRequest {
  return { competition_id: "" };
}

export const ListCompetitionUsersRequest: MessageFns<ListCompetitionUsersRequest> = {
  encode(message: ListCompetitionUsersRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.competition_id !== "") {
      writer.uint32(10).string(message.competition_id);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): ListCompetitionUsersRequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListCompetitionUsersRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.competition_id = reader.string();
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ListCompetitionUsersRequest {
    return { competition_id: isSet(object.competition_id) ? globalThis.String(object.competition_id) : "" };
  },

  toJSON(message: ListCompetitionUsersRequest): unknown {
    const obj: any = {};
    if (message.competition_id !== "") {
      obj.competition_id = message.competition_id;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ListCompetitionUsersRequest>, I>>(base?: I): ListCompetitionUsersRequest {
    return ListCompetitionUsersRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<ListCompetitionUsersRequest>, I>>(object: I): ListCompetitionUsersRequest {
    const message = createBaseListCompetitionUsersRequest();
    message.competition_id = object.competition_id ?? "";
    return message;
  },
};

function createBaseGetUserRequest(): GetUserRequest {
  return { id: "" };
}

export const GetUserRequest: MessageFns<GetUserRequest> = {
  encode(message: GetUserRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): GetUserRequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetUserRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.id = reader.string();
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetUserRequest {
    return { id: isSet(object.id) ? globalThis.String(object.id) : "" };
  },

  toJSON(message: GetUserRequest): unknown {
    const obj: any = {};
    if (message.id !== "") {
      obj.id = message.id;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetUserRequest>, I>>(base?: I): GetUserRequest {
    return GetUserRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetUserRequest>, I>>(object: I): GetUserRequest {
    const message = createBaseGetUserRequest();
    message.id = object.id ?? "";
    return message;
  },
};

function createBaseCreateUserRequest(): CreateUserRequest {
  return { email: "", firstname: "", lastname: "", image_url: "" };
}

export const CreateUserRequest: MessageFns<CreateUserRequest> = {
  encode(message: CreateUserRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.email !== "") {
      writer.uint32(10).string(message.email);
    }
    if (message.firstname !== "") {
      writer.uint32(18).string(message.firstname);
    }
    if (message.lastname !== "") {
      writer.uint32(26).string(message.lastname);
    }
    if (message.image_url !== "") {
      writer.uint32(34).string(message.image_url);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): CreateUserRequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateUserRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.email = reader.string();
          continue;
        }
        case 2: {
          if (tag !== 18) {
            break;
          }

          message.firstname = reader.string();
          continue;
        }
        case 3: {
          if (tag !== 26) {
            break;
          }

          message.lastname = reader.string();
          continue;
        }
        case 4: {
          if (tag !== 34) {
            break;
          }

          message.image_url = reader.string();
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreateUserRequest {
    return {
      email: isSet(object.email) ? globalThis.String(object.email) : "",
      firstname: isSet(object.firstname) ? globalThis.String(object.firstname) : "",
      lastname: isSet(object.lastname) ? globalThis.String(object.lastname) : "",
      image_url: isSet(object.image_url) ? globalThis.String(object.image_url) : "",
    };
  },

  toJSON(message: CreateUserRequest): unknown {
    const obj: any = {};
    if (message.email !== "") {
      obj.email = message.email;
    }
    if (message.firstname !== "") {
      obj.firstname = message.firstname;
    }
    if (message.lastname !== "") {
      obj.lastname = message.lastname;
    }
    if (message.image_url !== "") {
      obj.image_url = message.image_url;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateUserRequest>, I>>(base?: I): CreateUserRequest {
    return CreateUserRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<CreateUserRequest>, I>>(object: I): CreateUserRequest {
    const message = createBaseCreateUserRequest();
    message.email = object.email ?? "";
    message.firstname = object.firstname ?? "";
    message.lastname = object.lastname ?? "";
    message.image_url = object.image_url ?? "";
    return message;
  },
};

function createBaseUpdateUserRequest(): UpdateUserRequest {
  return { id: "", firstname: "", lastname: "", image_url: "" };
}

export const UpdateUserRequest: MessageFns<UpdateUserRequest> = {
  encode(message: UpdateUserRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.firstname !== "") {
      writer.uint32(18).string(message.firstname);
    }
    if (message.lastname !== "") {
      writer.uint32(26).string(message.lastname);
    }
    if (message.image_url !== "") {
      writer.uint32(34).string(message.image_url);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): UpdateUserRequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateUserRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.id = reader.string();
          continue;
        }
        case 2: {
          if (tag !== 18) {
            break;
          }

          message.firstname = reader.string();
          continue;
        }
        case 3: {
          if (tag !== 26) {
            break;
          }

          message.lastname = reader.string();
          continue;
        }
        case 4: {
          if (tag !== 34) {
            break;
          }

          message.image_url = reader.string();
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): UpdateUserRequest {
    return {
      id: isSet(object.id) ? globalThis.String(object.id) : "",
      firstname: isSet(object.firstname) ? globalThis.String(object.firstname) : "",
      lastname: isSet(object.lastname) ? globalThis.String(object.lastname) : "",
      image_url: isSet(object.image_url) ? globalThis.String(object.image_url) : "",
    };
  },

  toJSON(message: UpdateUserRequest): unknown {
    const obj: any = {};
    if (message.id !== "") {
      obj.id = message.id;
    }
    if (message.firstname !== "") {
      obj.firstname = message.firstname;
    }
    if (message.lastname !== "") {
      obj.lastname = message.lastname;
    }
    if (message.image_url !== "") {
      obj.image_url = message.image_url;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<UpdateUserRequest>, I>>(base?: I): UpdateUserRequest {
    return UpdateUserRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<UpdateUserRequest>, I>>(object: I): UpdateUserRequest {
    const message = createBaseUpdateUserRequest();
    message.id = object.id ?? "";
    message.firstname = object.firstname ?? "";
    message.lastname = object.lastname ?? "";
    message.image_url = object.image_url ?? "";
    return message;
  },
};

function createBaseDeleteUserRequest(): DeleteUserRequest {
  return { id: "" };
}

export const DeleteUserRequest: MessageFns<DeleteUserRequest> = {
  encode(message: DeleteUserRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): DeleteUserRequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteUserRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.id = reader.string();
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DeleteUserRequest {
    return { id: isSet(object.id) ? globalThis.String(object.id) : "" };
  },

  toJSON(message: DeleteUserRequest): unknown {
    const obj: any = {};
    if (message.id !== "") {
      obj.id = message.id;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<DeleteUserRequest>, I>>(base?: I): DeleteUserRequest {
    return DeleteUserRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<DeleteUserRequest>, I>>(object: I): DeleteUserRequest {
    const message = createBaseDeleteUserRequest();
    message.id = object.id ?? "";
    return message;
  },
};

function createBaseUserResponse(): UserResponse {
  return {
    id: "",
    email: "",
    firstname: "",
    lastname: "",
    image_url: "",
    created_at: undefined,
    updated_at: undefined,
  };
}

export const UserResponse: MessageFns<UserResponse> = {
  encode(message: UserResponse, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.email !== "") {
      writer.uint32(26).string(message.email);
    }
    if (message.firstname !== "") {
      writer.uint32(34).string(message.firstname);
    }
    if (message.lastname !== "") {
      writer.uint32(42).string(message.lastname);
    }
    if (message.image_url !== "") {
      writer.uint32(50).string(message.image_url);
    }
    if (message.created_at !== undefined) {
      Timestamp.encode(message.created_at, writer.uint32(58).fork()).join();
    }
    if (message.updated_at !== undefined) {
      Timestamp.encode(message.updated_at, writer.uint32(66).fork()).join();
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): UserResponse {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUserResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.id = reader.string();
          continue;
        }
        case 3: {
          if (tag !== 26) {
            break;
          }

          message.email = reader.string();
          continue;
        }
        case 4: {
          if (tag !== 34) {
            break;
          }

          message.firstname = reader.string();
          continue;
        }
        case 5: {
          if (tag !== 42) {
            break;
          }

          message.lastname = reader.string();
          continue;
        }
        case 6: {
          if (tag !== 50) {
            break;
          }

          message.image_url = reader.string();
          continue;
        }
        case 7: {
          if (tag !== 58) {
            break;
          }

          message.created_at = Timestamp.decode(reader, reader.uint32());
          continue;
        }
        case 8: {
          if (tag !== 66) {
            break;
          }

          message.updated_at = Timestamp.decode(reader, reader.uint32());
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): UserResponse {
    return {
      id: isSet(object.id) ? globalThis.String(object.id) : "",
      email: isSet(object.email) ? globalThis.String(object.email) : "",
      firstname: isSet(object.firstname) ? globalThis.String(object.firstname) : "",
      lastname: isSet(object.lastname) ? globalThis.String(object.lastname) : "",
      image_url: isSet(object.image_url) ? globalThis.String(object.image_url) : "",
      created_at: isSet(object.created_at) ? fromJsonTimestamp(object.created_at) : undefined,
      updated_at: isSet(object.updated_at) ? fromJsonTimestamp(object.updated_at) : undefined,
    };
  },

  toJSON(message: UserResponse): unknown {
    const obj: any = {};
    if (message.id !== "") {
      obj.id = message.id;
    }
    if (message.email !== "") {
      obj.email = message.email;
    }
    if (message.firstname !== "") {
      obj.firstname = message.firstname;
    }
    if (message.lastname !== "") {
      obj.lastname = message.lastname;
    }
    if (message.image_url !== "") {
      obj.image_url = message.image_url;
    }
    if (message.created_at !== undefined) {
      obj.created_at = fromTimestamp(message.created_at).toISOString();
    }
    if (message.updated_at !== undefined) {
      obj.updated_at = fromTimestamp(message.updated_at).toISOString();
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<UserResponse>, I>>(base?: I): UserResponse {
    return UserResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<UserResponse>, I>>(object: I): UserResponse {
    const message = createBaseUserResponse();
    message.id = object.id ?? "";
    message.email = object.email ?? "";
    message.firstname = object.firstname ?? "";
    message.lastname = object.lastname ?? "";
    message.image_url = object.image_url ?? "";
    message.created_at = (object.created_at !== undefined && object.created_at !== null)
      ? Timestamp.fromPartial(object.created_at)
      : undefined;
    message.updated_at = (object.updated_at !== undefined && object.updated_at !== null)
      ? Timestamp.fromPartial(object.updated_at)
      : undefined;
    return message;
  },
};

function createBaseListUsersResponse(): ListUsersResponse {
  return { users: [] };
}

export const ListUsersResponse: MessageFns<ListUsersResponse> = {
  encode(message: ListUsersResponse, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    for (const v of message.users) {
      UserResponse.encode(v!, writer.uint32(10).fork()).join();
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): ListUsersResponse {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListUsersResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.users.push(UserResponse.decode(reader, reader.uint32()));
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ListUsersResponse {
    return {
      users: globalThis.Array.isArray(object?.users) ? object.users.map((e: any) => UserResponse.fromJSON(e)) : [],
    };
  },

  toJSON(message: ListUsersResponse): unknown {
    const obj: any = {};
    if (message.users?.length) {
      obj.users = message.users.map((e) => UserResponse.toJSON(e));
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ListUsersResponse>, I>>(base?: I): ListUsersResponse {
    return ListUsersResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<ListUsersResponse>, I>>(object: I): ListUsersResponse {
    const message = createBaseListUsersResponse();
    message.users = object.users?.map((e) => UserResponse.fromPartial(e)) || [];
    return message;
  },
};

function createBaseGetPresignedURLRequest(): GetPresignedURLRequest {
  return { file_name: "", content_type: "" };
}

export const GetPresignedURLRequest: MessageFns<GetPresignedURLRequest> = {
  encode(message: GetPresignedURLRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.file_name !== "") {
      writer.uint32(10).string(message.file_name);
    }
    if (message.content_type !== "") {
      writer.uint32(18).string(message.content_type);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): GetPresignedURLRequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetPresignedURLRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.file_name = reader.string();
          continue;
        }
        case 2: {
          if (tag !== 18) {
            break;
          }

          message.content_type = reader.string();
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetPresignedURLRequest {
    return {
      file_name: isSet(object.file_name) ? globalThis.String(object.file_name) : "",
      content_type: isSet(object.content_type) ? globalThis.String(object.content_type) : "",
    };
  },

  toJSON(message: GetPresignedURLRequest): unknown {
    const obj: any = {};
    if (message.file_name !== "") {
      obj.file_name = message.file_name;
    }
    if (message.content_type !== "") {
      obj.content_type = message.content_type;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetPresignedURLRequest>, I>>(base?: I): GetPresignedURLRequest {
    return GetPresignedURLRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetPresignedURLRequest>, I>>(object: I): GetPresignedURLRequest {
    const message = createBaseGetPresignedURLRequest();
    message.file_name = object.file_name ?? "";
    message.content_type = object.content_type ?? "";
    return message;
  },
};

function createBaseGetPresignedURLResponse(): GetPresignedURLResponse {
  return { presigned_url: "", image_url: "" };
}

export const GetPresignedURLResponse: MessageFns<GetPresignedURLResponse> = {
  encode(message: GetPresignedURLResponse, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.presigned_url !== "") {
      writer.uint32(10).string(message.presigned_url);
    }
    if (message.image_url !== "") {
      writer.uint32(18).string(message.image_url);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): GetPresignedURLResponse {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetPresignedURLResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.presigned_url = reader.string();
          continue;
        }
        case 2: {
          if (tag !== 18) {
            break;
          }

          message.image_url = reader.string();
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetPresignedURLResponse {
    return {
      presigned_url: isSet(object.presigned_url) ? globalThis.String(object.presigned_url) : "",
      image_url: isSet(object.image_url) ? globalThis.String(object.image_url) : "",
    };
  },

  toJSON(message: GetPresignedURLResponse): unknown {
    const obj: any = {};
    if (message.presigned_url !== "") {
      obj.presigned_url = message.presigned_url;
    }
    if (message.image_url !== "") {
      obj.image_url = message.image_url;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetPresignedURLResponse>, I>>(base?: I): GetPresignedURLResponse {
    return GetPresignedURLResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetPresignedURLResponse>, I>>(object: I): GetPresignedURLResponse {
    const message = createBaseGetPresignedURLResponse();
    message.presigned_url = object.presigned_url ?? "";
    message.image_url = object.image_url ?? "";
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends globalThis.Array<infer U> ? globalThis.Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function toTimestamp(date: Date): Timestamp {
  const seconds = Math.trunc(date.getTime() / 1_000);
  const nanos = (date.getTime() % 1_000) * 1_000_000;
  return { seconds, nanos };
}

function fromTimestamp(t: Timestamp): Date {
  let millis = (t.seconds || 0) * 1_000;
  millis += (t.nanos || 0) / 1_000_000;
  return new globalThis.Date(millis);
}

function fromJsonTimestamp(o: any): Timestamp {
  if (o instanceof globalThis.Date) {
    return toTimestamp(o);
  } else if (typeof o === "string") {
    return toTimestamp(new globalThis.Date(o));
  } else {
    return Timestamp.fromJSON(o);
  }
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}

export interface MessageFns<T> {
  encode(message: T, writer?: BinaryWriter): BinaryWriter;
  decode(input: BinaryReader | Uint8Array, length?: number): T;
  fromJSON(object: any): T;
  toJSON(message: T): unknown;
  create<I extends Exact<DeepPartial<T>, I>>(base?: I): T;
  fromPartial<I extends Exact<DeepPartial<T>, I>>(object: I): T;
}
