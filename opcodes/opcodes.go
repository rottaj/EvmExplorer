package opcodes

type Opcode struct {
	Op        byte
	Mnemonic  string
	StaticGas int
	// dynamicGas func()?
	// If dynamicGas exists: use map StringToDynamicGasFunc to get function
	// Just like using StringToOpcode
}

// Create mapping of mneumonics. Check process step (string) with mapping.

// 0x00 - 0x0B
var (
	STOP       = Opcode{Op: 0x00, Mnemonic: "STOP", StaticGas: 0} // Stops execution & refunds all remaining gas
	ADD        = Opcode{Op: 0x01, Mnemonic: "ADD", StaticGas: 3}
	MUL        = Opcode{Op: 0x02, Mnemonic: "MUL", StaticGas: 5}
	SUB        = Opcode{Op: 0x03, Mnemonic: "SUB", StaticGas: 3}
	DIV        = Opcode{Op: 0x04, Mnemonic: "DIV", StaticGas: 5}
	SDIV       = Opcode{Op: 0x05, Mnemonic: "SDIV", StaticGas: 5}
	MOD        = Opcode{Op: 0x06, Mnemonic: "MOD", StaticGas: 5}
	SMOD       = Opcode{Op: 0x07, Mnemonic: "SMOD", StaticGas: 5}
	ADDMOD     = Opcode{Op: 0x08, Mnemonic: "ADDMOD", StaticGas: 8}
	MULMOD     = Opcode{Op: 0x09, Mnemonic: "MULMOD", StaticGas: 8}
	EXP        = Opcode{Op: 0x0A, Mnemonic: "EXP", StaticGas: 10}
	SIGNEXTEND = Opcode{Op: 0x0B, Mnemonic: "SIGNEXTEND", StaticGas: 5}
)

// 0x10 - 0x1D

var (
	LT     = Opcode{Op: 0x10, Mnemonic: "LT", StaticGas: 3}
	GT     = Opcode{Op: 0x11, Mnemonic: "GT", StaticGas: 3}
	SLT    = Opcode{Op: 0x12, Mnemonic: "SLT", StaticGas: 3}
	SGT    = Opcode{Op: 0x13, Mnemonic: "SGT", StaticGas: 3}
	EQ     = Opcode{Op: 0x14, Mnemonic: "EQ", StaticGas: 3}
	ISZERO = Opcode{Op: 0x15, Mnemonic: "ISZERO", StaticGas: 3}
	AND    = Opcode{Op: 0x16, Mnemonic: "AND", StaticGas: 3}
	OR     = Opcode{Op: 0x17, Mnemonic: "OR", StaticGas: 3}
	XOR    = Opcode{Op: 0x18, Mnemonic: "XOR", StaticGas: 3}
	NOT    = Opcode{Op: 0x19, Mnemonic: "NOT", StaticGas: 3}
	BYTE   = Opcode{Op: 0x1A, Mnemonic: "BYTE", StaticGas: 3}
	SHL    = Opcode{Op: 0x1B, Mnemonic: "SHL", StaticGas: 3}
	SHR    = Opcode{Op: 0x1C, Mnemonic: "SHR", StaticGas: 3}
	SAR    = Opcode{Op: 0x1D, Mnemonic: "SAR", StaticGas: 3}
)

// 0x20 - 0x3F

var (
	KECCAK256      = Opcode{Op: 0x20, Mnemonic: "KECCAK256", StaticGas: 30}
	ADDRESS        = Opcode{Op: 0x30, Mnemonic: "ADDRESS", StaticGas: 2}
	BALANCE        = Opcode{Op: 0x31, Mnemonic: "BALANCE", StaticGas: 100}
	ORIGIN         = Opcode{Op: 0x32, Mnemonic: "ORIGIN", StaticGas: 2}
	CALLER         = Opcode{Op: 0x33, Mnemonic: "CALLER", StaticGas: 2}
	CALLVALUE      = Opcode{Op: 0x34, Mnemonic: "CALLVALUE", StaticGas: 2}
	CALLDATALOAD   = Opcode{Op: 0x35, Mnemonic: "CALLDATALOAD", StaticGas: 3}
	CALLDATASIZE   = Opcode{Op: 0x36, Mnemonic: "CALLDATASIZE", StaticGas: 2}
	CALLDATACOPY   = Opcode{Op: 0x37, Mnemonic: "CALLDATACOPY", StaticGas: 3}
	CODESIZE       = Opcode{Op: 0x38, Mnemonic: "CODESIZE", StaticGas: 2}
	CODECOPY       = Opcode{Op: 0x39, Mnemonic: "CODECOPY", StaticGas: 3}
	GASPRICE       = Opcode{Op: 0x3A, Mnemonic: "GASPRICE", StaticGas: 2}
	EXTCODESIZE    = Opcode{Op: 0x3B, Mnemonic: "EXTCODESIZE", StaticGas: 100}
	EXTCODECOPY    = Opcode{Op: 0x3C, Mnemonic: "EXTCODECOPY", StaticGas: 100}
	RETURNDATASIZE = Opcode{Op: 0x3D, Mnemonic: "RETURNDATASIZE", StaticGas: 2}
	RETURNDATACOPY = Opcode{Op: 0x3E, Mnemonic: "RETURNDATACOPY", StaticGas: 3}
	EXTCODEHASH    = Opcode{Op: 0x3F, Mnemonic: "EXTCODEHASH", StaticGas: 100}
)

