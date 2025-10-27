---
outline: [2, 3]
---

# OlaresManifest Specification

Every **Olares Application Chart** should include a `OlaresManifest.yaml` file in the root directory. `OlaresManifest.yaml` provides all the essential information about an Olares App. Both the **Olares Market protocol** and the Olares depend on this information to distribute and install applications.

:::info NOTE
Latest Olares Manifest version: `0.9.0`
- Added a `conflict` field in `options` to declare incompatible applications
- Removed `analytics` field in `options`
- Modified the format of the `tailscale` section
- Added a `allowedOutboundPorts` field to allow non-http protocol external access through the specified port
- Modified the format of the `ports` section

:::
:::details Changelog
`0.8.3`
- Add a `mandatory` field in the `dependencies` section for dependent applications required for the installation
- Add `tailscaleAcls` section to permit applications to open specified ports via Tailscale

`0.8.2`
- Add a `runAsUser` option to force the app to run under non root user

`0.8.1`
- Add a `ports` section to specify exposed ports for UDP or TCP

`0.7.1`
- Add new `authLevel` value `internal`
- Change `spec`>`language` to `spec`>`locale` and support i18n
  :::

Here's an example of what a `OlaresManifest.yaml` file might look like:

::: details OlaresManifest.yaml Example

```Yaml
olaresManifest.version: '0.8.0'
olaresManifest.type: app
metadata:
  name: helloworld
  title: Hello World
  description: app helloworld
  icon: https://app.cdn.olares.com/appstore/default/defaulticon.webp
  version: 0.0.1
  categories:
  - Utilities
entrances:
- name: helloworld
  port: 8080
  title: Hello World
  host: helloworld
  icon: https://app.cdn.olares.com/appstore/default/defaulticon.webp
  authLevel: private
permission:
  appCache: true
  appData: true
  userData:
  - Home/Documents/
  - Home/Pictures/
  - Home/Downloads/BTDownloads/
spec:
  versionName: '0.0.1'
  featuredImage: https://link.to/featured_image.webp
  promoteImage:
  - https://link.to/promote_image1.webp
  - https://link.to/promote_image2.webp
  fullDescription: |
    A full description of your app.
  upgradeDescription: |
    Describe what is new in this upgraded version.
  developer: Developer's Name
  website: https://link.to.your.website
  sourceCode: https://link.to.sourceCode
  submitter: Submitter's Name
  language:
  - en
  doc: https://link.to.documents
  supportArch:
  - amd64
  limitedCpu: 1000m
  limitedMemory: 1000Mi
  requiredCpu: 50m
  requiredDisk: 50Mi
  requiredMemory: 12Mi

options:
  dependencies:
  - name: olares
    type: system
    version: '>=0.1.0'
```
:::

## olaresManifest.type

- Type: `string`
- Accepted Value: `app`, `recommend`, `middleware`

Olares currently supports 3 types of applications, each requiring different fields. This document uses `app` as an example to explain each field. For information on other types, please refer to the corresponding configuration guide.
- [Recommend Configuration Guide](recommend.md)

:::info Example
```Yaml
olaresManifest.type: app
```
:::

## olaresManifest.version

- Type: `string`

As Olares evolves, the configuration specification of `OlaresManifest.yaml` may change. You can identify whether these changes will affect your application by checking the `olaresManifest.version`. The `olaresManifest.version` consists of three integers separated by periods.

- An increase in the **first digit** indicates the introduction of incompatible configuration items. Applications that haven't updated their `OlaresManifest.yaml` will be unable to distribute or install.
- An increase in the **second digit** signifies changes in the mandatory fields for distribution and installation. However, the Olares remains compatible with the application distribution and installation of previous configuration versions. We recommend developers to promptly update and upgrade the application's `OlaresManifest.yaml` file.
- A change in the **third digit** does not affect the application's distribution and installation.

Developers can use 1-3 digit version numbers to indicate the application's configuration version. Here are some examples of valid versions:
```Yaml
olaresManifest.version: 1
olaresManifest.version: 1.1.0
olaresManifest.version: '2.2'
olaresManifest.version: "3.0.122"
```

## Metadata

Basic information about the app shown in the system and Olares Market.

:::info Example
```Yaml
metadata:
  name: nextcloud
  title: Nextcloud
  description: The productivity platform that keeps you in control
  icon: https://app.cdn.olares.com/appstore/nextcloud/icon.png
  version: 0.0.2
  categories:
  - Utilities
  - Productivity
```
:::

### name

