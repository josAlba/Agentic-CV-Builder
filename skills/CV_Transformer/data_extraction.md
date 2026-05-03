# CV Transformer: Extraction & Consolidation

This process transforms unstructured Markdown files into a single source of truth in JSON format.

## Data Source
Data is obtained from the `references/[PROFILE_NAME]/` folder (modular files per section). Files must follow the guidelines in **[markdown_structure.md](./markdown_structure.md)**.

## Extraction Rules
1. **Personal Profile**: Extract from the original `README.md`, cleaning names if the target is a public repository.
2. **Experience**: Convert each `experiencia_*.md` file into an object in the `experience` list.
3. **Achievements**: Bullet points are mapped to the `achievements` list.
4. **Skills**: Classify into `technical`, `soft`, and `tools`.
5. **Technologies**: Extract from the technology sections used in each job.

## JSON Schema (cv_data.json)
The resulting file must strictly follow the structure defined in the root schema to ensure compatibility with the Go generator.

**Key Instruction:** 
> "Analyze changes in the source Markdown files and sync the `cv_data.json` while maintaining chronological hierarchy (newest first)."
