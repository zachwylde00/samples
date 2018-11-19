
This page provides a step-by-step tutorial to deploy a sample app with Xooa's blockchain-as-a-service (BaaS).

# Overview

This repository contains a blockchain smart contract, also known as chaincode. Use the Xooa console to deploy it.

Xooa provides a permanent cloud end-point for the smart contract, enabling cloud-to-cloud integration, while retaining blockchain's peer-to-peer capabilities.

## Deploy the smart contract


1. Log in to the Xooa blockchain console at https://xooa.com/blockchain.
2. Go to **Apps**>**Deploy New**. If you didn't log in with your GitHub account, you will need to do it now.
3. Find the Github repository with the smart contract you want to deploy.  For example, use **Xooa/samples** to deploy one of the sample provided by Xooa. Tap **Select**, and then **Next**.

<img src="https://github.com/Xooa/samples/blob/master/images/deploy.png" alt="deploy" width="500px"/>

4. Select the Smart Contract you want to deploy, and then click **Deploy**.

5. Relax:  Xooa is doing the blockchain heavy lifting. You will be redirected to app dashboard when the deployment completes.

6.  Navigate to **Identities** tab, click or tap **Add New**, enter name for Identity  and set permissions to Read+Write. 

7. Copy and store the **API Token** value. You need it to authorize API requests. API Token cannot be dispalyed after you closed the window, but it may get regenerated. 

___


## The Xooa app dashboard 

The dashboard consists of the following tabs:
<dl>
  <dt>Details</dt>
  <dd>Information about your app.</dd>
<dt>Identities</dt>
  <dd>Create or delete an identity. You can also specify access rights to different identities, thus controlling the access of endpoints for your smart contract</dd>
   <dt>Activities</dt>
  <dd>History of your app's activities.</dd>
      <dt>Logs</dt>
 <dd>View smart contract events from the last 10 minutes.</dd>
  <dt>Manage</dt>
 <dd>Delete your app or update from github.</dd>
</dl>
___

## Explore the end-points for the smart Contract

1. Go to the **Details** tab, click **Explore API's**.

2. Enter **API Token** in the field in navigation pane. This is used to authenticate all API calls.

3. Go to the **Smart Contract > Invoke Smart Contract Function** from the navigation pane.

<img src="https://github.com/Xooa/samples/blob/master/images/invoke.png" alt="Invoke" width="300px"/>

4. In the **fcn** field, enter the Smart Contract function name you wish to Invoke. 
For example, if using this repo **get-set** smart contract to store data in blockchain `set` should be entred as the **fcn** field value 

5. In the **body** field, enter the data in the format expected by the smart contract. 
For example, if using this repo **get-set** smart contract to store data in blockchain ledger using the **set** function, the body should be entered as: 

    `[ "<key>", "<value>" ]`

6. Click or tap  **try**.
 * A response code of **200** indicates successful function call of the smart contract.
 * A response code of **202** indicates that your request is queued for processing. Final processing outcome may be obtained through **Result** end point by using **resultId** obtained in **response body**.
 * A response code of **400** indicates that you have a malformed request. Check the **body** and **fcn** fields.
 * A response code of **401** indicates that either you forgot to enter API Token you have entered invalid API token. While you cannot recover lost API  Token you can always generate a new one from Xooa console - app dashboard - **Identities** tab - **actions**

7. To view your transaction ain the blockchain Ledger, from your Xooa dashboard, go to **Ledger**.

8.  Go to the **Transactions** tab.
You can expand the data field to see your transactions.
