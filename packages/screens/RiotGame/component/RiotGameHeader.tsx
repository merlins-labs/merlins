import { useRoute } from "@react-navigation/native";
import React from "react";
import { Linking, StyleSheet, TouchableOpacity, View } from "react-native";
import { ScrollView } from "react-native-gesture-handler";
import { SvgProps } from "react-native-svg";

import cartSVG from "../../../../assets/icons/cart.svg";
import diamondSVG from "../../../../assets/icons/diamond.svg";
import fightSVG from "../../../../assets/icons/fight.svg";
import filmSVG from "../../../../assets/icons/film.svg";
import gameBoxSVG from "../../../../assets/icons/game-box.svg";
import inventorySVG from "../../../../assets/icons/inventory.svg";
import trophiesSVG from "../../../../assets/icons/trophies.svg";
import { BrandText } from "../../../components/BrandText";
import FlexRow from "../../../components/FlexRow";
import { NetworkSelector } from "../../../components/NetworkSelector/NetworkSelector";
import { SVG } from "../../../components/SVG";
import { Separator } from "../../../components/Separator";
import { ConnectWalletButton } from "../../../components/TopMenu/ConnectWalletButton";
import { TopLogo } from "../../../components/navigation/components/TopLogo";
import { SpacerRow } from "../../../components/spacer";
import { useForceNetworkKind } from "../../../hooks/useForceNetworkKind";
import { NetworkKind } from "../../../networks";
import {
  RootStackParamList,
  useAppNavigation,
} from "../../../utils/navigation";
import {
  neutral33,
  neutralA3,
  yellowDefault,
} from "../../../utils/style/colors";
import { fontMedium16 } from "../../../utils/style/fonts";
import {
  headerHeight,
  headerMarginHorizontal,
  layout,
} from "../../../utils/style/layout";
import { PickByValue } from "../../../utils/types/helper";

type MenuItem = {
  id: string;
  name: string;
  route?:
    | keyof PickByValue<RootStackParamList, undefined>
    | "RiotGameMarketplace";
  externalRoute?: string;
  iconSVG: React.FC<SvgProps>;
};

const MENU_ITEMS: MenuItem[] = [
  { id: "fight", name: "Fight", route: "RiotGameEnroll", iconSVG: fightSVG },
  {
    id: "inventory",
    name: "Inventory",
    route: "RiotGameInventory",
    iconSVG: inventorySVG,
  },
  {
    id: "breeding",
    name: "Breeding",
    route: "RiotGameBreeding",
    iconSVG: gameBoxSVG,
  },
  {
    id: "leaderboard",
    name: "Leaderboard",
    route: "RiotGameLeaderboard",
    iconSVG: trophiesSVG,
  },
  {
    id: "rarity",
    name: "Rarity",
    externalRoute:
      "https://fury-live.io/allnfts.php?collid=fury1j08452mqwadp8xu25kn9rleyl2gufgfjnv0sn8dvynynakkjukcq3vtuv2",
    iconSVG: diamondSVG,
  },
  {
    id: "marketplace",
    name: "Marketplace",
    route: "RiotGameMarketplace",
    iconSVG: cartSVG,
  },
  {
    id: "memories",
    name: "Memories",
    route: "RiotGameMemories",
    iconSVG: filmSVG,
  },
];

type RiotGameHeaderProps = {
  hideMenu?: boolean;
};

export const RiotGameHeader: React.FC<RiotGameHeaderProps> = ({
  hideMenu = false,
}) => {
  const navigation = useAppNavigation();
  const { name: routeName } = useRoute();
  useForceNetworkKind(NetworkKind.Cosmos);

  const onMenuItemClick = (item: MenuItem) => {
    if (item.externalRoute) {
      Linking.openURL(item.externalRoute);
    } else if (item.route) {
      // @ts-expect-error
      navigation.navigate(item.route);
    }
  };

  return (
    <View style={styles.outerContainer}>
      <View style={styles.innerContainer}>
        <View>
          <TopLogo />
        </View>

        <ScrollView
          showsHorizontalScrollIndicator={false}
          horizontal
          style={styles.menu}
        >
          {!hideMenu && (
            <FlexRow>
              {MENU_ITEMS.map((menuItem) => {
                // Enroll and Fight are under the same menu item
                let _routeName = routeName;
                if (_routeName === "RiotGameFight")
                  _routeName = "RiotGameEnroll";

                const color =
                  menuItem.route === _routeName ? yellowDefault : neutralA3;

                return (
                  <TouchableOpacity
                    onPress={() => onMenuItemClick(menuItem)}
                    style={{ marginRight: layout.padding_x4 }}
                    key={menuItem.id}
                  >
                    <FlexRow style={{ alignItems: "center" }}>
                      <SVG
                        width={16}
                        height={16}
                        color={color}
                        source={menuItem.iconSVG}
                        style={{ marginRight: layout.padding_x1 }}
                      />
                      <BrandText style={[fontMedium16, { color }]}>
                        {menuItem.name}
                      </BrandText>
                    </FlexRow>
                  </TouchableOpacity>
                );
              })}
            </FlexRow>
          )}
        </ScrollView>

        <View style={styles.section}>
          <SpacerRow size={1.5} />
          <Separator horizontal color={neutral33} />
          <SpacerRow size={1.5} />
          <NetworkSelector forceNetworkKind={NetworkKind.Cosmos} />
          <SpacerRow size={1.5} />
          <ConnectWalletButton
            style={{ marginRight: headerMarginHorizontal }}
          />
        </View>
      </View>
    </View>
  );
};

const styles = StyleSheet.create({
  outerContainer: {
    height: headerHeight,
    width: "100%",
    zIndex: 1000,
  },
  innerContainer: {
    flex: 1,
    borderBottomWidth: 1,
    borderColor: neutral33,
    flexDirection: "row",
    alignItems: "center",
  },
  menu: {
    paddingHorizontal: layout.padding_x4,
  },
  section: {
    flexDirection: "row",
    height: "100%",
    alignItems: "center",
  },
});