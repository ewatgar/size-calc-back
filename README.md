# size-calc-back

## Setup database

### Run containers

```bash
cd config/db && podman-compose up -d
```

### UI pgAdmin

Wait for it to load first after running it

```
localhost:5050
```

### Delete containers

```bash
cd config/db && podman-compose down
```

If you want to delete the persistent volumes too:

```bash
cd config/db && podman-compose down -v
```
