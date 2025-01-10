export function addLikeDislikeListeners(post, postId) {

    const likeButton = post.querySelector('.like');
    const dislikeButton = post.querySelector('.dislike');

    likeButton.addEventListener('click', () => handleReact(likeButton,dislikeButton, postId, "like", "post"));
    dislikeButton.addEventListener('click', () => handleReact(dislikeButton,likeButton, postId, "dislike", "post"));

}

export async function handleReact(button, follow , postId, type , target_Type) {
    // Send API request
    const response = await fetch(`/api/react/${postId}/${type}/${target_Type}`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
    });

    if (!response.ok) {// user is not logged in
        showRegistrationModal(); 
    }else{ // only update the like if no error
        interactiveLike(button, follow)
    }
}

function interactiveLike(button , follow ){
    const add = button.querySelector(".count");
    const subtract = follow.querySelector(".count");

    // Parse the current count from the button's span text
    let count = parseInt(add.textContent, 10);

    if (button.getAttribute("data-clicked") === "false") {

        count += 1; add.textContent = count; // Update the displayed count
        button.setAttribute("data-clicked", "true");
        button.style.backgroundColor = '#15F5BA'
        follow.style.backgroundColor = 'white'

        if (follow.getAttribute("data-clicked") === "true") {
            count -= 1; subtract.textContent = count; // Update the displayed count
            follow.setAttribute("data-clicked", "false");
            follow.style.backgroundColor = 'white'
        }
    }else if (button.getAttribute("data-clicked") === "true") {
        count -= 1; add.textContent = count; // Update the displayed count
        button.setAttribute("data-clicked", "false");
        button.style.backgroundColor = 'white'
    }
}

function showRegistrationModal() {
    const dialog = document.createElement('dialog');
    // Create message
    const message = document.createElement('p');
    message.textContent = 'You need to be logged in to react. Please register or log in to continue.';

    // Create register button
    const registerButton = document.createElement('button');
    registerButton.textContent = 'Register Now';

    // Create login button
    const loginButton = document.createElement('button');
    loginButton.textContent = 'Login';


    // Add event listeners
    registerButton.addEventListener('click', () => {
        window.location.href = '/register'; // Replace with your registration URL
    });
    loginButton.addEventListener('click', () => {
        window.location.href = '/login'; // Replace with your login URL
    });

    // Close dialog when clicking outside
    dialog.addEventListener('click', (event) => {
        if (event.target === dialog) {
            dialog.close();
        }
    }); 

    // Append content to dialog
    dialog.appendChild(message);
    dialog.appendChild(registerButton);
    dialog.appendChild(loginButton);

    // Append dialog to the body
    document.body.appendChild(dialog);

    // Show the dialog
    dialog.showModal();
}