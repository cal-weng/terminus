# Manage Olares with LarePass

The LarePass app allows you to easily manage your Olares device. You can upgrade Olares，monitor system status, manage network connections, perform remote controls, and access key device information from your phone.

This guide walks you through the core management features available in LarePass.

## Prerequisites

Before you begin, ensure the following:

- You have a valid Olares ID and an activated Olares device.
- Your Olares device is powered on and connected to a network.
- Your phone and Olares device are on the same local network.
- Your current account has administrator permissions.

## Access Olares management

To access the Olares management page:

1. Open LarePass app, and go to **Settings**.
2. Tap your Olares ID to enter the Olares management interface.

## Upgrade Olares

:::warning Olares admin required
Only Olares admin can perform system updates. Updates will apply to all members within the same Olares cluster.
:::

When a new version of Olares is available, you’ll see a `Found new version` prompt on the **Olares management** page.

![Upgrade Olares](/images/manual/larepass/olares-upgrade.png#bordered)

1. Tap on the version prompt, for example, `Found new version 1.12.0-20250728`.
2. Select an upgrade method:
    - **Download and upgrade**: Olares will download the update package and install it immediately. The system is temporarily unavailable during the upgrade process.
    - **Download only**: Only downloads the update package. You can trigger the upgrade later when convenient, and Olares remains available during the download.

3. Confirm the option. A progress indicator on LarePass will show the operation status. If you selected **Download only**, you need to tap **Upgrade now** to start upgrade after download completes.
4. Wait for the upgrade to complete. Olares will resume normal operation when the upgrade finishes. You can verify the update by checking the **System version** shown in the device card at the top of the **Olares management** page.

## Remote device control

In the upper-right corner of the Olares management page, tap the <i class="material-symbols-outlined">power_settings_new</i> icon to access remote control options:

 ![Device control](/images/manual/larepass/device-control.png)

- **Restart device** – Reboots the device. Status will show `Restarting`... and revert to `Olares Running` after approximately 5–8 minutes.
- **Remote shutdown** – Powers off the device. Status will display `Powered off`. You must turn it on manually afterward.

## Configure network

Tap **Wi-Fi configuration** to view or modify the current network settings of your Olares device.

### Switch from wired to wireless network

If Olares was activated over Ethernet, you can switch to the Wi-Fi on the same network:

![Wi-Fi switch](/images/manual/larepass/switch-wifi.jpg)

1. Tap the **Wi-Fi configuration** option to enter the network selection page.
2. Tap the Wi-Fi network from the list. If the network is password-protected, enter the password and tap **Confirm**.
3. Once connected, the network switches to Wi-Fi automatically. The transition takes approximately 5 minutes. The Olares status will change to `IP changing` before it reverts to `Olares Running`.

You can switch back to the wired network following the same steps.

::: tip Wired network recommended
To ensure an optimal and stable connection, we recommend using a wired network whenever possible.
:::

### Update IP address

If your Olares device moves to a different network:

1. Connect the Olares device to your wired network via Ethernet and power it on. Ensure your phone is connected to the Wi-Fi for that same network.
2. Open LarePass on your phone and go to the **Olares management** page.
3. LarePass will automatically scan Olares device within the network. When found, Olares will appear as `IP changing` in LarePass.
4. Once IP update finishes, the status will revert to `Olares running`. This process may take 5–10 minutes.

### Set Wi-Fi via Bluetooth

If you can't connect your Olares to a wired network during activation, or if Olares is on a different wired network than your phone, LarePass won't be able to find it. This can prevent you from completing activation or device management. In such cases, use the **Bluetooth network setup** feature to connect your Olares to your phone's Wi-Fi network.
 ![Bluetooth network](/images/manual/larepass/bluetooth-network.png)

1. On the **Olares not found** page, tap the **Bluetooth network setup** option. LarePass will use your phone's Bluetooth to scan for the nearby Olares device.
2. When your device appears in the list, tap **Network setup**.
3. Select the Wi-Fi network your phone is currently connected to. If it's password-protected, enter the password and tap **Confirm**.

    ::: tip Note
    If you select a Wi-Fi network different from the one your phone is on, LarePass won't be able to detect your Olares device after it connects.
    :::

4. Olares will start switching the network.Once the process is complete, a success message will appear. If you return to the Bluetooth network setup page, you'll see that Olares' IP address has changed to your phone's Wi-Fi subnet. 
   ::: tip Note
   The process takes longer if your Olares was activated earlier as the network switch affects more services.
   :::
5. Go back to the device scan page and tap **Discover nearby Olares** to find your device. You can now proceed with [device activation](activate-olares.md) or device management as instructed in this document.

## Uninstall Olares

This will reset your device to the prepared installation phase, where you can scan the LAN to re-install and activate Olares. 

![Uninstall Olares](/images/manual/larepass/restore-to-factory.png)

::: warning Proceed with caution
This will permanently delete all your accounts info and data. Proceed with caution.
:::

1. On the **Olares management** page, tap **Restore to Factory Settings**.

2. Review the risk prompt and enter your local LarePass lock screen password. If not set, you will be prompted to set one first.

3. Wait for the uninstallation to complete. You will return to the Olares login screen when it's done.