// 0x40 - 0x5B

var (
	BLOCKHASH   = Opcode{Op: 0x40, Mnemonic: "BLOCKHASH", StaticGas: 20}
	COINBASE    = Opcode{Op: 0x41, Mnemonic: "COINBASE", StaticGas: 2}
	TIMESTAMP   = Opcode{Op: 0x42, Mnemonic: "TIMESTAMP", StaticGas: 2}
	NUMBER      = Opcode{Op: 0x43, Mnemonic: "NUMBER", StaticGas: 2}
	DIFFICULTY  = Opcode{Op: 0x44, Mnemonic: "DIFFICULTY", StaticGas: 2}
	PREVRANDAO  = Opcode{Op: 0x44, Mnemonic: "PREVRANDAO", StaticGas: 2}
	GASLIMIT    = Opcode{Op: 0x45, Mnemonic: "GASLIMIT", StaticGas: 2}
	CHAINID     = Opcode{Op: 0x46, Mnemonic: "CHAINID", StaticGas: 2}
	SELFBALANCE = Opcode{Op: 0x47, Mnemonic: "SELFBALANCE", StaticGas: 5}
	BASEFEE     = Opcode{Op: 0x48, Mnemonic: "BASEFEE", StaticGas: 2}
	POP         = Opcode{Op: 0x50, Mnemonic: "POP", StaticGas: 2}
	MLOAD       = Opcode{Op: 0x51, Mnemonic: "MLOAD", StaticGas: 3}
	MSTORE      = Opcode{Op: 0x52, Mnemonic: "MSTORE", StaticGas: 3}
	MSTORE8     = Opcode{Op: 0x53, Mnemonic: "MSTORE8", StaticGas: 3}
	SLOAD       = Opcode{Op: 0x54, Mnemonic: "SLOAD", StaticGas: 100}
	SSTORE      = Opcode{Op: 0x55, Mnemonic: "SSTORE", StaticGas: 100}
	JUMP        = Opcode{Op: 0x56, Mnemonic: "JUMP", StaticGas: 8}
	JUMPI       = Opcode{Op: 0x57, Mnemonic: "JUMPI", StaticGas: 10}
	PC          = Opcode{Op: 0x58, Mnemonic: "PC", StaticGas: 2}
	MSIZE       = Opcode{Op: 0x59, Mnemonic: "MSIZE", StaticGas: 2}
	GAS         = Opcode{Op: 0x5A, Mnemonic: "GAS", StaticGas: 2}
	JUMPDEST    = Opcode{Op: 0x5B, Mnemonic: "JUMPDEST", StaticGas: 1}
)

// 0x60 - 0x7F (PUSH)

