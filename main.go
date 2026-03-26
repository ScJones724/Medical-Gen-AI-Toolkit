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

// The Intercom (Loading Messages)
var loadingMessages = []string{
	"☕ Taking a quick coffee break before rounds...",
	"✍️ Double-checking the intern's notes...",
	"🏥 Navigating the hospital corridors...",
	"🔦 Performing a quick systemic review...",
	"🔍 Squinting at the consultant's handwriting...",
	"🥪 Swallowing a sandwich between patients...",
	"📜 Checking if the lab results finally dropped...",
	"🩺 Just one more patient... I promise.",
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
    // 1. Pick a random loading message
    rand.Seed(time.Now().UnixNano())
    loadMsg := loadingMessages[rand.Intn(len(loadingMessages))]
    
    // 2. Simulate "thinking" time (800ms for extra drama)
    time.Sleep(800 * time.Millisecond)

    // 3. Pick the clinical fact
    tip := clinicalTips[rand.Intn(len(clinicalTips))]

    // 4. Send back the Tip + the loading message persona
    // We are adding the 'loadMsg' at the bottom in a small italic font
    fmt.Fprintf(w, `
    <div class="p-6 bg-white/20 border-l-8 border-white/40 text-white italic rounded-xl shadow-inner transition-all duration-500">
        %s
        <div class="text-[10px] mt-4 text-blue-200 font-black uppercase tracking-[0.3em] opacity-70 animate-pulse">
            Status: <span class="border-r-2 border-blue-200 pr-1">%s</span>
        </div>
    </div>
`, tip, loadMsg)
}