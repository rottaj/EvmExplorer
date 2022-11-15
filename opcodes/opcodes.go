package opcodes

type Opcode struct {
	Op       byte
	Mnemonic string
	//minGas int

}

// Create mapping of mneumonics. Check process step (string) with mapping.

// 0x00 - 0x0B
var (
	STOP       = Opcode{Op: 0x00, Mnemonic: "STOP"}
	ADD        = Opcode{Op: 0x01, Mnemonic: "ADD"}
	MUL        = Opcode{Op: 0x02, Mnemonic: "MUL"}
	SUB        = Opcode{Op: 0x03, Mnemonic: "SUB"}
	DIV        = Opcode{Op: 0x04, Mnemonic: "DIV"}
	SDIV       = Opcode{Op: 0x05, Mnemonic: "SDIV"}
	MOD        = Opcode{Op: 0x06, Mnemonic: "MOD"}
	SMOD       = Opcode{Op: 0x07, Mnemonic: "SMOD"}
	ADDMOD     = Opcode{Op: 0x08, Mnemonic: "ADDMOD"}
	MULMOD     = Opcode{Op: 0x09, Mnemonic: "MULMOD"}
	EXP        = Opcode{Op: 0x0A, Mnemonic: "EXP"}
	SIGNEXTEND = Opcode{Op: 0x0B, Mnemonic: "SIGNEXTEND"}
)

// 0x10 - 0x1D

var (
	LT     = Opcode{Op: 0x10, Mnemonic: "LT"}
	GT     = Opcode{Op: 0x11, Mnemonic: "GT"}
	SLT    = Opcode{Op: 0x12, Mnemonic: "SLT"}
	SGT    = Opcode{Op: 0x13, Mnemonic: "SGT"}
	EQ     = Opcode{Op: 0x14, Mnemonic: "EQ"}
	ISZERO = Opcode{Op: 0x15, Mnemonic: "ISZERO"}
	AND    = Opcode{Op: 0x16, Mnemonic: "AND"}
	OR     = Opcode{Op: 0x17, Mnemonic: "OR"}
	XOR    = Opcode{Op: 0x18, Mnemonic: "XOR"}
	NOT    = Opcode{Op: 0x19, Mnemonic: "NOT"}
	BYTE   = Opcode{Op: 0x1A, Mnemonic: "BYTE"}
	SHL    = Opcode{Op: 0x1B, Mnemonic: "SHL"}
	SHR    = Opcode{Op: 0x1C, Mnemonic: "SHR"}
	SAR    = Opcode{Op: 0x1D, Mnemonic: "SAR"}
)

// 0x20 - 0x3F

var (
	KECCAK256      = Opcode{Op: 0x20, Mnemonic: "KECCAK256"}
	ADDRESS        = Opcode{Op: 0x30, Mnemonic: "ADDRESS"}
	BALANCE        = Opcode{Op: 0x31, Mnemonic: "BALANCE"}
	ORIGIN         = Opcode{Op: 0x32, Mnemonic: "ORIGIN"}
	CALLER         = Opcode{Op: 0x33, Mnemonic: "CALLER"}
	CALLVALUE      = Opcode{Op: 0x34, Mnemonic: "CALLVALUE"}
	CALLDATALOAD   = Opcode{Op: 0x35, Mnemonic: "CALLDATALOAD"}
	CALLDATASIZE   = Opcode{Op: 0x36, Mnemonic: "CALLDATASIZE"}
	CALLDATACOPY   = Opcode{Op: 0x37, Mnemonic: "CALLDATACOPY"}
	CODESIZE       = Opcode{Op: 0x38, Mnemonic: "CODESIZE"}
	CODECOPY       = Opcode{Op: 0x39, Mnemonic: "CODECOPY"}
	GASPRICE       = Opcode{Op: 0x3A, Mnemonic: "GASPRICE"}
	EXTCODESIZE    = Opcode{Op: 0x3B, Mnemonic: "EXTCODESIZE"}
	EXTCODECOPY    = Opcode{Op: 0x3C, Mnemonic: "EXTCODECOPY"}
	RETURNDATASIZE = Opcode{Op: 0x3D, Mnemonic: "RETURNDATASIZE"}
	RETURNDATACOPY = Opcode{Op: 0x3F, Mnemonic: "RETURNDATACOPY"}
)