var (
	PUSH1  = Opcode{Op: 0x60, Mnemonic: "PUSH1", StaticGas: 3}
	PUSH2  = Opcode{Op: 0x61, Mnemonic: "PUSH2", StaticGas: 3}
	PUSH3  = Opcode{Op: 0x62, Mnemonic: "PUSH3", StaticGas: 3}
	PUSH4  = Opcode{Op: 0x63, Mnemonic: "PUSH4", StaticGas: 3}
	PUSH5  = Opcode{Op: 0x64, Mnemonic: "PUSH5", StaticGas: 3}
	PUSH6  = Opcode{Op: 0x65, Mnemonic: "PUSH6", StaticGas: 3}
	PUSH7  = Opcode{Op: 0x66, Mnemonic: "PUSH7", StaticGas: 3}
	PUSH8  = Opcode{Op: 0x67, Mnemonic: "PUSH8", StaticGas: 3}
	PUSH9  = Opcode{Op: 0x68, Mnemonic: "PUSH9", StaticGas: 3}
	PUSH10 = Opcode{Op: 0x69, Mnemonic: "PUSH10", StaticGas: 3}
	PUSH11 = Opcode{Op: 0x6A, Mnemonic: "PUSH11", StaticGas: 3}
	PUSH12 = Opcode{Op: 0x6B, Mnemonic: "PUSH12", StaticGas: 3}
	PUSH13 = Opcode{Op: 0x6C, Mnemonic: "PUSH13", StaticGas: 3}
	PUSH14 = Opcode{Op: 0x6D, Mnemonic: "PUSH14", StaticGas: 3}
	PUSH15 = Opcode{Op: 0x6E, Mnemonic: "PUSH15", StaticGas: 3}
	PUSH16 = Opcode{Op: 0x6F, Mnemonic: "PUSH16", StaticGas: 3}
	PUSH17 = Opcode{Op: 0x70, Mnemonic: "PUSH17", StaticGas: 3}
	PUSH18 = Opcode{Op: 0x71, Mnemonic: "PUSH18", StaticGas: 3}
	PUSH19 = Opcode{Op: 0x72, Mnemonic: "PUSH19", StaticGas: 3}
	PUSH20 = Opcode{Op: 0x73, Mnemonic: "PUSH20", StaticGas: 3}
	PUSH21 = Opcode{Op: 0x74, Mnemonic: "PUSH21", StaticGas: 3}
	PUSH22 = Opcode{Op: 0x75, Mnemonic: "PUSH22", StaticGas: 3}
	PUSH23 = Opcode{Op: 0x76, Mnemonic: "PUSH23", StaticGas: 3}
	PUSH24 = Opcode{Op: 0x77, Mnemonic: "PUSH24", StaticGas: 3}
	PUSH25 = Opcode{Op: 0x78, Mnemonic: "PUSH25", StaticGas: 3}
	PUSH26 = Opcode{Op: 0x79, Mnemonic: "PUSH26", StaticGas: 3}
	PUSH27 = Opcode{Op: 0x7A, Mnemonic: "PUSH27", StaticGas: 3}
	PUSH28 = Opcode{Op: 0x7B, Mnemonic: "PUSH28", StaticGas: 3}
	PUSH29 = Opcode{Op: 0x7C, Mnemonic: "PUSH29", StaticGas: 3}
	PUSH30 = Opcode{Op: 0x7D, Mnemonic: "PUSH30", StaticGas: 3}
	PUSH31 = Opcode{Op: 0x7E, Mnemonic: "PUSH31", StaticGas: 3}
	PUSH32 = Opcode{Op: 0x7F, Mnemonic: "PUSH32", StaticGas: 3}
)

// 0x80 -

var (
	DUP1  = Opcode{Op: 0x80, Mnemonic: "DUP1", StaticGas: 3}
	DUP2  = Opcode{Op: 0x81, Mnemonic: "DUP2", StaticGas: 3}
	DUP3  = Opcode{Op: 0x82, Mnemonic: "DUP3", StaticGas: 3}
	DUP4  = Opcode{Op: 0x83, Mnemonic: "DUP4", StaticGas: 3}
	DUP5  = Opcode{Op: 0x84, Mnemonic: "DUP5", StaticGas: 3}
	DUP6  = Opcode{Op: 0x85, Mnemonic: "DUP6", StaticGas: 3}
	DUP7  = Opcode{Op: 0x86, Mnemonic: "DUP7", StaticGas: 3}
	DUP8  = Opcode{Op: 0x87, Mnemonic: "DUP8", StaticGas: 3}
	DUP9  = Opcode{Op: 0x88, Mnemonic: "DUP9", StaticGas: 3}
	DUP10 = Opcode{Op: 0x89, Mnemonic: "DUP10", StaticGas: 3}
	DUP11 = Opcode{Op: 0x8A, Mnemonic: "DUP11", StaticGas: 3}
	DUP12 = Opcode{Op: 0x8B, Mnemonic: "DUP12", StaticGas: 3}
	DUP13 = Opcode{Op: 0x8C, Mnemonic: "DUP13", StaticGas: 3}
	DUP14 = Opcode{Op: 0x8D, Mnemonic: "DUP14", StaticGas: 3}
	DUP15 = Opcode{Op: 0x8E, Mnemonic: "DUP15", StaticGas: 3}
	DUP16 = Opcode{Op: 0x8F, Mnemonic: "DUP16", StaticGas: 3}
)

