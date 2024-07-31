#!/bin/bash

pkg_dir="pkg"

if [ ! -d "$pkg_dir" ]; then
  mkdir "$pkg_dir"
fi

for dir in */; do
  if [ "$dir" != "pkg/" ]; then
    mv "$dir" "$pkg_dir"
  fi
done