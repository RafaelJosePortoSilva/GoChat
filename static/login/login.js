document.getElementById("loginForm").addEventListener("submit", async function(event) {
    event.preventDefault();

    const username = document.getElementById("username").value;
    const password = document.getElementById("password").value;

    try {
        const response = await fetch("/login/", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({ username, password })
        });

        if (!response.ok) {
            let errorMessage = "An unexpected error occurred.";
            try {
                const errorResponse = await response.json();
                errorMessage = errorResponse.message || errorMessage;
            } catch (parseError) {
                console.error("Error parsing JSON:", parseError);
            }
            document.getElementById("error-message").innerText = errorMessage;
        }
    } catch (error) {
        console.error("Error logging in:", error);
        document.getElementById("error-message").innerText = "An error occurred during login.";
    }
});