// 0x40 - 0x5B

var (
	BLOCKHASH   = Opcode{Op: 0x40, Mnemonic: "BLOCKHASH"}
	COINBASE    = Opcode{Op: 0x41, Mnemonic: "COINBASE"}
	TIMESTAMP   = Opcode{Op: 0x42, Mnemonic: "TIMESTAMP"}
	NUMBER      = Opcode{Op: 0x43, Mnemonic: "NUMBER"}
	DIFFICULTY  = Opcode{Op: 0x44, Mnemonic: "DIFFICULTY"}
	GASLIMIT    = Opcode{Op: 0x45, Mnemonic: "GASLIMIT"}
	CHAINID     = Opcode{Op: 0x46, Mnemonic: "CHAINID"}
	SELFBALANCE = Opcode{Op: 0x47, Mnemonic: "SELFBALANCE"}
	BASEFEE     = Opcode{Op: 0x48, Mnemonic: "BASEFEE"}
	POP         = Opcode{Op: 0x50, Mnemonic: "POP"}
	MLOAD       = Opcode{Op: 0x51, Mnemonic: "MLOAD"}
	MSTORE      = Opcode{Op: 0x52, Mnemonic: "MSTORE"}
	MSTORE8     = Opcode{Op: 0x53, Mnemonic: "MSTORE8"}
	SLOAD       = Opcode{Op: 0x54, Mnemonic: "SLOAD"}
	SSTORE      = Opcode{Op: 0x55, Mnemonic: "SSTORE"}
	JUMP        = Opcode{Op: 0x56, Mnemonic: "JUMP"}
	JUMPI       = Opcode{Op: 0x57, Mnemonic: "JUMPI"}
	PC          = Opcode{Op: 0x58, Mnemonic: "PC"}
	MSIZE       = Opcode{Op: 0x59, Mnemonic: "MSIZE"}
	GAS         = Opcode{Op: 0x5A, Mnemonic: "GAS"}
	JUMPDEST    = Opcode{Op: 0x5B, Mnemonic: "JUMPDEST"}
)

// 0x60 - 0x7F (PUSH)

