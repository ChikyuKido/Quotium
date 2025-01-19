document.addEventListener("DOMContentLoaded", function () {
    // Define the elements
    const loginButton = document.querySelector('.button.is-light');
    const registerButton = document.querySelector('.button.is-primary');
    const profileInfo = document.createElement('div');
    profileInfo.classList.add('profile-info');

    // Function to make the request to /api/v1/user/profile
    function fetchUserProfile() {
        fetch('/api/v1/user/profile', {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
                // You might need to add authentication tokens here depending on your setup
            }
        })
            .then(response => {
                if (!response.ok) {
                    // If the response is not 200 OK, show login and register buttons
                    loginButton.style.display = 'block';
                    registerButton.style.display = 'block';
                    return Promise.reject('User not authenticated');
                }
                return response.json();
            })
            .then(data => {
                // If request is successful, hide login/register buttons and show user data
                loginButton.style.display = 'none';
                registerButton.style.display = 'none';

                // Create and display the profile information
                const username = document.createElement('strong');
                username.textContent = data.username;
                const email = document.createElement('small');
                email.textContent = data.email;

                profileInfo.appendChild(username);
                profileInfo.appendChild(document.createElement('br'));
                profileInfo.appendChild(email);

                // Append the profile info to the navbar or wherever you'd like to display it
                const navbarEnd = document.querySelector('.navbar-end .navbar-item');
                navbarEnd.appendChild(profileInfo);
            })
            .catch(error => {
                console.error('Error fetching user profile:', error);
            });
    }

    // Call the function to fetch user profile on page load
    fetchUserProfile();
});
