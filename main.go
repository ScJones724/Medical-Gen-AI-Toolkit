package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// Our Knowledge Base (Clinical Tips)
var clinicalTips = []string{
    // Existing ones
    "Pediatrics: Always calculate drug dosages based on weight (mg/kg) to avoid toxicity.",
    "Hypertension: A single high reading isn't a diagnosis; confirm with serial measurements.",
    "Hydration: For severe dehydration, the skin pinch test (turgor) is a key clinical sign.",
    "Pharmacology: Ensure the 'Five Rights' before administering any medication.",
    
    // Add these:
    "Hand Hygiene: Wash hands for at least 20 seconds before and after patient contact.",
    "Informed Consent: Always explain risks, benefits, and alternatives before procedures.",
    "Documentation: If it wasn't documented, it wasn't done — always chart thoroughly.",
    "Pain Assessment: Pain is what the patient says it is, not what you assume.",
    "Allergies: Verify allergy status before prescribing any new medication.",
    "Patient Identification: Use two identifiers (name + DOB) before any intervention.",
    "Fever: In adults, fever is >38.0°C (100.4°F); treat symptoms, not just the number.",
    "Antibiotics: They work for bacterial infections only — not for viral colds or flu.",
}

func main() {
	// These are the "rules" of our clinic:
	http.HandleFunc("/", handleIndex)      // Home page
	http.HandleFunc("/get-tip", handleGetTip) // When user clicks 'Generate'

	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	// Send the user to the front desk (index.html)
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, nil)
}

func handleGetTip(w http.ResponseWriter, r *http.Request) {
	// Simulate "thinking" time (500ms)
	time.Sleep(500 * time.Millisecond)

	// Pick a random clinical fact
	rand.Seed(time.Now().UnixNano())
	tip := clinicalTips[rand.Intn(len(clinicalTips))]

	// Send back the tip formatted as HTML
	fmt.Fprintf(w, "<div class='p-4 bg-blue-100 border-l-4 border-blue-500 text-blue-700 italic'>%s</div>", tip)
}