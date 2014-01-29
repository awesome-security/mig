/* Mozilla InvestiGator Agent

Version: MPL 1.1/GPL 2.0/LGPL 2.1

The contents of this file are subject to the Mozilla Public License Version
1.1 (the "License"); you may not use this file except in compliance with
the License. You may obtain a copy of the License at
http://www.mozilla.org/MPL/

Software distributed under the License is distributed on an "AS IS" basis,
WITHOUT WARRANTY OF ANY KIND, either express or implied. See the License
for the specific language governing rights and limitations under the
License.

The Initial Developer of the Original Code is
Mozilla Corporation
Portions created by the Initial Developer are Copyright (C) 2014
the Initial Developer. All Rights Reserved.

Contributor(s):
Julien Vehent jvehent@mozilla.com [:ulfr]

Alternatively, the contents of this file may be used under the terms of
either the GNU General Public License Version 2 or later (the "GPL"), or
the GNU Lesser General Public License Version 2.1 or later (the "LGPL"),
in which case the provisions of the GPL or the LGPL are applicable instead
of those above. If you wish to allow use of your version of this file only
under the terms of either the GPL or the LGPL, and not to allow others to
use your version of this file under the terms of the MPL, indicate your
decision by deleting the provisions above and replace them with the notice
and other provisions required by the GPL or the LGPL. If you do not delete
the provisions above, a recipient may use your version of this file under
the terms of any one of the MPL, the GPL or the LGPL.
*/

package main

import(
	"mig"
	"time"
)

var LOGGINGCONF = mig.Logging{
	Mode:	"stdout",	// stdout | file | syslog
	Level:	"debug",	// debug | info | ...
	//File:	"/tmp/migagt.log",
	//Host:	"syslog_hostname",
	//Port:	514,
	//Protocol: "udp",
}

// location of the rabbitmq server
var AMQPBROKER string = "amqp://guest:guest@172.21.1.1:5672/"

// frequency at which the agent sends keepalive messages
var HEARTBEATFREQ string = "300s"

// timeout after which a module run is killed
var MODULETIMEOUT time.Duration = 300 * time.Second

// PGP public key that is authorized to sign actions
var PUBLICPGPKEYS = [...]string{
`
-----BEGIN PGP PUBLIC KEY BLOCK-----
Version: GnuPG v1.4.15 (GNU/Linux). Name: opsec+mig@mozilla.com

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
=bih3
-----END PGP PUBLIC KEY BLOCK-----
`,
`
-----BEGIN PGP PUBLIC KEY BLOCK-----
Version: SKS 1.1.4. Name: jvehent+mig@mozilla.com

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
D6fA7lwj/BFick6UyApfvw==
=DW5+
-----END PGP PUBLIC KEY BLOCK-----
`}
