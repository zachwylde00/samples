
This page provides a step-by-step tutorial to integrate a Quickbooks app with Xooa's blockchain-as-a-service (BaaS).

The repository used in this example is <https://github.com/Xooa/samples>

# Overview

This repository contains the blockchain smart contract (sometimes also called "chaincode"). You'll deploy the smart contract via the Xooa console.

Using Xooa, provide a permanent cloud end-point for a QuickBooks app, enable cloud-to-cloud integration while maintaining the peer-to-peer capabilities of blockchain.

## Deploy the QuickBooks smart contract 

 
1. Log in to Xooa at <https://xooa.com/blockchain>

2. Go to **Apps**, then **Deploy New**.

3. Select **Xooa Git Repos**, and then click **Next**.

4. Select the branch (usually **master**) and **Xooa-quickBooks** as the Smart Contract, and then click **Deploy**.

5. Relax:  Xooa is doing the blockchain heavy lifting. You'll be redirected to app dashboard when the deployment completes.

6. Copy and save the **API Token** shown in the prompt. You will require this to authorize API requests to the Xooa app. You can also regenerate token from the **Identities** tab.

___

## Connect your company to Xooa Blockchain

1. Go to <https://xooa.com/integrations/quickbooks>. Click **Sign in with Intuit**.

2. Click **Connect to QuickBooks**.

3. Select the company you want to configure and invoke updates to *Xooa smart contract*.

4. Click **Connect**.

5. Enter **API token** you have noted earlier. Enter **function name** to call when any change happens.

	> The app is now configured and ready to receive any event from *QuickBooks* and logs to *Xooa blockchain*.