package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

var supportedLanguages = map[string]bool{
	// Langues supportées au format ISO 639-1
	"af": true, "ar": true, "az": true, "be": true, "bg": true, "bn": true,
	"bs": true, "ca": true, "cs": true, "cy": true, "da": true, "de": true,
	"el": true, "en": true, "es": true, "et": true, "fa": true, "fi": true,
	"fr": true, "gl": true, "gu": true, "he": true, "hi": true, "hr": true,
	"hu": true, "hy": true, "id": true, "is": true, "it": true, "ja": true,
	"jv": true, "ka": true, "kk": true, "km": true, "kn": true, "ko": true,
	"ku": true, "ky": true, "lt": true, "lv": true, "mk": true, "ml": true,
	"mn": true, "mr": true, "ms": true, "my": true, "ne": true, "nl": true,
	"no": true, "pa": true, "pl": true, "ps": true, "pt": true, "ro": true,
	"ru": true, "si": true, "sk": true, "sl": true, "sq": true, "sr": true,
	"sv": true, "sw": true, "ta": true, "te": true, "th": true, "tl": true,
	"tr": true, "uk": true, "ur": true, "uz": true, "vi": true, "zh": true,
}

func main() {
	// Paramètres en ligne de commande
	inputFile := flag.String("input", "", "Chemin du fichier .mkv à transcrire")
	language := flag.String("lang", "fr", "Code de langue ISO 639-1 (par ex: fr)")
	flag.Parse()

	// Récupération de la clé d'API depuis la variable d'environnement
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Println("Erreur : Veuillez définir la clé API dans la variable d'environnement OPENAI_API_KEY")
		return
	}

	// Validation du code de langue ISO 639-1
	if !supportedLanguages[*language] {
		fmt.Printf("Erreur : le code de langue '%s' n'est pas supporté ou invalide.\n", *language)
		fmt.Println("Utilisez un code de langue ISO 639-1 valide, par ex: 'fr' pour français, 'en' pour anglais.")
		return
	}

	fmt.Printf("Début de la conversion du fichier %s en .mp3...\n", *inputFile)

	// Conversion du fichier .mkv en .mp3
	outputFile := changeExtension(*inputFile, ".mp3")
	if err := convertToMP3(*inputFile, outputFile); err != nil {
		fmt.Printf("Erreur lors de la conversion en mp3 : %v\n", err)
		return
	}
	defer os.Remove(outputFile)

	fmt.Println("Conversion en .mp3 terminée avec succès.")

	// Appel de l'API OpenAI Whisper pour la transcription
	fmt.Println("Début de l'appel à l'API OpenAI Whisper pour la transcription...")
	transcription, err := callOpenAIWhisperAPI(apiKey, outputFile, *language)
	if err != nil {
		fmt.Printf("Erreur lors de l'appel à l'API OpenAI : %v\n", err)
		return
	}

	// Sauvegarde de la transcription
	err = os.WriteFile("transcript.txt", []byte(transcription), 0644)
	if err != nil {
		fmt.Printf("Erreur lors de l'enregistrement de la transcription : %v\n", err)
		return
	}
	fmt.Println("Transcription enregistrée dans transcript.txt")

	// Appel de l'API OpenAI GPT-4 pour le compte-rendu
	fmt.Println("Début de l'appel à l'API OpenAI GPT-4 pour le compte-rendu...")
	resume, err := callOpenAIGPT4(apiKey, transcription)
	if err != nil {
		fmt.Printf("Erreur lors de l'appel à GPT-4 : %v\n", err)
		return
	}

	// Sauvegarde du compte-rendu
	err = os.WriteFile("resume.md", []byte(resume), 0644)
	if err != nil {
		fmt.Printf("Erreur lors de l'enregistrement du compte-rendu : %v\n", err)
		return
	}
	fmt.Println("Compte-rendu enregistré dans resume.md")
}

// Convertit un fichier .mkv en .mp3 à l'aide de ffmpeg
func convertToMP3(inputFile, outputFile string) error {
	cmd := exec.Command("ffmpeg", "-i", inputFile, "-ar", "16000", "-ac", "1", "-acodec", "libmp3lame", "-y", outputFile)
	fmt.Printf("Commande ffmpeg : %s\n", cmd.String())
	return cmd.Run()
}

