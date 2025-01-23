document.addEventListener("DOMContentLoaded", function () {
    const loginButton = document.querySelector('.button.is-light');
    const registerButton = document.querySelector('.button.is-primary');
    const profileInfo = document.createElement('div');
    profileInfo.classList.add('profile-info');

    function fetchUserProfile() {
        fetch('/api/v1/user/profile', {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
            }
        })
            .then(response => {
                if (!response.ok) {
                    loginButton.style.display = 'block';
                    registerButton.style.display = 'block';
                    return Promise.reject('User not authenticated');
                }
                return response.json();
            })
            .then(data => {
                loginButton.style.display = 'none';
                registerButton.style.display = 'none';

                const username = document.createElement('strong');
                username.textContent = data.username;
                const email = document.createElement('small');
                email.textContent = data.email;

                profileInfo.appendChild(username);
                profileInfo.appendChild(document.createElement('br'));
                profileInfo.appendChild(email);

                const navbarEnd = document.querySelector('.navbar-end .navbar-item');
                navbarEnd.appendChild(profileInfo);
            })
            .catch(error => {
                console.error('Error fetching user profile:', error);
            });
    }

    fetchUserProfile();
});
