# Route based rendering logic

For this project, I wanted to use route based rendering. To achieve this, I've used the request context to determine the route, and render the appropriate templates based on the route and request kind.

This does come with some complexity and principles.

1. Template building:
   The API must build templates upon API initialization.

2. Template rendering:
   The API must be able to render full pages, fragments and components based on the route and request kind.

3. Template management:
   - The file- and template name must always align.
   - Page content templates are organized in a hierarchical structure based on their route (`/routes`).
   - Page layouts are defined in the `/layouts` directory.
     - Nested layouts... ?
   - Components are defined in the `/components` directory.

You can see an example of how this works in the `web` directory.

```
web
| index.html
| layout.html
|
|__ components
|   | navbar.html
|   | count.html
|
|__ routes
    | index.html
    |
    |__ blog
    |   | index.html
    |
    |__ form
    |   | index.html
```

For this project, I've decided to use the `html/template` package to manage and manipulate templates, and the `Echo` framework to handle routing.

The render function is part of the echo context and has some limitations. Because I don't want any request context logic in the renderer, but rather in the API middleware/ handler the `Render` function must hold all generated templates in a single struct variable. To still have some freedom we will make this struct hold a `templates map[string]*template.Template`, which is a key-value store where the key is the template name and the value is the template itself. If we are smart about how we build this template map, we can leverage request context route information to dynamically render templates.

File conventions: (Inspired by Next.js App Router)

- A base layout file: `layout.html` must be created at the root of the `/routes/` directory. This layout file holds the common elements of a html site layout. It is used as a base layout for all pages. This base layout file holds an "index" template block, which is used to render the content of the page.

- A main `index.html`, which holds the page content, wrapped in a "index" template block.

- Since components are reusable pieces of HTML that can be used across multiple pages, they are not defined on a specific route but rather in a seperate `/components` directory.

# Template building.

To support route based page rendering, the templates are built and stored in an key/value map named `templates`. An advantage of this is that we will be able to directly render a full page when someone directly hits the `/blog` route.

For example:

- Page content (fragment) will be stored as: `index`.
- The `/blog` page content will be stored as: `blog/index`.
- The full `/blog` page will be stored as: `blog/page`
- Fragments, which are because of the convention coupled to the page route, are stored as: `blog/{fragmentName}`.

When we navigate via a hx-swap or get request, we only need to render a certain fragment.

# Timed animations

Let's update the form to use class-tools:

```/dev/null/add-to-list-form.html#L1-18
{{ define "add_to_list_form" }}
<form
    id="todo-form"
    class="add-to-list-form"
    hx-post="/list/add"
    hx-target="#todo-list-container"
    hx-swap="outerHTML"
>
    <input
        id="form-item-input"
        type="text"
        name="item"
        class="form-input"
        placeholder="Add items to list"
        required
        classes-add="success-input delay:0"
        classes-remove="success-input delay:1000"
    />
    <button type="submit" class="add-button" id="add-button">ADD</button>
</form>
{{ end }}
```

3. Add the new styles to your CSS. I notice you're using SASS/SCSS (because of the source map comment). Let's add these new styles:

```/dev/null/new-styles.css#L1-12
.form-input {
    /* ... your existing styles ... */
    transition: all 0.3s ease;  /* Add this for smooth transitions */
}

.success-input {
    border: 2px solid #4d6862 !important;
    background-color: #84a59d !important;
    color: rgb(82.39, 112.61, 105.28) !important;
    box-shadow: inset 2px 4px 0px #4d6862 !important;
    transform: scale(1.02);
}
```

Here's how it works:

1. When the form is submitted successfully:
   - The `success-input` class is immediately added (`delay:0`)
   - The input will change to the success state colors
   - After 1000ms (1 second), the class is automatically removed (`delay:1000`)
   - The input smoothly transitions back to its original state

The transition works like this:

- When you add an item: Normal state → Success state (immediate)
- 1 second pause
- Success state → Normal state (smooth transition)

No JavaScript is required! It's all handled by:

- HTMX class-tools extension for the class swapping
- CSS transitions for the smooth animations

The benefits of this approach:

1. Zero JavaScript
2. Declarative in the HTML
3. Smooth transitions handled by CSS
4. Follows HTMX's HTML-first philosophy
5. Automatically reverts after the specified delay

Would you like me to help you implement this in your codebase or would you like to see any modifications to the colors/timing?
