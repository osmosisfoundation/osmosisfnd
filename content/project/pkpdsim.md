+++
# Date this page was created.
date = "2016-04-27"

# Project title.
title = "pkpdsim"

# Project summary to display on homepage.
summary = "pkpdsim is a R library for simulation of PKPD models defined as ODE systems"

# Optional image to display on homepage (relative to `static/img/` folder).
image_preview = "pkpdsim.png"

# Optional image to display on project detail page (relative to `static/img/` folder).
image = "pkpdsim.png"

# Tags: can be used for filtering projects.
# Example: `tags = ["r", "simulation", "pkpd"]`
tags = ["pkpdsim","r", "simulation", "pkpd"]

# Optional external URL for project (replaces project detail page).
external_link = "https://github.com/ronkeizer/PKPDsim"

# Does the project detail page use math formatting?
math = true

+++

PKPDsim is an R package for numerical integration of ODE systems, in particular pharmacokinetic-pharmacodynamic (PK-PD) mixed-effects models.

In pharmacometrics, models are often defined as systems of ordinary differential equations (ODEs). Although solving ODEs numerically in R is relatively straightforward using e.g. the deSolve library, the implementation of e.g. infusions and complex dosing regimens as well as the incorporation of random effects is cumbersome. Outside of R, a tool like Berkeley Madonna provides excellent interactivity features and is fast, but is much inferior to R regarding plotting functionality, cumbersome regarding implementation of dose regimens and multi-level variability, and also not open source.

In short, the issues above were the impetus for developing the PKPDsim library in R. For fast numerical integration of the ODEs, the module uses the Boost C++ library under the hood.

PKPDsim is released under the MIT open source license.
