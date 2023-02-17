# Incrementer

Un semplice programma che mi incrementa la versione delle referenze nei 
miei javascript dove si trova la seguente stringa:

    import store from './store/index.js?version=102'
e cerca di incrementare la versione, in questo caso 102, se il file 
index.js è stato modificato dopo una certa data. 
Il cambiamento della versione è poi inoltrato a tutti i nodi superiori
fino al root. 

## Motivazione
Il browser esegue una cache degli _import js_. Quindi se la mia app (es. cup-service) 
cambia un file, per esmepio store/index.js ed esgue un update del file index.html 
quando pubblico una nuova versione, per il browser non è
sufficiente per eseguire un reload del file main.js. Solo se cambia la stringa del comando _import_
allora viene fatto il fetch sul server. Per questo ho bisogno di questo tool per cambiare automaticamente
la stringa in main.js da 

    import store from './store/index.js?version=102'
a

    import store from './store/index.js?version=103'

in modo recursivo fino al root.