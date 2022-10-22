<h3 align="center">
<a href="#"><img src="https://i.ibb.co/19t6QjB/logoza-ru-1.png" alt="Memory Castle Project" width="400"></a>
<br>
A simple application for those who never FORGET
<br>
</h3>
[![Coverage Status](https://coveralls.io/repos/github/radixiura/Meca-Go/badge.svg?branch=main)](https://coveralls.io/github/radixiura/Meca-Go?branch=main)
-----
  
<p align="center">
  <a href="#"><img src="https://img.shields.io/github/release/cossacklabs/acra.svg" alt="GitHub release"></a>
  <a href="#"><img src="https://coveralls.io/repos/github/radixiura/Meca-Go/badge.svg?branch=main" alt='Coverage Status' /></a>
  <a href="#"><img class="badge" tag="github.com/cossacklabs/acra" src="https://goreportcard.com/badge/github.com/cossacklabs/acra"></a>
  <a href="#"><img src='https://godoc.org/github.com/cossacklabs/acra?status.svg'  alt='godoc'/></a>
</p>

<br>

| [Meca Examples](#) | [Documentation](#) | [Installation](#) | [Feedback](#) |
| ---- | ---- | ---- | ---- |

## What is MECA Project?
Meca - lallalala.

Meca includes [](https://www.infoq.com/articles/ale-software-architects/) for data fields, multi-layered access control, database leakage prevention, and intrusion detection capabilities in one suite. Acra was specifically designed for distributed apps (web, server-side and mobile) that store data in one or many databases / datastores.

<table><thead><tr><th>Perfect cases for using Meca</th>
<th>For this, who</th></tr></thead>
<tbody><tr><td>You want to read memories of random people</td>
<td rowspan=3><ul>
<li>Cares about own memory</li>
<li>Likes to read</li>
<li>SaaS</li>
</tr><tr><td>You want to leave your step in a history!</td>
</tr><tr><td>You like to be more intelligent</td>
</tr></tbody></table>

## How does MECA work?
Acra consists of several services and utilities. Acra services allow you to construct infinitely sophisticated data flows that are perfectly suited to your exact infrastructure. Depending on your architecture and use case, you might need to deploy only basic services or all of them.

* **Security enforcement components**: services where "encryption happens". One of them is required: AcraServer, AcraTranslator, AnyProxy, or client-side SDKs.
* **Key storage:** datastores where Acra keeps encrypted keys: Redis, table in your database, any KV store. One of them is required.
* **Master key storage:** KMS, Vault. One of them is strongly recommended.
* **Additional services and utils:** key management utils, data migration scripts, transport security service, policy management tools. Any of them are optional.

Refer to [Acra-in-depth / Architecture](https://docs.cossacklabs.com/acra/acra-in-depth/architecture/) to learn more about Acra components. Refer to [Acra-in-depth / Data flow](https://docs.cossacklabs.com/acra/acra-in-depth/data-flow/) to see more typical Acra-based dataflows and deployments.

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

## Contributing to us
If you’d like to contribute your code or provide any other kind of input to Acra, you’re very welcome. Your starting point for contributing [is here](#).

## License
MECA is totally free!

## Contacts
If you want to ask a technical question, feel free to raise an [Issue](https://github.com/radixiura/Meca-Go/issues) or write to [radix_iura@protonmail.com](mailto:radix_iura@protonmail.com).
