
// list peers
function lsp() {
    admin.peers.forEach(function(x){console.log(x.enode)})
}

// list accounts
function lsa() {
    personal.listAccounts.forEach(function(x){console.log(x, eth.getBalance(x)/1e18)})
}

// unlock account
function unlock(x) {
    personal.unlockAccount(x)
}

// transfer
function transfer(from, to, value) {
    tx={"from": from, "to": to, "value": value*1e18 }
    eth.sendTransaction(tx)
}
