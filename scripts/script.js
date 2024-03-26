function openNav() {
    var panel = document.getElementById("word-tree");
    var openbtn = document.getElementById("openBtn");
    panel.style.transform = "translateX(0)";
    openbtn.style.backgroundColor = "#2d3d50";

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

var openbtn = document.getElementById("openBtn");

openbtn.addEventListener("mouseover", function() {
    // Change background color when mouse hovers over the button
    openbtn.style.backgroundColor = "#333333";
});

// Restore original background color when mouse moves out
openbtn.addEventListener("mouseout", function() {
    // Restore the original background color
    openbtn.style.backgroundColor = "#2d3d50"; // or any other color code you want to set as default
});

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

function dropDownToggle(event) {
    // Prevent the default behavior of the button
    event.preventDefault();

    // Toggle the visibility of the options container
    document.getElementById("optionsContainer").classList.toggle("show");
}

// Close the dropdown menu if the user clicks outside of it
document.addEventListener("click", function(event) {
    var dropdownButton = document.getElementById("dropdown-button");
    var optionsContainer = document.getElementById("optionsContainer");

    // Check if the clicked element is not the dropdown button or its descendant
    if (
        event.target !== optionsContainer &&
        !optionsContainer.contains(event.target)
    ) {
        // Close the dropdown menu if it's currently open
        if (optionsContainer.classList.contains("show")) {
            optionsContainer.classList.remove("show");
        }
    }
});
