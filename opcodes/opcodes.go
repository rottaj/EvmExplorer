package opcodes

type Opcode struct {
	Op        byte
	Mnemonic  string
	staticGas int
	// dynamicGas func()?
	// If dynamicGas exists: use map StringToDynamicGasFunc to get function
	// Just like using StringToOpcode
}

// Create mapping of mneumonics. Check process step (string) with mapping.

// 0x00 - 0x0B
var (
	STOP       = Opcode{Op: 0x00, Mnemonic: "STOP", staticGas: 0} // Stops execution & refunds all remaining gas
	ADD        = Opcode{Op: 0x01, Mnemonic: "ADD", staticGas: 3}
	MUL        = Opcode{Op: 0x02, Mnemonic: "MUL", staticGas: 5}
	SUB        = Opcode{Op: 0x03, Mnemonic: "SUB", staticGas: 3}
	DIV        = Opcode{Op: 0x04, Mnemonic: "DIV", staticGas: 5}
	SDIV       = Opcode{Op: 0x05, Mnemonic: "SDIV", staticGas: 5}
	MOD        = Opcode{Op: 0x06, Mnemonic: "MOD", staticGas: 5}
	SMOD       = Opcode{Op: 0x07, Mnemonic: "SMOD", staticGas: 5}
	ADDMOD     = Opcode{Op: 0x08, Mnemonic: "ADDMOD", staticGas: 8}
	MULMOD     = Opcode{Op: 0x09, Mnemonic: "MULMOD", staticGas: 8}
	EXP        = Opcode{Op: 0x0A, Mnemonic: "EXP", staticGas: 10}
	SIGNEXTEND = Opcode{Op: 0x0B, Mnemonic: "SIGNEXTEND", staticGas: 5}
)

// 0x10 - 0x1D

var (
	LT     = Opcode{Op: 0x10, Mnemonic: "LT", staticGas: 3}
	GT     = Opcode{Op: 0x11, Mnemonic: "GT", staticGas: 3}
	SLT    = Opcode{Op: 0x12, Mnemonic: "SLT", staticGas: 3}
	SGT    = Opcode{Op: 0x13, Mnemonic: "SGT", staticGas: 3}
	EQ     = Opcode{Op: 0x14, Mnemonic: "EQ", staticGas: 3}
	ISZERO = Opcode{Op: 0x15, Mnemonic: "ISZERO", staticGas: 3}
	AND    = Opcode{Op: 0x16, Mnemonic: "AND", staticGas: 3}
	OR     = Opcode{Op: 0x17, Mnemonic: "OR", staticGas: 3}
	XOR    = Opcode{Op: 0x18, Mnemonic: "XOR", staticGas: 3}
	NOT    = Opcode{Op: 0x19, Mnemonic: "NOT", staticGas: 3}
	BYTE   = Opcode{Op: 0x1A, Mnemonic: "BYTE", staticGas: 3}
	SHL    = Opcode{Op: 0x1B, Mnemonic: "SHL", staticGas: 3}
	SHR    = Opcode{Op: 0x1C, Mnemonic: "SHR", staticGas: 3}
	SAR    = Opcode{Op: 0x1D, Mnemonic: "SAR", staticGas: 3}
)

// 0x20 - 0x3F

