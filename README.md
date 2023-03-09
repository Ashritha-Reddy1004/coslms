## LEAVE MANAGEMET SYSTEM USING COSMOS ##


**To generate proto**

    $ make proto-gen


**To build lmsapp**

    $ make build

**To add keys**

    $ ./lmsd keys add [key_name]

**To list all keys**

    $ ./lmsd keys list

**Initialize the testnet with chain-id and some validator name**

This will create the configurations required for the testnet in daemon home directory.

    $ ./lmsd init --chain-id testnet myvalidator

**Add some genesis accounts and tokens.**

 We can do this via keyname (if the exists on your local machine) or with address

    $ ./lmsd add-genesis-account validator-key 1000000000stake
    $ ./lmsd add-genesis-account $(./simd keys show mykey1 -a) 10000000000stake


**Create gentx (create validator genesis transaction)**

    $ ./lmsd gentx validator-key --chain-id testnet

**Collect gentxs**

    $ ./lmsd collect-gentxs

**Start the node**

    $ ./lmsd start

**To generate lms module commands**

    $ ./build/lmsd tx coslms

**To  generate transaction sub-commands**
   
    $ ./build/lmsd tx

**To generate query sub-commands**
    $ ./build/lmsd query

**To execute transaction commands**
    
    $ lmsd tx [command]

**To execute query commands**
   
    $ lmsd query [command]

**To know more about transaction commands**
   
    $ lmsd tx [command] --help

**To know more about query commands**
   
    $ lmsd query [command] --help