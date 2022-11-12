package opcodes

type Opcode struct {
	op       byte
	mnemonic string
	//minGas int

}

// 0x00 - 0x0B
var (
	STOP       = Opcode{op: 0x00, mnemonic: "STOP"}
	ADD        = Opcode{op: 0x01, mnemonic: "ADD"}
	MUL        = Opcode{op: 0x02, mnemonic: "MUL"}
	SUB        = Opcode{op: 0x03, mnemonic: "SUB"}
	DIV        = Opcode{op: 0x04, mnemonic: "DIV"}
	SDIV       = Opcode{op: 0x05, mnemonic: "SDIV"}
	MOD        = Opcode{op: 0x06, mnemonic: "MOD"}
	SMOD       = Opcode{op: 0x07, mnemonic: "SMOD"}
	ADDMOD     = Opcode{op: 0x08, mnemonic: "ADDMOD"}
	MULMOD     = Opcode{op: 0x09, mnemonic: "MULMOD"}
	EXP        = Opcode{op: 0x0A, mnemonic: "EXP"}
	SIGNEXTEND = Opcode{op: 0x0B, mnemonic: "SIGNEXTEND"}
)

// 0x10 - 0x1D

var (
	LT     = Opcode{op: 0x10, mnemonic: "LT"}
	GT     = Opcode{op: 0x11, mnemonic: "GT"}
	SLT    = Opcode{op: 0x12, mnemonic: "SLT"}
	SGT    = Opcode{op: 0x13, mnemonic: "SGT"}
	EQ     = Opcode{op: 0x14, mnemonic: "EQ"}
	ISZERO = Opcode{op: 0x15, mnemonic: "ISZERO"}
	AND    = Opcode{op: 0x16, mnemonic: "AND"}
	OR     = Opcode{op: 0x17, mnemonic: "OR"}
	XOR    = Opcode{op: 0x18, mnemonic: "XOR"}
	NOT    = Opcode{op: 0x19, mnemonic: "NOT"}
	BYTE   = Opcode{op: 0x1A, mnemonic: "BYTE"}
	SHL    = Opcode{op: 0x1B, mnemonic: "SHL"}
	SHR    = Opcode{op: 0x1C, mnemonic: "SHR"}
	SAR    = Opcode{op: 0x1D, mnemonic: "SAR"}
)

// 0x20 - 0x3F

var (
	SHA3           = Opcode{op: 0x20, mnemonic: "SHA3"}
	ADDRESS        = Opcode{op: 0x30, mnemonic: "ADDRESS"}
	BALANCE        = Opcode{op: 0x31, mnemonic: "BALANCE"}
	ORIGIN         = Opcode{op: 0x32, mnemonic: "ORIGIN"}
	CALLER         = Opcode{op: 0x33, mnemonic: "CALLER"}
	CALLVALUE      = Opcode{op: 0x34, mnemonic: "CALLVALUE"}
	CALLDATALOAD   = Opcode{op: 0x35, mnemonic: "CALLDATALOAD"}
	CALLDATASIZE   = Opcode{op: 0x36, mnemonic: "CALLDATASIZE"}
	CALLDATACOPY   = Opcode{op: 0x37, mnemonic: "CALLDATACOPY"}
	CODESIZE       = Opcode{op: 0x38, mnemonic: "CODESIZE"}
	CODECOPY       = Opcode{op: 0x39, mnemonic: "CODECOPY"}
	GASPRICE       = Opcode{op: 0x3A, mnemonic: "GASPRICE"}
	EXTCODESIZE    = Opcode{op: 0x3B, mnemonic: "EXTCODESIZE"}
	EXTCODECOPY    = Opcode{op: 0x3C, mnemonic: "EXTCODECOPY"}
	RETURNDATASIZE = Opcode{op: 0x3D, mnemonic: "RETURNDATASIZE"}
	RETURNDATACOPY = Opcode{op: 0x3F, mnemonic: "RETURNDATACOPY"}
)

// 0x40 - 0x5B

var (
	BLOCKHASH   = Opcode{op: 0x40, mnemonic: "BLOCKHASH"}
	COINBASE    = Opcode{op: 0x41, mnemonic: "COINBASE"}
	TIMESTAMP   = Opcode{op: 0x42, mnemonic: "TIMESTAMP"}
	NUMBER      = Opcode{op: 0x43, mnemonic: "NUMBER"}
	DIFFICULTY  = Opcode{op: 0x44, mnemonic: "DIFFICULTY"}
	GASLIMIT    = Opcode{op: 0x45, mnemonic: "GASLIMIT"}
	CHAINID     = Opcode{op: 0x46, mnemonic: "CHAINID"}
	SELFBALANCE = Opcode{op: 0x47, mnemonic: "SELFBALANCE"}
	BASEFEE     = Opcode{op: 0x48, mnemonic: "BASEFEE"}
	POP         = Opcode{op: 0x50, mnemonic: "POP"}
	MLOAD       = Opcode{op: 0x51, mnemonic: "MLOAD"}
	MSTORE      = Opcode{op: 0x52, mnemonic: "MSTORE"}
	MSTORE8     = Opcode{op: 0x53, mnemonic: "MSTORE8"}
	SLOAD       = Opcode{op: 0x54, mnemonic: "SLOAD"}
	SSTORE      = Opcode{op: 0x55, mnemonic: "SSTORE"}
	JUMP        = Opcode{op: 0x56, mnemonic: "JUMP"}
	JUMPI       = Opcode{op: 0x57, mnemonic: "JUMPI"}
	PC          = Opcode{op: 0x58, mnemonic: "PC"}
	MSIZE       = Opcode{op: 0x59, mnemonic: "MSIZE"}
	GAS         = Opcode{op: 0x5A, mnemonic: "GAS"}
	JUMPDEST    = Opcode{op: 0x5B, mnemonic: "JUMPDEST"}
)