var (
	KECCAK256      = Opcode{Op: 0x20, Mnemonic: "KECCAK256", staticGas: 30}
	ADDRESS        = Opcode{Op: 0x30, Mnemonic: "ADDRESS", staticGas: 2}
	BALANCE        = Opcode{Op: 0x31, Mnemonic: "BALANCE", staticGas: 100}
	ORIGIN         = Opcode{Op: 0x32, Mnemonic: "ORIGIN", staticGas: 2}
	CALLER         = Opcode{Op: 0x33, Mnemonic: "CALLER", staticGas: 2}
	CALLVALUE      = Opcode{Op: 0x34, Mnemonic: "CALLVALUE", staticGas: 2}
	CALLDATALOAD   = Opcode{Op: 0x35, Mnemonic: "CALLDATALOAD", staticGas: 3}
	CALLDATASIZE   = Opcode{Op: 0x36, Mnemonic: "CALLDATASIZE", staticGas: 2}
	CALLDATACOPY   = Opcode{Op: 0x37, Mnemonic: "CALLDATACOPY", staticGas: 3}
	CODESIZE       = Opcode{Op: 0x38, Mnemonic: "CODESIZE", staticGas: 2}
	CODECOPY       = Opcode{Op: 0x39, Mnemonic: "CODECOPY", staticGas: 3}
	GASPRICE       = Opcode{Op: 0x3A, Mnemonic: "GASPRICE", staticGas: 2}
	EXTCODESIZE    = Opcode{Op: 0x3B, Mnemonic: "EXTCODESIZE", staticGas: 100}
	EXTCODECOPY    = Opcode{Op: 0x3C, Mnemonic: "EXTCODECOPY", staticGas: 100}
	RETURNDATASIZE = Opcode{Op: 0x3D, Mnemonic: "RETURNDATASIZE", staticGas: 2}
	RETURNDATACOPY = Opcode{Op: 0x3E, Mnemonic: "RETURNDATACOPY", staticGas: 3}
	EXTCODEHASH    = Opcode{Op: 0x3F, Mnemonic: "EXTCODEHASH", staticGas: 100}
)

// 0x40 - 0x5B

var (
	BLOCKHASH   = Opcode{Op: 0x40, Mnemonic: "BLOCKHASH", staticGas: 20}
	COINBASE    = Opcode{Op: 0x41, Mnemonic: "COINBASE", staticGas: 2}
	TIMESTAMP   = Opcode{Op: 0x42, Mnemonic: "TIMESTAMP", staticGas: 2}
	NUMBER      = Opcode{Op: 0x43, Mnemonic: "NUMBER", staticGas: 2}
	DIFFICULTY  = Opcode{Op: 0x44, Mnemonic: "DIFFICULTY", staticGas: 2}
	PREVRANDAO  = Opcode{Op: 0x44, Mnemonic: "PREVRANDAO", staticGas: 2}
	GASLIMIT    = Opcode{Op: 0x45, Mnemonic: "GASLIMIT", staticGas: 2}
	CHAINID     = Opcode{Op: 0x46, Mnemonic: "CHAINID", staticGas: 2}
	SELFBALANCE = Opcode{Op: 0x47, Mnemonic: "SELFBALANCE", staticGas: 5}
	BASEFEE     = Opcode{Op: 0x48, Mnemonic: "BASEFEE", staticGas: 2}
	POP         = Opcode{Op: 0x50, Mnemonic: "POP", staticGas: 2}
	MLOAD       = Opcode{Op: 0x51, Mnemonic: "MLOAD", staticGas: 3}
	MSTORE      = Opcode{Op: 0x52, Mnemonic: "MSTORE", staticGas: 3}
	MSTORE8     = Opcode{Op: 0x53, Mnemonic: "MSTORE8", staticGas: 3}
	SLOAD       = Opcode{Op: 0x54, Mnemonic: "SLOAD", staticGas: 100}
	SSTORE      = Opcode{Op: 0x55, Mnemonic: "SSTORE", staticGas: 100}
	JUMP        = Opcode{Op: 0x56, Mnemonic: "JUMP", staticGas: 8}
	JUMPI       = Opcode{Op: 0x57, Mnemonic: "JUMPI", staticGas: 10}
	PC          = Opcode{Op: 0x58, Mnemonic: "PC", staticGas: 2}
	MSIZE       = Opcode{Op: 0x59, Mnemonic: "MSIZE", staticGas: 2}
	GAS         = Opcode{Op: 0x5A, Mnemonic: "GAS", staticGas: 2}
	JUMPDEST    = Opcode{Op: 0x5B, Mnemonic: "JUMPDEST", staticGas: 1}
)

// 0x60 - 0x7F (PUSH)

