document.addEventListener('DOMContentLoaded', function() {
    const pageTitle = document.title.trim();
    const menuLinks = document.querySelectorAll('#mainMenu a');

    menuLinks.forEach(link => {
        if (link.textContent.trim() === pageTitle) {
            link.style.textDecoration = 'underline';
        }
    });
});
