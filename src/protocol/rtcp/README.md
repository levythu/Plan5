# Plan5 - Reversable TCP

RTCP has identical interfaces and actions compared to `net.TCP`, while with builtin support to a server hiding somewhere with a reverse proxy on public discover point.

## Protocol Description

RTCP builds itself upon TCP. While traditional TCP has two roles (server and client), RTCP has three more besides them:

- *Direct Server* is the traditional TCP server. To make it accessible by everyone, it must have public address (i.e. public IP) in the Internet.
- *Direct Client* is the traditional TCP client. Once you connects to a direct server, you get a direct client.
- *Discover Point (DP)* serves as the discover point and reverse proxy for Indirect Server. An Indirect Client talks to DP in order to communicate with the Indirect Server. DP can be multiplexed by specifying *ISName* of Indirect Server.
- *Indirect Server* is a server who doesn't have public address. To make it accessible by others, it connects to DP to get overlaid connection from indirect client.
- *Indirect Client* is the client for indirect server.

Direct server and client works exactly the same way with traditional TCP. However, indirect ones work quite differently. Generally, all session (TCP connection) starts with exchanging a message containing magic character, version number. If it is not compatible, the connection aborts. Then, indirect C/S talks to DP for the purpose

The connection between DP and indirect server is long-term, which coalesces all indirect sessions from all clients. However, for congestion control purposes, the tcp connection pool can be adujsted.

### Message Type
Messages mostly happens on indirect sessions. Each message starts with a prefix line specifying the type, and then a JSON string line for message body.

#### RMsgAloha (Type = 0)
Message for initializing the connection. Format:

```javascript
{
  "V": 10000,   // a number, specifying the protocol version
  "R": /*TYPE-ROLE*/,
  "N": "WhoIAm",
}

TYPE-ROLE = [RRoleDS = 0, RRoleDC = 1, RRoleIS = 2, RRoleIC = 3, RRoleDP = 4,
             RAbortVersion = 5]
```

#### RMsgConnect (Type = 1)
```javascript
{
  "N": "WhoToConnect" /* Who to connect for IC->DP, Who connects this for DP->IS */
}
```

#### RMsgAck (Type = 2)
```javascript
{
  "X": 0123, /* Session number created by IS, DP will remember it for the session, OR -1 for failure */
  "R": "Failure Reasion"
}
```

#### RMsgData (Type = 3)
```javascript
{
  "X": 0123, /* Session number */
  "C": "Content"
}
```

#### RMsgClose (Type = 4)
```javascript
{
  "X": 0123, /* Session number */
  "R": "Reason for closing"
}
```


### Indirect Session:
```
IS --Aloha--> DP
DP --Aloha--> IS
# Listener Created
IC --Aloha--> DP
DP --Aloha--> IS

IC --Connect--> DP
DP --Connect--> IS
# IS Setting up new session, assign session number
IS --Ack--> DP
DP --Ack--> IC
# Connection established

IS/IC --Data--> DP
DP --Data-->IC/IS

IS/IC --Close--> DP
DP --Close--> IS
DP --Close--> IC
```
