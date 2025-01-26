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

document.addEventListener('DOMContentLoaded', () => {
    const navbarBurgers = Array.prototype.slice.call(document.querySelectorAll('.navbar-burger'), 0);

    if (navbarBurgers.length > 0) {
        navbarBurgers.forEach(el => {
            el.addEventListener('click', () => {
                const target = el.dataset.target;
                const $target = document.getElementById(target);
                el.classList.toggle('is-active');
                $target.classList.toggle('is-active');
            });
        });
    }
});
