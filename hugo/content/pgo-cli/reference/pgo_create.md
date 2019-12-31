---
title: "pgo create"
---
## pgo create

Create a Postgres Operator resource

### Synopsis

CREATE allows you to create a new Operator resource. For example:
    pgo create cluster
    pgo create pgbouncer
    pgo create pgouser
    pgo create pgorole
    pgo create policy
    pgo create namespace
    pgo create user

```
pgo create [flags]
```

### Options

```
  -h, --help   help for create
```

### Options inherited from parent commands

```
      --apiserver-url string     The URL for the PostgreSQL Operator apiserver.
      --debug                    Enable debugging when true.
      --disable-tls              Disable TLS authentication to the Postgres Operator.
      --exclude-os-trust         Exclude CA certs from OS default trust store
  -n, --namespace string         The namespace to use for pgo requests.
      --pgo-ca-cert string       The CA Certificate file path for authenticating to the PostgreSQL Operator apiserver.
      --pgo-client-cert string   The Client Certificate file path for authenticating to the PostgreSQL Operator apiserver.
      --pgo-client-key string    The Client Key file path for authenticating to the PostgreSQL Operator apiserver.
```

### SEE ALSO

* [pgo](/pgo-cli/reference/pgo/)	 - The pgo command line interface.
* [pgo create cluster](/pgo-cli/reference/pgo_create_cluster/)	 - Create a PostgreSQL cluster
* [pgo create namespace](/pgo-cli/reference/pgo_create_namespace/)	 - Create a namespace
* [pgo create pgbouncer](/pgo-cli/reference/pgo_create_pgbouncer/)	 - Create a pgbouncer 
* [pgo create pgorole](/pgo-cli/reference/pgo_create_pgorole/)	 - Create a pgorole
* [pgo create pgouser](/pgo-cli/reference/pgo_create_pgouser/)	 - Create a pgouser
* [pgo create policy](/pgo-cli/reference/pgo_create_policy/)	 - Create a SQL policy
* [pgo create schedule](/pgo-cli/reference/pgo_create_schedule/)	 - Create a cron-like scheduled task
* [pgo create user](/pgo-cli/reference/pgo_create_user/)	 - Create a PostgreSQL user

###### Auto generated by spf13/cobra on 27-Dec-2019