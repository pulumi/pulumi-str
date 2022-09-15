import * as str from "@pulumi/str";

const out = str.replaceOutput({
  string: "Foo",
  old: "o",
  new: "O",
});

out.result.apply((s: string) => {
  if (s !== "FOO") {
    throw "Unhappy";
  }
  return s;
});

export const s = out;
