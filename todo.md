### TODO:

- wider dialog in taller-dialogue?
- actually make use of the distinct title style colour feature.

---
Great, thank you! Anyhow, let's go back to the topic that we'd begun discussing before I noticed this problem... I think it might be nice to be ableto set a different colour than just the regular text colour for tool output in general in thems, just to be alble to make the tool output stand out from the other text a bit more if desired, lets call it oolOutput if that name is available.

We should probably take a similar approach to this that we just took for the last colour property that we added in feat/distinct-title-colour: use toolOutput if it's in the theme's JSON, otherwise fall back to the current colour (I think you'd said that was the "text" colour).

Let's do the usual song and dance routine: go implement this in a new feat/tool-output-colour branch, push it to origin, and then come back to the integration branch and merge that new branch in then do the usual rebuild and reinstall routine.
