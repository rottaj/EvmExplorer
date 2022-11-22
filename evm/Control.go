package evm

//"github.com/rottaj/EvmExplorer/opcodes"

// Not adding running param in evm since this program is meant for static analysis
// Computing gas & incrementing program counter
func (evm *Evm) stop() { // stops execution of evm
	evm.Pc += 1
}

// need to add jumpdest & jump.
// Keep jumpdest as evm param?
