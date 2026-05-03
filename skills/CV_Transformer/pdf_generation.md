# PDF Generator: Execution & Rendering

This process takes the `cv_data.json` file and produces the final PDF document using the Go engine.

## Execution Requirements
- The `cv_data.json` file must be up to date.
- The command must be executed from the `src/` folder.

## Generation Commands
To launch the generation, use the following command:

```bash
go run main.go -template [TEMPLATE_NAME]
```

## Template Selection
- **tech**: (Recommended) Modern design in orange and violet.
- **modern**: Classic-modern two-column design.
- **default**: Traditional linear design.

**Golden Rule:** If the user does not specify a template, the system MUST use the **default** template.

## PDF Parameters
- Format: A4.
- Rendering: Chromedp (High Fidelity).
- Output Location: Project root.
