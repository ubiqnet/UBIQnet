package loader

import (
	pluginbadgerds "ubiqnet/plugin/plugins/badgerds"
	pluginiplddagjose "ubiqnet/plugin/plugins/dagjose"
	pluginflatfs "ubiqnet/plugin/plugins/flatfs"
	pluginfxtest "ubiqnet/plugin/plugins/fxtest"
	pluginipldgit "ubiqnet/plugin/plugins/git"
	pluginlevelds "ubiqnet/plugin/plugins/levelds"
	pluginpeerlog "ubiqnet/plugin/plugins/peerlog"
)

// DO NOT EDIT THIS FILE
// This file is being generated as part of plugin build process
// To change it, modify the plugin/loader/preload.sh

func init() {
	Preload(pluginipldgit.Plugins...)
	Preload(pluginiplddagjose.Plugins...)
	Preload(pluginbadgerds.Plugins...)
	Preload(pluginflatfs.Plugins...)
	Preload(pluginlevelds.Plugins...)
	Preload(pluginpeerlog.Plugins...)
	Preload(pluginfxtest.Plugins...)
}