// Change l'extension du fichier
func changeExtension(filename, newExt string) string {
	return filename[:len(filename)-len(filepath.Ext(filename))] + newExt
}

// Appel à l'API Whisper d'OpenAI pour obtenir la transcription
func callOpenAIWhisperAPI(apiKey, audioFile, language string) (string, error) {
	file, err := os.Open(audioFile)
	if err != nil {
		return "", fmt.Errorf("erreur d'ouverture du fichier : %v", err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(audioFile))
	if err != nil {
		return "", fmt.Errorf("erreur de création de la requête multipart : %v", err)
	}
	if _, err = io.Copy(part, file); err != nil {
		return "", fmt.Errorf("erreur de copie du fichier dans la requête : %v", err)
	}

	writer.WriteField("model", "whisper-1")
	writer.WriteField("language", language)
	writer.Close()

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/audio/transcriptions", body)
	if err != nil {
		return "", fmt.Errorf("erreur lors de la création de la requête HTTP : %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("erreur d'envoi de la requête HTTP : %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("l'API a retourné une erreur : %s - %s", resp.Status, string(respBody))
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("erreur de lecture de la réponse : %v", err)
	}

	return string(respBody), nil
}

// Appel à l'API GPT-4 pour générer un compte-rendu
func callOpenAIGPT4(apiKey, transcription string) (string, error) {
	// Créer le corps de la requête JSON pour GPT-4
	requestBody, err := json.Marshal(map[string]interface{}{
		"model": "gpt-4o",
		"messages": []map[string]string{
			{"role": "system", "content": "Vous êtes un assistant expert en création de compte-rendu de réunions."},
			{"role": "user", "content": fmt.Sprintf(`Vous êtes un assistant expert en création de compte-rendu de réunions au format Markdown. Voici une transcription de réunion en texte brut :

%s

Veuillez générer un compte-rendu structuré et détaillé de cette réunion au format Markdown, avec les éléments suivants :

- Un **titre principal** intitulé "Compte-Rendu de la Réunion".
- Une **section "Date"** mentionnant la date de la réunion.
- Une **section "Participants"** listant les noms des participants à la réunion.
- Une **section "Objet"** décrivant brièvement l'objectif principal de la réunion.
- Une liste des **points discutés**, où chaque point inclut :
  - Un **titre** décrivant le sujet principal de la discussion.
  - Une **sous-section "Contexte"** expliquant le contexte ou le problème abordé.
  - Une **sous-section "Discussion"** résumant les échanges ou réflexions de l'équipe.
  - Une **sous-section "Décision" ou "Proposition"** mentionnant les décisions prises ou les propositions formulées pour résoudre le problème.
- Une **section "Actions et Suivi"** avec les actions assignées et les personnes responsables, organisées sous forme de liste.
- Une **section "Conclusion"** résumant les principaux points et actions à suivre.

Formattez chaque section avec des titres et sous-titres en Markdown et veillez à structurer les informations pour une lecture rapide et efficace.`, transcription)},
		},
	})
	if err != nil {
		return "", fmt.Errorf("erreur lors de la création du JSON : %v", err)
	}

	// Préparer et envoyer la requête HTTP à l'API GPT-4
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(requestBody))
	if err != nil {
		return "", fmt.Errorf("erreur lors de la création de la requête HTTP : %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("erreur d'envoi de la requête HTTP : %v", err)
	}
	defer resp.Body.Close()

	// Lire et renvoyer le corps de la réponse
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("erreur de lecture de la réponse : %v", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return "", fmt.Errorf("erreur lors de l'analyse du JSON : %v", err)
	}

	// Extraction du texte généré
	if choices, ok := result["choices"].([]interface{}); ok && len(choices) > 0 {
		if message, ok := choices[0].(map[string]interface{})["message"].(map[string]interface{}); ok {
			return message["content"].(string), nil
		}
	}

	return "", fmt.Errorf("réponse inattendue de l'API : %s", string(respBody))
}