var (
	PUSH1  = Opcode{Op: 0x60, Mnemonic: "PUSH1", staticGas: 3}
	PUSH2  = Opcode{Op: 0x61, Mnemonic: "PUSH2", staticGas: 3}
	PUSH3  = Opcode{Op: 0x62, Mnemonic: "PUSH3", staticGas: 3}
	PUSH4  = Opcode{Op: 0x63, Mnemonic: "PUSH4", staticGas: 3}
	PUSH5  = Opcode{Op: 0x64, Mnemonic: "PUSH5", staticGas: 3}
	PUSH6  = Opcode{Op: 0x65, Mnemonic: "PUSH6", staticGas: 3}
	PUSH7  = Opcode{Op: 0x66, Mnemonic: "PUSH7", staticGas: 3}
	PUSH8  = Opcode{Op: 0x67, Mnemonic: "PUSH8", staticGas: 3}
	PUSH9  = Opcode{Op: 0x68, Mnemonic: "PUSH9", staticGas: 3}
	PUSH10 = Opcode{Op: 0x69, Mnemonic: "PUSH10", staticGas: 3}
	PUSH11 = Opcode{Op: 0x6A, Mnemonic: "PUSH11", staticGas: 3}
	PUSH12 = Opcode{Op: 0x6B, Mnemonic: "PUSH12", staticGas: 3}
	PUSH13 = Opcode{Op: 0x6C, Mnemonic: "PUSH13", staticGas: 3}
	PUSH14 = Opcode{Op: 0x6D, Mnemonic: "PUSH14", staticGas: 3}
	PUSH15 = Opcode{Op: 0x6E, Mnemonic: "PUSH15", staticGas: 3}
	PUSH16 = Opcode{Op: 0x6F, Mnemonic: "PUSH16", staticGas: 3}
	PUSH17 = Opcode{Op: 0x70, Mnemonic: "PUSH17", staticGas: 3}
	PUSH18 = Opcode{Op: 0x71, Mnemonic: "PUSH18", staticGas: 3}
	PUSH19 = Opcode{Op: 0x72, Mnemonic: "PUSH19", staticGas: 3}
	PUSH20 = Opcode{Op: 0x73, Mnemonic: "PUSH20", staticGas: 3}
	PUSH21 = Opcode{Op: 0x74, Mnemonic: "PUSH21", staticGas: 3}
	PUSH22 = Opcode{Op: 0x75, Mnemonic: "PUSH22", staticGas: 3}
	PUSH23 = Opcode{Op: 0x76, Mnemonic: "PUSH23", staticGas: 3}
	PUSH24 = Opcode{Op: 0x77, Mnemonic: "PUSH24", staticGas: 3}
	PUSH25 = Opcode{Op: 0x78, Mnemonic: "PUSH25", staticGas: 3}
	PUSH26 = Opcode{Op: 0x79, Mnemonic: "PUSH26", staticGas: 3}
	PUSH27 = Opcode{Op: 0x7A, Mnemonic: "PUSH27", staticGas: 3}
	PUSH28 = Opcode{Op: 0x7B, Mnemonic: "PUSH28", staticGas: 3}
	PUSH29 = Opcode{Op: 0x7C, Mnemonic: "PUSH29", staticGas: 3}
	PUSH30 = Opcode{Op: 0x7D, Mnemonic: "PUSH30", staticGas: 3}
	PUSH31 = Opcode{Op: 0x7E, Mnemonic: "PUSH31", staticGas: 3}
	PUSH32 = Opcode{Op: 0x7F, Mnemonic: "PUSH32", staticGas: 3}
)

// 0x80 -

var (
	DUP1  = Opcode{Op: 0x80, Mnemonic: "DUP1", staticGas: 3}
	DUP2  = Opcode{Op: 0x81, Mnemonic: "DUP2", staticGas: 3}
	DUP3  = Opcode{Op: 0x82, Mnemonic: "DUP3", staticGas: 3}
	DUP4  = Opcode{Op: 0x83, Mnemonic: "DUP4", staticGas: 3}
	DUP5  = Opcode{Op: 0x84, Mnemonic: "DUP5", staticGas: 3}
	DUP6  = Opcode{Op: 0x85, Mnemonic: "DUP6", staticGas: 3}
	DUP7  = Opcode{Op: 0x86, Mnemonic: "DUP7", staticGas: 3}
	DUP8  = Opcode{Op: 0x87, Mnemonic: "DUP8", staticGas: 3}
	DUP9  = Opcode{Op: 0x88, Mnemonic: "DUP9", staticGas: 3}
	DUP10 = Opcode{Op: 0x89, Mnemonic: "DUP10", staticGas: 3}
	DUP11 = Opcode{Op: 0x8A, Mnemonic: "DUP11", staticGas: 3}
	DUP12 = Opcode{Op: 0x8B, Mnemonic: "DUP12", staticGas: 3}
	DUP13 = Opcode{Op: 0x8C, Mnemonic: "DUP13", staticGas: 3}
	DUP14 = Opcode{Op: 0x8D, Mnemonic: "DUP14", staticGas: 3}
	DUP15 = Opcode{Op: 0x8E, Mnemonic: "DUP15", staticGas: 3}
	DUP16 = Opcode{Op: 0x8F, Mnemonic: "DUP16", staticGas: 3}
)

