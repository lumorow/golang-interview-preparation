# Базовые знания по сетям

## TCP and UDP. In which case is UDP preferable?
***
`TCP` is a transport protocol for data transfer in `TCP/IP` networks, which preliminarily establishes a connection to the network.
`UDP` is a transport protocol that transmits datagram messages without the need to establish a connection in the `IP` network.

The difference between the `TCP` and `UDP` protocols is the so-called “delivery guarantee”. `TCP` requires a response from the client,
to whom the data packet is delivered, confirmation of delivery, and for this it needs a pre-established connection.
Also, the `TCP` protocol is considered reliable, while `UDP` has even received the name “unreliable datagram protocol”.
`TCP` eliminates data loss, duplication and mixing of packets, and delays. `UDP` allows all this, and the connection
it doesn't need it to work. Processes to which data is sent via `UDP` must make do with what they receive,
even with losses. `TCP` controls the congestion of the connection, `UDP` does not control anything except the integrity of received
datagram.
On the other hand, thanks to such non-selectivity and lack of control, `UDP` delivers data packets (datagrams)
much faster, therefore for applications that are designed for high bandwidth and fast exchange,
`UDP` can be considered the optimal protocol. These include online and browser games, as well as viewing programs
streaming video and applications for video communication (or voice): the loss of a packet, complete or partial, does not change anything,
It is not necessary to repeat the request, but the download is much faster. The `TCP` protocol, as more reliable,
successfully used even in email programs, allowing you to control not only traffic, but also the length of the message
and traffic exchange speed.

## What is NAT?
***
Networks are usually designed using private `IP` addresses.
These addresses are `10.0.0.0/8`, `172.16.0.0/12` and `192.168.0.0/16`. These private addresses are used within an organization or site,
to allow devices to communicate locally and not be routed over the internet. To allow the device to
private `IPv4` address to access devices and resources outside the local network, the private address must first
be transferred to a publicly accessible address.

And just `NAT` converts private addresses into public ones. This allows a device with a private `IPv4` address
access resources outside of his private network. `NAT` in combination with `IPv4` private addresses has proven useful
method of storing public `IPv4` addresses. One public `IPv4` address can be used by hundreds, even
thousands of devices, each with a private `IPv4` address. `NAT` has the added benefit of
in adding a degree of privacy and security to the network as it hides internal `IPv4` addresses from external ones
networks.

## What are HTTP and HTTPS, what are their differences?
***
`HTTP` (from the English HyperText Transfer Protocol - hypertext transfer protocol) is an application data transfer protocol
online. Currently used to obtain information from websites. The `HTTP` protocol is based on the use
“client-server” technologies: the client sending the request is the initiator of the connection; the server receiving the request
executes it and sends the result to the client.

`HTTPS` (from the English HyperText Transfer Protocol Secure - secure hypertext transfer protocol) is an extension
`HTTP` protocol, supporting encryption using `SSL` and `TLS` cryptographic protocols.
### What is the difference between HTTP and HTTPS:

`HTTPS` is not a separate data transfer protocol, but is an extension of the `HTTP` protocol with an encryption add-on;
data transmitted via the `HTTP` protocol is not protected, `HTTPS` ensures the confidentiality of information by encrypting it;
`HTTP` uses port 80, `HTTPS` uses port 443.

## What are SSL and TLS, are there any differences between them?
***
`SSL` or Secure Sockets Layer was the original name of the protocol developed by `Netscape`
in the mid-90s. SSL 1.0 was never publicly available, and version 2.0 had serious flaws.
The `SSL 3.0` protocol, released in 1996, was completely redesigned and set the tone for the next stage of development.

When the next version of the protocol was released in 1999, it was standardized by a special design working group
network of the Internet and gave it a new name: transport layer security, or `TLS`. As the `TLS` documentation says,
“the difference between this protocol and S`SL 3.0` is not critical.” `TLS` and `SSL` form a constantly updated series of protocols,
and they are often combined under the name `SSL`/`TLS`.
The `TLS` protocol encrypts Internet traffic of any kind. The most common type is web traffic. You know when your browser establishes a TLS connection -
if the link in the address bar starts with "https".
`TLS` is also used by other applications, such as mail and teleconferencing systems.
The most secure encryption method is asymmetric encryption. This requires 2 keys, 1 public and 1 private. These are files with information,
most often very large numbers. The mechanism is complex, but simply put, you can use a public key to encrypt data, but you need a private key to decrypt it.
The two keys are linked using a complex mathematical formula that is difficult to hack.
Because asymmetric encryption uses complex mathematical calculations, it requires a lot of computing resources. `TLS` solves this problem
by using asymmetric encryption only at the beginning of the session to encrypt communications between the server and client.
The server and client must agree on a single session key, which the two of them will use to encrypt data packets.

## Resources
***
- [How the Web Works: A Primer for Newcomers to Web Development (or anyone, really)](https://www.freecodecamp.org/news/how-the-web-works-a-primer-for-newcomers-to-web-development-or-anyone-really-b4584e63585c#.7l3tokoh1)
- [How the Web Works Part II: Client-Server Model & the Structure of a Web Application](https://medium.com/free-code-camp/how-the-web-works-part-ii-client-server-model-the-structure-of-a-web-application-735b4b6d76e3#.e6tmj8112)

## README.md
***
- eng [English](https://github.com/lumorow/golang-interview-preparation/blob/main/Networking/README.md)
- ru [Русский](https://github.com/lumorow/golang-interview-preparation/blob/main/Networking/readme/README.ru.md)