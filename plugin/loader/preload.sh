#!/usr/bin/env bash

DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

LIST="$DIR/preload_list"

to_preload() {
  awk 'NF' $LIST | sed '/^#/d'
  if [[ -n "$IPFS_PLUGINS" ]]; then
      for plugin in $IPFS_PLUGINS; do
          echo "$plugin ubiqnet/plugin/plugins/$plugin *"
      done
  fi
}

cat <<EOL
package loader

import (
EOL

to_preload | while read -r name path num; do
	echo "plugin$name \"$path\""
done | sort -u

cat <<EOL
)


// DO NOT EDIT THIS FILE
// This file is being generated as part of plugin build process
// To change it, modify the plugin/loader/preload.sh

func init() {
EOL

to_preload | while read -r name path num; do
	case "$num" in
		'*') echo "	Preload(plugin$name.Plugins...)" ;; # All plugins
		*) echo "	Preload(plugin$name.Plugins[$num])" ;; # A specific plugin
	esac
done

echo "}"