- Type: `string`
- Accepted Value: `[a-z][a-z0-9]?`

App’s namespace in Olares, lowercase alphanumeric characters only. It can be up to 30 characters, and needs to be consistent with `FolderName` and `name` field in `Chart.yaml`.

### title

- Type: `string`

The title of your app title shown in the Olares Market.  Must be within `30` characters.

### description

- Type: `string`

A short description appears below app title in the Olares Market.

### icon

- Type: `url`

Your app icon that appears in the Olares Market.

The app's icon must be a `PNG` or `WEBP` format file, up to `512 KB`, with a size of `256x256 px`.

### version

- Type: `string`

The **Chart Version** of the application. It should be incremented each time the content in the **Chart** changes. It should follow the [Semantic Versioning 2.0.0](https://semver.org/) and needs to be consistent with the `version` field in `Chart.yaml`.

### categories

- Type: `list<string>`
- Accepted Value: `Blockchain`, `Utilities`, `Social Network`, `Entertainment`, `Productivity`

Used to display your app on different category page in Olares Market.

## Entrances

The number of entrances through which to access the app.  You must specify at least 1 access method, with a maximum of 10 allowed.

:::info Example
```Yaml
entrances:
- name: a
  host: firefox
  port: 3000
  title: Firefox
  authLevel: public
  invisible: false
- name: b
  host: firefox
  port: 3001
  title: admin
```
:::

### name

- Type: `string`
- Accepted Value: `[a-z]([-a-z0-9]*[a-z0-9])?`

  Name of the Entrance. It can be up to `63` characters, and needs to be unique in an app.

### port

- Type: `int`
- Accepted Value: `0-65535`

### host

- Type: `string`
- Accepted Value: `[a-z]([-a-z0-9]*[a-z0-9])?`

  Ingress name of current entrance, lowercase alphanumeric characters and `-` only. It can be up to `63` characters.

### title

- Type: `string`

Title that appears in the Olares desktop after installed. It can be up to `30` characters.

### icon

- Type: `url`
- Optional

Icon that appears in the Olares desktop after installed. The app's icon must be a `PNG` or `WEBP` format file, up to `512 KB`, with a size of `256x256 px`.

### authLevel

- Type: `string`
- Accepted Value: `public`, `private`, `internal`
- Default: `private`
- Optional

Specify the authentication level of the entrance.
- **Public**: Accessible by anyone on the Internet without restrictions.
- **Private**: Requires authorization for access from both internal and external networks.
- **Internal**: Requires authorization for access from external networks. No authentication is required when accessing from within the internal network (via LAN/VPN).

### invisible

- Type: `boolean`
- Default: `false`
- Optional

When `invisible` is `true`, the entrance will not be displayed on the Olares desktop.

### openMethod

- Type: `string`
- Accepted Value: `default`, `iframe`, `window`
- Default: `default`
- Optional

Explicitly defines how to open this entrance in Desktop.

The `iframe` creates a new window within the desktop window through an iframe. The `window` opens a new tab in the browser. The `default` follows the system setting, which is `iframe` by default.

### windowPushState
- Type: `boolean`
- Default: `false`
- Optional

When embedding the application in an iframe on the desktop, the application's URL may change dynamically. Due to browser's same-origin policy, the desktop (parent window) cannot directly detect these changes in the iframe URL. Consequently, if you reopen the application tab, it will display the initial URL instead of the updated one.

To ensure a seamless user experience, you can enable this option by setting it to true. This action prompts the gateway to automatically inject the following code into the iframe. This code sends an event to the parent window (desktop) whenever the iframe's URL changes. As a result, the desktop can track URL changes and open the correct page.

::: details Code
```Javascript
<script>
  (function () {
    if (window.top == window) {
        return;
    }
    const originalPushState = history.pushState;
    const pushStateEvent = new Event("pushstate");
    history.pushState = function (...args) {
      originalPushState.apply(this, args);
      window.dispatchEvent(pushStateEvent);
    };
    window.addEventListener("pushstate", () => {
      window.parent.postMessage(
        {type: "locationHref", message: location.href},
        "*"
      );
    });
  })();
</script>
```
:::

## Ports

Specify exposed ports

:::info Example
```Yaml
ports:
- name: rdp-tcp             # Name of the entrance that provides service
  host: windows-svc         # Ingress name of the entrance that provides service
  port: 3389                # Port of the entrance that provides service
  exposePort: 46879         # The port to be exposed can only be assigned to one application at a time within the cluster.
  addToTailscaleAcl: true   # Automatically added to Tailscle's ACL
```
:::

Olares will expose the ports you specify for an application, which are accessible via the application domain name in the local network, for example: `84864c1f.your_olares_id.olares.com:46879`. For each port you expose, Olares configures both TCP and UDP with the same port number.

When the `addToTailscaleAcl` field is set to `true`, the port will be automatically added to the Tailscale's ACL, and there is no need to configure in the tailscale section.

:::info NOTE
The exposed ports can only be accessed on the local network or through a VPN.
:::

## Permission

:::info Example
```Yaml
permission:
  appCache: true
  appData: true
  userData:
  - /Home/  
  sysData:
  - dataType: legacy_prowlarr
    appName: prowlarr
    port: 9696
    group: api.prowlarr
    version: v2
    ops:
    - All
```
:::

### appCache

- Type: `boolean`
- Optional

Whether the app requires read and write permission to the `Cache` folder. If `.Values.userspace.appCache` is used in the deployment YAML, then `appCache` must be set to `true`.

### appData

- Type: `boolean`
- Optional

Whether the app requires read and write permission to the `Data` folder. If `.Values.userspace.appData` is used in the deployment YAML, then `appData` must be set to `true`.

### userData

- Type: `list<string>`
- Optional

Whether the app requires read and write permission to user's `Home` folder. List all directories that the application needs to access under the user's `Home`. All `userData` directory configured in the deployment YAML, must be included here.

### sysData

- Type: `list<map>`
- Optional

Declare the list of APIs that this app needs to access.

:::info Example
```Yaml
  sysData:
  - group: service.bfl
    dataType: app
    version: v1
    ops:
    - InstallDevApp
  - dataType: legacy_prowlarr
    appName: prowlarr
    port: 9696
    group: api.prowlarr
    version: v2
    ops:
    - All
```
:::

All system API [providers](../advanced/provider.md) are list below:
| Group | version | dataType | ops |
| ----------- | ----------- | ----------- | ----------- |
| service.appstore | v1 | app | InstallDevApp, UninstallDevApp
| message-dispatcher.system-server | v1 | event | Create, List
| service.desktop | v1 | ai_message | AIMessage
| service.did | v1 | did | ResolveByDID, ResolveByName, Verify
| api.intent | v1 | legacy_api | POST
| service.intent | v1 | intent | RegisterIntentFilter, UnregisterIntentFilter, SendIntent, QueryIntent, ListDefaultChoice, CreateDefaultChoice, RemoveDefaultChoice, ReplaceDefaultChoice
| service.message | v1 | message | GetContactLogs, GetMessages, Message
| service.notification | v1 | message | Create
| service.notification | v1 | token | Create
| service.search | v1 | search | Input, Delete, InputRSS, DeleteRSS, QueryRSS, QuestionAI
| secret.infisical | v1 | secret | CreateSecret, RetrieveSecret
| secret.vault | v1 | key | List, Info, Sign

## Tailscale
- Type: `map`
- Optional

Allow applications to add Access Control Lists (ACL) in Tailscale to open specified ports.

:::info Example
```Yaml
tailscale:
  acls:
  - proto: tcp
    dst:
    - "*:46879"
  - proto: "" # Optional. If not specified, all supported protocols will be allowed.
    dst:
    -  "*:4557"    
```
:::

## Spec
Additional information about the application, primarily used for display in the Olares Market.

:::info Example
```Yaml
spec:
  namespace: os-system 
  # optional. Install the app to a specified namespace, e.g. os-system, user-space, user-system
  
  versionName: '10.8.11' 
  # The version of the application that this chart contains. It is recommended to enclose the version number in quotes. This value corresponds to the appVersion field in the `Chart.yaml` file. Note that it is not related to the `version` field.

  featuredImage: https://app.cdn.olares.com/appstore/jellyfin/promote_image_1.jpg
  # The featured image is displayed when the app is featured in the Market.

  promoteImage:
  - https://app.cdn.olares.com/appstore/jellyfin/promote_image_1.jpg
  - https://app.cdn.olares.com/appstore/jellyfin/promote_image_2.jpg
  - https://app.cdn.olares.com/appstore/jellyfin/promote_image_3.jpg
  fullDescription: |
    Jellyfin is the volunteer-built media solution that puts you in control of your media. Stream to any device from your own server, with no strings attached. Your media, your server, your way.
  upgradeDescription: |
    upgrade descriptions
  developer: Jellyfin
  website: https://jellyfin.org/
  doc: https://jellyfin.org/docs/
  sourceCode: https://github.com/jellyfin/jellyfin
  submitter: Olares
  locale:
  - en-US
  - zh-CN
  # List languages and regions supported by this app

  requiredMemory: 256Mi
  requiredDisk: 128Mi
  requiredCpu: 0.5
  # Specifies the minimum resources required to install and run the application. Once the app is installed, the system will reserve these resources to ensure optimal performance.

  limitedDisk: 256Mi
  limitedCpu: 1
  limitedMemory: 512Mi
  # Specifies the maximum resource limits for the application. If the app exceeds these limits, it will be temporarily suspended to prevent system overload and ensure stability.

  legal:
  - text: Community Standards
    url: https://jellyfin.org/docs/general/community-standards/
  - text: Security policy
    url: https://github.com/jellyfin/jellyfin/security/policy
  license:
  - text: GPL-2.0
    url: https://github.com/jellyfin/jellyfin/blob/master/LICENSE
  supportClient:
  - android: https://play.google.com/store/apps/details?id=org.jellyfin.mobile
  - ios: https://apps.apple.com/us/app/jellyfin-mobile/id1480192618
```
:::

### i18n

To add multi-language support for your app in Olares Market:

1. Create an `i18n` folder in the Olares Application Chart root directory.
2. In the `i18n` folder, create separate subdirectories for each supported locale.
3. In each locale subdirectory, place a localized version of the `OlaresManifest.yaml` file.

Olares Market will automatically display the content of the corresponding "OlaresManifest.yaml" file based on users' locale settings.
:::info Example
```
.
├── Chart.yaml
├── README.md
├── OlaresManifest.yaml
├── i18n
│   ├── en-US
│   │   └── OlaresManifest.yaml
│   └── zh-CN
│       └── OlaresManifest.yaml
├── owners
├── templates
│   └── deployment.yaml
└── values.yaml
```
:::
Currently, you can add i18n content for the following fields:
```Yaml
metadata:
  description:
  title:
spec:
  fullDescription:
  upgradeDescription:
```

### supportArch
- Type: `list<string>`
- Accepted Value: `amd64`, `arm64`
- Optional

Specifies the CPU architecture supported by the application. Currently only `amd64` and `arm64` are available.

:::info Example
```yaml
spec:
  supportArch:
  - amd64
  - arm64
```
:::

:::info NOTE
Olares does not support mixed-architecture clusters for now.
:::

### onlyAdmin
- Type: `boolean`
- Default: `false`
- Optional

When set to `true`, only the admin can install this app.

### runAsUser
- Type: `boolean`
- Optional

When set to `true`, Olares forces the application to run under user ID `1000` (as a non-root user).

## Middleware
- Type: `map`
- Optional

The Olares provides highly available middleware services. Developers do not need to install middleware repeatedly. Just simply add required middleware here, You can then directly use the corresponding middleware information in the application's deployment YAML file.

Use the `scripts` field to specify scripts that should be executed after the database is created. Additionally, use the `extension` field to add the corresponding extension in the  database.

:::info Example
```Yaml
middleware:
  postgres:
    username: immich
    databases:
    - name: immich
      extensions:
      - vectors
      - earthdistance
      scripts:
      - BEGIN;                                           
      - ALTER DATABASE $databasename SET search_path TO "$user", public, vectors;
      - ALTER SCHEMA vectors OWNER TO $dbusername;
      - COMMIT;
      # The OS provides two variables, $databasename and $dbusername, which will be replaced by Olares Application Runtime when the command is executed.
  redis:
    password: password
    namespace: db0
  mongodb:
    username: chromium
    databases:
    - name: chromium
      script:
      - 'db.getSiblingDB("$databasename").myCollection.insertOne({ x: 111 });'
      # Please make sure each line is a complete query.
```
:::

Use the middleware information in deployment YAML

```yaml
- name: DB_POSTGRESDB_DATABASE # The database name you configured in OlaresManifest, specified in middleware.postgres.databases[i].name
  value: {{ .Values.postgres.databases.<dbname> }}
- name: DB_POSTGRESDB_HOST
  value: {{ .Values.postgres.host }}
- name: DB_POSTGRESDB_PORT
  value: "{{ .Values.postgres.port }}"
- name: DB_POSTGRESDB_USER
  value: {{ .Values.postgres.username }}
- name: DB_POSTGRESDB_PASSWORD
  value: {{ .Values.postgres.password }}


# For mongodb, the corresponding value is as follows
host --> {{ .Values.mongodb.host }}
port --> "{{ .Values.mongodb.port }}"  # The port and password in the yaml file need to be enclosed in double quotes.
username --> {{ .Values.mongodb.username }}
password --> "{{ .Values.mongodb.password }}" # The port and password in the yaml file need to be enclosed in double quotes.
databases --> "{{ .Values.mongodb.databases }}" # The value type of database is a map. You can get the database using {{ .Values.mongodb.databases.<dbname> }}. The <dbname> is the name you configured in OlaresManifest, specified in middleware.mongodb.databases[i].name


# For Redis, the corresponding value is as follows
host --> {{ .Values.redis.host }}
port --> "{{ .Values.redis.port }}"
password --> "{{ .Values.redis.password }}"

```

## Options

Configure system-related options here.

### policies
- Type: `map`
- Optional

Define detailed access control for subdomains of the app.

:::info Example
```yaml
options:
  policies:
    - uriRegex: /$
      level: two_factor
      oneTime: false
      validDuration: 3600s
      entranceName: gitlab
```
:::

### clusterScoped
- Type: `map`
- Optional

Whether this app is installed for all users in an Olares cluster.

:::info Example For Server
```yaml
metadata:
  name: gitlab
options:
  appScope:
    clusterScoped: true
    appRef:
      - gitlabclienta #app name of clients
      - gitlabclientb
```
:::

:::info Example For Client
```yaml
metadata:
  name: gitlabclienta
options:
  dependencies:
    - name: olares
      version: ">=0.3.6-0"
      type: system
    - name: gitlab #app name of server
      version: ">=0.0.1"
      type: application
      mandatory: true
```
:::


### dependencies
- Type: `list<map>`

Specify the dependencies and requirements for your application. It includes other applications that your app depends on, as well as any specific operating system (OS) version requirements.

If this application requires other dependent applications for proper installation, you should set the `mandatory` field to `true`.

:::info Example
```yaml
options:
  dependencies:
    - name: olares
      version: ">=1.0.0-0"
      type: system
    - name: mongodb
      version: ">=6.0.0-0"
      type: middleware
      mandatory: true # Set this field to true if the dependency needs to be installed first.
```
:::

### websocket
- Type: `map`
- Optional

Enable websocket for the app. Refer to [websocket](../advanced/websocket.md) for more information.

:::info Example
```yaml
options:
  websocket:
    url: /ws/message
    port: 8888
```
:::

### resetCookie
- Type: `map`
- Optional

If the app requires cookies, please enable this feature. Refer to [cookie](../advanced/cookie.md) for more information.

:::info Example
```yaml
options:
  resetCookie:
    enabled: true
```
:::

### upload
- Type: `map`
- Optional

The Olares Application Runtime includes a built-in file upload component designed to simplify the file upload process in your application. Refer to [upload](../advanced/file-upload.md) for more information.

:::info Example
```yaml
upload:
  # The types of files that are allowed to be uploaded, * stands for any type, The type of the uploaded file must be in the list.
  fileType:
    - pdf
  # The path of 'dest' must be a mountPath
  dest: /appdata
  # The maximum size of file, in bytes
  limitedSize: 3729747942
```
:::

### mobileSupported
- Type: `boolean`
- Default: `false`
- Optional

Determine whether the application is compatible with mobile web browsers and can be displayed on the mobile version of Olares Desktop. Enable this option if the app is optimized for mobile web browsers. This will make the app visible and accessible on the mobile version of Olares Desktop.

:::info Example
```yaml
mobileSupported: true
```
:::

### oidc
- Type: `map`
- Optional

The Olares includes a built-in OpenID Connect authentication component to simplify identity verification of users. Enable this option to use OpenID in your app.
```yaml
# OpenID related variables in yaml
{{ .Values.oidc.client.id }}
{{ .Values.oidc.client.secret }}
{{ .Values.oidc.issuer }}
```

:::info Example
```yaml
oidc:
  enabled: true
  redirectUri: /path/to/uri
  entranceName: navidrome
```
:::

### apiTimeout
- Type: `int`
- Optional

Specifies the timeout limit for API providers in seconds. The default value is `15`. Use `0` to allow an unlimited API connection.

:::info Example
```yaml
apiTimeout: 0
```
:::



### allowedOutboundPorts
- Type: `map`
- Optional

​​The specified ports will be opened to allow external access via non-HTTP protocols, such as SMTP.

:::info Example
```yaml
allowedOutboundPorts:
  - 465
  - 587
```
:::   