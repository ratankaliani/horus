# Horus
**Horus** is a blockchain built using Cosmos SDK and Tendermint.

## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Web Frontend

Ignite CLI has scaffolded a Vue.js-based web app in the `vue` directory. Run the following commands to install dependencies and start the app:

```
cd vue
npm install
npm run serve
```

The frontend app is built using the `@starport/vue` and `@starport/vuex` packages. For details, see the [monorepo for Ignite front-end development](https://github.com/ignite/web).

### Ignite Relayer

The Ignite relayer will connect multiple instances of Horus chains running.

To install the latest version of the Ignite CLI (and thereby Ignite Relayer), execute the following command on your machine:

```
curl https://get.ignite.com/cli@nightly! | bash
```

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.ignite.com/username/planet@latest! | sudo bash
```
`username/planet` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Demo Flow

## Flow

Setup & pre-shutdown demonstration!

1. Boot up three chains, Earth, Mars and Horus
    1. `ignite chain serve -c earth.yml`
    2. `ignite chain serve -c mars.yml`
    3. `ignite chain serve -c horus.yml`
2. Set up two IBC relayer connections, Earth → Mars and Earth → Horus
    1. `rm -rf ~/.ignite/relayer`
    2. 
    
    ignite relayer configure -a \
    --source-rpc "[http://0.0.0.0:26657](http://0.0.0.0:26657/)" \
    --source-faucet "[http://0.0.0.0:4500](http://0.0.0.0:4500/)" \
    --source-port "blog" \
    --source-version "blog-1" \
    --source-gasprice "0.0000025stake" \
    --source-prefix "cosmos" \
    --source-gaslimit 300000 \
    --target-rpc "[http://0.0.0.0:26661](http://0.0.0.0:26661/)" \
    --target-faucet "[http://0.0.0.0:4502](http://0.0.0.0:4502/)" \
    --target-port "blog" \
    --target-version "blog-1" \
    --target-gasprice "0.0000025stake" \
    --target-prefix "cosmos" \
    --target-gaslimit 300000
    
    c. 
    
    ignite relayer configure -a \
    --source-rpc "[http://0.0.0.0:26657](http://0.0.0.0:26657/)" \
    --source-faucet "[http://0.0.0.0:4500](http://0.0.0.0:4500/)" \
    --source-port "blog" \
    --source-version "blog-1" \
    --source-gasprice "0.0000025stake" \
    --source-prefix "cosmos" \
    --source-gaslimit 300000 \
    --target-rpc "[http://0.0.0.0:26659](http://0.0.0.0:26659/)" \
    --target-faucet "[http://0.0.0.0:4501](http://0.0.0.0:4501/)" \
    --target-port "blog" \
    --target-version "blog-1" \
    --target-gasprice "0.0000025stake" \
    --target-prefix "cosmos" \
    --target-gaslimit 300000
    
    d. `ignite relayer connect`
    
3. Send an IBC post from Earth to Mars by Alice
    1. Post 1: `planetd tx blog send-ibc-post blog channel-1 "Hello" "Hello Mars, I'm Alice from Earth" --from alice --chain-id earth --home ~/.earth`
4. Show that Mars stores the IBC post and Earth receives the ack
    1. Received on Mars: `planetd q blog list-post --node tcp://localhost:26659`
    2. Ack on Earth: `planetd q blog list-sent-post`

Now a user, Mallory will send a malicious message, which will break the invariant set by Earth’s governance.

1. Send an IBC post from Earth to Mars by Mallory (malicious)
    1. Post 1: `planetd tx blog send-ibc-post blog channel-1 "Some Malicious Message" "This is an exploit! (Hidden)" --from mallory --chain-id earth --home ~/.earth`
2. Show that Mars stores the IBC post and Earth receives the ack
    1. Received on Mars: `planetd q blog list-post --node tcp://localhost:26659`
    2. Ack on Earth: `planetd q blog list-sent-post`

We now know some malicious activity has occurred that has broken an invariant that Horus was set to protect on the consumer chain, Earth. Now a Horus validator, now seeing that this activity has occurred (via a ping to an RPC endpoint on their own full node), will pause IBC bridging outbound from Earth. This will pause IBC for all accounts, not just Mallory but also for accounts like Alice.

1. Shut down Earth sending IBC posts via the Eye Interchain Account on Horus via BridgeShutdown message
    1. `planetd tx blog send-ibc-horus-action blog channel-0 "No more IBC for Earth" "BridgeShutdown" --from eye --chain-id horus --home ~/.horus --node tcp://localhost:26661`
        1. Update this to Horus
2. Show that Mallory can no longer send IBC message to Mars
    1. Post 2: `planetd tx blog send-ibc-post blog channel-1 "Exploit Entendre" "Are you ready for Part 2?" --from mallory --chain-id earth --home ~/.earth`
3. Show that Alice can no longer send IBC message to Mars
    1. Post 2: `planetd tx blog send-ibc-post blog channel-1 "Hello Pt. 2" "Hello Earth, I'm Bob from Mars" --from alice --chain-id earth --home ~/.`earth

1. List Sent Actions on Horus
    1. `planetd q blog list-sent-action --node tcp://localhost:26661`

Now, after checking that the validators of the consumer chain have confirmed that normal chain activity has been restored (no more malicious action), Horus can resume bridging outbound from Earth.

1. Re-start Earth IBC bridge with Alice account on Horus via BridgeRestore message
    1. `planetd tx blog send-ibc-horus-action blog channel-0 "The Return of the IBC" "BridgeRestore" --from eye --chain-id horus --home ~/.horus --node tcp://localhost:26661`
        1. Update this to be from Horus
    2. `planetd q blog list-sent-action --node tcp://localhost:26661`
2. Send IBC post from Earth to Mars by Alice
    1. Post 2: `planetd tx blog send-ibc-post blog channel-1 "Please be back online" "I hope you can store this!" --from alice --chain-id earth --home ~/.earth`
3. Show Mars stores the IBC post
    1. Post 2 now added: `planetd q blog list-post --node tcp://localhost:26659`

As we’ve seen in the demo, chains like Earth can set invariants via governance, which when broken, can be acted upon by Horus validators. In doing so, Horus provides a flexible and modular solution to Cosmos chain security. Thank you!
