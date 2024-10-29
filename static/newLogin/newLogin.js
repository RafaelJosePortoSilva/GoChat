document.getElementById("loginForm").addEventListener("submit", async function(event) {
    event.preventDefault();

    const username = document.getElementById("username").value;
    const password = document.getElementById("password").value;

    try {
        const response = await fetch("/login/new", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({ username, password })
        });

        if (!response.ok) {
            const errorMessage = await response.json();
            document.getElementById("error-message").innerText = errorMessage.message;
        }
    } catch (error) {
        console.error("Error logging in:", error);
        document.getElementById("error-message").innerText = "An error occurred during register.";
    }
});
