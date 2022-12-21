# Incremental(ish) Ingress Controller

A ground-up rewrite of KIC. Removes the traditional KIC store and parser in
favor of a Koko integration.

Unlike KIC, controllers are responsible for translating Kubernetes resources
into Kong resources, rather than simply inserting copies of resources into the
store. These resources and their dependencies are then added to a Koko store
via direct Golang calls to the Koko admin API.

The integrated Koko instance is responsible for registering hybrid Kong data
plane instances and broadcasting configuration to them over wRPC.

## Kong repo template instructions

- ~Create template from repository~
- ~From the new repository settings page enable "Automatically delete head
  branches" as well as "Allow auto-merge"~
- ~From the new repository branches page create branch protection rule for
  `main` that requires "pre-commit" to pass as well as "Require a pull request
  before merging"~
- ~Following the [CODEOWNERS
  SYNTAX](https://docs.github.com/en/repositories/managing-your-repositorys-settings-and-features/customizing-your-repository/about-code-owners#codeowners-syntax)
  guidelines, update the new repository CODEOWNERS file~
- Following our [Github bot guidline
  documentation](https://konghq.atlassian.net/wiki/spaces/ENGEN/pages/2720268304/How+To+-+Github+Automation+Guidelines)
  add a github and dependabot secret for AUTO_MERGE_TOKEN
- Open a pull request on the new repository that seeds the secret baseline file
  `detect-secrets scan > .secrets.baseline` as well as a sensible README.md
- **Update** the .github/template-sync.yml file in
  [kong/template-generic](https://github.com/Kong/template-generic) repository
  with the **cloned repository name** to enable template sync changes

## Architecture overview

inc-ingress-controller is based off the standard Kubebuilder template.

Koko is wrapped in a shim that implements the controller-runtime Runnable
interface. This shim allows IIC to provision the Koko store separately, so that
it can then expose the store to controllers.

Controllers are responsible for adding resources and their dependencies to the
Koko store. The Ingress controller, for example, when receiving an Ingress
request, will create both the routes _and_ services, certificates, plugins,
etc. requested by that Ingress, and sets up watches (not yet implemented) for
the related resources. This should limit the number of ingested resources
considerably versus KIC, similar to how KIC currently limits ingested Secrets.

IIC uses an in-memory SQLite instance for Koko's store. Upon gaining
leadership, IIC starts the Koko runnable and populates an empty database from
scratch. This approach is similar to KIC's DB-less operation, but supports
multiple Kong instances with a single controller.

### Not yet implemented

The initial implementation of IIC is an extremely limited PoC. Although it
demonstrates a basic end-to-end implementation that exposes a Kong hybrid
control plane service, loads configuration from Kubernetes resources,
translates Kubernetes configuration to Kong configuration, and uploads it to
any Kong instance that registers, many critical features are not implemented:

- Ingresses and Services are the _only_ supported resources. IIC does not even
  populate Kong upstreams and targets from Service Endpoints.
- Resources are added and updated, but not deleted.
- There is no configuration. Authentication relies on a hard-coded shared mTLS
  certificate.
- Class filters are non-existent.

### Implementation hacks

Koko is not designed for library use, and exposes almost no public packages.
This limitation requires hacks for integration into IIC:

- This repository contains a snapshot of the entire Koko codebase (taken
  ~2022-12-12) under internal/koko. Changes within this are fairly minimal
  (though there are some leftover changes from earlier integration attempts),
  namely changing the package import paths to use the local copy. This
  circumvents the internal restriction, but is not remotely sustainable.
- Koko is largely run as a single unit under controller-runtime. Its individual
  components (the persister, store, relay server, event server, and admin
  server) are started by this single unit using the original gang thread
  management. It may make sense to run these separately via individual
  controller-runtime runnables, but this requires refactoring them and adding
  compatibility layers to preserve the existing Koko interfaces for upstream.
  The service setup is fairly tightly-coupled to the shared resources created
  in Koko's cmd/run, and ripping them out piecemeal to set up indidivual
  packages resulted in dependency hell when initially attempted.
- IIC implements its own in-memory API surface for the admin API client, as the
  existing one very much wants to live inside its typical HTTP server. This
  implementation should live in Koko, not its clients.
- To provide access to the Koko store to IIC's controllers, IIC instantiates
  and sets up the store (along with some dependencies) on its own, and then
  injects this pre-built store into Koko's setup.

Broadly, to integrate without the hack, Koko would need to create exported
packages that can set up its various services and a package for store CRUD
operations. These are effectively the requirements laid out in the Jira tickets
KOKO-669, KOKO-670, and KOKO-671.
