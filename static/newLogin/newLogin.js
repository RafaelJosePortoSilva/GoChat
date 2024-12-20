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
            let errorMessage = "An unexpected error occurred.";
            try {
                const errorResponse = await response.json();
                errorMessage = errorResponse.message || errorMessage;
            } catch (parseError) {
                console.error("Error parsing JSON:", parseError);
            }
            document.getElementById("error-message").innerText = errorMessage;
        } else {
            try {
                const successResponse = await response.json();
                const message = successResponse.message || "Registration successful!";
                document.getElementById("error-message").innerText = message;
        
                window.location.href = "/login";
            } catch (parseError) {
                console.error("Error parsing JSON:", parseError);
                document.getElementById("error-message").innerText = "An error occurred, but the registration might be successful.";
            }
        }
        
    } catch (error) {
        console.error("Error logging in:", error);
        document.getElementById("error-message").innerText = "An error occurred during register.";
    }
});
