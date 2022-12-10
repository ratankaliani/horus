import {
  Badge,
  Box,
  Divider,
  Heading,
  ListItem,
  SimpleGrid,
  Stack,
  Text,
  List,
  UnorderedList,
  useBreakpointValue,
  useColorModeValue,
} from "@chakra-ui/react";
import loadConfig from "next/dist/server/config";
import * as React from "react";
import useSWR from "swr";
import axios from "axios";
import { time } from "console";

interface Props {
  label: string;
  value: string;
}

interface Post {
  id: string;
  title: string;
  content: string;
  creator: string;
}

interface SentPost {
  id: string;
  title: string;
  content: string;
  chain: string;
  creator: string;
}

interface SentAction {
  id: string;
  title: string;
  actionID: string;
  action: string;
  chain: string;
  creator: string;
}

function useLatestBlock(network) {
  const fetcher = (url: string) => axios(url).then((r) => r.data);
  switch (network) {
    case "Earth": {
      const { data, error } = useSWR(
        "http://0.0.0.0:1317/planet/blog/sent_post",
        fetcher,
        { refreshInterval: 100 }
      );

      return data;
    }
    case "Mars": {
      const { data, error } = useSWR(
        "http://0.0.0.0:1318/planet/blog/post",
        fetcher,
        { refreshInterval: 100 }
      );

      return data;
    }
    case "Horus": {
      const { data, error } = useSWR(
        "http://0.0.0.0:1319/planet/blog/sent_action",
        fetcher,
        { refreshInterval: 100 }
      );

      return data;
    }
  }
}

function getRelevantData(data, label) {
  if (data == null) {
    return "";
  }
  switch(label) {
    case "Earth": {
      const listItems = data.SentPost.map((sp: SentPost) =>
      <ListItem key={sp.id} value = {sp.id + ": " + sp.title}/>)
      return {"title": "Sent Messages", data: listItems}
    }
    case "Mars": {
      const listItems = data.Post.map((p: Post) =>
      <ListItem key={p.id} value = {p.id + ": " + p.title}/>)
      return {"title": "Received Messages", data: listItems}
    }
    case "Horus": {
      const listItems = data.SentAction.map((action: SentPost) =>
      <ListItem key={action.id} value = {action.id + ": " + action.title}/>)
      return {"title": "Actions", data: listItems}
    }
  }
}

function getListItems(data, label) {
  if (data == null) {
    return "";
  }
  switch(label) {
    case "Earth": {
      const listItems = data.SentPost.map((sp: SentPost) =>
      <ListItem>{sp.id + ": " + sp.title}</ListItem>)
      return listItems
    }
    case "Mars": {
      const listItems = data.Post.map((p: Post) =>
      <ListItem>{p.id + ": " + p.title}</ListItem>)
      return listItems
    }
    case "Horus": {
      const listItems = data.SentAction.map((action: SentPost) =>
      <ListItem>{action.id + ": " + action.title}</ListItem>)
      return listItems
    }
  }
}

export const Inbox = (props: Props) => {
  const { label, value, ...boxProps } = props;
  const data = useLatestBlock(label);
  const formattedData = getRelevantData(data, label)
  // var displayData;
  // if (data == null) {
  //   displayData = "";
  // } else {
  //   displayData = Object.keys(data)[0];

  // }

  // console.log("This is the data", label, displayData)
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
            {label + (label == "Mars" ? " Inbox" : " Outbox")}
          </Heading>
        </Box>
        <Divider />
        <Heading size="md" textAlign="center">
          {formattedData ? formattedData["title"] : "Loading..."}
        </Heading>
        <UnorderedList>
          {formattedData ? getListItems(data, label) : "Loading..."}
        </UnorderedList>
      </Stack>
    </Box>
    
  );
};
