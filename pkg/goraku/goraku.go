package goraku

import (
	"github.com/nlopes/slack"
	"log"
)

type Bot interface {
	// Reply to message with text.
	Reply(message *slack.MessageEvent, text string)
	// Reply to message using thread, attachments, and so on.
	ReplyWithOptions(message *slack.MessageEvent, options ...slack.MsgOption) error
}

type Goraku struct {
	Bot
	API           *slack.Client
	Me            *slack.UserDetails
	Team          *slack.Team
	rtm           *slack.RTM
	pluginManager *pluginManager
}

func NewSlackBot(slackApiToken string, options ...slack.Option) *Goraku {
	client := slack.New(slackApiToken, options...)
	return &Goraku{
		API:           client,
		pluginManager: newPluginManager(),
	}
}

// Reply to message with text.
// Use ReplyWIthOptions for complicated message.
func (g *Goraku) Reply(message *slack.MessageEvent, text string) {
	outgoing := g.rtm.NewOutgoingMessage(text, message.Channel)
	g.rtm.SendMessage(outgoing)
}

// Reply to message using thread, attachments, and so on.
func (g *Goraku) ReplyWithOptions(message *slack.MessageEvent, options ...slack.MsgOption) error {
	_, _, err := g.rtm.PostMessage(message.Channel, options...)
	return err
}

// Start RTM connection.
func (g *Goraku) Start(options ...slack.RTMOption) {
	g.rtm = g.API.NewRTM(options...)
	go g.rtm.ManageConnection()

	for msg := range g.rtm.IncomingEvents {
		log.Print("Event Received: ")
		switch ev := msg.Data.(type) {
		case *slack.HelloEvent:
			// Ignore hello
		case *slack.ConnectedEvent:
			log.Printf("Connected: %v\n", ev)
			info := g.rtm.GetInfo()
			g.Me = info.User
			g.Team = info.Team
		case *slack.DisconnectedEvent:
			log.Printf("Disconnected: %v\n", ev)
			break
		case *slack.MessageEvent:
			g.pluginManager.processMessageEvent(g, ev)
			//ctx.Plugins.ExecPlugins(ctx.responseEvent(ev))
		case *slack.PresenceChangeEvent:
			log.Printf("Presence Change: %v\n", ev)
		case slack.LatencyReport:
			log.Printf("Current latency: %v\n", ev.Value)
		case *slack.RTMError:
			log.Printf("Error: %d - %s\n", ev.Code, ev.Msg)
		default:
			// Ignore other events..
			log.Printf("Unexpected: %+v\n", ev)
		}
	}
}

// Add goraku plugin to manager
func (g *Goraku) AddPlugin(plugin Plugin) {
	g.pluginManager.addPlugin(plugin)
}