var (
	PUSH1  = Opcode{Op: 0x60, Mnemonic: "PUSH1"}
	PUSH2  = Opcode{Op: 0x61, Mnemonic: "PUSH2"}
	PUSH3  = Opcode{Op: 0x62, Mnemonic: "PUSH3"}
	PUSH4  = Opcode{Op: 0x63, Mnemonic: "PUSH4"}
	PUSH5  = Opcode{Op: 0x64, Mnemonic: "PUSH5"}
	PUSH6  = Opcode{Op: 0x65, Mnemonic: "PUSH6"}
	PUSH7  = Opcode{Op: 0x66, Mnemonic: "PUSH7"}
	PUSH8  = Opcode{Op: 0x67, Mnemonic: "PUSH8"}
	PUSH9  = Opcode{Op: 0x68, Mnemonic: "PUSH9"}
	PUSH10 = Opcode{Op: 0x69, Mnemonic: "PUSH10"}
	PUSH11 = Opcode{Op: 0x6A, Mnemonic: "PUSH11"}
	PUSH12 = Opcode{Op: 0x6B, Mnemonic: "PUSH12"}
	PUSH13 = Opcode{Op: 0x6C, Mnemonic: "PUSH13"}
	PUSH14 = Opcode{Op: 0x6D, Mnemonic: "PUSH14"}
	PUSH15 = Opcode{Op: 0x6E, Mnemonic: "PUSH15"}
	PUSH16 = Opcode{Op: 0x6F, Mnemonic: "PUSH16"}
	PUSH17 = Opcode{Op: 0x70, Mnemonic: "PUSH17"}
	PUSH18 = Opcode{Op: 0x71, Mnemonic: "PUSH18"}
	PUSH19 = Opcode{Op: 0x72, Mnemonic: "PUSH19"}
	PUSH20 = Opcode{Op: 0x73, Mnemonic: "PUSH20"}
	PUSH21 = Opcode{Op: 0x74, Mnemonic: "PUSH21"}
	PUSH22 = Opcode{Op: 0x75, Mnemonic: "PUSH22"}
	PUSH23 = Opcode{Op: 0x76, Mnemonic: "PUSH23"}
	PUSH24 = Opcode{Op: 0x77, Mnemonic: "PUSH24"}
	PUSH25 = Opcode{Op: 0x78, Mnemonic: "PUSH25"}
	PUSH26 = Opcode{Op: 0x79, Mnemonic: "PUSH26"}
	PUSH27 = Opcode{Op: 0x7A, Mnemonic: "PUSH27"}
	PUSH28 = Opcode{Op: 0x7B, Mnemonic: "PUSH28"}
	PUSH29 = Opcode{Op: 0x7C, Mnemonic: "PUSH29"}
	PUSH30 = Opcode{Op: 0x7D, Mnemonic: "PUSH30"}
	PUSH31 = Opcode{Op: 0x7E, Mnemonic: "PUSH31"}
	PUSH32 = Opcode{Op: 0x7F, Mnemonic: "PUSH32"}
)

// 0x80 -

var (
	DUP1  = Opcode{Op: 0x80, Mnemonic: "DUP1"}
	DUP2  = Opcode{Op: 0x81, Mnemonic: "DUP2"}
	DUP3  = Opcode{Op: 0x82, Mnemonic: "DUP3"}
	DUP4  = Opcode{Op: 0x83, Mnemonic: "DUP4"}
	DUP5  = Opcode{Op: 0x84, Mnemonic: "DUP5"}
	DUP6  = Opcode{Op: 0x85, Mnemonic: "DUP6"}
	DUP7  = Opcode{Op: 0x86, Mnemonic: "DUP7"}
	DUP8  = Opcode{Op: 0x87, Mnemonic: "DUP8"}
	DUP9  = Opcode{Op: 0x88, Mnemonic: "DUP9"}
	DUP10 = Opcode{Op: 0x89, Mnemonic: "DUP10"}
	DUP11 = Opcode{Op: 0x8A, Mnemonic: "DUP11"}
	DUP12 = Opcode{Op: 0x8B, Mnemonic: "DUP12"}
	DUP13 = Opcode{Op: 0x8C, Mnemonic: "DUP13"}
	DUP14 = Opcode{Op: 0x8D, Mnemonic: "DUP14"}
	DUP15 = Opcode{Op: 0x8E, Mnemonic: "DUP15"}
	DUP16 = Opcode{Op: 0x8F, Mnemonic: "DUP16"}
)

var (
	SWAP1  = Opcode{Op: 0x90, Mnemonic: "SWAP1"}
	SWAP2  = Opcode{Op: 0x91, Mnemonic: "SWAP2"}
	SWAP3  = Opcode{Op: 0x92, Mnemonic: "SWAP3"}
	SWAP4  = Opcode{Op: 0x93, Mnemonic: "SWAP4"}
	SWAP5  = Opcode{Op: 0x94, Mnemonic: "SWAP5"}
	SWAP6  = Opcode{Op: 0x95, Mnemonic: "SWAP6"}
	SWAP7  = Opcode{Op: 0x96, Mnemonic: "SWAP7"}
	SWAP8  = Opcode{Op: 0x97, Mnemonic: "SWAP8"}
	SWAP9  = Opcode{Op: 0x98, Mnemonic: "SWAP9"}
	SWAP10 = Opcode{Op: 0x99, Mnemonic: "SWAP10"}
	SWAP11 = Opcode{Op: 0x9A, Mnemonic: "SWAP11"}
	SWAP12 = Opcode{Op: 0x9B, Mnemonic: "SWAP12"}
	SWAP13 = Opcode{Op: 0x9C, Mnemonic: "SWAP13"}
	SWAP14 = Opcode{Op: 0x9D, Mnemonic: "SWAP14"}
	SWAP15 = Opcode{Op: 0x9E, Mnemonic: "SWAP15"}
	SWAP16 = Opcode{Op: 0x9F, Mnemonic: "SWAP16"}
)

