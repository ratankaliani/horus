import {
  Badge,
  Box,
  Heading,
  SimpleGrid,
  Stack,
  Text,
  useBreakpointValue,
  useColorModeValue,
} from "@chakra-ui/react";
import loadConfig from "next/dist/server/config";
import * as React from "react";
import useSWR from "swr";

interface Props {
  label: string;
  value: string;
}

function useLatestBlock(network) {
  const fetcher = (...args) => fetch(...args).then((res) => res.json());

  switch (network) {
    case "Earth": {
      const { data, error } = useSWR(
        "http://0.0.0.0:1317/cosmos/base/tendermint/v1beta1/blocks/latest",
        fetcher,
        { refreshInterval: 100 }
      );

      return data;
    }
    case "Mars": {
      const { data, error } = useSWR(
        "http://0.0.0.0:1318/cosmos/base/tendermint/v1beta1/blocks/latest",
        fetcher,
        { refreshInterval: 100 }
      );

      return data;
    }
    case "Horus": {
      const { data, error } = useSWR(
        "http://0.0.0.0:1319/cosmos/base/tendermint/v1beta1/blocks/latest",
        fetcher,
        { refreshInterval: 100 }
      );

      return data;
    }
  }
}

export const Inbox = (props: Props) => {
  const { label, value, ...boxProps } = props;
  const data = useLatestBlock(label);
  return (
    <Box
      px={{ base: "4", md: "6" }}
      py={{ base: "5", md: "6" }}
      bg="white"
      borderRadius="lg"
      boxShadow={useColorModeValue("sm", "sm-dark")}
      {...boxProps}
    >
      <Stack>
        <Box>
          <Heading size="lg">
            {label + (label == "Horus" ? " Outbox" : " Inbox")}
          </Heading>
        </Box>
      </Stack>
    </Box>
  );
};
