<template>
  <div style="overflow:auto; height: 500px;">
  <MDBTable border="primary">
    <thead>
      <tr>
        <th scope="col">Chain</th>
        <th scope="col">Transaction Hash</th>
        <th scope="col">Timestamp</th>
        <th scope="col">Invariant Broken</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="todo in todos" :key="todo.id">
        <td>{{todo.chain}}</td>
        <td>{{todo.tx}}</td>
        <td>{{todo.timestamp}}</td>
        <td>{{todo.invariantbroken}}</td>
      </tr>
    </tbody>
  </MDBTable>
</div>
</template>

<script>
   import {
    MDBTable,
    MDBBtn,
    MDBBadge
  } from 'mdb-vue-ui-kit';
  export default {
    components: {
      MDBTable,
    },
    data() {
    return {
      todos: [
      { id: 1, chain: "Osmosis", tx: 'https://www.mintscan.io/evmos/txs/4B9E3084A7B5BF8D619A29F3F9C200519879A227BCE3F17575EA635E3706C348', timestamp:'1669695924', invariantbroken:'false'},
      { id: 2, chain: "Evmos", tx: 'https://www.mintscan.io/osmosis/txs/87FF62825875D3FF640BD29E6DFA19E41C9B432A849DC0C98D2709FA986F39FD', timestamp:'1669695924', invariantbroken:'false'},
      { id: 3, chain: "Osmosis", tx: 'https://www.mintscan.io/evmos/txs/4B9E3084A7B5BF8D619A29F3F9C200519879A227BCE3F17575EA635E3706C348', timestamp:'1669695924', invariantbroken:'true'},
      { id: 4, chain: "Osmosis", tx: 'https://www.mintscan.io/evmos/txs/4B9E3084A7B5BF8D619A29F3F9C200519879A227BCE3F17575EA635E3706C348', timestamp:'1669695924', invariantbroken:'false'},
      { id: 5, chain: "Evmos", tx: 'https://www.mintscan.io/osmosis/txs/87FF62825875D3FF640BD29E6DFA19E41C9B432A849DC0C98D2709FA986F39FD', timestamp:'1669695924', invariantbroken:'false'},
      { id: 6, chain: "Osmosis", tx: 'https://www.mintscan.io/evmos/txs/4B9E3084A7B5BF8D619A29F3F9C200519879A227BCE3F17575EA635E3706C348', timestamp:'1669695924', invariantbroken:'true'},
      { id: 1, chain: "Osmosis", tx: 'https://www.mintscan.io/evmos/txs/4B9E3084A7B5BF8D619A29F3F9C200519879A227BCE3F17575EA635E3706C348', timestamp:'1669695924', invariantbroken:'false'},
      { id: 2, chain: "Evmos", tx: 'https://www.mintscan.io/osmosis/txs/87FF62825875D3FF640BD29E6DFA19E41C9B432A849DC0C98D2709FA986F39FD', timestamp:'1669695924', invariantbroken:'false'},
      { id: 3, chain: "Osmosis", tx: 'https://www.mintscan.io/evmos/txs/4B9E3084A7B5BF8D619A29F3F9C200519879A227BCE3F17575EA635E3706C348', timestamp:'1669695924', invariantbroken:'true'},
      { id: 4, chain: "Osmosis", tx: 'https://www.mintscan.io/evmos/txs/4B9E3084A7B5BF8D619A29F3F9C200519879A227BCE3F17575EA635E3706C348', timestamp:'1669695924', invariantbroken:'false'},
      { id: 5, chain: "Evmos", tx: 'https://www.mintscan.io/osmosis/txs/87FF62825875D3FF640BD29E6DFA19E41C9B432A849DC0C98D2709FA986F39FD', timestamp:'1669695924', invariantbroken:'false'},
      { id: 6, chain: "Osmosis", tx: 'https://www.mintscan.io/evmos/txs/4B9E3084A7B5BF8D619A29F3F9C200519879A227BCE3F17575EA635E3706C348', timestamp:'1669695924', invariantbroken:'true'},
      ]
    }
  },
  };

</script>

<style scoped>
</style>
