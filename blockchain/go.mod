module github.com/picfight/pfcd/blockchain

require (
	github.com/picfight/pfcd/blockchain/stake v1.0.1
	github.com/picfight/pfcd/chaincfg v1.1.1
	github.com/picfight/pfcd/chaincfg/chainhash v1.0.1
	github.com/picfight/pfcd/database v1.0.1
	github.com/picfight/pfcd/pfcec v0.0.0-20180801202239-0761de129164
	github.com/picfight/pfcd/pfcec/edwards v0.0.0-20180721031028-5369a485acf6
	github.com/picfight/pfcd/pfcec/secp256k1 v1.0.0
	github.com/picfight/pfcd/pfcjson v1.0.0
	github.com/picfight/pfcd/pfcutil v1.1.1
	github.com/picfight/pfcd/gcs v1.0.1
	github.com/picfight/pfcd/txscript v1.0.1
	github.com/picfight/pfcd/wire v1.1.0
	github.com/decred/slog v1.0.0
)

replace (
	github.com/picfight/pfcd/blockchain/stake => ./stake
	github.com/picfight/pfcd/chaincfg => ../chaincfg
	github.com/picfight/pfcd/chaincfg/chainhash => ../chaincfg/chainhash
	github.com/picfight/pfcd/database => ../database
	github.com/picfight/pfcd/pfcec => ../pfcec
	github.com/picfight/pfcd/pfcec/edwards => ../pfcec/edwards
	github.com/picfight/pfcd/pfcec/secp256k1 => ../pfcec/secp256k1
	github.com/picfight/pfcd/pfcjson => ../pfcjson
	github.com/picfight/pfcd/pfcutil => ../pfcutil
	github.com/picfight/pfcd/gcs => ../gcs
	github.com/picfight/pfcd/txscript => ../txscript
	github.com/picfight/pfcd/wire => ../wire
)
