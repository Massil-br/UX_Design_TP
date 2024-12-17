let page = 1;
const limit = 14;
const container = document.getElementById('posts-container');

async function loadPosts(){
    const response = await fetch(`/api/posts?page=${page}&limit=${limit}`);
    const posts = await response.json();
    console.log(posts); 

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

    page++;  
}

window.addEventListener('scroll', () => {
    
    if (window.innerHeight + window.scrollY >= document.body.offsetHeight) {
        loadPosts();
    }
});


loadPosts();
