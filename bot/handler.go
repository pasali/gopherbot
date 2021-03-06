package bot

import (
	"encoding/json"
	"fmt"
	"strings"
)

// an empty object type for passing a Handler to the connector.
type handler struct{}

/* Handle incoming messages and other callbacks from the connector. */

// GetLogLevel returns the bot's current loglevel, mainly for the
// connector to make it's own decision about logging
func (h handler) GetLogLevel() LogLevel {
	return getLogLevel()
}

// GetInstallPath gets the path to the bot's install dir -
// the location of default configuration and stock external plugins.
func (h handler) GetInstallPath() string {
	robot.RLock()
	defer robot.RUnlock()
	return robot.installPath
}

// GetLocalPath gets the path to the bot's install dir -
// the location of default configuration and stock external plugins.
func (h handler) GetLocalPath() string {
	robot.RLock()
	defer robot.RUnlock()
	return robot.localPath
}

// ChannelMessage accepts an incoming channel message from the connector.
func (h handler) IncomingMessage(channelName, userName, messageFull string) {
	// When command == true, the message was directed at the bot
	isCommand := false
	logChannel := channelName
	var message string

	robot.RLock()
	for _, user := range robot.ignoreUsers {
		if strings.EqualFold(userName, user) {
			Log(Debug, "Ignoring user", userName)
			robot.RUnlock()
			return
		}
	}
	robot.RUnlock()
	if robot.preRegex != nil {
		matches := robot.preRegex.FindAllStringSubmatch(messageFull, -1)
		if matches != nil && len(matches[0]) == 2 {
			isCommand = true
			message = matches[0][1]
		}
	}
	if !isCommand && robot.postRegex != nil {
		matches := robot.postRegex.FindAllStringSubmatch(messageFull, -1)
		if matches != nil && len(matches[0]) == 3 {
			isCommand = true
			message = matches[0][1] + matches[0][2]
		}
	}
	if !isCommand {
		message = messageFull
	}
	if len(channelName) == 0 { // true for direct messages
		isCommand = true
		logChannel = "(direct message)"
	}
	Log(Trace, fmt.Sprintf("Command \"%s\" in channel \"%s\"", message, logChannel))
	handleMessage(isCommand, channelName, userName, message)
}

// GetProtocolConfig unmarshals the connector's configuration data into a provided struct
func (h handler) GetProtocolConfig(v interface{}) error {
	robot.RLock()
	err := json.Unmarshal(protocolConfig, v)
	robot.RUnlock()
	return err
}

// GetBrainConfig unmarshals the brain's configuration data into a provided struct
func (h handler) GetBrainConfig(v interface{}) error {
	robot.RLock()
	err := json.Unmarshal(brainConfig, v)
	robot.RUnlock()
	return err
}

// GetElevateConfig unmarshals the brain's configuration data into a provided struct
func (h handler) GetElevateConfig(v interface{}) error {
	robot.RLock()
	err := json.Unmarshal(elevateConfig, v)
	robot.RUnlock()
	return err
}

// Log logs a message to the robot's log file (or stderr)
func (h handler) Log(l LogLevel, v ...interface{}) {
	Log(l, v...)
}

// Connectors that support it can call SetFullName; otherwise it can
// be configured in gobot.conf.
func (h handler) SetFullName(n string) {
	Log(Debug, "Setting full name to: "+n)
	robot.Lock()
	robot.fullName = n
	robot.Unlock()
	updateRegexes()
}

// Connectors that support it can call SetName; otherwise it should
// be configured in gobot.conf.
func (h handler) SetName(n string) {
	Log(Debug, "Setting name to: "+n)
	robot.Lock()
	robot.name = n
	ignoring := false
	for _, name := range robot.ignoreUsers {
		if strings.EqualFold(n, name) {
			ignoring = true
			break
		}
	}
	if !ignoring {
		robot.ignoreUsers = append(robot.ignoreUsers, strings.ToLower(n))
	}
	robot.Unlock()
	updateRegexes()
}
