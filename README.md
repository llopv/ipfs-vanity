# ipfs-vanity

Before looking for a vanity address, you may want to look to how to [attach a
DNS record to IPNS address](https://github.com/ipfs/examples/tree/master/examples/websites).

**WARNING:** This is experimental software and it is not optimized. Use it at
your own risk.

ipfs-vanity lets you create vanity keyspairs for IPFS. You can customize the
first letters of your IPNS address.

When run, it will write the bytes of the serialized private key to stdout.

Use the `-name` option to specify the letters you want after the `Qm` part.

By default, a 2048 bit RSA key will be generated. The keysize can be changed by
specifying the `-bitsize` option, and the key type can be changed by specifying
the `-type` option (currently only RSA is implemented).

## Installation
```
$ go get github.com/llopv/ipfs-vanity
```

## Usage
```
$ ipfs-vanity -name IPFS > my.key
Generating some 4096 bit RSA keys...
.................................................. 500
.................................................. 1000
.................................................. 1500
.................................................. 2000
.................................................. 2500
.................................................. 3000
...
Success!
ID for generated key: QmIPFSbevfCChmpuhbNMtnnghiQozMfd4ZL2grNu996p4s
```
