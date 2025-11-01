# Virtual Studio XR - Project Progress

## Week 1 – Project Setup & Basic WebXR Scene ✅

**Goal**: Set up a modern TS project that can render a 3D scene and enter VR with a single button.

### Tasks:
- [x] Create Vite + TypeScript project structure
- [x] Install Three.js and @types/three
- [x] Create folder structure (scene/, loop/, utils/)
- [x] Create `src/main.ts` - app entrypoint with VRButton
- [x] Create `src/scene/setupScene.ts` - scene, camera, lights
- [x] Create `src/scene/createObjects.ts` - spinning cube
- [x] Create `src/loop/renderLoop.ts` - XR-aware animation loop
- [x] Create `src/utils/resizeHandler.ts` - window resize handler
- [x] Enable WebXR (renderer.xr.enabled = true)
- [x] Add VRButton to DOM
- [x] Run dev server and verify: grey background, spinning cube, VR button

**Status**: ✅ Complete - Verified working with screenshot

---

## Week 2 – Interaction & Controllers ⏳

**Goal**: Add input to the scene - VR controller → raycast → select object → change it.

### Tasks:
- [ ] Create `src/managers/` folder
- [ ] Create `src/managers/InteractionManager.ts`
  - [ ] Set up XR controller via `renderer.xr.getController(0)`
  - [ ] Implement raycasting with THREE.Raycaster
  - [ ] Add selectstart event listener
  - [ ] Implement color change on object selection
- [ ] Create `src/scene/createMultipleObjects.ts`
  - [ ] Add cube mesh
  - [ ] Add sphere mesh
  - [ ] Add cone mesh
  - [ ] Return array of interactable objects
- [ ] Update `src/main.ts` to use InteractionManager
- [ ] Register created objects as interactable
- [ ] Test: Point at objects with controller and change color on select

**Expected Result**: Scene has 3 objects (cube, sphere, cone) that change color when selected with controller

---

## Week 3 – AR, Lighting, and Physics ⏳

**Goal**: Add AR mode and basic physics so objects can rest on a surface and look more real.

### Tasks:
- [ ] Install cannon-es physics engine (`npm install cannon-es`)
- [ ] Create `src/physics/` folder
- [ ] Create `src/physics/PhysicsManager.ts`
  - [ ] Set up CANNON.World with gravity
  - [ ] Create ground plane
  - [ ] Implement addMeshWithPhysics method
  - [ ] Implement update method to sync physics → Three.js
  - [ ] Create shape from mesh helper method
- [ ] Create `src/ar/` folder
- [ ] Create `src/ar/setupAR.ts`
  - [ ] Enable AR session using ARButton
  - [ ] Configure required features (hit-test)
- [ ] Update `src/main.ts` (Week 3 version)
  - [ ] Add AR button
  - [ ] Create PhysicsManager instance
  - [ ] Register all objects with physics
  - [ ] Update render loop to step physics world
- [ ] Test: Objects should fall and rest on ground plane
- [ ] Test: AR button appears (requires HTTPS + supported device)

**Expected Result**: Objects have physics simulation, AR mode available

---

## Week 4 – XR UI, Scene Persistence, Deployment ⏳

**Goal**: Make it feel like a real app with 3D UI, object creation, and scene saving.

### Tasks:
- [ ] Create `src/state/` folder
- [ ] Create `src/state/SceneStateManager.ts`
  - [ ] Define SavedObject interface
  - [ ] Implement save method (to localStorage)
  - [ ] Implement load method (from localStorage)
  - [ ] Handle object type, position, and color
- [ ] Create `src/ui/` folder
- [ ] Create `src/ui/XRMenu.ts`
  - [ ] Create 3D panel background
  - [ ] Add "Add" button mesh
  - [ ] Add "Reset" button mesh
  - [ ] Add "Save" button mesh
  - [ ] Define XRMenuCallbacks type
  - [ ] Return menu group positioned in space
- [ ] Update `src/main.ts` (Week 4 version)
  - [ ] Create SceneStateManager instance
  - [ ] Create XR menu with callbacks
  - [ ] Implement onAdd callback (create new mesh)
  - [ ] Implement onReset callback (clear scene)
  - [ ] Implement onSave callback (persist to localStorage)
  - [ ] Add menu to scene
  - [ ] Optional: Load saved scene on startup
- [ ] Test: Add objects dynamically
- [ ] Test: Save and load scene state
- [ ] Test: Reset clears scene

**Expected Result**: Working 3D UI menu, ability to add/reset/save scene

---

## Deployment 📦

### Tasks:
- [ ] Run `npm run build` to create production build
- [ ] Test production build locally (`npm run preview`)
- [ ] Choose deployment platform:
  - [ ] Option A: Vercel
  - [ ] Option B: Netlify
  - [ ] Option C: GitHub Pages
- [ ] Configure HTTPS (required for WebXR on devices)
- [ ] Deploy and test on actual XR device
- [ ] Verify VR mode works
- [ ] Verify AR mode works (if device supports)

---

## Optional Enhancements 🚀

### Future Ideas:
- [ ] Add hand tracking support
- [ ] Add networking/multiplayer (WebRTC / Colyseus)
- [ ] Improve UI with `three-mesh-ui` library
- [ ] Add GLTF model loading instead of primitive shapes
- [ ] Add teleportation locomotion
- [ ] Add haptic feedback
- [ ] Add spatial audio
- [ ] Add custom shaders/materials
- [ ] Add shadows and advanced lighting
- [ ] Add post-processing effects

---

## Notes

- **Current Status**: Week 1 Complete ✅
- **Next Step**: Begin Week 2 - Interaction & Controllers
- **Dev Server**: Running on http://localhost:5173/
- **XR Testing**: Requires HTTPS and WebXR-capable device (e.g., Meta Quest 3)

