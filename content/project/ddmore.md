+++
# Date this page was created.
date = "2016-04-27"

# Project title.
title = "DDMore"

# Project summary to display on homepage.
summary = "ddmore enables access to curated and shared knowledge for the benefit of model-informed drug discovery, development and usage"

# Optional image to display on homepage (relative to `static/img/` folder).
image_preview = "ddmore.png"

# Optional image to display on project detail page (relative to `static/img/` folder).
image = "ddmore.png"

# Tags: can be used for filtering projects.
# Example: `tags = ["r", "simulation", "pkpd"]`
tags = ["ddmore","r", "simulation", "pkpd"]

# Optional external URL for project (replaces project detail page).
external_link = "https://github.com/metrumresearchgroup/mrgsolve"

# Does the project detail page use math formatting?
math = true

+++

mrgsolve facilitates simulation in R from hierarchical, ordinary differential equation (ODE) based models typically employed in drug development. The modeler creates a model specification file consisting of R and C++ code that is parsed, compiled, and dynamically loaded into the R session. Input data are passed in and simulated data are returned as R objects, so disk access is never required during the simulation cycle after compiling.
