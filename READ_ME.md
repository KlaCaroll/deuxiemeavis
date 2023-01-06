## Pré-requis

- Go
- Docker
- Docker-compose

# Rendu du test technique

Merci pour le temps que vous y accorderez

# Lancement de l'API

Lancez la base de données à l'aide de la commande suivante :

```sh
docker-compose up -d
```

depuis le dossier api  lancer la commande :

```sh
go run .
```

# Lancement des endpoints


Pour le 1er endpoint lancer la commande :
```sh
http :8080/v2/doctors 
```


Pour le 2nd endpoint lancer la commande :
```sh
http :8080/v2/requestslist status='new'
```
ou
```sh
http :8080/v2/requestslist status='archived'
```

Pour le 3eme endpoint lancer la commande :
```sh
http :8080/v2/newrequest patient_id:=4 doctor_id:=3 disease_id:=3 diagnosis='Lorem ipsum'
```

Pour le 4eme et dernier endpoint lancer la commande :
```sh
http :8080/v2/secondopinion id=1 second_opinion='Lorem ipsum' 
```