var (
	LOG0 = Opcode{Op: 0xA0, Mnemonic: "LOG0"}
	LOG1 = Opcode{Op: 0xA1, Mnemonic: "LOG1"}
	LOG2 = Opcode{Op: 0xA2, Mnemonic: "LOG2"}
	LOG3 = Opcode{Op: 0xA2, Mnemonic: "LOG3"}
	LOG4 = Opcode{Op: 0xA2, Mnemonic: "LOG4"}
)

var (
	CREATE       = Opcode{Op: 0xF0, Mnemonic: "CREATE"}
	CALL         = Opcode{Op: 0xF1, Mnemonic: "CALL"}
	CALLCODE     = Opcode{Op: 0xF2, Mnemonic: "CALLCODE"}
	RETURN       = Opcode{Op: 0xF3, Mnemonic: "RETURN"}
	DELEGATECALL = Opcode{Op: 0xF4, Mnemonic: "DELEGATECALL"}
	CREATE2      = Opcode{Op: 0xF5, Mnemonic: "CREATE2"}
	STATICCALL   = Opcode{Op: 0xFA, Mnemonic: "STATICCALL"}
	REVERT       = Opcode{Op: 0xFD, Mnemonic: "REVERT"}
	INVALID      = Opcode{Op: 0xFE, Mnemonic: "INVALID"}
	SELFDESTRUCT = Opcode{Op: 0xFF, Mnemonic: "SELFDESTRUCT"}
)

