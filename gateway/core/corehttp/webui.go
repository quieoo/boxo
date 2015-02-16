package corehttp

// TODO: move to IPNS
const WebUIPath = "/ipfs/QmXdu7HWdV6CUaUabd9q2ZeA4iHZLVyDRj3Gi4dsJsWjbr"

// this is a list of all past webUI paths.
var WebUIPaths = []string{
	WebUIPath,
	"/ipfs/QmaaqrHyAQm7gALkRW8DcfGX3u8q9rWKnxEMmf7m9z515w",
	"/ipfs/QmSHDxWsMPuJQKWmVA1rB5a3NX2Eme5fPqNb63qwaqiqSp",
	"/ipfs/QmctngrQAt9fjpQUZr7Bx3BsXUcif52eZGTizWhvcShsjz",
}

var WebUIOption = RedirectOption("webui", WebUIPath)
