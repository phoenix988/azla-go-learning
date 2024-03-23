function openNav() {
    var panel = document.getElementById("word-tree");
    panel.style.transform = "translateX(0)";

    if (window.innerWidth <= 600) {
        // Adjust this threshold as needed
        panel.style.width = "340px";
    } else {
        panel.style.width = "500px";
    }
}

function closeNav() {
    var panel = document.getElementById("word-tree");
    panel.style.transform = "translateX(-250px)";

    // Check if the screen width is less than or equal to a certain threshold (e.g., for phones)
    if (window.innerWidth <= 600) {
        // Adjust this threshold as needed
        panel.style.width = "0";
    } else {
        panel.style.width = "250px";
    }
}

document
    .getElementById("loginForm")
    .addEventListener("submit", function(event) {
        event.preventDefault(); // Prevent default form submission

        // Get form data
        var formData = new FormData(this);

        // Send AJAX request to server
        var xhr = new XMLHttpRequest();
        xhr.open("POST", "/login", true);
        xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");

        xhr.onreadystatechange = function() {
            if (xhr.readyState === XMLHttpRequest.DONE) {
                if (xhr.status === 200) {
                    // Request was successful, handle response
                    var response = JSON.parse(xhr.responseText);
                    if (response.success) {
                        // Redirect or perform other actions for successful login
                        window.location.href = "/dashboard";
                    } else {
                        // Show pop-up for incorrect password
                        alert("Incorrect password");
                    }
                } else {
                    // Handle other HTTP status codes
                    console.error("Error: " + xhr.status);
                }
            }
        };

        // Send form data
        xhr.send(new URLSearchParams(formData));
    });

// Function to hide the message after a certain amount of time
function hideMessageAfterTimeout(id) {
    setTimeout(function() {
        var message = document.getElementById(id);
        if (message) {
            message.style.display = "none"; // Hide the message
        }
    }, 5000); // 5000 milliseconds = 5 seconds
}

function hideWindowAfterTimeout(id) {
    setTimeout(function() {
        var message = document.getElementById(id);
        if (message) {
            message.style.display = "none"; // Hide the message
        }
    }, 5000); // 5000 milliseconds = 5 seconds
}
