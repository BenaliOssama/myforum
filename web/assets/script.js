import { RenderPost } from "./rendring.js"
import {handleReact} from "./likes.js"

export const GetData = async (postIds) => {
    let target = []
    try{
        if (postIds === false) {
            postIds = [];
            let response = await fetch('http://localhost:8080/api/posts');
            if (!response.ok) throw new Error("Network response was not ok");
            let lastPostId = await response.json();
            for (let postId = 1; postId <= lastPostId; postId++) {
                postIds.push(postId);
            }
        }
        for (let i = postIds.length - 1; i >= 0; i--) {
            let link = `http://localhost:8080/api/posts?id=${postIds[i]}`;
            let postResponse = await fetch(link);
            if (!postResponse.ok) throw new Error("Failed to fetch post data");
            let post = await postResponse.json();
            if (post.PostId !== 0) {
                target.push(post)
                RenderPost(target)
            }
        }
    } catch (err) {
        console.error("Error fetching data:", err);
    }
};

export const getComment = async (post, id) => {
    try {
        const res = await fetch(`http://localhost:8080/api/comments?post=${id}`)
        if (res.ok) {
            const allComment = await res.json()
            if (allComment) {
                for (let comment of allComment) {
                    const com = document.createElement('div');
                    com.classList.add('comment');
                    com.innerHTML = `
                        <div class="one_comment">
                            <p><i class="fa fa-user"></i> ${comment.user_name}:<i> ${comment.content}</i> </p> 
                            <div class="actions">
                                <button data-clicked="${comment.clicked}"  class="com_like" 
                                style="background-color: ${comment.clicked ? '#15F5BA' : 'white'};">
                                    <i class="fas fa-thumbs-up"></i> <span class="count">${comment.like_count}</span>
                                </button>

                                <button data-clicked="${comment.disclicked}"  class="com_dislike" 
                                style="background-color: ${comment.disclicked ? '#15F5BA' : 'white'};">
                                    <i class="fas fa-thumbs-down"></i> <span class="count">${comment.dislike_count}</span>
                                </button>
                            </div>
                        <div>
                    `;

                    post.insertAdjacentElement('beforeend', com);

                    // Add event listeners for like and dislike buttons
                    const likeButton = com.querySelector('.com_like');
                    const dislikeButton = com.querySelector('.com_dislike');

                    likeButton.addEventListener('click', async () => {
                        await handleReact(likeButton, dislikeButton,comment.comment_id, 'like', "comment");
                    });

                    dislikeButton.addEventListener('click', async () => {
                        await handleReact(dislikeButton, likeButton, comment.comment_id, 'dislike', "comment");
                    });
                }
            }
        }
    } catch (error) {
        console.error(error);
    }
}

export const logoutEvent = (log) => {
    log.addEventListener('click', async () => {
        try {
            const response = await fetch('http://localhost:8080/api/logout', {
                method: 'POST',
                credentials: 'include'
            });

            if (response.ok) {
                console.log('Logged out successfully');
                window.location.href = '/';
            } else {
                console.error('Logout failed');
            }
        } catch (error) {
            console.error('Error logging out:', error);
        }
    });
}