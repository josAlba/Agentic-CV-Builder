# SKILL: CV Engine Management

This skill is the control center for managing and generating your professional Curriculum Vitae. It is divided into three specialized modules:

## Skill Modules

1.  **[skill.md](./skill.md)** (Master Index): Acts as the main entry point and explains the overall workflow.
2.  **[data_extraction.md](./data_extraction.md)**: Contains all technical logic to transform your Markdown files into structured JSON, with cleaning and hierarchy rules.
3.  **[pdf_generation.md](./pdf_generation.md)**: Details how to run the Go engine, available templates (`default`, `modern`, `tech`), and the default template rule.
4.  **[markdown_structure.md](./markdown_structure.md)**: Defines the required format for the source Markdown files.

## Master Workflow
1. **Update**: Edit your source files in `references/[PROFILE_NAME]/`.
2. **Sync**: Use the **Extraction** module to update `cv_data.json`.
3. **Generate**: Use the **Generation** module to create the final PDF document.

---
*This modular system allows for scaling the curriculum with new templates or data sources while maintaining consistency.*
