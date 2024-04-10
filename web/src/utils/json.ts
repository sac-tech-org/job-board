// Credit to this article: https://hackernoon.com/mastering-type-safe-json-serialization-in-typescript

export type JSONPrimitive = string | number | boolean | null | undefined;

export type JSONValue =
  | JSONPrimitive
  | JSONValue[]
  | {
      [key: string]: JSONValue;
    };

export type NotAssignableToJson = bigint | symbol | Function;

export type JSONCompatible<T> = unknown extends T
  ? never
  : {
      [P in keyof T]: T[P] extends JSONValue
        ? T[P]
        : T[P] extends NotAssignableToJson
          ? never
          : JSONCompatible<T[P]>;
    };

function toJsonValue<T>(value: JSONCompatible<T>): JSONValue {
  return value;
}

export function safeJsonStringify<T>(data: JSONCompatible<T>) {
  return JSON.stringify(data);
}

function safeJsonParse(text: string): unknown {
  return JSON.parse(text);
}
