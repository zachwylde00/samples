
This page provides a step-by-step tutorial to integrate a blockchain SmartApp with Xooa's blockchain-as-a-service (BaaS).

The repository used in this example is <https://github.com/Xooa/samples>

# Overview

This repository contains the blockchain smart contract (sometimes also called "chaincode") and SmartThings SmartApps. You'll deploy the smart contract via the Xooa console and the SmartApps via the SmartThings IDE. The Blockchain Event Logger SmartApp allows you to log SmartThings events to blockchain. The Blockchain Event Viewer SmartApp allows you to retrieve and view SmartThings events logged in a blockchain. 

Using Xooa, provide a permanent cloud end-point for a SmartThings app, enable cloud-to-cloud integration while maintaining the peer-to-peer capabilities of blockchain.

## Deploy the SmartThings smart contract 

 
1. Log in to Xooa at <https://xooa.com/blockchain>

2. Navigate to **Apps**, then **Deploy New**.

3. Select **Xooa Git Repos**, and then click **Next**.

4. Select the branch (usually **master**) and **Xooa-SmartThings** as the Smart Contract, and then click **Deploy**.

5. Relax:  Xooa is doing the blockchain heavy lifting. You'll be redirected to app dashboard when the deployment completes.

6. Copy and save the **API Token** shown in the prompt. You will require this to authorize API requests from the SmartThings smartapp.

	To regenerate **API Token**, go to `Identitites` tab. Click on **Actions** against the identity ID. Click on **Regenerate API Token**. Click **Generate**.

___

## Blockchain Event Logger SmartApp Setup (SmartThings IDE)

1. Log in with your Samsung SmartThings account to the SmartThings IDE at <https://graph.api.smartthings.com/>.

2. From the menu, click **My SmartApps**.

You now need to publish the app.  You can do this with or without GitHub integration:

**Without GitHub integration:**

1. Click **New SmartApp**.

2. Copy the *groovy* script from the **smartthings** folder of the  **samples** repo: <https://raw.githubusercontent.com/Xooa/samples/master/smartthings/blockchain-event-logger.groovy>
    
3. On the **From code** tab, paste the code.

4. Click **Create**.

5. Click **Save** and then click **Publish > For me**.

**With GitHub integration** (if you haven't already set up GitHub to work with SmartThings, here is the community FAQ on the subject <https://community.smartthings.com/t/faq-github-integration-how-to-add-and-update-from-repositories/39046>)

1. From the menu, click **My SmartApps** and then click **Settings**.   
2. Click **Add new repository**.
3.  Add the GitHub repo to your IDE with the following parameters:
    * `Owner`: xooa
    * `Name`: smartthings-xooa
    * `Branch`: master
4. Click **Save**.
5. Click **Update from Repo**.
6. Click **samples (master)**.
7. Select **blockchain-event-logger.groovy** from the **New (only in GitHub)** column.
8.  Select **Publish**.
9.  Click  **Execute Update**.


There are two apps available for SmartThings in the Google Play store. We recommend the classic app over the new app.

## Blockchain Event Logger SmartApp Setup (Smartphone)
Before you begin, ensure that you:

* Have the SmartThings app installed on your phone
* Have at least one location and one device defined
* Are using the same login ID as your developer account


### SmartThings Classic App (Preferred)

1. Open the SmartThings app and tap **Automation** in the bottom navigation bar.

2. Tap the SmartApps tab on top.

3. Tap **Add a SmartApp**.

4. Scroll to the last entry and then tap **My Apps**.

5. Tap the `Blockchain Event Logger` app.

6. Select the devices you want logging on your blockchain.

7. Enter the **Xooa API token** you noted earlier.

9. Click **Save**.

### SmartThings New App

1. Open the SmartThings app and tap **Automations** in the bottom navigation bar.

2. Tap **Add** (Android) or **+** (iOS).

3. If prompted, select the location you want to add the app to.

4. Tap **Done** (iOS).

5. Find `Blockchain Event Logger`, usually it will appear last and may take a few seconds to appear.

6. Tap it to continue setting it up.

7. Select which devices you want to log in the Xooa blockchain.

8. Enter the **Xooa API token** you noted earlier.

## Blockchain Event Viewer SmartApp Setup (SmartThings IDE)

Follow the same steps as `Blockchain Event Logger SmartApp Setup (SmartThings IDE)` except:

1. For section **Without GitHub integration**, copy this *groovy* script from the **smartthings** folder of the  **samples** repo: <https://raw.githubusercontent.com/Xooa/samples/master/smartthings/blockchain-event-viewer.groovy>

2. For section **With GitHub integration**,in step 7, choose `blockchain-event-viewer.groovy` from **samples** GitHub repo.

3. Skip `Blockchain Event Logger SmartApp Setup (Smartphone)` steps.

### SmartThings classic app (Preferred)

1. Open the SmartThings app and tap **Automation** in the bottom navigation bar.

2. Tap the SmartApps tab on top.

3. Tap **Add a SmartApp**.

4. Scroll to the bottom and tap **My Apps**.

5. Tap the `Blockchain Event Viewer` app.

6. Enter the **Xooa API token** you noted earlier.

7. Click **Next** to proceed to view the devices logging to the blockchain.

8. Click any device to view the past logged events for that device.

9. Input the date for which you want to view the past logged events.(Last logged date is preset)

10. Click **Save** to store **Xooa API Token** with SmartApp for future uses.

### SmartThings new app

1. Open the SmartThings app and tap **Automations** in the bottom navigation bar.

2. Tap **Add**(in android) or **+**(in IOS).

3. If prompted, select the location you want to add the app to.

4. Tap **Done**(in IOS).

5. Find `Blockchain Event Viewer`, usually it appears at the bottom of the page and may take a few seconds to appear.

6. Tap the app.

7. Enter the **Xooa API token** you noted earlier.

8. Click **Next** to proceed to view devices logging to the blockchain.

9. Click on any device to view the past logged events for that device.

10. Input the date for which you want to view the past logged events.(Latest logged date is preset)

11. Click **Save** to store the **Xooa API Token** with SmartApp for future uses.
