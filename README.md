Med-Pulse AI: High-Concurrency Clinical Knowledge Toolkit
Developed by Fridah Mumbi Njung'e (ScJones724)
1. Title & Objective
Title: "Med-Pulse: A High-Concurrency Clinical AI Knowledge Toolkit"

Technology Chosen: Go (Golang), htmx, and Tailwind CSS.

Why? As a Clinical Officer and Data Scientist, I needed a stack that was fast enough for real-time clinical use but simple enough to deploy without heavy JavaScript frameworks. Go handles thousands of concurrent users (high-concurrency), while htmx allows for a "Single-Page App" feel with zero page refreshes.

End Goal: To provide a "dashing," responsive web interface where clinicians can instantly summon clinical pearls and evidence-based tips.

2. Quick Summary of the Technology
What is it? Go is a compiled, high-performance language by Google. htmx is a library that allows access to modern browser features (AJAX) directly in HTML attributes.

Where is it used? Used in high-traffic systems like Cloudflare and Netflix for backend speed, and by modern web developers who want to avoid the "bloat" of React or Angular.

Real-world Example: High-performance dashboards and real-time data monitoring tools.

3. System Requirements
OS: Windows 10/11 (amd64).

Tools: Visual Studio Code (VS Code), Go Compiler v1.26.1.

Dependencies: None (Go Standard Library). Frontend uses CDN links for htmx and Tailwind CSS for a "zero-install" client experience.

4. Installation & Setup Instructions
Initialize Project:

Bash
mkdir medical-toolkit
cd medical-toolkit
go mod init medical-toolkit
Verify Environment: Ensure the Go compiler is in your System PATH.

