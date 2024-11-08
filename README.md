# MeetingResumeGenerator (French)

## Description

MeetingResumeGenerator est une CLI (écrite en Go) qui permet de convertir un fichier `.mkv` en `.mp3`, de transcrire son contenu audio en texte à l'aide de l'API OpenAI Whisper ou Replicate (avec le modèle `whisper-diarization`), puis de générer un compte-rendu structuré à l'aide de l'API OpenAI GPT-4. Cela est particulièrement utile pour obtenir des transcriptions et des synthèses de réunions, conférences ou discussions.

## Fonctionnalités

- Conversion de fichiers `.mkv` en `.mp3` à l'aide de `ffmpeg`.
- Transcription audio en texte à l'aide de l'API OpenAI Whisper.
- Transcription audio avec diarisation (identification des locuteurs) à l'aide de l'API Replicate.
- Génération d'un compte-rendu au format Markdown à l'aide de l'API OpenAI GPT-4.
- Prise en charge de plusieurs langues (codes ISO 639-1).

## Prérequis

- **ffmpeg** : pour la conversion de fichiers `.mkv` en `.mp3`. Assurez-vous d'avoir `ffmpeg` installé et accessible via la ligne de commande.
- **Clé d'API OpenAI** : une clé API valide est nécessaire pour utiliser les services OpenAI. Vous devez définir la variable d'environnement `OPENAI_API_KEY`.
- **Clé d'API Replicate** : une clé API valide est nécessaire pour utiliser les services Replicate pour la diarisation. Vous devez définir la variable d'environnement `REPLICATE_API_KEY`.
- **Go** : pour compiler et exécuter le projet.

## Installation

1. Clonez le dépôt :
   ```sh
   git clone git@github.com:yoanbernabeu/MeetingResumeGenerator.git
   cd MeetingResumeGenerator
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

2. (Optionnel) Définissez votre clé d'API Replicate en tant que variable d'environnement pour la diarisation :
   ```sh
   export REPLICATE_API_KEY="votre_cle_api_replicate"
   ```

3. Exécutez la commande suivante pour convertir un fichier `.mkv`, le transcrire et générer un compte-rendu :
   ```sh
   meetingresumegenerator -input chemin/vers/votre_fichier.mkv -lang fr [-diarization nombre_de_locuteurs]
   ```

   - `-input` : Spécifiez le chemin du fichier `.mkv` que vous voulez transcrire.
   - `-lang` : Spécifiez le code ISO 639-1 de la langue de l'audio (ex: `fr` pour français, `en` pour anglais). Par défaut, la langue est `fr`.
   - `-diarization` : (Optionnel) Spécifiez le nombre de locuteurs pour la diarisation. Si cette option est utilisée, l'API Replicate sera appelée pour la transcription avec identification des locuteurs.

4. Les fichiers de transcription (`transcript.txt`) et de compte-rendu (`resume.md`) seront générés dans le répertoire courant.

## Exemple sans diarisation

Supposons que vous ayez un fichier nommé `meeting.mkv` que vous voulez transcrire en français.

```sh
meetingresumegenerator -input meeting.mkv -lang fr
```

Cela va :
1. Convertir `meeting.mkv` en `meeting.mp3`.
2. Transcrire l'audio avec l'API OpenAI Whisper et sauvegarder le résultat dans `transcript.txt`.
3. Générer un compte-rendu formaté en Markdown dans `resume.md`.

### Exemple avec diarisation

Supposons que vous ayez un fichier audio avec deux locuteurs et que vous souhaitez identifier les locuteurs dans la transcription.
Pour utiliser la diarisation avec le modèle Replicate [thomasmol/whisper-diarization](https://replicate.com/thomasmol/whisper-diarization), exécutez la commande suivante :

```sh
meetingresumegenerator -input chemin/vers/votre_fichier.mkv -lang fr -diarization 2
```

Dans cet exemple, `-diarization 2` indique que l'audio contient deux locuteurs. Le modèle Replicate sera utilisé pour transcrire l'audio avec identification des locuteurs.

Cela va :
1. Convertir le fichier `.mkv` en `.mp3`.
2. Transcrire l'audio avec l'API Replicate et sauvegarder le résultat dans `transcript.txt`.
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

C'est avant tout un outil personnel, codé en grande partie avec ChatGPT pour répondre à un besoin spécifique. Si vous souhaitez contribuer, n'hésitez pas à ouvrir une issue ou une pull request, mais aucune garantie n'est donnée quant à la fusion des modifications 😀.

## Licence

Ce projet est sous licence MIT. Voir le fichier [LICENSE](LICENSE.md) pour plus d'informations.