// Create mapping of strings -> structs
var StringToOpcode = map[string]Opcode{
	"STOP":           STOP,
	"ADD":            ADD,
	"MUL":            MUL,
	"SUB":            SUB,
	"DIV":            DIV,
	"SDIV":           SDIV,
	"MOD":            MOD,
	"SMOD":           SMOD,
	"ADDMOD":         ADDMOD,
	"MULMOD":         MULMOD,
	"EXP":            EXP,
	"SIGNEXTEND":     SIGNEXTEND,
	"LT":             LT,
	"GT":             GT,
	"SLT":            SLT,
	"SGT":            SGT,
	"EQ":             EQ,
	"ISZERO":         ISZERO,
	"AND":            AND,
	"OR":             OR,
	"XOR":            XOR,
	"NOT":            NOT,
	"BYTE":           BYTE,
	"SHL":            SHL,
	"SHR":            SHR,
	"SAR":            SAR,
	"SHA3":           KECCAK256, // Add SHA3 for legacy Opcodes
	"KECCAK256":      KECCAK256,
	"ADDRESS":        ADDRESS,
	"BALANCE":        BALANCE,
	"ORIGIN":         ORIGIN,
	"CALLER":         CALLER,
	"CALLVALUE":      CALLVALUE,
	"CALLDATALOAD":   CALLDATALOAD,
	"CALLDATASIZE":   CALLDATASIZE,
	"CALLDATACOPY":   CALLDATACOPY,
	"CODESIZE":       CODESIZE,
	"CODECOPY":       CODECOPY,
	"GASPRICE":       GASPRICE,
	"EXTCODESIZE":    EXTCODESIZE,
	"EXTCODECOPY":    EXTCODECOPY,
	"RETURNDATASIZE": RETURNDATASIZE,
	"RETURNDATACOPY": RETURNDATACOPY,
	"BLOCKHASH":      BLOCKHASH,
	"COINBASE":       COINBASE,
	"TIMESTAMP":      TIMESTAMP,
	"NUMBER":         NUMBER,
	"DIFFICULTY":     DIFFICULTY,
	"GASLIMIT":       GASLIMIT,
	"CHAINID":        CHAINID,
	"SELFBALANCE":    SELFBALANCE,
	"BASEFEE":        BASEFEE,
	"POP":            POP,
	"MLOAD":          MLOAD,
	"MSTORE":         MSTORE,
	"MSTORE8":        MSTORE8,
	"SLOAD":          SLOAD,
	"SSTORE":         SSTORE,
	"JUMP":           JUMP,
	"JUMPI":          JUMPI,
	"PC":             PC,
	"MSIZE":          MSIZE,
	"GAS":            GAS,
	"JUMPDEST":       JUMPDEST,
	"PUSH1":          PUSH1,
	"PUSH2":          PUSH2,
	"PUSH3":          PUSH3,
	"PUSH4":          PUSH4,
	"PUSH5":          PUSH5,
	"PUSH6":          PUSH6,
	"PUSH7":          PUSH7,
	"PUSH8":          PUSH8,
	"PUSH9":          PUSH9,
	"PUSH10":         PUSH10,
	"PUSH11":         PUSH11,
	"PUSH12":         PUSH12,
	"PUSH13":         PUSH13,
	"PUSH14":         PUSH14,
	"PUSH15":         PUSH15,
	"PUSH16":         PUSH16,
	"PUSH17":         PUSH17,
	"PUSH18":         PUSH18,
	"PUSH19":         PUSH19,
	"PUSH20":         PUSH20,
	"PUSH21":         PUSH21,
	"PUSH22":         PUSH22,
	"PUSH23":         PUSH23,
	"PUSH24":         PUSH24,
	"PUSH25":         PUSH25,
	"PUSH26":         PUSH26,
	"PUSH27":         PUSH27,
	"PUSH28":         PUSH28,
	"PUSH29":         PUSH29,
	"PUSH30":         PUSH30,
	"PUSH31":         PUSH31,
	"PUSH32":         PUSH32,
	"DUP1":           DUP1,
	"DUP2":           DUP2,
	"DUP3":           DUP3,
	"DUP4":           DUP4,
	"DUP5":           DUP5,
	"DUP6":           DUP6,
	"DUP7":           DUP7,
	"DUP8":           DUP8,
	"DUP9":           DUP9,
	"DUP10":          DUP10,
	"DUP11":          DUP11,
	"DUP12":          DUP12,
	"DUP13":          DUP13,
	"DUP14":          DUP14,
	"DUP15":          DUP15,
	"DUP16":          DUP16,
	"SWAP1":          SWAP1,
	"SWAP2":          SWAP2,
	"SWAP3":          SWAP3,
	"SWAP4":          SWAP4,
	"SWAP5":          SWAP5,
	"SWAP6":          SWAP6,
	"SWAP7":          SWAP7,
	"SWAP8":          SWAP8,
	"SWAP9":          SWAP9,
	"SWAP10":         SWAP10,
	"SWAP11":         SWAP11,
	"SWAP12":         SWAP12,
	"SWAP13":         SWAP13,
	"SWAP14":         SWAP14,
	"SWAP15":         SWAP15,
	"SWAP16":         SWAP16,
	"LOG0":           LOG0,
	"LOG1":           LOG1,
	"LOG2":           LOG2,
	"LOG3":           LOG3,
	"LOG4":           LOG4,
	"CREATE":         CREATE,
	"CALL":           CALL,
	"CALLCODE":       CALLCODE,
	"RETURN":         RETURN,
	"DELEGATECALL":   DELEGATECALL,
	"CREATE2":        CREATE2,
	"STATICCALL":     STATICCALL,
	"REVERT":         REVERT,
	"INVALID":        INVALID,
	"SELFDESTRUCT":   SELFDESTRUCT,
}
