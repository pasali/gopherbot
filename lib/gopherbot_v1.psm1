# Return values for robot method calls
Enum BotRet
{
    Ok = 0
    UserNotFound = 1
    ChannelNotFound = 2
    AttributeNotFound = 3
    FailedUserDM = 4
    FailedChannelJoin = 5
    DatumNotFound = 6
    DatumLockExpired = 7
    DataFormatError = 8
    BrainFailed = 9
    InvalidDatumKey = 10
    InvalidDblPtr = 11
    InvalidCfgStruct = 12
    NoConfigFound = 13
    RetryPrompt = 14
    ReplyNotMatched = 15
    UseDefaultValue = 16
    TimeoutExpired = 17
    Interrupted = 18
    MatcherNotFound = 19
    NoUserEmail = 20
    NoBotEmail = 21
    MailError = 22
}

# Plugin return values / exit codes, return values from CallPlugin
Enum PlugRet
{
    Normal = 0
    Success = 1
    Fail = 2
    MechanismFail = 3
    ConfigurationError = 4
}

function enc64([String] $msg) {
    [OutputType([String])]
    $b  = [System.Text.Encoding]::UTF8.GetBytes($msg)
    $enc = [System.Convert]::ToBase64String($b)
    return "base64:$enc"
}

function dec64([String] $msg) {
    [OutputType([String])]
    [String[]] $parts = $msg.Split(":")
    if ($parts[0] -eq "base64") {
        $b  = [System.Convert]::FromBase64String($parts[1])
        return [System.Text.Encoding]::UTF8.GetString($b)
    } else {
        return $msg
    }
}

class Attribute {
    # Properties
    [String] $Attr
    [BotRet] $Ret

    Attribute([PSCustomObject] $obj) {
        $a = dec64($obj.Attribute)
        $this.Attr = $a
        $this.Ret = $obj.RetVal -As [BotRet]
    }

    [String] ToString() {
        return $this.Attr
    }
}

class Reply {
    [String] $Reply
    [BotRet] $Ret

    Reply([String] $rep, [Int] $ret) {
        $this.Reply = $rep
        $this.Ret = $ret -As [BotRet]
    }

    [String] ToString() {
        return $this.Reply
    }
}

class BotFuncCall {
    [String] $FuncName
    [String] $User
    [String] $Channel
    [String] $Format
    [String] $PluginID
    [PSCustomObject] $FuncArgs

    BotFuncCall([String] $fn, [String] $u, [String] $c, [String] $fmt, [String] $p, [PSCustomObject] $funcArgs ) {
        $this.FuncName = $fn
        $this.User = $u
        $this.Channel = $c
        $this.Format = $fmt
        $this.PluginID = $p
        $this.FuncArgs = $funcArgs
    }
}

class Robot
{
    # Properties
    [String] $Channel
    [String] $User
    hidden [String] $PluginID

    # Constructor
    Robot([String] $channel, [String] $user, [String] $pluginid) {
        $this.Channel = $channel
        $this.User = $user
        $this.PluginID = $pluginid
    }

    [Robot] Direct() {
        return [Robot]::new("", $this.User, $this.PluginID)
    }

    Pause([single] $seconds) {
        Start-Sleep $seconds
    }

    [bool] CheckAdmin() {
        return $this.Call("CheckAdmin", $null).Boolean -As [bool]
    }

    [bool] Elevate([bool] $immediate) {
        $funcArgs = [PSCustomObject]@{ Immediate=$immediate }
        return $this.Call("Elevate", $funcArgs).Boolean -As [bool]
    }

    [bool] Elevate() {
        return $this.Elevate($FALSE)
    }

    [int] RandomInt([int] $i) {
        return Get-Random -Maximum $i
    }

    [string] RandomString([String[]] $sarr) {
        $l = $sarr.Count
        $i = Get-Random -Maximum $l
        return $sarr[$i]
    }

    [PSCustomObject] GetPluginConfig() {
        return $this.Call("GetPluginConfig", $null)
    }

