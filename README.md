<div align="center">

# Olares - Open-Source SelfHosted Alternative to Public Clouds <!-- omit in toc -->

[![Mission](https://img.shields.io/badge/Mission-Let%20people%20own%20their%20data%20again-purple)](#)<br />
[![Last Commit](https://img.shields.io/github/last-commit/beclab/olares)](https://github.com/beclab/olares/commits/main)
![Build Status](https://github.com/beclab/olares/actions/workflows/release-daily.yaml/badge.svg)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/beclab/olares)](https://github.com/beclab/olares/releases)
[![GitHub Repo stars](https://img.shields.io/github/stars/beclab/olares?style=social)](https://github.com/beclab/olares/stargazers)
[![Discord](https://img.shields.io/badge/Discord-7289DA?logo=discord&logoColor=white)](https://discord.com/invite/BzfqrgQPDK)
[![License](https://img.shields.io/badge/License-Olares-darkblue)](https://github.com/beclab/olares/blob/main/LICENSE.md)

<p>
  <a href="./README.md"><img alt="Readme in English" src="https://img.shields.io/badge/English-FFFFFF"></a>
  <a href="./README_CN.md"><img alt="Readme in Chinese" src="https://img.shields.io/badge/简体中文-FFFFFF"></a>
</p>

</div>

![cover](https://file.bttcdn.com/github/terminus/desktop-dark.jpeg)

*Build your local AI assistants, sync data across places, self-hosted your workspace, stream your own media, and more——all in a true home cloud made possible by Olares.*
<p align="center">
  <a href="https://olares.xyz">Website</a> ·
  <a href="https://docs.olares.xyz">Documentation</a> ·
  <a href="https://olares.xyz/larepass">Download LarePass</a> ·
  <a href="https://github.com/beclab/apps">Olares Apps</a> ·
  <a href="https://space.olares.xyz">Olares Space</a>
</p>

## Table of Contents <!-- omit in toc -->
- [Introduction](#introduction)
- [Motivation and design](#motivation-and-design)
- [Tech stacks](#tech-stacks)
- [Features](#features)
- [Feature comparison](#feature-comparison)
- [Getting started](#getting-started)
- [Project navigation](#project-navigation)
- [Contributing to Olares](#contributing-to-olares)
- [Community \& contact](#community--contact)
- [Staying ahead](#staying-ahead)
- [Special thanks](#special-thanks)
  
## Introduction

Transform your edge device into a sovereign cloud with Olares - a free, self-hosted alternative to public clouds. Powered by Kubernetes, Olares brings cloud-level capabilities to your home without compromising privacy or ease of use. By securely storing your data and accessing your self-hosted services from anywhere via Olares, you gain complete control over your digital life. 

Typical use cases of Olares include:

🤖 **Local AI**: Host and run world-class open-source AI models locally, including large language models, image generation, and speech recognition. Create custom AI assistants that integrate seamlessly with your personal data and applications, all while ensuring enhanced privacy and control. <br>

💻**Personal data repository**: Securely store, sync, and manage your photos, documents, and important files in a unified storage and access anywhere. <br>

🛠️ **Self-hosted workspace**: Create a free, powerful workspace for your team or family with open source selfhosted alternatives. <br>

🎥 **Private media server**: Host your own streaming services with your personal media collections. <br>

🏡 **Smart Home Hub**: Create a central control point for your IoT devices and home automation. <br>

🤝 **User-owned decentralized social media**: Easily install decentralized social media apps such as Mastodon, Ghost, and WordPress on Olares, allowing you to build a personal brand without the risk of being banned or paying platform commissions.<br>

📚 **Learning platform**: Explore self-hosting, container orchestration, and cloud technologies hands-on.

## Motivation and design

We believe the current state of the internet, where user data is centralized and exploited by monopolistic corporations, is deeply flawed. Our goal is to empower individuals with true data ownership and control.

Olares provides a next-generation decentralized Internet framework consisting of the following three integral components:  

- **Snowinning Protocol**: A decentralized identity and reputation system that integrates decentralized identifiers (DIDs), verifiable credentials (VCs), and reputation data. 
- **Olares OS**: An one-stop self-hosted operating system running on edge devices, allowing users to host their own data and applications.  
- **LarePass**: A comprehensive client software that securely bridges users to their Olares systems. It offers remote access, identity and device management, data storage, and productivity tools, providing a seamless interface for all Olares interactions.  

## Tech stacks

 Public clouds have IaaS, PaaS, and SaaS layers. Olares provides open-source alternatives to these layers.

  ![Tech Stacks](https://file.bttcdn.com/github/terminus/v2/tech-stack-olares.jpeg)

## Features

Olares offers a wide array of features designed to enhance security, ease of use, and development flexibility:

- **Enterprise-grade security**: Simplified network configuration using Tailscale, Headscale, Cloudflare Tunnel, and FRP.
- **Secure and permissionless application ecosystem**: Sandboxing ensures application isolation and security.
- **Unified file system and database**: Automated scaling, backups, and high availability.
- **Single sign-on**: Log in once to access all applications within Olares with a shared authentication service.
- **AI capabilities**: Comprehensive solution for GPU management, local AI model hosting, and private knowledge bases while maintaining data privacy.
- **Built-in applications**: Includes file manager, sync drive, vault, reader, app market, settings, and dashboard.
- **Seamless anywhere access**: Access your devices from anywhere using dedicated clients for mobile, desktop, and browsers.
- **Development tools**: Comprehensive development tools for effortless application development and porting.

Here are some screenshots from the UI for a sneak peek:

| Desktop–AI-Powered Personal Desktop     |  **Files**–A Secure Home to Your Data
| :--------: | :-------: |
| ![Desktop](https://file.bttcdn.com/github/terminus/v2/desktop.jpg) | ![Files](https://file.bttcdn.com/github/terminus/v2/files.jpg) |
| **Vault–1Password alternative**|**Market–App ecosystem in your control** |
| ![vault](https://file.bttcdn.com/github/terminus/v2/vault.jpg) | ![market](https://file.bttcdn.com/github/terminus/v2/market.jpg) |
|**Wise–Your digital secret garden** | **Settings–Managing Olares efficiently** |
| ![settings](https://file.bttcdn.com/github/terminus/v2/wise.jpg) | ![](https://file.bttcdn.com/github/terminus/v2/settings.jpg) |
|**Dashboard–constant Olares monitoring**  | **Profile–Your homepage on decentralized network** |
| ![dashboard](https://file.bttcdn.com/github/terminus/v2/dashboard.jpg) | ![profile](https://file.bttcdn.com/github/terminus/v2/profile.jpg) |
| **Devbox–Developing, debugging, and deploying**|**Controlhub–Managing Kubernetes clusters easily**  |
| ![Devbox](https://file.bttcdn.com/github/terminus/v2/devbox.jpg) | ![Controlhub](https://file.bttcdn.com/github/terminus/v2/controlhub.jpg)|

</div>

## Feature comparison 

To help you understand how Olares stands out in the landscape, we've created a comparison table that highlights its features alongside those of other leading solutions in the market.

**Legend:** 

- 🚀: **Auto**, indicates that the system completes the task automatically.
- ✅: **Yes**, indicates that users without a developer background can complete the setup through the product's UI prompts.
- 🛠️: **Manual Configuration**, indicates that even users with an engineering background need to refer to tutorials to complete the setup.
- ❌:  **No**, indicates that the feature is not supported.

| | Olares | Synology | TrueNAS | CasaOS | Unraid |
| --- | --- | --- | --- | --- | --- |
| Source Code License | Olares License | Closed | GPL 3.0 | Apache 2.0 | Closed |
| Built On | Kubernetes | Linux | Kubernetes | Docker | Docker |
| Multi-Node | ✅   | ❌   | ✅   | ❌   | ❌   |
| Build-in Apps | ✅ (Rich desktop apps) | ✅ (Rich desktop apps) | ❌ (CLI) | ✅ (Simple desktop apps) | ✅ (Dashboard) |
| Free Domain Name | ✅   | ✅   | ❌   | ❌   | ❌   |
| Auto SSL Certificate | 🚀  | ✅   | 🛠️ | 🛠️ | 🛠️ |
| Reverse Proxy | 🚀  | ✅   | 🛠️ | 🛠️ | 🛠️ |
| VPN Management | 🚀  | 🛠️ | 🛠️ | 🛠️ | 🛠️ |
| Graded App Entrance | 🚀  | 🛠️ | 🛠️ | 🛠️ | 🛠️ |
| Multi-User Management | ✅ User management <br>🚀 Resource isolation | ✅ User management<br>🛠️ Resource isolation | ✅ User management<br>🛠️ Resource isolation | ❌   | ✅ User management  <br>🛠️ Resource isolation |
| Single Login for All Apps | 🚀  | ❌   | ❌   | ❌   |  ❌   |
| Cross-Node Storage | 🚀 (Juicefs+<br>MinIO) | ❌   | ❌   | ❌   | ❌   |
| Database Solution | 🚀 (Built-in cloud-native solution) | 🛠️ | 🛠️ | 🛠️ | 🛠️ |
| Disaster Recovery | 🚀 (MinIO's [**Erasure Coding**](https://min.io/docs/minio/linux/operations/concepts/erasure-coding.html)**)** | ✅ RAID | ✅ RAID | ✅ RAID | ✅ Unraid Storage |
| Backup | ✅ App Data  <br>✅ User Data | ✅ User Data | ✅ User Data | ✅ User Data | ✅ User Data |
| App Sandboxing | ✅   | ❌   | ❌ (K8S's namespace) | ❌   | ❌   |
| App Ecosystem | ✅ (Official + third-party) | ✅ (Majorly official apps) | ✅ (Official + third-party submissions) | ✅ Majorly official apps | ✅ (Community app market) |
| Developer Friendly | ✅ IDE  <br>✅ CLI  <br>✅ SDK  <br>✅ Doc | ✅ CLI  <br>✅ SDK  <br>✅ Doc | ✅ CLI  <br>✅ Doc | ✅ CLI  <br>✅ Doc | ✅ Doc |
| Local LLM Hosting | 🚀  | 🛠️ | 🛠️ | 🛠️ | 🛠️ |
| Local LLM app development | 🚀 | 🛠️ | 🛠️ | 🛠️ | 🛠️ |
| Client Platforms | ✅ Android  <br>✅ iOS  <br>✅ Windows  <br>✅ Mac  <br>✅ Chrome Plugin | ✅ Android  <br>✅ iOS | ❌   | ❌   | ❌   |
| Client Functionality | ✅ (All-in-one client app) | ✅ (14 separate client apps) | ❌   | ❌   |  ❌   |

## Getting started

Refer to [Getting Started Guide](https://docs.olares.xyz/manual/get-started/) to spin up your Olares on Linux, Windows, Mac, or Raspberry Pi.

## Project navigation

Olares consists of numerous code repositories publicly available on GitHub. The current repository is responsible for the final compilation, packaging, installation, and upgrade of the operating system, while specific changes mostly take place in their corresponding repositories.

The following table lists the project directories under Olares and their corresponding repositories. Find the one that interests you:

<details>
<summary><b>Framework components</b></summary>

| Directory | Repository | Description |
| --- | --- | --- |
| [frameworks/app-service](https://github.com/beclab/olares/tree/main/frameworks/app-service) | <https://github.com/beclab/app-service> | A system framework component that provides lifecycle management and various security controls for all apps in the system. |
| [frameworks/backup-server](https://github.com/beclab/olares/tree/main/frameworks/backup-server) | <https://github.com/beclab/backup-server> | A system framework component that provides scheduled full or incremental cluster backup services. |
| [frameworks/bfl](https://github.com/beclab/olares/tree/main/frameworks/bfl) | <https://github.com/beclab/bfl> | Backend For Launcher (BFL), a system framework component serving as the user access point and aggregating and proxying interfaces of various backend services. |
| [frameworks/GPU](https://github.com/beclab/olares/tree/main/frameworks/GPU) | <https://github.com/grgalex/nvshare> | GPU sharing mechanism that allows multiple processes (or containers running on Kubernetes) to securely run on the same physical GPU concurrently, each having the whole GPU memory available. |
| [frameworks/l4-bfl-proxy](https://github.com/beclab/olares/tree/main/frameworks/l4-bfl-proxy) | <https://github.com/beclab/l4-bfl-proxy> | Layer 4 network proxy for BFL. By prereading SNI, it provides a dynamic route to pass through into the user's Ingress. |
| [frameworks/osnode-init](https://github.com/beclab/olares/tree/main/frameworks/osnode-init) | <https://github.com/beclab/osnode-init> | A system framework component that initializes node data when a new node joins the cluster. |
| [frameworks/system-server](https://github.com/beclab/olares/tree/main/frameworks/system-server) | <https://github.com/beclab/system-server> | As a part of system runtime frameworks, it provides a mechanism for security calls between apps. |
| [frameworks/tapr](https://github.com/beclab/olares/tree/main/frameworks/tapr) | <https://github.com/beclab/tapr> | Olares Application Runtime components. |

<b>System-Level Applications and Services</b>

</details>

<details>
<summary><b>System-Level Applications and Services</b></summary>

| Directory | Repository | Description |
| --- | --- | --- |
| [apps/analytic](https://github.com/beclab/olares/tree/main/apps/analytic) | <https://github.com/beclab/analytic> | Developed based on [Umami](https://github.com/umami-software/umami), Analytic is a simple, fast, privacy-focused alternative to Google Analytics. |
| [apps/market](https://github.com/beclab/olares/tree/main/apps/market) | <https://github.com/beclab/market> | This repository deploys the front-end part of the application market in Olares. |
| [apps/market-server](https://github.com/beclab/olares/tree/main/apps/market-server) | <https://github.com/beclab/market> | This repository deploys the back-end part of the application market in Olares. |
| [apps/argo](https://github.com/beclab/olares/tree/main/apps/argo) | <https://github.com/argoproj/argo-workflows> | A workflow engine for orchestrating container execution of local recommendation algorithms. |
| [apps/desktop](https://github.com/beclab/olares/tree/main/apps/desktop) | <https://github.com/beclab/desktop> | The built-in desktop application of the system. |
| [apps/devbox](https://github.com/beclab/olares/tree/main/apps/devbox) | <https://github.com/beclab/devbox> | An IDE for developers to port and develop Olares applications. |
| [apps/LarePass](https://github.com/beclab/olares/tree/main/apps/LarePass) | <https://github.com/beclab/LarePass> | A free alternative to 1Password and Bitwarden for teams and enterprises of any size Developed based on [Padloc](https://github.com/padloc/padloc). It serves as the client that helps you manage DID, Olares ID, and Olares devices. |
| [apps/files](https://github.com/beclab/olares/tree/main/apps/files) | <https://github.com/beclab/files> | A built-in file manager modified from [Filebrowser](https://github.com/filebrowser/filebrowser), providing management of files on Drive, Sync, and various Olares physical nodes. |
| [apps/notifications](https://github.com/beclab/olares/tree/main/apps/notifications) | <https://github.com/beclab/notifications> | The notifications system of Olares |
| [apps/profile](https://github.com/beclab/olares/tree/main/apps/profile) | <https://github.com/beclab/profile> | Linktree alternative in Olares|
| [apps/rsshub](https://github.com/beclab/olares/tree/main/apps/rsshub) | <https://github.com/beclab/rsshub> | A RSS subscription manager based on [RssHub](https://github.com/DIYgod/RSSHub). |
| [apps/settings](https://github.com/beclab/olares/tree/main/apps/settings) | <https://github.com/beclab/settings> | Built-in system settings. |
| [apps/system-apps](https://github.com/beclab/olares/tree/main/apps/system-apps) | <https://github.com/beclab/system-apps> | Built based on the _kubesphere/console_ project, system-service provides a self-hosted cloud platform that helps users understand and control the system's runtime status and resource usage through a visual Dashboard and feature-rich ControlHub. |
| [apps/wizard](https://github.com/beclab/olares/tree/main/apps/wizard) | <https://github.com/beclab/wizard> | A wizard application to walk users through the system activation process. |
</details>

<details>
<summary><b>Third-party Components and Services</b></summary>

| Directory | Repository | Description |
| --- | --- | --- |
| [third-party/authelia](https://github.com/beclab/olares/tree/main/third-party/authelia) | <https://github.com/beclab/authelia> | An open-source authentication and authorization server providing two-factor authentication and single sign-on (SSO) for your applications via a web portal. |
| [third-party/headscale](https://github.com/beclab/olares/tree/main/third-party/headscale) | <https://github.com/beclab/headscale> | An open source, self-hosted implementation of the Tailscale control server in Olares to manage Tailscale in LarePass across different devices. |
| [third-party/infisical](https://github.com/beclab/olares/tree/main/third-party/infisical) | <https://github.com/beclab/infisical> | An open-source secret management platform that syncs secrets across your teams/infrastructure and prevents secret leaks. |
| [third-party/juicefs](https://github.com/beclab/olares/tree/main/third-party/juicefs) | <https://github.com/beclab/juicefs-ext> | A distributed POSIX file system built on top of Redis and S3, allowing apps on different nodes to access the same data via POSIX interface. |
| [third-party/ks-console](https://github.com/beclab/olares/tree/main/third-party/ks-console) | <https://github.com/kubesphere/console> | Kubesphere console that allows for cluster management via a Web GUI. |
| [third-party/ks-installer](https://github.com/beclab/olares/tree/main/third-party/ks-installer) | <https://github.com/beclab/ks-installer-ext> | Kubesphere installer component that automatically creates Kubesphere clusters based on cluster resource definitions. |
| [third-party/kube-state-metrics](https://github.com/beclab/olares/tree/main/third-party/kube-state-metrics) | <https://github.com/beclab/kube-state-metrics> | kube-state-metrics (KSM) is a simple service that listens to the Kubernetes API server and generates metrics about the state of the objects. |
| [third-party/notification-mananger](https://github.com/beclab/olares/tree/main/third-party/notification-manager) | <https://github.com/beclab/notification-manager-ext> | Kubesphere's notification management component for unified management of multiple notification channels and custom aggregation of notification content. |
| [third-party/predixy](https://github.com/beclab/olares/tree/main/third-party/predixy) | <https://github.com/beclab/predixy> | Redis cluster proxy service that automatically identifies available nodes and adds namespace isolation. |
| [third-party/redis-cluster-operator](https://github.com/beclab/olares/tree/main/third-party/redis-cluster-operator) | <https://github.com/beclab/redis-cluster-operator> | A cloud-native tool for creating and managing Redis clusters based on Kubernetes. |
| [third-party/seafile-server](https://github.com/beclab/olares/tree/main/third-party/seafile-server) | <https://github.com/beclab/seafile-server> | The backend service of Seafile (Sync Drive) for handling data storage. |
| [third-party/seahub](https://github.com/beclab/olares/tree/main/third-party/seahub) | <https://github.com/beclab/seahub> | The front-end and middleware service of Seafile (Sync Drive) for handling file sharing, data synchronization, etc. |
| [third-party/tailscale](https://github.com/beclab/olares/tree/main/third-party/tailscale) | <https://github.com/tailscale/tailscale> | Tailscale has been integrated in LarePass of all platforms. |
</details>

<details>
<summary><b>Additional libraries and components</b></summary>

| Directory | Repository | Description |
| --- | --- | --- |
| [build/installer](https://github.com/beclab/olares/tree/main/build/installer) |     | The template for generating the installer build. |
| [build/manifest](https://github.com/beclab/olares/tree/main/build/manifest) |     | Installation build image list template. |
| [libs/fs-lib](https://github.com/beclab/olares/tree/main/libs) | <https://github.com/beclab/fs-lib> | The SDK library for the iNotify-compatible interface implemented based on JuiceFS. |
| [scripts](https://github.com/beclab/olares/tree/main/scripts) |     | Assisting scripts for generating the installer build. |
</details>

## Contributing to Olares

We are welcoming contributions in any form:

- If you want to develop your own applications on Olares, refer to:<br>
https://docs.olares.xyz/developer/develop/


- If you want to help improve Olares, refer to:<br>
https://docs.olares.xyz/developer/contribute/olares.html

## Community & contact

* [**Github Discussion**](https://github.com/beclab/olares/discussions). Best for sharing feedback and asking questions.
* [**GitHub Issues**](https://github.com/beclab/olares/issues). Best for filing bugs you encounter using Olares and submitting feature proposals. 
* [**Discord**](https://discord.com/invite/BzfqrgQPDK). Best for sharing anything Olares.

## Staying ahead  

Star the Olares project to receive instant notifications about new releases and updates.

 
![star us](https://file.bttcdn.com/github/terminus/olares.git.v2.gif)
 

## Special thanks 

The Olares project has incorporated numerous third-party open source projects, including: [Kubernetes](https://kubernetes.io/), [Kubesphere](https://github.com/kubesphere/kubesphere), [Padloc](https://padloc.app/), [K3S](https://k3s.io/), [JuiceFS](https://github.com/juicedata/juicefs), [MinIO](https://github.com/minio/minio), [Envoy](https://github.com/envoyproxy/envoy), [Authelia](https://github.com/authelia/authelia), [Infisical](https://github.com/Infisical/infisical), [Dify](https://github.com/langgenius/dify), [Seafile](https://github.com/haiwen/seafile),[HeadScale](https://headscale.net/), [tailscale](https://tailscale.com/), [Redis Operator](https://github.com/spotahome/redis-operator), [Nitro](https://nitro.jan.ai/), [RssHub](http://rsshub.app/), [predixy](https://github.com/joyieldInc/predixy), [nvshare](https://github.com/grgalex/nvshare), [LangChain](https://www.langchain.com/), [Quasar](https://quasar.dev/), [TrustWallet](https://trustwallet.com/), [Restic](https://restic.net/), [ZincSearch](https://zincsearch-docs.zinc.dev/), [filebrowser](https://filebrowser.org/), [lego](https://go-acme.github.io/lego/), [Velero](https://velero.io/), [s3rver](https://github.com/jamhall/s3rver), [Citusdata](https://www.citusdata.com/).
