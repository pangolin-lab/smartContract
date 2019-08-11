# smartContract
smartcontract on ethereum



abigen --abi ProtonManager.abi --pkg service --type ProtonManager --out ProtonManager.go
abigen --abi PangolinManager.abi --pkg ethInterface --type PangolinManager --out PangolinManager.go
abigen --abi SimpleProtonManager.abi --pkg service --type SimpleProtonManager --out SimpleProtonManager.go