    [PSCustomObject] Call([String] $fname, [PSCustomObject] $funcArgs, [String] $format) {
        $bfc = [BotFuncCall]::new($fname, $this.User, $this.Channel, $format, $this.PluginID, $funcArgs)
        $fc = ConvertTo-Json $bfc
        # if ($fname -ne "Log") { $this.Log("Debug", "DEBUG - Sending: $fc") }
        $r = Invoke-WebRequest -URI "$Env:GOPHER_HTTP_POST/json" -Method Post -UseBasicParsing -Body $fc
        $c = $r.Content
        # if ($fname -ne "Log") { $this.Log("Debug", "DEBUG - Got back: $c") }
        return ConvertFrom-Json $c
    }

    [PSCustomObject] Call([String] $fname, [PSCustomObject] $funcArgs) {
        return $this.Call($fname, $funcArgs, "variable")
    }

    [PlugRet] CallPlugin([String] $plugName, [String[]]$plugArgs) {
        $funcArgs = [PSCustomObject]@{ PluginName=$plugName }
        $ret = $this.Call("CallPlugin", $funcArgs)
        if ([PlugRet]$ret.PlugRetVal -ne "Success") {
            return [PlugRet]$ret.PlugRetVal
        }
        if ( $ret.InterpreterPath -match "powershell" ) {
            $plugPath = $ret.PluginPath -replace ' ','` '
        } else {
            $plugPath = $ret.PluginPath
        }
        $plugArgs = [Array]$plugPath + $plugArgs
        $Env:GOPHER_PLUGIN_ID = $ret.PluginID
        $proc = Start-Process -FilePath $ret.InterpreterPath -ArgumentList $plugArgs -NoNewWindow -PassThru -Wait
        $Env:GOPHER_PLUGIN_ID = $this.PluginID
        return [PlugRet]$proc.ExitCode
    }

    [PSCustomObject] CheckoutDatum([String] $key, [Bool] $rw) {
        $funcArgs = [PSCustomObject]@{ Key=$key; RW=$rw }
        $ret = $this.Call("CheckoutDatum", $funcArgs)
        $ret | Add-Member -NotePropertyName Key -NotePropertyValue $key
        return $ret
    }

    CheckinDatum([PSCustomObject] $mem){
        $funcArgs = [PSCustomObject]@{ Key=$mem.Key; Token=$mem.LockToken }
        $this.Call("CheckinDatum", $funcArgs)
    }

    [BotRet] UpdateDatum([PSCustomObject] $mem){
        $funcArgs = [PSCustomObject]@{ Key=$mem.Key; Token=$mem.LockToken; Datum=$mem.Datum }
        $ret = $this.Call("UpdateDatum", $funcArgs)
        return $ret.RetVal -As [BotRet]
    }

    [BotRet] Remember([String] $key, [String] $value){
        $funcArgs = [PSCustomObject]@{ Key=$key; Value=$value }
        $ret = $this.Call("Remember", $funcArgs)
        return $ret.RetVal -As [BotRet]
    }

    [BotRet] RememberContext([String] $context, [String] $value){
        return $this.Remember("context:"+$context, $value)
    }

    [String] Recall([String] $key){
        $funcArgs = [PSCustomObject]@{ Key=$key }
        $ret = $this.Call("Recall", $funcArgs)
        return $ret.StrVal
    }

    [Attribute] GetSenderAttribute([String] $attr) {
        $funcArgs = [PSCustomObject]@{ Attribute=$attr }
        $ret = $this.Call("GetSenderAttribute", $funcArgs)
        return [Attribute]::new($ret)
    }

    [Attribute] GetUserAttribute([String] $user, [String] $attr) {
        $funcArgs = [PSCustomObject]@{ User=$user; Attribute=$attr }
        $ret = $this.Call("GetUserAttribute", $funcArgs)
        return [Attribute]::new($ret)
    }

