
fs = require 'fs'
hat = require 'hat'
{exec} = require 'child_process'

config    = require './config'

log4js    = require 'log4js'
log       = log4js.getLogger("[#{config.name}]")

# HELPERS
createTmpDir = require './createtmpdir'

safeUsers = []

createTmpFile =(username, command, callback)->
  tmpFile = "/Users/#{username}/.tmp/tmp_#{hat()}.sh"
  fs.writeFile tmpFile,command,'utf-8',(err)->
    if err
      callback? err
    else
      callback? null,tmpFile

prepareForBashExecute =(options, callback)->
  {username,command} = options
  createTmpDir username, (err)->
    if err
      log.error err
      callback err
    else
      log.debug "[OK] executing command for #{username}: #{command}"
      createTmpFile username, command, (err, tmpfile)->
        if err then callback? err
        else callback null, tmpfile

execute = (options,callback)->

  {filename,command,username} = options
  if filename
    execStr = "chown #{username} #{filename} && su -l #{username} -c 'sh #{filename}'"
    unlink = yes
  else if command
    execStr = "su -l #{username} -c '#{command}'"
    unlink = no
  else
    log.error "execute can only work with provided .filename or .command"

  cmd = exec execStr, {maxBuffer: 1024*1024}, (err,stdout,stderr)->
    respond {err,stdout,stderr},callback
    fs.unlink filename if unlink is yes
    log.debug "executed", truncateOutput execStr

truncateOutput = (output)->

  if output.length > 300
    "#{output[0...300]} ...[truncated output]"
  else
    output

respond = (options,callback)->

  if Array.isArray options
    [err,stdout,stderr] = options
  else
    {err,stdout,stderr} = options

  if err?
    callback {code: err.code, message: err.message, stderr: stderr}, stdout
    if stdout
      log.warn "[WARNING]", err, truncateOutput(stderr), truncateOutput stdout
    else
      log.error "[ERROR]", err, truncateOutput stderr
  else
    # log.info "[OK] command \"#{command}\" executed for user #{username}"
    callback? null,stdout

containsAnyChar =(haystack, needle)->
  # for char in needle
  #   if ~haystack.indexOf char
  #     return yes
  # no
  a = createRegExp needle
  return a.test haystack


createRegExp = do ->
  memos = {}
  (str, flags)->
    return memos[str] if memos[str]?
    special = /(\.|\^|\$|\*|\+|\?|\(|\)|\[|\{|\\|\|)/
    memos[str] = RegExp '('+str.split('').map(
      (char)-> char.replace special, (_, foundChar)-> '\\'+foundChar
    ).join('|')+')', flags

checkUid = (options, createSystemUser, callback)->
  #
  # This methid will check user's uid
  #
  #
  # options =
  #   username : String  #username of the unix user
  #
  {username, nrOfRecursion} = options

  if username in safeUsers
    # log.info "Username '#{username}' is already safe"
    callback? null
  else
    getuid = exec "/usr/bin/id -u #{username}", (err,stdout,stderr)=>
      if err?
        log.error "[ERROR] unable to get user's UID: #{stderr}"
        if nrOfRecursion is 1
          callback?  "[ERROR] unable to get user's UID, can't create user: #{stderr}"
        else
          createSystemUser? {username,fullName:username,password:hat()},(err,res)=>
            unless err
              log.info "User is just created, run the command again, it will work this time."
              checkUid {username,nrOfRecursion:1},callback
            else
              log.error "CANT CREATE THIS USER"
              callback?  "[ERROR] unable to get user's UID, can't create user: #{stderr}"


      else if stdout < config.minAllowedUid
        stdout = stdout.replace(/(\r\n|\n|\r)/gm," ")
        log.error e = "[ERROR]  min UID for commands is #{config.minAllowedUid}, your #{stdout}"
        callback? e
      else
        stdout = stdout.replace(/(\r\n|\n|\r)/gm," ")
        log.debug "[OK] func:checkUid: user's #{username} UID #{stdout} allowed"
        safeUsers.push username
        callback? null

module.exports =(options, callback)->
  #
  # this method will execute any system command inside user's env
  #
  # options =
  #   username : String #username of the unix user
  #   command  : String #bash command
  #
  # START
  # if command.contains anyOf ";&|><*?`$(){}[]!#'"
    # 1- create tmpdir /Users/[username]/.tmp
    # 2- create tmpfile /Users/[username]/.tmp/tmp_[uniqid].txt
    # 3- write the command inside the tmp file
    # 4- bash execute the tmp file with su -l [username]
    # 5- delete the tmpfile.
  # else
  #   1- execute command with su -l username -c 'command'

  {username,command} = options

  # log.debug "func:executeCommand: executing command #{command}"
  checkUid options, @createSystemUser?.bind(@), (error)->
    if error?
      callback? error
    else
      chars = ";&|><*?`$(){}[]!#"
      if containsAnyChar command,chars
        log.debug "exec in a file",command
        prepareForBashExecute options, (err, filename)->
          unless err
            execute {filename,username},callback
          else
            callback err
            log.error err
      else
        execute {command,username},callback
        # log.debug "exec directly",command

  # # Send memory usage to librato
  # {argv} = require 'optimist'
  # KONFIG = require('koding-config-manager').load("main.#{argv.c}")
  # if KONFIG.librato?.push
  #   os = require "os"

  #   # Post to Librato
  #   librato = require("librato-metrics").createClient(
  #     email: KONFIG.librato.email
  #     token: KONFIG.librato.token
  #   )
  #   data = counters: [
  #     name: 'kite.sharedhosting.execute'
  #     value: 1
  #   ]
  #   librato.post '/metrics', data, (err, response) ->
  #     if err
  #       console.log "Librato - Can't push stats: " + err
