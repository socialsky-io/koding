package main

import "runtime"

const (
	KiteVersion = "0.0.1"

	// Name is the user facing name for this binary. Internally we call it
	// klientctl to avoid confusion.
	Name = "kd"

	// KlientName is the user facing name for klient.
	KlientName = "KD Daemon"

	// KlientAddress is url of locally running klient to connect to send
	// user commands.
	KlientAddress = "http://127.0.0.1:56789/kite"

	// KiteHome is full path to the kite key that we will use to authenticate
	// to the given klient.
	KiteHome = "/etc/kite"

	// KlientDirectory is full path to directory that holds klient. Note that
	// this gets created and deleted as needed. **Never** make this something important,
	// like /usr/local/bin
	KlientDirectory = "/opt/kite/klient"

	// KlientctlDirectory is full path to directory that holds klientctl.
	KlientctlDirectory = "/usr/local/bin"

	// KlientctlBinName is the bin named that will be stored in the KlientctlDirectory.
	KlientctlBinName = "kd"

	// KontrolUrl is the url to connect to authenticate local klient and get
	// list of machines.
	KontrolUrl = "https://koding.com/kontrol/kite"

	// Version is the current version of klientctl. This number is used
	// by CheckUpdate to determine if current version is behind or equal to latest
	// version on S3 bucket.
	Version = 9

	osName = runtime.GOOS

	// S3UpdateLocation is publically accessible url to check for new updates.
	S3UpdateLocation = "https://koding-kd.s3.amazonaws.com/latest-version.txt"

	// S3KlientctlPath is publically accessible url for latest version of klient.
	// Each OS has its own version of binary, identifiable by OS suffix.
	S3KlientPath = "https://koding-kd.s3.amazonaws.com/klient-" + osName

	// S3KlientctlPath is publically accessible url for latest version of
	// klientctl. Each OS has its own version of binary, identifiable by suffix.
	S3KlientctlPath = "https://koding-kd.s3.amazonaws.com/klientctl-" + osName

	// SSHDefaultKeyDir is the default directory that stores users ssh key pairs.
	SSHDefaultKeyDir = ".ssh"

	// SSHDefaultKeyDir is the default name of the ssh key pair.
	SSHDefaultKeyName = "kd-ssh-key"
)