var (
	SWAP1  = Opcode{Op: 0x90, Mnemonic: "SWAP1", staticGas: 3}
	SWAP2  = Opcode{Op: 0x91, Mnemonic: "SWAP2", staticGas: 3}
	SWAP3  = Opcode{Op: 0x92, Mnemonic: "SWAP3", staticGas: 3}
	SWAP4  = Opcode{Op: 0x93, Mnemonic: "SWAP4", staticGas: 3}
	SWAP5  = Opcode{Op: 0x94, Mnemonic: "SWAP5", staticGas: 3}
	SWAP6  = Opcode{Op: 0x95, Mnemonic: "SWAP6", staticGas: 3}
	SWAP7  = Opcode{Op: 0x96, Mnemonic: "SWAP7", staticGas: 3}
	SWAP8  = Opcode{Op: 0x97, Mnemonic: "SWAP8", staticGas: 3}
	SWAP9  = Opcode{Op: 0x98, Mnemonic: "SWAP9", staticGas: 3}
	SWAP10 = Opcode{Op: 0x99, Mnemonic: "SWAP10", staticGas: 3}
	SWAP11 = Opcode{Op: 0x9A, Mnemonic: "SWAP11", staticGas: 3}
	SWAP12 = Opcode{Op: 0x9B, Mnemonic: "SWAP12", staticGas: 3}
	SWAP13 = Opcode{Op: 0x9C, Mnemonic: "SWAP13", staticGas: 3}
	SWAP14 = Opcode{Op: 0x9D, Mnemonic: "SWAP14", staticGas: 3}
	SWAP15 = Opcode{Op: 0x9E, Mnemonic: "SWAP15", staticGas: 3}
	SWAP16 = Opcode{Op: 0x9F, Mnemonic: "SWAP16", staticGas: 3}
)

var (
	LOG0 = Opcode{Op: 0xA0, Mnemonic: "LOG0", staticGas: 375}
	LOG1 = Opcode{Op: 0xA1, Mnemonic: "LOG1", staticGas: 750}
	LOG2 = Opcode{Op: 0xA2, Mnemonic: "LOG2", staticGas: 1125}
	LOG3 = Opcode{Op: 0xA3, Mnemonic: "LOG3", staticGas: 1500}
	LOG4 = Opcode{Op: 0xA4, Mnemonic: "LOG4", staticGas: 1875}
)

var (
	CREATE       = Opcode{Op: 0xF0, Mnemonic: "CREATE", staticGas: 32000}
	CALL         = Opcode{Op: 0xF1, Mnemonic: "CALL", staticGas: 100}
	CALLCODE     = Opcode{Op: 0xF2, Mnemonic: "CALLCODE", staticGas: 100}
	RETURN       = Opcode{Op: 0xF3, Mnemonic: "RETURN", staticGas: 0}
	DELEGATECALL = Opcode{Op: 0xF4, Mnemonic: "DELEGATECALL", staticGas: 100}
	CREATE2      = Opcode{Op: 0xF5, Mnemonic: "CREATE2", staticGas: 32000}
	STATICCALL   = Opcode{Op: 0xFA, Mnemonic: "STATICCALL", staticGas: 100}
	REVERT       = Opcode{Op: 0xFD, Mnemonic: "REVERT", staticGas: 0}
	INVALID      = Opcode{Op: 0xFE, Mnemonic: "INVALID", staticGas: 0} // NaN (CONSUMES ALL GAS WITHOUT REFUND)
	SELFDESTRUCT = Opcode{Op: 0xFF, Mnemonic: "SELFDESTRUCT", staticGas: 5000}
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
	if op == PUSH1 {
		return true
	} else {
		return false
	}
}
