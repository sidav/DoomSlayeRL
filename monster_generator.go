package main

import "DoomSlayeRL/routines"

func rmg_getRandomValue(a []string) string {
	return a[routines.Random(len(a))]
}

func rmg_generateName(pref, mid, suff bool) string {
	var prefixes = []string {
		"Unseen", "Awful", "Infernal", "Fearsome", "Wrathful", "Fiery", "Crushing", "Cybernetic", "Subject #451:",
		"Unwilling", "Forgotten",
	}
	var syllables = []string {
		"Ber", "Char", "Zar", "Imp", "Vile", "Infer", "Rev", "Reve", "Raxx", "Arch", "For", "In", "Suc", "Arach",
	}
	var middles = []string {
		"no", "er", "ke", "ker", "ling", "ath", "hel", "zor", "vile", "cubus", "ron", "for", "ven", "lar", "tacu",
	}
	var endings = []string {
		"er", "ker", "ling", "ath", "hel", "zor", "vile", "cubus", "ron", "for", "fer", "rer", "tron",
	}
	var suffixes = []string {
		"of Doom", "of Gehenna", "the Hellish", "of Abaddon", "the Smiting", "The Ripper", "The UAC Experiment",
		"the Eternal", "Unnamed", "of Hell", "The Torn Soul", "The Aeons Torn",
	}
	name := ""
	if pref {
		name += rmg_getRandomValue(prefixes) + " "
	}
	name += rmg_getRandomValue(syllables)
	if mid {
		name += rmg_getRandomValue(middles)
	}
	name += rmg_getRandomValue(endings)
	if suff {
		name += " "+rmg_getRandomValue(suffixes)
	}
	return name
}