var (
	SWAP1  = Opcode{Op: 0x90, Mnemonic: "SWAP1", StaticGas: 3}
	SWAP2  = Opcode{Op: 0x91, Mnemonic: "SWAP2", StaticGas: 3}
	SWAP3  = Opcode{Op: 0x92, Mnemonic: "SWAP3", StaticGas: 3}
	SWAP4  = Opcode{Op: 0x93, Mnemonic: "SWAP4", StaticGas: 3}
	SWAP5  = Opcode{Op: 0x94, Mnemonic: "SWAP5", StaticGas: 3}
	SWAP6  = Opcode{Op: 0x95, Mnemonic: "SWAP6", StaticGas: 3}
	SWAP7  = Opcode{Op: 0x96, Mnemonic: "SWAP7", StaticGas: 3}
	SWAP8  = Opcode{Op: 0x97, Mnemonic: "SWAP8", StaticGas: 3}
	SWAP9  = Opcode{Op: 0x98, Mnemonic: "SWAP9", StaticGas: 3}
	SWAP10 = Opcode{Op: 0x99, Mnemonic: "SWAP10", StaticGas: 3}
	SWAP11 = Opcode{Op: 0x9A, Mnemonic: "SWAP11", StaticGas: 3}
	SWAP12 = Opcode{Op: 0x9B, Mnemonic: "SWAP12", StaticGas: 3}
	SWAP13 = Opcode{Op: 0x9C, Mnemonic: "SWAP13", StaticGas: 3}
	SWAP14 = Opcode{Op: 0x9D, Mnemonic: "SWAP14", StaticGas: 3}
	SWAP15 = Opcode{Op: 0x9E, Mnemonic: "SWAP15", StaticGas: 3}
	SWAP16 = Opcode{Op: 0x9F, Mnemonic: "SWAP16", StaticGas: 3}
)

var (
	LOG0 = Opcode{Op: 0xA0, Mnemonic: "LOG0", StaticGas: 375}
	LOG1 = Opcode{Op: 0xA1, Mnemonic: "LOG1", StaticGas: 750}
	LOG2 = Opcode{Op: 0xA2, Mnemonic: "LOG2", StaticGas: 1125}
	LOG3 = Opcode{Op: 0xA3, Mnemonic: "LOG3", StaticGas: 1500}
	LOG4 = Opcode{Op: 0xA4, Mnemonic: "LOG4", StaticGas: 1875}
)

var (
	CREATE       = Opcode{Op: 0xF0, Mnemonic: "CREATE", StaticGas: 32000}
	CALL         = Opcode{Op: 0xF1, Mnemonic: "CALL", StaticGas: 100}
	CALLCODE     = Opcode{Op: 0xF2, Mnemonic: "CALLCODE", StaticGas: 100}
	RETURN       = Opcode{Op: 0xF3, Mnemonic: "RETURN", StaticGas: 0}
	DELEGATECALL = Opcode{Op: 0xF4, Mnemonic: "DELEGATECALL", StaticGas: 100}
	CREATE2      = Opcode{Op: 0xF5, Mnemonic: "CREATE2", StaticGas: 32000}
	STATICCALL   = Opcode{Op: 0xFA, Mnemonic: "STATICCALL", StaticGas: 100}
	REVERT       = Opcode{Op: 0xFD, Mnemonic: "REVERT", StaticGas: 0}
	INVALID      = Opcode{Op: 0xFE, Mnemonic: "INVALID", StaticGas: 0} // NaN (CONSUMES ALL GAS WITHOUT REFUND)
	SELFDESTRUCT = Opcode{Op: 0xFF, Mnemonic: "SELFDESTRUCT", StaticGas: 5000}
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
	"EXTCODEHASH":    EXTCODEHASH,
	"BLOCKHASH":      BLOCKHASH,
	"COINBASE":       COINBASE,
	"TIMESTAMP":      TIMESTAMP,
	"NUMBER":         NUMBER,
	"DIFFICULTY":     DIFFICULTY, // ADD DIFFICULTY for pre-merge Opcodes (EIP-4399)
	"PREVRANDAO":     PREVRANDAO,
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

// Check Operation Codes

func IsPush(op Opcode) bool {
	switch op {
	case PUSH1, PUSH2, PUSH3, PUSH4, PUSH5, PUSH6, PUSH7, PUSH8, PUSH9, PUSH10, PUSH11, PUSH12, PUSH13, PUSH14, PUSH15, PUSH16, PUSH17, PUSH18, PUSH19, PUSH20, PUSH21, PUSH22, PUSH23, PUSH24, PUSH25, PUSH26, PUSH27, PUSH28, PUSH29, PUSH30, PUSH31, PUSH32:
		return true

	default:
		return false
	}
}

func IsSwap(op Opcode) bool {
	switch op {
	case SWAP1, SWAP2, SWAP3, SWAP4, SWAP5, SWAP6, SWAP7, SWAP8, SWAP9, SWAP10, SWAP11, SWAP12, SWAP13, SWAP14, SWAP15, SWAP16:
		return true
	default:
		return false
	}
}

func IsDup(op Opcode) bool {
	switch op {
	case DUP1, DUP2, DUP3, DUP4, DUP5, DUP6, DUP7, DUP8, DUP9, DUP10, DUP11, DUP12, DUP13, DUP14, DUP15, DUP16:
		return true
	default:
		return false
	}
}
