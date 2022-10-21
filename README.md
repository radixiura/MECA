<h3 align="center">
  <a href="#"><img src="https://i.ibb.co/cbDzFRL/castle.jpg" alt="Memory Castle Project" width="400"></a>
  
  <br>
  Memory Castle. <b>Place your dreams in a safe zone!</b>
  <br>
</h3>

-----

<p align="center">
  <a href="https://github.com/cossacklabs/acra/releases"><img src="https://img.shields.io/github/release/cossacklabs/acra.svg" alt="GitHub release"></a>
  <a href="https://circleci.com/gh/cossacklabs/acra"><img src="https://circleci.com/gh/cossacklabs/acra/tree/master.svg?style=shield" alt="Circle CI"></a>
  <a href='https://coveralls.io/github/cossacklabs/themis'><img src='https://coveralls.io/repos/github/cossacklabs/themis/badge.svg?branch=master' alt='Coverage Status' /></a>
  <a href='https://goreportcard.com/report/github.com/cossacklabs/acra'><img class="badge" tag="github.com/cossacklabs/acra" src="https://goreportcard.com/badge/github.com/cossacklabs/acra"></a>
  <a href='https://godoc.org/github.com/cossacklabs/acra'><img src='https://godoc.org/github.com/cossacklabs/acra?status.svg'  alt='godoc'/></a>
  <br/><a href="https://github.com/cossacklabs/acra/releases/latest"><img src="https://img.shields.io/badge/Server%20Platforms-Ubuntu%20%7C%20Debian%20%7C%20CentOS-green.svg" alt="Server platforms"></a>
 
</p>
<br>


