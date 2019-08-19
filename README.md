# smartContract
smartcontract on ethereum



abigen --abi ProtonManager.abi --pkg service --type ProtonManager --out ProtonManager.go
abigen --abi PangolinManager.abi --pkg ethInterface --type PangolinManager --out PangolinManager.go
abigen --abi SimpleProtonManager.abi --pkg service --type SimpleProtonManager --out SimpleProtonManager.go
abigen --abi PPNAdmin.abi --pkg eth --type PPNAdmin --out PPNAdmin.go
abigen --abi PPNToken.abi --pkg eth --type PPNToken --out PPNToken.go