Run Server: ```bash
go run main.go

Access App: Open browser to http://localhost:8080.

5. Minimal Working Example
The app uses a Go backend to serve a random clinical fact. When the user clicks the button, htmx sends a request to the server, which responds with a small HTML fragment that is instantly swapped into the UI.

Go
// Backend logic for the "Consultation"
http.HandleFunc("/get-tip", func(w http.ResponseWriter, r *http.Request) {
    tips := []string{"Hypertension: Confirm with serial readings.", "Allergies: Verify status before RX."}
    // Simulate 'Clinical Reasoning' time
    time.Sleep(800 * time.Millisecond)
    fmt.Fprintf(w, "<div class='tip-box'>%s</div>", tips[rand.Intn(len(tips))])
})
Expected Output: A "Glassmorphism" styled card that pulses while "Paging the Senior AI Registrar" and then displays a clinical pearl.

6. AI Prompt Journal (Assessment Section)
Prompt Session 1: Architecture & Layman Explanations
Prompt 1.1:

"Help me build a Gen AI Toolkit using Go and htmx. I'm a clinician with some coding experience but I don't really understand backend vs frontend. Explain the code in layman's terms that I can actually understand and present to other clinicians."

AI Response Summary:

The AI used the "Receptionist vs. Doctor" analogy to explain the architecture:

The Receptionist (Frontend - HTML/htmx):

Greets patients (users) when they arrive at the clinic (website)

Takes their request (clicking the button)

Says "Please wait in the waiting room" while the doctor prepares

Calls the patient back when the doctor is ready (displays the tip)

The Doctor (Backend - Go):

Stays in the consultation room (the server)

Never sees patients directly

Reviews the file (processes the request)

Consults their medical knowledge (clinical tips array)

Writes a prescription (formats the tip as HTML)

Sends it to the receptionist to give to the patient

The Communication System (htmx):

The receptionist doesn't need to rebuild the whole clinic every time

They just hand the prescription to the patient without making everyone else leave the waiting room

That's why the page doesn't refresh - only the tip container changes

The Code Explained Line by Line:

go
// This is like the clinic's operating hours
func main() {
    // "When someone walks in the front door, send them to reception"
    http.HandleFunc("/", handleIndex)
    
    // "When someone presses the 'Get Advice' button, call the doctor"
    http.HandleFunc("/get-tip", handleGetTip)
    
    // "Open the clinic doors on port 8080"
    http.ListenAndServe(":8080", nil)
}
html
<!-- This is the reception desk -->
<button hx-post="/get-tip"    <!-- "Call the doctor when clicked" -->
        hx-target="#result">  <!-- "Put the response here" -->
    Get Clinical Tip
</button>

<div id="result"></div>       <!-- The waiting room where tips appear -->
Why This Matters for a Clinician:
The AI explained that this separation of concerns mirrors clinical practice:

Frontend = Patient-facing (simple, clear, intuitive)

Backend = Clinical decision-making (complex, evidence-based, secure)

htmx = The clinical assistant who updates the patient without disrupting the whole clinic

Evaluation:
Extremely Helpful 

This analogy transformed abstract programming concepts into something I could grasp immediately. More importantly, it gave me the language to explain the project during my presentation to other clinicians who don't code. I was able to say, "The receptionist handles the waiting room, the doctor handles the medicine, and they don't need to rebuild the hospital every time someone asks a question."

The AI also anticipated questions I'd get from clinicians:

"Is it secure?" Yes, because the doctor (backend) never exposes clinical logic to the patient

"Is it fast?" Yes, because only the tip changes, not the whole page

"Can I trust it?" The tips are hardcoded initially, but can be connected to verified medical databases

This foundational understanding made every subsequent coding session make sense.

Prompt Session 2: Troubleshooting Environment Variables
Prompt 2.1:

"I installed Go on Windows 11 but when I type 'go version' in the terminal, I get 'go is not recognized as an internal or external command, operable program or batch file.' I've restarted my computer twice and I'm stuck. Help me fix this."

AI Response Summary:

The AI immediately recognized this as a PATH environment variable issue and provided a detailed walkthrough:

Step 1: Verify Go Actually Installed

bash
# Check if the files exist
dir "C:\Program Files\Go\bin"
# Expected output: go.exe, gofmt.exe, etc.
Step 2: Locate Where Windows Looks for Programs
The AI explained that Windows has a "phonebook" called PATH - a list of folders it checks when you type a command. If Go's folder isn't in that phonebook, Windows has no idea where to find go.exe.

Step 3: Add Go to System PATH (Not User PATH)
The AI emphasized using System PATH instead of User PATH because:

System PATH affects all users and all terminals

User PATH only affects your profile and doesn't always apply to admin terminals

Step-by-Step Instructions:

Press Windows Key + R, type sysdm.cpl, press Enter

Click Advanced tab → Environment Variables

Under System Variables (bottom box), find Path → Click Edit

Click New → Add C:\Program Files\Go\bin

Click OK on all dialogs (this is critical - closing just one dialog doesn't save)

Step 4: Verify Without Restarting

bash
# Open a NEW PowerShell window (old ones still have old PATH)
# Test with:
go version
# Expected: go version go1.21.5 windows/amd64
Why Restarting Didn't Work:
The AI explained that Windows loads PATH at login. If you only restarted the terminal (not the whole computer), the changes wouldn't apply. But if you restarted the computer and it still didn't work, the PATH was likely added to User PATH instead of System PATH.

Alternative Solution if Still Broken:

bash
# Check current PATH
$env:Path -split ';' | Select-String "Go"

# If nothing shows, manually set PATH for current session
$env:Path += ";C:\Program Files\Go\bin"
Prevention for Future:
The AI suggested using Windows Package Manager for future installations:

bash
# This automatically handles PATH
winget install GoLang.Go
Evaluation:
Critical 

This was the first major blocker. Without this fix, the project would have ended before it started. The AI didn't just say "fix your PATH" - it explained why PATH matters, which PATH to use, and how to verify it worked.

The distinction between System and User PATH was particularly valuable. I had previously added Go to User PATH and wondered why it didn't work in admin terminals. The AI's explanation that System PATH affects all contexts solved this mystery.

Reflection:
This experience taught me that environment variables are like clinical pathways - if the system doesn't know where to find the resources (medication, specialists, tools), nothing works. The AI helped me "map the pathway" correctly.

Prompt Session 3: GitHub Identity & Branch Conflict
Prompt 3.1:

"I'm having a GitHub crisis at 9 PM with my deadline looming. Two problems:

I'm stuck in an account loop. Every time I try to push, it wants to authenticate as 'Jones254-oss' which is an organization account I don't have access to anymore. I'm logged into my personal GitHub in the browser but Git keeps using the wrong account.

When I try to push, I get 'src refspec main does not match any'. I think it's a branch name issue.

Please help me fix this before midnight!"

AI Response Summary:

The AI recognized this as a credential conflict combined with a branch naming mismatch, and provided an emergency triage plan:

Problem 1: Account Loop / Wrong Git User

Root Cause:
Git caches credentials in the Windows Credential Manager. Even though you logged out of the organization account in the browser, Git still had stored credentials for that account.

Emergency Fix:

bash
# Option A: Use Incognito for temporary fix
# Open VS Code in incognito mode or clear Git credentials:
git config --global --unset credential.helper
# Now Git will ask for credentials fresh on next push
Option B: Clear Windows Credential Manager

Press Windows Key → Type "Credential Manager"

Click "Windows Credentials"

Scroll to "Generic Credentials"

Find and delete any entries containing git:https://github.com

Close VS Code, reopen, and push again

Option C: Use Personal Access Token (Recommended)

bash
# Generate token at: GitHub Settings → Developer settings → Personal access tokens
# Use token as password when prompted
Why This Happened:
The AI explained that Git doesn't use browser sessions - it uses stored credentials. If you ever authenticated with the organization account, Git "remembers" it forever until manually cleared.

Problem 2: Branch Name Mismatch (src refspec main does not match any)

Root Cause:
Your local branch is named master (Git's old default) but GitHub expects main (new default).

Fix:

bash
# Check current branch
git branch
# Output: * master (your local is master)

# Rename local branch to main
git branch -M main

# Now push
git push -u origin main
If You Already Have Commits on master:

bash
# Push master to main on remote (alternative approach)
git push -u origin master:main
Prevention for Future Projects:

bash
# Set Git to use 'main' as default for all new repos
git config --global init.defaultBranch main
Complete Emergency Script:

bash
# Step 1: Clear credentials
git config --global --unset credential.helper

# Step 2: Check and rename branch
git branch -M main

# Step 3: Push with fresh credentials
git push -u origin main
# When prompted, use your personal GitHub credentials or token
What If Nothing Works?
The AI provided a nuclear option:

bash
# Remove remote and re-add
git remote remove origin
git remote add origin https://github.com/yourusername/yourrepo.git
git push -u origin main
Evaluation:
Very Helpful 

This saved me from a 12 AM deadline panic. The AI correctly diagnosed two separate issues:

Credential conflict (organization account stuck in cache)

Branch naming mismatch (master vs main)

The "Incognito mode" suggestion was the quickest fix - it allowed me to push without clearing all credentials globally. I used the branch rename command, pushed successfully, and met my deadline.

What I Learned:

Git credentials are cached independently of browser sessions

master and main are just names, but they must match between local and remote

The -M flag forces branch rename (vs -m which requires the branch to exist)

GitHub's UI confusion (showing one account while Git uses another) is common when credentials get crossed

Reflection:
This was the most stressful moment of the project. The AI's calm, step-by-step approach walked me through the fix without making me feel like I was in over my head. It treated this like a clinical emergency: assess the patient (Git), identify the two problems (credential conflict + branch mismatch), treat the most critical first (credentials), then address the secondary issue (branch name).

7. Common Issues & Fixes
Issue: Terminal failed to recognize Go after installation.

Fix: Manually updated the Windows System PATH and restarted the terminal session.

Issue: error: src refspec main does not match any.

Fix: Local Git initialized with master while GitHub expected main. Used git branch -M main to rename and sync.

Issue: GitHub "Owner" dropdown wouldn't show the personal account.

Fix: Used a direct URL parameter (?owner=ScJones724) and an Incognito window to force-select the correct profile.

8. References
Go Language Documentation

htmx Official Reference

Tailwind CSS Playbook