| [Acra Engineering Examples](https://github.com/cossacklabs/acra-engineering-demo) | [Documentation and tutorials](https://docs.cossacklabs.com/acra/) | [Installation](https://github.com/cossacklabs/acra#installation-and-launch) | [Acra feedback](#acra-feedback) |
| ---- | ---- | ---- | ---- |


## What is MECA Project?
Meca - lallalala.

Meca provides u to memory

![meca_logo](https://user-images.githubusercontent.com/69742759/197078765-cdb91170-23ae-42ba-b3d1-b148abbd5d33.jpg)

<table><thead><tr><th>Perfect Acra-compatible applications</th>
<th>Typical industries</th></tr></thead>
<tbody><tr><td>Web and mobile apps that store data in a centralised database or object storage</td>
<td rowspan=3><ul>
<li>Healthcare, patient apps</li>
<li>Finance, fintech, neobanking</li>
<li>SaaS</li>
<li>Critical infrastructures</li>
<li>Apps with > 1000 users</li></ul></td>
</tr><tr><td>IoT apps that collect telemetry and process data in the cloud</td>
</tr><tr><td>High-load data processing apps</td>
</tr></tbody></table>

Acra gives you tools for encrypting each sensitive data record (data field, database cell, json) before storing them in the database / file storage. And then decrypting them in a secure compartmented area (on Acra side). Acra allows to encrypt data as early as possible and operate on encrypted data.

Acra's [cryptographic design](https://docs.cossacklabs.com/acra/acra-in-depth/security-design/) ensures that no secret (password, key, etc.) leaked from the application or database will be sufficient for decryption of the protected data. Acra minimises the leakage scope, detects unauthorised behavior, and prevents the leakage, informing operators of the incident underway.

This is [Acra Community Edition](https://www.cossacklabs.com/acra/#pricing), it's free for commercial and non-commercial use, forever.



### Major security features

<table><tbody><tr><tr><td><li><a href="https://docs.cossacklabs.com/acra/security-controls/encryption/">Application-level encryption</a></li></td><td> encryption on client-side and/or Acra-side – each data field is encrypted using unique encryption keys.</td>
</tr><tr><td><li>Selective encryption </li></td><td>you select which columns to encrypt to balance good security and performance.</td>
</tr><tr><td><li><a href="https://docs.cossacklabs.com/acra/acra-in-depth/cryptography-and-key-management/" target=_blank>Fast and reliable crypto</a></li></td><td>two crypto-envelopes: <a href="https://docs.cossacklabs.com/acra/acra-in-depth/data-structures/">AcraBlocks and AcraStructs</a>.<br>AcraBlocks are fast symmetric containers, use them by default.<br>AcraStructs are asymmetric containers, use them for client-side encryption.</td>
</tr><tr><td><li><a href="https://docs.cossacklabs.com/acra/security-controls/searchable-encryption/">Searchable encryption</a></li></td><td>search through encrypted data without decryption. Designed for <i>exact</i> queries, based on AES-GCM and blind index.</td>
</tr><tr><td><li><a href="https://docs.cossacklabs.com/acra/security-controls/masking/">Masking / anonymization</a></li></td><td>use full or partial masking to remove or mask sensitive data.</td>
</tr><tr><td><li><a href="https://docs.cossacklabs.com/acra/security-controls/tokenization/">Tokenization</a></li></td><td>substitute sensitive data with a token and match it to original only when needed.</td>
</tr><tr><td><li><a href="https://docs.cossacklabs.com/acra/security-controls/key-management/">Basic key management tools</a></li></td><td>built-in tools for key generation, export, backup, rotation, etc.</td>
</tr><tr><td><li><a href="https://docs.cossacklabs.com/acra/security-controls/sql-firewall/">Blocking suspicious SQL queries</a></li></td><td>through a built-in SQL firewall.</td>
</tr><tr><td><li><a href="https://docs.cossacklabs.com/acra/security-controls/intrusion-detection/">Intrusion detection</a></li></td><td> using poison records (honey tokens) to warn about suspicious behaviour.</td>
</tr><tr><td><li><a href="https://docs.cossacklabs.com/acra/security-controls/key-management/operations/rotation/">Key rotation without data re-encryption</a> ᵉ </li></td><td rowspan=3>available for <a href="https://www.cossacklabs.com/acra/#pricing" target="_blank">Acra Enterprise</a> users.</td>
</tr><tr><td><li><a href="https://docs.cossacklabs.com/acra/configuring-maintaining/key-storing/kms-integration/">KMS support</a> ᵉ </li></td>
</tr><tr><td><li><a href="https://docs.cossacklabs.com/acra/security-controls/security-logging-and-events/audit-logging/">Cryptographically protected audit log</a> ᵉ </li></td>
</tr></tbody></table>

Acra delivers different layers of defense for different parts and stages of the data lifecycle. This is what **defence in depth** is – an independent set of security controls aimed at mitigating multiple risks in case of an attacker crossing the outer perimeter. 

### Multiple ways to integrate

<table><tbody>
<tr><td><li> <a href="https://docs.cossacklabs.com/acra/acra-in-depth/architecture/acraserver/" target=_blank>AcraServer</a>: transparent SQL proxy </li></td><td> all Acra features packed into a database proxy that parses traffic between an app and a database and applies security functions where appropriate. </td></tr>
<tr><td><li> <a href="https://docs.cossacklabs.com/acra/acra-in-depth/architecture/acratranslator/" target=_blank>AcraTranslator</a>: encryption-as-a-service API </li></td><td> API server, that exposes most of Acra’s features as HTTP / gRPC API with traffic protection. </td></tr>
<tr><td><li> <a href="https://docs.cossacklabs.com/acra/security-controls/transport-security/acra-connector/" target=_blank>AcraConnector</a>: transport authentication and encryption </li></td><td> optional client-side service for authentication and transport encryption.</td></tr>
<tr><td><li> <a href="https://docs.cossacklabs.com/acra/acra-in-depth/architecture/anyproxy/" target=_blank>AnyProxy</a>: use Acra with any database / datastore via SDK ᵉ </li></td><td rowspan=5>available for <a href="https://www.cossacklabs.com/acra/#pricing" target="_blank">Acra Enterprise</a> users.</td></tr>
<tr><td><li> <a href="https://docs.cossacklabs.com/acra/acra-in-depth/architecture/sdks/acrawriter/" target=_blank>AcraWriter</a>: SDK for client-side encryption ᵉ</li></td>
<tr><td><li> <a href="https://docs.cossacklabs.com/acra/acra-in-depth/architecture/sdks/acrareader/" target=_blank>AcraReader</a>: SDK for client-side decryption ᵉ</li></td></tr>
<tr><td><li> <a href="https://docs.cossacklabs.com/acra/acra-in-depth/architecture/sdks/acratranslator-sdk/" target=_blank>SDK for AcraTranslator</a>: client-side SDK that encapsulates AcraTranslator's API ᵉ </li></td></tr>
<tr><td><li> <a href="https://docs.cossacklabs.com/acra/guides/integrating-acra-translator-into-new-infrastructure/http_api/#bulk-processing-api-enterprise" target=_blank>Bulk API for AcraTranslator</a> ᵉ </li></td></tr>
</tbody></table>


## Cryptography

Acra relies on our cryptographic library [Themis](https://www.cossacklabs.com/themis/), which implements high-level cryptosystems based on the best available [open-source implementations](https://docs.cossacklabs.com/themis/crypto-theory/cryptography-donors/) of the [most reliable ciphers](https://docs.cossacklabs.com/themis/architecture/soter/). Acra strictly doesn't contain self-made cryptographic primitives or obscure ciphers. 

To deliver its unique guarantees, Acra relies on the combination of well-known ciphers and smart key management scheme. See [Cryptography and key management](https://docs.cossacklabs.com/acra/acra-in-depth/cryptography-and-key-management/).

<table><tbody>
<tr><td> Default crypto-primitive source </td><td> OpenSSL </td></tr>
<tr><td> Supported crypto-primitive sources ᵉ<td> BoringSSL, LibreSSL, FIPS-compliant, GOST-compliant, HSM</td></tr>
<tr><td> Storage encryption (<a href="https://docs.cossacklabs.com/acra/acra-in-depth/data-structures/acrablock/" target=_blank>AcraBlocks</a>) </td><td> AES-256-GCM + AES-256-GCM </td></tr>
<tr><td> Storage encryption (<a href="https://docs.cossacklabs.com/acra/acra-in-depth/data-structures/acrastruct/" target=_blank>AcraStructs</a>) </td><td> AES-256-GCM + ECDH </td></tr>
<tr><td> <a href="https://docs.cossacklabs.com/acra/security-controls/transport-security/" target=_blank>Transport encryption</a> </td><td> TLS v1.2+ or Themis Secure Session </td></tr>
<tr><td> <a href="https://docs.cossacklabs.com/acra/acra-in-depth/architecture/key-storage-and-kms/" target=_blank>KMS integration</a> ᵉ</td><td> Amazon KMS, Google Cloud Platform KMS, HashiCorp Vault, Keywhiz, etc </td></tr>
</tbody></table>

ᵉ — available in the [Enterprise version of Acra](https://www.cossacklabs.com/acra/#pricing/) only. [Drop us an email](mailto:sales@cossacklabs.com) to get a full list of features and a quote.

## How does Acra work?

Acra consists of several services and utilities. Acra services allow you to construct infinitely sophisticated data flows that are perfectly suited to your exact infrastructure. Depending on your architecture and use case, you might need to deploy only basic services or all of them.

* **Security enforcement components**: services where "encryption happens". One of them is required: AcraServer, AcraTranslator, AnyProxy, or client-side SDKs.
* **Key storage:** datastores where Acra keeps encrypted keys: Redis, table in your database, any KV store. One of them is required.
* **Master key storage:** KMS, Vault. One of them is strongly recommended.
* **Additional services and utils:** key management utils, data migration scripts, transport security service, policy management tools. Any of them are optional.

Refer to [Acra-in-depth / Architecture](https://docs.cossacklabs.com/acra/acra-in-depth/architecture/) to learn more about Acra components. Refer to [Acra-in-depth / Data flow](https://docs.cossacklabs.com/acra/acra-in-depth/data-flow/) to see more typical Acra-based dataflows and deployments.


### Protecting data in SQL databases using AcraServer

Let's see [the simplest dataflow with AcraServer](https://docs.cossacklabs.com/acra/acra-in-depth/data-flow/#simplest-version-with-sql-proxy). 

AcraServer works as transparent encryption/decryption proxy with SQL databases. The application doesn't know that the data is encrypted before it gets to the database, the database also doesn't know that someone has encrypted the data. That's why we often call this mode a "transparent encryption".

<p align="center"><img src="https://github.com/cossacklabs/acra/wiki/Images/readme/app-as-db.png" alt="Server-side encryption and decryption using AcraServer" width="1600"></p>

_You have a client-side application that talks to the SQL database. You add AcraServer in the middle, working as SQL proxy, and point application to it._

**This is what the process of writing and reading the data to/from a database looks like:**

1. You deploy AcraServer and configure it: connection to the database, TLS certificates, select which fields to encrypt, mask or tokenise, enable SQL request firewall, etc.
2. Once AcraServer is deployed, it is ready to accept SQL requests.
3. You point the client-side application to the AcraServer instead of the SQL database.
4. On receiving SQL queries from the app, AcraServer parses each query and performs security operations: encryption, masking, tokenisation. To know which values to change, AcraServer uses a configuration file where you have described which columns should be encrypted, masked, tokenised.
5. After performing the operation, AcraServer passes the modified queries to the database, and the database response – back to the client application. Suppose you select to encrypt the email field: it means that original string is encrypted into [cryptographic container](https://docs.cossacklabs.com/acra/acra-in-depth/data-structures/) and sent to the database as binary data.
6. When the client application wants to read the data, it sends a SELECT query to the AcraServer that sends it to the database. 
7. Upon retrieving the database response, AcraServer tries to decrypt, demask, detokenise specified fields, and returns them to the application.
8. Application receives data in plaintext.

Except for data processing operations, AcraServer also analyses SQL queries: blocks the unwanted ones using the built-in configurable SQL firewall, detect SQL injections using poison records, sends logs and metrics, and alerts your Ops team in suspicious cases.

Check out the [Guide: Integrating AcraServer into infrastructure](https://docs.cossacklabs.com/acra/guides/integrating-acra-server-into-infrastructure/) to learn more about AcraServer features and how to use them.

### Protecting data in any file storage using AcraTranslator

Let's see the [simplest dataflow with AcraTranslator](https://docs.cossacklabs.com/acra/acra-in-depth/data-flow/#simplest-version-with-api-service).

AcraTranslator works as Encryption-as-a-Service using HTTP and gRPC API. The application sends API request to the AcraTranslator with data fields and operations (encryption, decryption, tokenisation, detokenisation, etc). The application is responsible for storing the encrypted data in the database (NoSQL, KV store, SQL, AWS S3 – any) and communicating with AcraTranslator to decrypt it back.

AcraTranslator and AcraServer are fully independent server-side components and can be used together or separately depending on your infrastructure.

<p align="center"><img src="https://github.com/cossacklabs/acra/wiki/Images/readme/app-at-app-db.png" alt="Server-side encryption and decryption using AcraTranslator" width="700"></p>

_You have a client-side application that knows which fields to encrypt, decrypt, tokenise, and where to store them. You add AcraTranslator, and teach the application to perform API calls to use it._

**This is what the process of writing and reading the data to/from a database looks like:**

1. You deploy AcraTranslator in your infrastructure and configure TLS certificates.
2. Once AcraTranslator is deployed, it is ready to accept API requests.
3. Your application calls AcraTranslator and sends data fields and operations on them (encryption, decryption, tokenisation, detokenisation).
4. On receiving API requests, AcraTranslator performs the required operation and sends the result back to the app. Suppose the app sends the "email" field and "encrypt" operation. In that case, the original string is encrypted into [cryptographic container](https://docs.cossacklabs.com/acra/acra-in-depth/data-structures/) and sent back to the app as binary data.
5. The application takes encrypted data and stores it in the database/datastore.
6. Once the application needs to get plaintext data, it reads encrypted data from the database/datastore, and sends an API request to the AcraTranslator. Suppose the app sends the "email" field and "decrypt" operation. In that case, the original data (binary blob) is decrypted to a string and sent to the app back. 

Except for data processing operations, AcraTranslator also analyses API queries: detects intrusions using poison records, sends logs and metrics, and alerts your Ops team in suspicious cases.

Check out the [Guide: Integrating AcraTranslator into infrastructure](https://docs.cossacklabs.com/acra/guides/integrating-acra-translator-into-new-infrastructure/) to learn more about AcraServer features and how to use them.


## Installation and launch

See [Getting started](https://docs.cossacklabs.com/acra/getting-started/) to learn how to [install Acra](https://docs.cossacklabs.com/acra/getting-started/installing/) or to [try Acra without coding](https://docs.cossacklabs.com/acra/getting-started/trying/).

Requirements: Linux or macOS with installed Docker.

| ⚙️ [Run Acra Example Projects](https://github.com/cossacklabs/acra-engineering-demo) ⚙️ |
|---|

## Documentation and tutorials

The most recent versions of the documentation, tutorials, and demos for Acra are available on the official [Cossack Labs Documentation Server](https://docs.cossacklabs.com/acra/).

To gain an initial understanding of Acra, you might want to:

- [What is Acra](https://docs.cossacklabs.com/acra/what-is-acra/) to get an overview of things.
- Acra's [security controls](https://docs.cossacklabs.com/acra/security-controls/) to learn more about encryption, masking, tokenisation, SQL firewall, intrusion detections, etc.
- Typical [dataflows](https://docs.cossacklabs.com/acra/acra-in-depth/data-flow/) that shows which Acra components you need and what are the Pros and Cons of each combination.
- Read the notes on Acra's [architecture](https://docs.cossacklabs.com/acra/acra-in-depth/architecture/) and [security design](https://docs.cossacklabs.com/acra/acra-in-depth/security-design/) to understand better what you get when you use Acra and what is the threat model that Acra operates in.

You can also check out the speaker slides for the following talks by Cossack Labs engineers:
- ["Encryption Without Magic, Risk Management Without Pain"](https://speakerdeck.com/vixentael/encryption-without-magic-risk-management-without-pain) by [Anastasiia Voitova](https://github.com/vixentael).
- ["Data encryption for Ruby web applications"](https://speakerdeck.com/shad/data-encryption-for-ruby-web-applications) by [Dmytro Shapovalov](https://github.com/shadinua).
- ["Building SQL firewall(AcraCensor): insights from developers"](https://speakerdeck.com/storojs72/building-sql-firewall-insights-from-developers) by [Artem Storozhuk](https://github.com/storojs72).

## Example projects

| ⚙️ [Run Acra Example Projects](https://github.com/cossacklabs/acra-engineering-demo) ⚙️ |
|---|


## Security consulting

It takes more than just getting cryptographic code to compile to secure the sensitive data. Acra won't make you “compliant out of the box” with all the modern security regulations, and no other tool will.

[We help companies](https://www.cossacklabs.com/solutions/security-strategy-advisory/) plan their data security strategy by auditing, assessing data flow, and classifying the data, enumerating the risks. We do the hardest, least-attended part of reaching the compliance – turning it from the “cost of doing business” into the “security framework that prevents risks”.


## Contributing to us

If you’d like to contribute your code or provide any other kind of input to Acra, you’re very welcome. Your starting point for contributing [is here](https://docs.cossacklabs.com/acra/contributing-and-community/).

## Acra feedback

If you are an Acra user, please leave a [short feedback](https://bit.ly/acra-feedback).

## License

Acra Community Edition is licensed as Apache 2 open-source software.


## Contacts

If you want to ask a technical question, feel free to raise an [Issue](https://github.com/cossacklabs/acra/issues) or write to [dev@cossacklabs.com](mailto:dev@cossacklabs.com).

To talk to the business wing of Cossack Labs Limited, drop us an email to [info@cossacklabs.com](mailto:info@cossacklabs.com).
   
[![Blog](https://img.shields.io/badge/blog-cossacklabs.com-7a7c98.svg)](https://cossacklabs.com/) [![Twitter CossackLabs](https://img.shields.io/badge/twitter-cossacklabs-fbb03b.svg)](https://twitter.com/cossacklabs) [![DEV CossackLabs](https://img.shields.io/badge/devto-%40cossacklabs-black.svg)](https://dev.to/cossacklabs/) [![Medium CossackLabs](https://img.shields.io/badge/medium-%40cossacklabs-orange.svg)](https://medium.com/@cossacklabs/)

Leave a star in GitHub, give a clap in Medium and share if you found this helpful.
