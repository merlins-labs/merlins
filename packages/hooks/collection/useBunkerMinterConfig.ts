import { useQuery } from "@tanstack/react-query";

import { MerlinsBunkerMinterQueryClient } from "../../contracts-clients/merlins-bunker-minter/MerlinsBunkerMinter.client";
import { mustGetNonSigningCosmWasmClient } from "../../networks";

export const useBunkerMinterConfig = (
  networkId: string | undefined,
  contractAddress: string | undefined,
  enabled?: boolean
) => {
  if (enabled === undefined) {
    enabled = true;
  }
  const { data, ...other } = useQuery(
    ["bunkerMinterConfig", networkId, contractAddress],
    async () => {
      if (!networkId || !contractAddress) {
        return undefined;
      }
      const cosmwasm = await mustGetNonSigningCosmWasmClient(networkId);
      const minterClient = new MerlinsBunkerMinterQueryClient(
        cosmwasm,
        contractAddress
      );
      return await minterClient.config();
    },
    {
      staleTime: Infinity,
      enabled: !!networkId && !!contractAddress && enabled,
    }
  );
  return { bunkerMinterConfig: data, ...other };
};
