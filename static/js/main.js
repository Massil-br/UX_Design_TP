let page = 1;
const limit = 10;
const container = document.getElementById('posts-container');

async function loadPosts(){
    const response = await fetch(`/api/posts?page=${page}&limit=${limit}`);
    const posts = await response.json();
    console.log(posts); // Check the structure of the posts

    posts.forEach(post => {
        const postDiv = document.createElement('div');
        postDiv.className = 'post';
        postDiv.innerHTML = `
            <h2>${post.name}</h2>
            <p>${post.description}</p>
            <p>Price: $${post.price.toFixed(2)}</p>
            <img src="${post.image_url}" alt="${post.name}"/>
        `;
        container.appendChild(postDiv);
    });

    page++;  // Increment page after loading the posts
}

// Event listener for infinite scrolling
window.addEventListener('scroll', () => {
    // If the user scrolls to the bottom of the page, load more posts
    if (window.innerHeight + window.scrollY >= document.body.offsetHeight) {
        loadPosts();
    }
});

// Load the first set of posts when the page loads
loadPosts();
