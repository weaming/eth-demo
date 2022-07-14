brew update
brew upgrade
brew tap ethereum/ethereum
brew install ethereum
brew install solidity

solc --abi Storage.sol -o build
abigen --abi build/Storage.abi --pkg main --type Storage --out Storage.go
solc --bin Storage.sol -o Storage.bin
abigen --abi build/Storage.abi --pkg main --type Storage --out Storage.go --bin Storage.bin/Storage.bin

# https://geth.ethereum.org/docs/interface/les
# --syncmode value (default: snap) Blockchain sync mode ("snap", "full" or "light")
geth --goerli --syncmode light --http --http.api "eth,debug"
# /Users/weaming/Library/Ethereum/goerli/geth.ipc

# https://geth.ethereum.org/docs/interface/managing-your-accounts
 geth account import <keyfile>  # 从私钥导入账户，设置密码并生成 keystore 文件