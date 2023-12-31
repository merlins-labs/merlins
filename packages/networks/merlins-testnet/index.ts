import { merlinsTestnetCurrencies } from "./currencies";
import { NetworkFeature, NetworkInfo, NetworkKind } from "../types";

const nameServiceContractAddress =
  "fury18etjzrma5604af50jjklk3wlkqcsxdrvmy6jzw5naw2t7kyv4rys3kpwky";

const riotContractAddressGen1 =
  "fury1afwrcs58afaka6ltynevwcvq8zhejr3ssn703c0hky5emh890vzsry5wp5";

export const merlinsTestnetNetwork: NetworkInfo = {
  id: "merlins-testnet",
  kind: NetworkKind.Cosmos,
  chainId: "merlins-testnet-v3",
  displayName: "Merlins Testnet",
  icon: "icons/networks/merlins.svg",
  features: [
    NetworkFeature.NFTMarketplace,
    NetworkFeature.Organizations,
    NetworkFeature.SocialFeed,
    NetworkFeature.UPP,
  ],
  currencies: merlinsTestnetCurrencies,
  txExplorer: "https://explorer.merlins.world/merlins-testnet/tx/$hash",
  accountExplorer:
    "https://explorer.merlins.world/merlins-testnet/account/$address",
  contractExplorer:
    "https://explorer.merlins.world/merlins-testnet/account/$address",
  idPrefix: "tesfury",
  testnet: true,
  backendEndpoint: "https://dapp-backend.testnet.merlins.world",
  addressPrefix: "fury",
  restEndpoint: "https://rest.fury.black",
  rpcEndpoint: "http://35.204.44.202:26657",
  stakeCurrency: "ufury",
  gasPriceStep: {
    low: 0.0,
    average: 0.025,
    high: 0.04,
  },
  cosmosFeatures: [
    "stargate",
    "ibc-transfer",
    "cosmwasm",
    "no-legacy-stdTx",
    "ibc-go",
  ],
  nameServiceContractAddress,
  nameServiceDefaultImage:
    "ipfs://bafkreieqcwmjcb64r42ygs6a4dswz63djzgayjn3rhzjber3e42cknawlm",
  nameServiceTLD: ".fury",
  vaultContractAddress:
    "fury13d26fzmvuvlvvxwt3ur0vkv8j6f28mf4nqfmxu3xk5r57wzxxglqut8ezk",
  riotContractAddressGen0:
    "fury162skshe30f43kv2q2rw6we2mu3pvz43lm2zrg4hq50jdd2fjjdjsvm8mc7",
  riotContractAddressGen1,
  riotSquadStakingContractAddressV1:
    "fury15wl5lw9pxpedcekgqr8vk5xfpkrsax3reyfkkey458t6a7vv0d5ssark6m",
  riotSquadStakingContractAddressV2:
    "fury1gzk7kets2dfr7rcmzunkem0a23pn9ktkpjf0elxp38x42zpddt6qhvsp9l",
  distributorContractAddress:
    "fury1ff0ze096zcy5va5xhl7zenwt5vca6z9n0l8du7q7qsd7ftrcz0cqvd0vgc",
  riotersFooterContractAddress:
    "fury1m6g3l3j5t988zwmp965hrhsxvm8jrkf2ulw4gmwj8hx8pd84tvcqeearsd",
  secondaryDuringMintList: [
    nameServiceContractAddress,
    riotContractAddressGen1,
  ],
  excludeFromLaunchpadList: [riotContractAddressGen1],
  socialFeedContractAddress:
    "fury19y3xr6lghw04tj2ets8y70mrynlnflejanl2ys3n3c5vaasj358shjsk20",
  daoCw20CodeId: 99,
  daoFactoryCodeId: 100,
  daoCoreCodeId: 101,
  daoPreProposeSingleCodeId: 102,
  daoProposalSingleCodeId: 103,
  daoVotingCw20StakedCodeId: 104,
  daoCw20StakeCodeId: 105,
  daoCw4GroupCodeId: 106,
  daoVotingCw4CodeId: 109,
  daoFactoryContractAddress:
    "fury1r29chp8ufwgx9u3wr4sfk050aardhkzwve7nht6y06gvlaqutr3qt83z6r",
  coreDAOAddress:
    "fury1dy5h9q9zue4swxe9mzracm8gudp0fcf2ncllch6pfq9d0fq0ftgq546063",
};
