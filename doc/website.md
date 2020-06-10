
Pfeile zeigen die Kontaktaufnahme an. Die Spitze des Pfeils zeigt auf den kontaktierten Service.
# Services 

## order
Koordiniert eine Bestellung. Reserviert zunächst einen Artikel, leitet dann die Bezahlung und Versand ein, bei Misserfolg wird die Reservierung storniert.

### Daten
enthält alle bisherigen Bestellungen (OrderId -> []{ArtikelId, Preis})

### Kommuniziert 
* synchron mit shipment, um einen Artikel, der bestellt werden soll, aber noch nicht bezahlt ist, zu reservieren
* synchron mit catalog, um Preise abzufragen
* synchron mit Payment: Die Artikeldaten werden 
* asynchron, um einen reservierten Artikel zu verschicken oder Reservierung zu stornieren

## payment
Koordininiert einen Bezahlvorgang. Fragt dazu Bezahldaten bei customer an und sendet diese mit dem zu zahlenden Betrag an einen Zahlungsdienstleister. Kontaktiert shipment, sobald die Bezahlung vollständig ist.

### Daten
speichert die bisherigen Bezahlvorgänge: OrderID -> {ArtikelID->Preis}

### Kommuniziert
* synchron mit customer. Die Kommunikation muss
abgeschlossen sein, bevor der Zahlungsdienstleister eingebunden werden kann
* (a)synchron mit dem Zahlungsdienstleister. Die Transaktion muss erfolgreich abgeschossen sein, bevor payment eine erfolgreiche Zahlung meldet. Dies kann gleich oder später erfolgen, letzteres sollte mit einem Timeout verbunden sein
* asynchron mit shipment, um den Eingang der Zahlung mitzuteilen

## shipment
Koordiniert den Versand eines Artikels. Reserviert Artikel im Lager für den Versand, sodass sie nocht für weitere Bestellungen vorgemerkt werden.

### Daten
Enthält die Versandhistorie: OrderID -> []{ArtikelID, Anzahl, Versandstatus}

### Kommuniziert
* synchron mit stock, um zu reservieren: hier soll sofort eine Antwort kommen
* synchron mit customer, um Kundendaten abzufragen. PHier soll zeitnah eine Antwort erfolgen, dies ist bei synchroner Kommunikation 
einfacher
* asynchron mit dem shipping service. Argumentation wie beim payment_service
* asynchron mit Stock, um den Bestand upzudaten (Reservierung storniert oder Bestand reduziert, da Gegenstand versand)

## customer
Verwaltet die Kundendaten

### Daten
Enthält die Kundendaten: Name, Adresse, Zahlungsinformationen, Kundenid

### Kommunikation
antwortet auf eingehende Anfragen

## stock
Verwaltet die Güterbestände

### Daten
Enthält Daten zu jedem Gut: Anzahl im Lager, Anzahl reserviert

### Kommuniziert
* aysnchron mit Katalog, wenn sich stock ändert

## catalog
Verwaltet die Artikeldaten

### Daten
ArtikelId, Beschreibung, Preis

## customerservice
zusätzlicher Service, um eine Stornierungen und Retouren abzuwickeln (optional)

### Daten
orderId zum tracken der aktuell bearbeiteten Retouren/Stornierungen

### Kommuniziert
* asynchron mit order, um Einträge zu aktualisieren
* asynchron mit payment, um Zahlungen rückabzuwickeln
* asynchron mit shipment, um Ersatzlieferungen einzuleiten

# Ablauf einer Bestellung

1. Client fragt catalog nach allen Artikeln
2. client registriert Kunden bei customer
3. client gibt für Kunden bei order eine Bestellung auf
4. order reserviert bei Stock den Artikel
	* vorher fragt order bei customer nach, ob der Kunde existiert
	* bei Fehler (Artikel|Kunde existiert nicht) bricht order ab
5.	order holt sich von catalog den Preis der/des Artikel(s)
6.	order leitet die Artikeldaten der Bestellung an payment weiter
7.	order gibt shipment den Sendeauftrag
8.	payment signalisiert shipment das Ende des Bezahlvorgangs
9.	shipment reagiert:
	* Abbruch: gibt den Artikel wieder frei (Stornierung der Reservierung)
	* OK: weist stock an, den Vorrat an Artikeln um die entsprechende Menge zu reduzieren		

