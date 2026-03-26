# Med-Pulse AI: High-Concurrency Clinical Knowledge Toolkit

**Developed by Fridah Mumbi Njung'e (ScJones724)**

[![Go Version](https://img.shields.io/badge/Go-1.26.1-00ADD8?style=flat-square&logo=go)](https://golang.org/)
[![HTMX](https://img.shields.io/badge/HTMX-2.0-3366CC?style=flat-square&logo=htmx)](https://htmx.org/)
[![Tailwind CSS](https://img.shields.io/badge/Tailwind_CSS-3.0-38B2AC?style=flat-square&logo=tailwind-css)](https://tailwindcss.com/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

---

## 1. Title & Objective

**Title:** *"Med-Pulse: A High-Concurrency Clinical AI Knowledge Toolkit"*

**Technology Chosen:** **Go (Golang)** , **htmx**, and **Tailwind CSS**.

**Why?**  
As a *Clinical Officer* and *Data Scientist*, I needed a stack that was **fast enough for real-time clinical use** but simple enough to deploy without heavy JavaScript frameworks. 
- **Go** handles thousands of concurrent users (high-concurrency).  
- **htmx** allows for a *Single-Page App* feel with zero page refreshes.

> **End Goal:** To provide a *"dashing,"* responsive web interface where clinicians can instantly summon clinical pearls and evidence-based tips.

---

## 2. Quick Summary of the Technology

**What is it?**  
- **Go** is a compiled, high-performance language developed by Google.  
- **htmx** is a library that allows access to modern browser features (AJAX) directly in HTML attributes.

**Where is it used?**  
- Used in high-traffic systems like **Cloudflare** and **Netflix** for backend speed.  
- Used by modern web developers who want to avoid the *"bloat"* of React or Angular.

**Real-world Example:**  
High-performance dashboards and real-time data monitoring tools.

---

## 3. System Requirements

- **OS:** Windows 10/11 (amd64).  
- **Tools:** Visual Studio Code (VS Code), Go Compiler v1.26.1.  
- **Dependencies:** None (Go Standard Library). Frontend uses CDN links for htmx and Tailwind CSS for a *"zero-install"* client experience.

---

## 4. Installation & Setup Instructions

### 1. Initialize Project

Open your terminal and run:

```
mkdir medical-toolkit
cd medical-toolkit
go mod init medical-toolkit
```

### 2. Verify Environment

Ensure the Go compiler is in your System PATH.  
*(If you run into issues, see the Common Issues section.)*

### 3. Run Server

```
go run main.go
```

### 4. Access App

Open your browser and navigate to: [http://localhost:8080](http://localhost:8080)

---

## 5. Minimal Working Example

The app uses a Go backend to serve a random clinical fact. When the user clicks the button, **htmx** sends a request to the server, which responds with a small HTML fragment that is instantly swapped into the UI.

### Go Backend Logic

```
http.HandleFunc("/get-tip", func(w http.ResponseWriter, r *http.Request) {
    tips := []string{
        "Hypertension: Confirm with serial readings.",
        "Allergies: Verify status before RX.",
    }
    time.Sleep(800 * time.Millisecond)
    fmt.Fprintf(w, "<div class='tip-card'>%s</div>", tips[rand.Intn(len(tips))])
})
```

### Frontend Interaction

```
<button 
    hx-post="/get-tip" 
    hx-target="#result" 
    hx-indicator="#loading"
    class="px-6 py-2 bg-blue-600 text-white rounded-lg">
    Get Clinical Tip
</button>
<div id="result"></div>
```

**Expected Output:**  
A *"Glassmorphism"* styled card that pulses while *"Paging the Senior AI Registrar"* and then displays a clinical pearl.

---

## 6. AI Prompt Journal (Assessment Section)

### Prompt Session 1: Architecture & Layman Explanations

> **Prompt:** *"Help me build a Gen AI Toolkit using Go and htmx. I'm a clinician with some coding experience but I don't really understand backend vs frontend. Explain the code in layman's terms that I can actually understand and present to other clinicians."*

**AI Response Summary:**  
The AI used the **Receptionist vs. Doctor** analogy:

- **The Receptionist (Frontend - HTML/htmx):** Handles the waiting room, keeps the UI calm and responsive.
- **The Doctor (Backend - Go):** Handles the clinical logic, evidence, and heavy processing.
- **The Communication System (htmx):** Acts as the clinical assistant who updates the patient without disrupting the whole clinic.

**Why This Matters for a Clinician:**  
> *"This analogy transformed abstract programming concepts into something I could grasp immediately."*

**Evaluation:** *Extremely Helpful* 

---

### Prompt Session 2: Troubleshooting Environment Variables

> **Prompt:** *"I installed Go on Windows 11 but when I type 'go version' in the terminal, I get 'go is not recognized as an internal or external command...'"*

**AI Response Summary:**  
The AI recognized this as a **PATH environment variable issue** and provided a walkthrough:

1. **Verify Go Installation:**
   ```
   dir "C:\Program Files\Go\bin"
   ```
2. **Add Go to System PATH:**  
   - Press `Windows + R`, type `sysdm.cpl`, go to *Advanced* → *Environment Variables*.  
   - Under *System Variables*, find `Path` → *Edit* → *New* → add `C:\Program Files\Go\bin`.  
   - Click *OK* on all dialogs.
3. **Verify:**
   ```
   go version
   ```
   Expected: `go version go1.21.5 windows/amd64`

**Evaluation:** *Critical* 🔧  
This fix prevented the project from stalling.

---

### Prompt Session 3: GitHub Identity & Branch Conflict

> **Prompt:** "I'm having a GitHub crisis at 9 PM with my deadline looming. Two problems:

 a). I'm stuck in an account loop. Every time I try to push, it wants to authenticate as 'Jones254-oss' which is an organization account I don't have access to anymore. I'm logged into my personal GitHub in the browser but Git keeps using the wrong account.

 b).When I try to push, I get 'src refspec main does not match any'. I think it's a branch name issue.

Please help me fix this before midnight!"

**AI Response Summary:**  
The AI treated this like a clinical emergency:

1. **Fix Credential Conflict:**
   ```
   git config --global --unset credential.helper
   git config --global user.name "ScJones724"
   git config --global user.email "your-email@example.com"
   ```
2. **Fix Branch Mismatch:**
   ```
   git branch -M main
   git push -u origin main
   ```

**Evaluation:** *Life-saving* 

---

## 7. Common Issues & Fixes

| Issue | Fix |
| :--- | :--- |
| **Terminal failed to recognize Go** | Manually updated the Windows *System PATH* and restarted terminal. |
| **error: src refspec main does not match any** | Used `git branch -M main` to rename local branch to match GitHub. |
| **GitHub "Owner" dropdown shows wrong account** | Used direct URL parameter (`?owner=ScJones724`) and Incognito window. |

---

## 8. References

- [Go Language Documentation](https://golang.org/doc/)
- [htmx Official Reference](https://htmx.org/docs/)
- [Tailwind CSS Playbook](https://tailwindcss.com/docs)

---

## 9. Code & Links for Quick Copy

### Clone the Repository
```
git clone https://github.com/ScJones724/medical-toolkit.git
cd medical-toolkit
```

### Run the Application
```
go run main.go
```

---

*Made with ❤️ by a Clinical Officer who codes.*
```
