## Generate contract clients for EVM

- Clone the repo `https://github.com/merlins-labsevm-nft-vault`
- Set the needed hardhat environment 
- Compiles the contracts with: `npx hardhat compile`
- Copy the needed generated files, ex: If the contract name is `MerlinsNft` then we need to:
    + copy `types/factories/<MerlinsNft__factory.ts` => `packages/evm-contracts-clients/merlins-nft` 
    + copy `types/<MerlinsNft.d.ts` => `packages/evm-contracts-clients/merlins-nft` 
- Adjust the import:
    + `import type { MerlinsNft, MerlinsNftInterface } from "../MerlinsNft";` => `import type { MerlinsNft, MerlinsNftInterface } from "./MerlinsNft";`
    + `import type { TypedEventFilter, TypedEvent, TypedListener } from "./common";` => `import type { TypedEventFilter, TypedEvent, TypedListener } from "../common";`
