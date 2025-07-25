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

## DevOps actions

### Release et version

**conventions:**
  - major: bris de l'interface. Sauf pour la version 0 dont l'interface est instable et peut bisée à tout moment
  - minor: ajout, retrait ou changement de fonctionnalité(s)
  - patch: bugfix, ou tout autre modification non fonctionelle

```sh
specruptiva release minor "Brève description du changement"
```

### API pour persister les schémas

#### Start API

```sh
cd /cmd/api
go run .
```

#### Documentation
https://documenter.getpostman.com/view/16523457/2sB34kDyLS#89d5d5b6-f19b-4430-ae18-58c232a06381