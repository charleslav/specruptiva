# specruptiva

Disruption par la specification. Ou outils de gestion centré sur les données.

## Usage

### CLI 

```sh 
# rend disponible la commande specruptiva
export PATH=$PATH:$PWD/scripts

# affiche l'aide
specruptiva --help

# devrait échouer (status code 1)
specruptiva validate test/pets.cue test/charlie.yml

# devrait réussir (status code 0)
specruptiva validate test/pets.cue test/fido.yml

# à partir de stdin
cat test/fido.yml | specruptiva validate  test/pets.cue
```

### API pour persister les schémas

```sh
# Start API
go run cmd/api/main.go
# ou bien
specruptiva start-api

# créer schema
specruptiva api-schema-create test/pets.cue

# lister tous les schemas
specruptiva api-schema-list

# afficher un schema
specruptiva api-schema-read 1

# remplace un schema
specruptiva api-schema-update 1 test/pets_v2.cue

# supprime un schema
specruptiva api-schema-delete 1
```

## DevOps actions

### Release et version

**conventions:**
  - major: bris de l'interface. Sauf pour la version 0 dont l'interface est instable et peut bisée à tout moment
  - minor: ajout, retrait ou changement de fonctionnalité(s)
  - patch: bugfix, ou tout autre modification non fonctionelle

```sh
specruptiva release minor "Brève description du changement"
```

### Test

```sh
. setup
# test l'api rest
specruptiva test-api
```