    [Attribute] GetBotAttribute([String] $attr) {
        $funcArgs = [PSCustomObject]@{ Attribute=$attr }
        $ret = $this.Call("GetBotAttribute", $funcArgs)
        return [Attribute]::new($ret)
    }

    [Reply] PromptForReply([String] $regexid, [String] $prompt) {
        return $this.PromptUserChannelForReply($regexid, $this.User, $this.Channel, $prompt)
    }

    [Reply] PromptUserForReply([String] $regexid, [String] $user, [String] $prompt) {
        return $this.PromptUserChannelForReply($regexid, $user, "", $prompt)
    }

    [Reply] PromptUserChannelForReply([String] $regexid, [String] $user, [String] $channel, [String] $prompt) {
        $funcArgs = [PSCustomObject]@{ RegexID=$regexid; User=$user; Channel=$channel; Prompt=$prompt }
        $ret = $null
        For ($i=0; $i -le 3; $i++) {
            $ret = $this.Call("PromptUserChannelForReply", $funcArgs)
            if ([int]$ret.RetVal -eq "RetryPrompt" ){ continue }
            $rep = dec64($ret.Reply)
            return [Reply]::new($rep, $ret.Ret -As [BotRet])
        }
        $rep = dec64($ret.Reply)
        if ($ret -eq "RetryPrompt" ) {
            return [Reply]::new($rep, [BotRet]("Interrupted"))
        }
        return [Reply]::new($rep.Reply, $ret.RetVal -As [BotRet])
    }

    Log([String] $level, [String] $message) {
        $funcArgs = [PSCustomObject]@{ Level=$level; Message=$message }
        $this.Call("Log", $funcArgs)
    }

    [BotRet] SendChannelMessage([String] $channel, [String] $msg, [String] $format="variable") {
        $funcArgs = [PSCustomObject]@{ Channel=$channel; Message=enc64($msg) }
        return $this.Call("SendChannelMessage", $funcArgs, $format).RetVal -As [BotRet]
    }

    [BotRet] SendChannelMessage([String] $channel, [String] $msg) {
        return $this.SendChannelMessage($channel, $msg, "variable")
    }

    [BotRet] SendUserMessage([String] $user, [String] $msg, [String] $format="variable") {
        $funcArgs = [PSCustomObject]@{ User=$user; Message=enc64($msg) }
        return $this.Call("SendUserMessage", $funcArgs, $format).RetVal -As [BotRet]
    }

    [BotRet] SendUserMessage([String] $user, [String] $msg) {
        return $this.SendUserMessage($user, $msg, "variable")
    }

    [BotRet] SendUserChannelMessage([String] $user, [String] $channel, [String] $msg, [String] $format="variable") {
        $funcArgs = [PSCustomObject]@{ User=$user; Channel=$channel; Message=enc64($msg) }
        return $this.Call("SendUserChannelMessage", $funcArgs, $format).RetVal -As [BotRet]
    }

    [BotRet] SendUserChannelMessage([String] $user, [String] $channel, [String] $msg) {
        return $this.SendUserChannelMessage($user, $channel, $msg, "variable")
    }

    [BotRet] Say([String] $msg, [String] $format) {
        if ($this.Channel -eq ""){
            return $this.SendUserMessage($this.User, $msg, $format)
        } else {
            return $this.SendChannelMessage($this.Channel, $msg, $format)
        }
    }

    [BotRet] Say([String] $msg) {
        return $this.Say($msg, "variable")
    }

    [BotRet] Reply([String] $msg, [String] $format = "variable") {
        if ($this.Channel -eq "") {
            return $this.SendUserMessage($this.User, $msg, $format)
        } else {
            return $this.SendUserChannelMessage($this.User, $this.Channel, $msg, $format)
        }
    }

    [BotRet] Reply([String] $msg) {
        return $this.Reply($msg, "variable")
    }
}

function Get-Robot() {
    return [Robot]::new($Env:GOPHER_CHANNEL, $Env:GOPHER_USER, $Env:GOPHER_PLUGIN_ID)
}

export-modulemember -function Get-Robot
