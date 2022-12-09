import Head from "next/head";
import {
  Box,
  Divider,
  Grid,
  Heading,
  Text,
  Stack,
  Container,
  Link,
  Button,
  Flex,
  Icon,
  SimpleGrid,
  useColorMode,
  useColorModeValue,
} from "@chakra-ui/react";
import { BsFillMoonStarsFill, BsFillSunFill } from "react-icons/bs";
import { Product, Dependency, WalletSection } from "../components";
import { dependencies, products } from "../config";
import SimpleSidebar from "../components/sidebar";
import { Stat } from "../components/stat";
import { useCallback } from "react";
import ReactFlow, {
  MiniMap,
  Controls,
  Background,
  useNodesState,
  useEdgesState,
  addEdge,
} from "reactflow";
import "reactflow/dist/style.css";
import { Inbox } from "../components/inbox";

export default function Home() {
  const stats = [{ label: "Horus" }, { label: "Earth" }, { label: "Mars" }];
  const inboxes = [{ label: "Horus" }, { label: "Earth" }, { label: "Mars" }];

  const initialNodes = [
    {
      id: "Mars",
      position: { x: -200, y: 0 },
      data: { label: "Mars" },
      connectable: false,
      draggable: false,
      selectable: false,
      deleteable: false,
      sourcePosition: "right",
      targetPosition: "top",
    },
    {
      id: "Horus",
      position: { x: 0, y: -100 },
      data: { label: "Horus" },
      connectable: false,
      draggable: false,
      selectable: false,
      deleteable: false,
      type: "input",
    },
    {
      id: "Earth",
      position: { x: 200, y: 0 },
      data: { label: "Earth" },
      connectable: false,
      draggable: false,
      selectable: false,
      deleteable: false,
      targetPosition: "left",
      type: "output",
    },
  ];

  const proOptions = { hideAttribution: true };
  const initialEdges = [
    {
      id: "m-h",
      source: "Horus",
      target: "Mars",
      animated: true,
      style: { stroke: "black" },
      label: "Monitoring",
    },
    {
      id: "e-h",
      source: "Horus",
      target: "Earth",
      animated: true,
      style: { stroke: "black" },
      label: "Monitoring",
    },
    {
      id: "m-e",
      source: "Mars",
      target: "Earth",
      animated: true,
      style: { stroke: "red" },
      label: "IBC channel-0",
    },
  ];

  const [nodes, setNodes, onNodesChange] = useNodesState(initialNodes);
  const [edges, setEdges, onEdgesChange] = useEdgesState(initialEdges);

  const onConnect = useCallback(
    (params) => setEdges((eds) => addEdge(params, eds)),
    []
  );

  return (
    <>
      <SimpleSidebar>
        <SimpleGrid columns={{ base: 1, md: 3 }} gap={{ base: "5", md: "6" }}>
          {stats.map(({ label, value }) => (
            <Stat key={label} label={label} value={value} />
          ))}
        </SimpleGrid>
        <SimpleGrid
          mt="5"
          columns={{ base: 1, md: 1 }}
          gap={{ base: "5", md: "6" }}
        >
          <Box
            w="100%"
            h="500px"
            color="white"
            px={{ base: "4", md: "6" }}
            py={{ base: "5", md: "6" }}
            bg="white"
            borderRadius="lg"
            boxShadow={useColorModeValue("sm", "sm-dark")}
          >
            <ReactFlow
              proOptions={proOptions}
              nodes={nodes}
              edges={edges}
              onNodesChange={onNodesChange}
              onEdgesChange={onEdgesChange}
              onConnect={onConnect}
              fitView
            >
              <Background />
            </ReactFlow>
          </Box>
        </SimpleGrid>
        <SimpleGrid
          mt="5"
          columns={{ base: 1, md: 3 }}
          gap={{ base: "5", md: "6" }}
        >
          {inboxes.map(({ label, value }) => (
            <Inbox key={label} label={label} value={value} />
          ))}
        </SimpleGrid>
      </SimpleSidebar>
    </>
  );
}
