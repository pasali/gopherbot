#!/usr/bin/ruby
require 'net/http'
require 'json'

# To install:
# 1) Copy this file to plugins/weather.rb
# 2) Enable in gopherbot.yaml like so:

#ExternalPlugins:
#- Name: weather
#  Path: plugins/weather.rb

# 3) Put your configuration in conf/plugins/weather.yaml:

#Config:
#  APIKey: <your openweathermap key>
#  TemperatureUnits: imperial # or 'metric'
#  Country: 'us' # or other ISO 3166 country code

# load the Gopherbot ruby library and instantiate the bot
require ENV["GOPHER_INSTALLDIR"] + '/lib/gopherbot_v1'
bot = Robot.new()

# 
defaultConfig = <<'DEFCONFIG'
---
Help:
- Keywords: [ "weather" ]
  Helptext: [ "(bot), weather in <city or zip code> - fetch the weather from OpenWeatherMap" ]
CommandMatchers:
- Command: weather
  Regex: '(?i:weather (?:in|for) ([\w ]+))'
DEFCONFIG

command = ARGV.shift()

case command
when "configure"
	puts defaultConfig
	exit
when "weather"
    c = bot.GetPluginConfig()
    uri = URI("http://api.openweathermap.org/data/2.5/weather?q=#{location},#{c["Country"]}&units=#{c["TemperatureUnits"]}&APPID=#{c["APIKey"]}"
    d = JSON::parse(Net::HTTP.get(uri))
    w = d["weather"][0]
    t = d["main"]
    bot.Say("The weather in #{d["name"]} is currently \"#{w["description"]}\" and #{t["temp"]} degrees, with a forecast low of #{t["temp_min"]} and high of #{t["temp_max"]}")
end