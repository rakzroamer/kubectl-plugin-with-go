# kubectl-plugin-with-go
Trying to create a kubectl plugin using Go and Client-go to extend kubectl via plugins to achieve ns kill or terminate

# Motivation
When working in the homelab or in dev mode - we tend to go for the easier options to get rid of k8s objects created in your playground namespace - ns. Think you were so busy creating lot many resources and want to remove resources at one 'Go', that's where namespace level - comes into picture. There are chances that k8s api doesn't respond or other reasons that underlying objects are stuck - this is where a namespace plugin to help avert certain conditions

## Anology
In the case of Azure cloud, everything is constructed under Az subscriptions. [Google - at Project level]. It's a convenient one stop option of get rid of all the resources created - obsiouvly there are options within 3 days to claim back ;). But the big idea is use a way of abstraction to "not to dwell at individual resources to delete" but to take care at higher level.

## kubectl plugins
Instead of using k kill ns nsname, this repo will walkthrough the process of creating kubectl plugin which interms using additional args and take care as a single command.

You may ask why don't we go with alias?. It can be, but here the intention is to explore kubectl plugins using Golang and curiosity to extend further is my big idea.