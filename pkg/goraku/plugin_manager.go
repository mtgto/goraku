package goraku

import "github.com/nlopes/slack"

type pluginManager struct {
	plugins []Plugin
}

func newPluginManager() *pluginManager {
	return &pluginManager{
		plugins: []Plugin{},
	}
}

func (pm *pluginManager) addPlugin(plugin Plugin) {
	pm.plugins = append(pm.plugins, plugin)
}

func (pm *pluginManager) processMessageEvent(g *Goraku, event *slack.MessageEvent) {
	for _, plugin := range pm.plugins {
		plugin.Hear(g, *event)
	}
}
