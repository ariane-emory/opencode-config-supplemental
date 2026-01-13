# New Feature Request:

### Feature Description: Poisonous Giant Snake Monster

The goal is to introduce a new monster to the game: a poisonous giant snake. This monster will add a new layer of challenge and strategy to the gameplay, leveraging the existing systems for monsters, hazards, and combat mechanics. The implementation of this feature must integrate seamlessly with the current codebase and adhere to the established design patterns and architecture of the game.

### Key Requirements

1. **Monster Attributes**:
   - The poisonous giant snake must have unique attributes such as health, attack power, movement speed, and poison damage. These attributes should be defined in the `monster` package, specifically in `monster.go` and `monster_test.go`.
   - The snake's poison effect should deal damage over time, requiring integration with the `magic` or `hazard` systems for status effects.

2. **Behavior and AI**:
   - The snake must exhibit distinct movement and attack patterns. For example, it could move in a slithering pattern and prioritize attacking the nearest player.
   - The AI logic for the snake should be implemented in the `monster` package, leveraging existing AI patterns or introducing new ones if necessary.

3. **Poison Mechanic**:
   - The poison effect should be a status ailment that reduces the player's health over several turns. This mechanic should integrate with the `hazard` or `magic` systems, as seen in files like `hazard.go`, `effects.go`, and their respective test files.
   - The poison effect should include visual or textual feedback to the player, such as a message in the `ui/message_log.go` or a status indicator in the `ui/stats_panel.go`.

4. **Integration with Existing Systems**:
   - The snake must be added to the monster factory in `factory.go` to ensure it can spawn in the game world.
   - The spawning logic should consider the snake's rarity and the appropriate dungeon levels for its appearance. This may involve modifying `dragon_spawn_frequency_test.go` or creating a new test for snake spawn frequency.
   - The snake's interactions with players, items, and other monsters must be tested to ensure compatibility with the game's combat and event systems.

5. **Testing and Validation**:
   - Comprehensive unit tests must be added to validate the snake's behavior, attributes, and interactions. These tests should be placed in `monster_test.go`, `hazard_test.go`, and other relevant test files.
   - Integration tests should ensure the snake functions correctly within the broader game systems. These tests may involve files like `dungeon_test.go` or `gameloop_test.go`.

6. **Art and Representation**:
   - The snake's visual representation in the terminal must align with the game's aesthetic. This may involve adding a new character or symbol to represent the snake in the `rendering/map.go` file.
   - Any animations or effects related to the snake, such as slithering or poison attacks, should be implemented in the `animation` package.

7. **Documentation**:
   - Update `STRUCTURE.md` to include any new files or significant changes to existing files.
   - Document the snake's attributes, behavior, and mechanics in a new or existing design document, such as `GIANT_SCORPION_IMPLEMENTATION_PLAN.md`.

### Additional Considerations

- Ensure the snake's poison mechanic is balanced and does not overly penalize the player. Playtesting may be required to fine-tune its difficulty.
- Consider reusability of the poison mechanic for other monsters or traps in the future.
- Maintain consistency with the game's existing lore and theme when describing the snake in the game's narrative or flavor text.

## IMPORTANT: Employ our standard methodology to maximize the odds of successful implementation!

So long as you proceed systematically, work hard, and adhere to our standard practices, your successful completion of the task is as good as guaranteed! Remember:

- Start by thinking the implementation of this feature through thoroughly. Then, you MUST break the the feature down into small steps to produce a detailed, step-by-step plan that you will use to implement this feature. Group the plan's steps into "phases": the code MUST continue to build correctly and all tests MUST pass after each phase is completed.
- Next, write the plan into an appropriately named new Markdown file in the project's root directory which includes checkboxes in which to mark the completion of each step.
- Proceed to systematically implement the plan that you just wrote in the Markdown file. You MUST check off each step you've completed in the Markdown file immediately as you complete it, you MAY NOT proceed to the next step until you have checked off the current step.
- Follow through and finish the job: you MUST continue complete the task! Keep working until every step in the Markdown file has been checked off and the entire plan has been completed. The code MUST build correctly and all tests MUST pass afterwards.
