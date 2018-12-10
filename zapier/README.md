
This page provides a step-by-step tutorial to integrate the Zapier app with Xooa's blockchain platform-as-a-service (PaaS).

**Project documentation:** <https://docs.xooa.com>

# Overview

This repository contains the smart contract (chaincode) for zapier. Use the Xooa console to deploy it.

Xooa provides a permanent cloud end-point for the zapier app, enabling cloud-to-cloud integration, while retaining blockchain's peer-to-peer capabilities.

## Deploy the Zapier smart contract 

1. Follow the instructions here: https://docs.xooa.com/start.html#deploy-the-smart-contract-app, selecting the **Xooa-Zapier** as the smart contract.

2. Record the **API Token** when it is shown: you will need it to authorize API requests from SmartApp.

   > **Tip:**  to regenerate the API token: 
   >
   > 1. Go to the **Identities** tab. 
   > 2. Next to the ID, select **Actions**.
   > 3. Select **Regenerate API Token**, and then select **Generate**.

___

## Set up the Zap

1. [Log in](https://zapier.com/app/login) to zapier with your Zapier credentials or sign up for a new account if you're a new user.

2. Click **Make a Zap!**.

### Set up the action zap

1. Search for the app you want to trigger action in *Xooa* for.

2. Choose the trigger you want to configure. Click **Save+Continue**.

3. Connect an account you've for that app. Click **Save+Continue**.

4. Click **Pull in Samples**. Click **Continue**.

5. Click **Add a Step** > **Action/Search**.

6. Search for *Xooa* app to choose as an action app.

7. Choose **Invoke smart contract** as Action for Xooa app. Click **Save+Continue**.

8. Click **Connect an Account**.

9. Enter Xooa Api Token you've noted earlier. Click **Yes, Continue**. Click **Save+Continue**.

10. Enter the function name you want to call from the smart contract. **SaveNewEvent** is the default smart contract function if you've not performed any changes in the smart contract.

11. Choose a unique ID among the fields fetched from the trigger app upon clicking on the icon on right to the *Key* field. A composite key can be entered if supported by smart contract.

12. Choose value you want to log with Xooa among the fields fetched from the trigger app upon clicking on the icon on right to the *Value* field.

13. Click **Continue**.

14. Click **Send Test To Xooa**. The test should be successful for you to proceed further.

15. Click **Finish**.

16. **Name your Zap** to access it or edit it in the future. Click on *off switch* to turn the zap on.

> The zap is now configured and ready to accept data and log into Xooa blockchain.


### Set up the trigger zap

1. Search for *Xooa* app to choose as a trigger app.

2. Choose **New Event** as trigger for Xooa app. Click **Save+Continue**.

3. Click **Connect an Account**.

4. Enter Xooa Api Token you've noted earlier. Click **Yes, Continue**. Click **Save+Continue**.

5. Enter the **Event Name** you want to subscribe to. You can use regex. Keep the field blank to subscribe to all events.

6. Click **Continue**.

7. Click **Pull in Samples**. Click **Continue**.

8. Click **Add a Step** > **Action/Search**.

9. Search for the app you want to take action upon receiving an event.

10. Choose the action you want to configure. Click **Save+Continue**.

11. Connect an account you've for that app. Click **Save+Continue**.

12. Set up the fields for the action. Click **Continue**.

13. **Send Test** to the action app.

14. Click **Finish**.

15. Name your Zap to access it or edit it in the future. Click on *off switch* to turn the zap on.

> The zap is now configured and ready to trigger upon events from Xooa blockchain