import atomSVG from "../../assets/icons/networks/cosmos-hub-circle.svg";
import osmosisSVG from "../../assets/icons/networks/osmosis-circle.svg";
import merlinsSVG from "../../assets/icons/networks/merlins-circle.svg";
import { TransactionAccount } from "../screens/WalletManager/types";

export const TEMP_IMAGE = "https://i.pravatar.cc/400?img=7";

export const DEPOSIT_WITHDARW_TYPE: { [key: string]: TransactionAccount } = {
  merlins: {
    account: "fury1grdjuukhq789y7z7w2uj2jjzl2gd3qpsulf7ud",
    name: "MERLINS",
    subtitle: "Network",
    icon: merlinsSVG,
    amount: 22.5552323,
  },
  osmosis: {
    account: "osmosis1grdjuukhq789y7z7w2uj2jjzl2gd3qpsulf7ud",
    name: "OSMOSIS",
    subtitle: "Network",
    icon: osmosisSVG,
    amount: 42.5552323,
  },
  atom: {
    account: "atom1grdjuukhq789y7z7w2uj2jjzl2gd3qpsulf7ud",
    name: "ATOM",
    subtitle: "Cosmos Hub",
    icon: atomSVG,
    amount: 10.5552323,
  },
};
