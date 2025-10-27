---
title: Activate and log in to Olares
description: Learn how to activate Olares for the first time, reactivate it after reinstallation, and complete secure two-factor login using the LarePass mobile app.
---

# Activate Olares

Olares uses your **Olares ID** and the **LarePass mobile app** to provide a secure and seamless authentication experience. This document guides you on how to activate Olares and complete two-factor verification during login using LarePass.

:::warning Same network required for admin users
To avoid activation failures, ensure that both your phone and the Olares device are connected to the **same network** when activating as an **admin user**.  
For **member users**, the same network requirement does **not** apply.  
:::

## Activate after one-line script installation

If you [installed Olares via the one-line script](../get-started/install-olares.md#install-olares) and completed the initial setup in the wizard:

![Activate Olares](/images/manual/larepass/activate-olares.png#bordered)

1. Open LarePass app.
2. Tap **Scan QR code** to scan the QR code on the Wizard page. 
3. Follow the on-screen instructions on LarePass to reset the login password for Olares. 

After successful activation, the LarePass app will automatically return to the home screen, and the Wizard will redirect you to the login page.

## Activate after ISO installation

If you installed Olares via ISO on PVE or are using an Olares hardware device with ISO pre-installed:

![ISO Activate](/images/manual/larepass/iso-activate.png#bordered)

1. Open LarePass app.

2. Tap **Discover nearby Olares**. Your Olares device should appear.

3. Tap **Install now** to finish the installation process.

4. Tap **Activate now** to activate Olares and complete initialization.

5. Follow the prompt to set the login password for Olares.

   ![ISO Activate](/images/manual/larepass/iso-activate-2.png#bordered)

Once complete, you can access Olares via the provided URL and your credentials.

:::tip Device not discovered?
If your phone cannot connect to the same network as the Olares device, LarePass will not be able to discover your Olares. In this case, use the [Bluetooth network configuration](manage-olares.md#set-wi-fi-via-bluetooth) feature to connect Olares to the same Wi-Fi as your phone, then repeat the activation process.
:::

## Reactivate Olares with the same Olares ID

If you have reinstalled Olares, the original instance becomes unavailable. You can reactivate the new installation using your existing Olares ID:

1. Open LarePass on your phone. You can see a red prompt: "No active Olares found".
2. Tap **Learn more** > **Reactivate** to enter the QR scan page.
3. Tap **Scan QR code** to scan the QR code on the wizard page and activate Olares.

## Two-factor verification with LarePass

When you log in to Olares, you will be promoted to complete the two-factor verification. You can either confirm the login directly in LarePass app or manually enter a 6-digit verification code.

- **To confirm login on LarePass**:
  ![2FA](/images/manual/larepass/second-confirmation.png#bordered)
  1. Open the login notification on your phone.
  2. In the message, click **Confirm** to complete the login process. 

- **To manually enter the verification code**:
  ![OTP](/images/manual/larepass/otp-larepass.jpg#bordered)

  1. On the Wizard page, select **Verify using one time password from LarePass**.
  2. Open LarePass on your phone and go to **Settings**.
  3. In the **My Olares** card, tap the authenticator to generate a one time verification code.
  4. Return to your Wizard page and enter the code to complete the login.

::: tip Note
The verification code is time-sensitive. Ensure you enter it before it expires. If it does, you will need to generate a new code.
:::

After successful verification, you'll be redirected to the Olares desktop.
