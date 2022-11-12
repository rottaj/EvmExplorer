package opcodes

type Opcode struct {
	op       byte
	mnemonic string
	//minGas int

}

// Create mapping of mneumonics. Check process step (string) with mapping.

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

// 0x60 - 0x7F (PUSH)

var (
	PUSH1  = Opcode{op: 0x60, mnemonic: "PUSH1"}
	PUSH2  = Opcode{op: 0x61, mnemonic: "PUSH2"}
	PUSH3  = Opcode{op: 0x62, mnemonic: "PUSH3"}
	PUSH4  = Opcode{op: 0x63, mnemonic: "PUSH4"}
	PUSH5  = Opcode{op: 0x64, mnemonic: "PUSH5"}
	PUSH6  = Opcode{op: 0x65, mnemonic: "PUSH6"}
	PUSH7  = Opcode{op: 0x66, mnemonic: "PUSH7"}
	PUSH8  = Opcode{op: 0x67, mnemonic: "PUSH8"}
	PUSH9  = Opcode{op: 0x68, mnemonic: "PUSH9"}
	PUSH10 = Opcode{op: 0x69, mnemonic: "PUSH10"}
	PUSH11 = Opcode{op: 0x6A, mnemonic: "PUSH11"}
	PUSH12 = Opcode{op: 0x6B, mnemonic: "PUSH12"}
	PUSH13 = Opcode{op: 0x6C, mnemonic: "PUSH13"}
	PUSH14 = Opcode{op: 0x6D, mnemonic: "PUSH14"}
	PUSH15 = Opcode{op: 0x6E, mnemonic: "PUSH15"}
	PUSH16 = Opcode{op: 0x6F, mnemonic: "PUSH16"}
	PUSH17 = Opcode{op: 0x70, mnemonic: "PUSH17"}
	PUSH18 = Opcode{op: 0x71, mnemonic: "PUSH18"}
	PUSH19 = Opcode{op: 0x72, mnemonic: "PUSH19"}
	PUSH20 = Opcode{op: 0x73, mnemonic: "PUSH20"}
	PUSH21 = Opcode{op: 0x74, mnemonic: "PUSH21"}
	PUSH22 = Opcode{op: 0x75, mnemonic: "PUSH22"}
	PUSH23 = Opcode{op: 0x76, mnemonic: "PUSH23"}
	PUSH24 = Opcode{op: 0x77, mnemonic: "PUSH24"}
	PUSH25 = Opcode{op: 0x78, mnemonic: "PUSH25"}
	PUSH26 = Opcode{op: 0x79, mnemonic: "PUSH26"}
	PUSH27 = Opcode{op: 0x7A, mnemonic: "PUSH27"}
	PUSH28 = Opcode{op: 0x7B, mnemonic: "PUSH28"}
	PUSH29 = Opcode{op: 0x7C, mnemonic: "PUSH29"}
	PUSH30 = Opcode{op: 0x7D, mnemonic: "PUSH30"}
	PUSH31 = Opcode{op: 0x7E, mnemonic: "PUSH31"}
	PUSH32 = Opcode{op: 0x7F, mnemonic: "PUSH32"}
)

// 0x80 -

var (
	DUP1  = Opcode{op: 0x80, mnemonic: "DUP1"}
	DUP2  = Opcode{op: 0x81, mnemonic: "DUP2"}
	DUP3  = Opcode{op: 0x82, mnemonic: "DUP3"}
	DUP4  = Opcode{op: 0x83, mnemonic: "DUP4"}
	DUP5  = Opcode{op: 0x84, mnemonic: "DUP5"}
	DUP6  = Opcode{op: 0x85, mnemonic: "DUP6"}
	DUP7  = Opcode{op: 0x86, mnemonic: "DUP7"}
	DUP8  = Opcode{op: 0x87, mnemonic: "DUP8"}
	DUP9  = Opcode{op: 0x88, mnemonic: "DUP9"}
	DUP10 = Opcode{op: 0x89, mnemonic: "DUP10"}
	DUP11 = Opcode{op: 0x8A, mnemonic: "DUP11"}
	DUP12 = Opcode{op: 0x8B, mnemonic: "DUP12"}
	DUP13 = Opcode{op: 0x8C, mnemonic: "DUP13"}
	DUP14 = Opcode{op: 0x8D, mnemonic: "DUP14"}
	DUP15 = Opcode{op: 0x8E, mnemonic: "DUP15"}
	DUP16 = Opcode{op: 0x8F, mnemonic: "DUP16"}
)

var (
	SWAP1  = Opcode{op: 0x90, mnemonic: "SWAP1"}
	SWAP2  = Opcode{op: 0x91, mnemonic: "SWAP2"}
	SWAP3  = Opcode{op: 0x92, mnemonic: "SWAP3"}
	SWAP4  = Opcode{op: 0x93, mnemonic: "SWAP4"}
	SWAP5  = Opcode{op: 0x94, mnemonic: "SWAP5"}
	SWAP6  = Opcode{op: 0x95, mnemonic: "SWAP6"}
	SWAP7  = Opcode{op: 0x96, mnemonic: "SWAP7"}
	SWAP8  = Opcode{op: 0x97, mnemonic: "SWAP8"}
	SWAP9  = Opcode{op: 0x98, mnemonic: "SWAP9"}
	SWAP10 = Opcode{op: 0x99, mnemonic: "SWAP10"}
	SWAP11 = Opcode{op: 0x9A, mnemonic: "SWAP11"}
	SWAP12 = Opcode{op: 0x9B, mnemonic: "SWAP12"}
	SWAP13 = Opcode{op: 0x9C, mnemonic: "SWAP13"}
	SWAP14 = Opcode{op: 0x9D, mnemonic: "SWAP14"}
	SWAP15 = Opcode{op: 0x9E, mnemonic: "SWAP15"}
	SWAP16 = Opcode{op: 0x9F, mnemonic: "SWAP16"}
)

var (
	LOGO  = Opcode{op: 0xA0, mnemonic: "LOGO"}
	LOGO1 = Opcode{op: 0xA1, mnemonic: "LOGO1"}
	LOGO2 = Opcode{op: 0xA2, mnemonic: "LOGO2"}
	LOGO3 = Opcode{op: 0xA2, mnemonic: "LOGO3"}
	LOGO4 = Opcode{op: 0xA2, mnemonic: "LOGO4"}
)

var (
	CREATE       = Opcode{op: 0xF0, mnemonic: "CREATE"}
	CALL         = Opcode{op: 0xF1, mnemonic: "CALL"}
	CALLCODE     = Opcode{op: 0xF2, mnemonic: "CALLCODE"}
	RETURN       = Opcode{op: 0xF3, mnemonic: "RETURN"}
	DELEGATECALL = Opcode{op: 0xF4, mnemonic: "DELEGATECALL"}
	CREATE2      = Opcode{op: 0xF5, mnemonic: "CREATE2"}
	STATICCALL   = Opcode{op: 0xFA, mnemonic: "STATICCALL"}
	REVERT       = Opcode{op: 0xFD, mnemonic: "REVERT"}
	INVALID      = Opcode{op: 0xFE, mnemonic: "INVALID"}
	SELFDESTRUCT = Opcode{op: 0xFF, mnemonic: "SELFDESTRUCT"}
)
