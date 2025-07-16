# specruptiva

Disruption par la specification. Ou outils de gestion centré sur les données.

## Usage

### CLI 

````
```sh 
# rend disponible la commande specruptiva
export PATH=$PATH:$PWD/scripts

# affiche l'aide
specruptiva --help

# devrait échouer (status code 1)
specruptiva validate test/pets.cue test/charlie.yml

# devrait réussir (status code 0)
specruptiva validate test/pets.cue test/fido.yml





```

