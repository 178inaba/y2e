#!/bin/bash

go build
eval "$(y2e _test/test.yaml)"
rm y2e

if [[ "${FOO}" != 'foo' ]]; then
  printf '${FOO} is %s, want %s.\n' "${FOO}" 'foo'
  exit 1
fi

if [[ "${BAR}" != '123' ]]; then
  printf '${BAR} is %s, want %s.\n' "${BAR}" '123'
  exit 1
fi

if [[ "${BAZ}" != 'true' ]]; then
  printf '${BAZ} is %s, want %s.\n' "${BAZ}" 'true'
  exit 1
fi

echo ok
