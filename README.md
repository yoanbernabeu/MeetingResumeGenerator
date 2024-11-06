# MeetingResumeGenerator (French)

## Description

MeetingResumeGenerator est une CLI (écrite en Go) qui permet de convertir un fichier `.mkv` en `.mp3`, de transcrire son contenu audio en texte à l'aide de l'API OpenAI Whisper, puis de générer un compte-rendu structuré à l'aide de l'API OpenAI GPT-4. Cela est particulièrement utile pour obtenir des transcriptions et des synthèses de réunions, conférences ou discussions.

## Fonctionnalités
- Conversion de fichiers `.mkv` en `.mp3` à l'aide de `ffmpeg`.
- Transcription audio en texte à l'aide de l'API OpenAI Whisper.
- Génération d'un compte-rendu au format Markdown à l'aide de l'API OpenAI GPT-4.
- Prise en charge de plusieurs langues (codes ISO 639-1).

## Prérequis

- **ffmpeg** : pour la conversion de fichiers `.mkv` en `.mp3`. Assurez-vous d'avoir `ffmpeg` installé et accessible via la ligne de commande.
- **Clé d'API OpenAI** : une clé API valide est nécessaire pour utiliser les services OpenAI. Vous devez définir la variable d'environnement `OPENAI_API_KEY`.
- **Go** : pour compiler et exécuter le projet.

## Installation

1. Clonez le dépôt :
   ```sh
   git clone <URL-du-depot>
   cd <nom-du-repertoire>
   ```
2. Compilez le programme :
   ```sh
   go build -o meetingresumegenerator
   ```
3. Installer le programme de manière globale :
   ```sh
   # Pour Linux et macOS
   sudo mv meetingresumegenerator /usr/local/bin/
   ```
4. Assurez-vous d'avoir `ffmpeg` installé sur votre système :
   ```sh
   # Pour Ubuntu/Debian
   sudo apt-get install ffmpeg

   # Pour macOS (via Homebrew)
   brew install ffmpeg
   ```

## Utilisation

1. Définissez votre clé d'API OpenAI en tant que variable d'environnement :
   ```sh
   export OPENAI_API_KEY="votre_cle_api_openai"
   ```

2. Exécutez la commande suivante pour convertir un fichier `.mkv`, le transcrire et générer un compte-rendu :
   ```sh
   meetingresumegenerator -input chemin/vers/votre_fichier.mkv -lang fr
   ```

   - `-input` : Spécifiez le chemin du fichier `.mkv` que vous voulez transcrire.
   - `-lang` : Spécifiez le code ISO 639-1 de la langue de l'audio (ex: `fr` pour français, `en` pour anglais). Par défaut, la langue est `fr`.

3. Les fichiers de transcription (`transcript.txt`) et de compte-rendu (`resume.md`) seront générés dans le répertoire courant.

## Exemple

Supposons que vous ayez un fichier nommé `meeting.mkv` que vous voulez transcrire en français.

```sh
meetingresumegenerator -input meeting.mkv -lang fr
```

Cela va :
1. Convertir `meeting.mkv` en `meeting.mp3`.
2. Transcrire l'audio avec l'API OpenAI Whisper et sauvegarder le résultat dans `transcript.txt`.
3. Générer un compte-rendu formaté en Markdown dans `resume.md`.

## Langues supportées

MeetingResumeGenerator prend en charge de nombreuses langues. Voici une liste non exhaustive des codes de langues ISO 639-1 supportés :

- `fr` : Français
- `en` : Anglais
- `es` : Espagnol
- `de` : Allemand
- `it` : Italien
- `ja` : Japonais
- `zh` : Chinois

Pour la liste complète des langues supportées, référez-vous au code source.

## Remarques

- Assurez-vous que votre fichier `.mkv` est de bonne qualité pour garantir une transcription précise.
- Le fichier audio est converti à une fréquence d'échantillonnage de 16000 Hz et un canal audio unique pour être compatible avec l'API OpenAI Whisper.

## Avertissements

- Les appels aux API OpenAI peuvent entraîner des coûts, assurez-vous de comprendre la tarification avant d'exécuter la CLI.
- Veuillez éviter de traiter des informations confidentielles ou sensibles à travers ces API, car les données sont envoyées à un service tiers.

## Contribution

Les contributions sont les bienvenues ! N'hésitez pas à ouvrir des issues pour signaler des bugs ou proposer de nouvelles fonctionnalités.

## Licence

Ce projet est sous licence MIT. Voir le fichier `LICENSE` pour plus d'informations.

## Contact

Pour toute question ou support, n'hésitez pas à contacter [votre.nom@example.com](mailto:votre.nom@example.com).
