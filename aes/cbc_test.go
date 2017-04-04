package aes

import (
	"cryptopals/model"
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
}

func Test_AesCbc_Encrypt_Decrypt(t *testing.T) {
	key := "WE RUST IN PEACE"

	iv := "0000000000000000"

	cbc := AesCbc{}

	encrypted, err := cbc.Encrypt(plain, key, iv)
	if err != nil {
		t.Fatal(err)
	}

	decrypted, err := cbc.Decrypt(encrypted, key, iv)
	if err != nil {
		t.Fatal(err)
	}

	if !model.Equals(decrypted, []byte(plain)) {
		t.Fatal("plain and decrypted do not match.")
	} else {
		log.Println("Test_AesCbc_Encrypt_Decrypt pass!")
	}
}

var plain = `Tremble you weakings, cower in fear
I am your ruler, land, sea and air
Immense in my girth, erect I stand tall
I am nuclear murderer I am Polaris
Ready to pounce at the touch of a button
My systems locked in on military gluttons

I rule on land, air and sea
Pass judgment on humanity
Winds blow from the bowels of hell
Will we give warnings, only time will tell
Satan rears his ugly head, to spit into the wind

I spread disease like a dog
Discharge my payload a mile high
Rotten egg air of death wrestles your nostrils
Launch the Polaris, the end doesn't scare us

When will this cease
The warheads will all rust in peace
Bomb shelters filled to the brim
Survival such a silly whim
World leaders sell missiles cheap
Your stomach turns, your flesh creeps

High priests of holocaust, fire from the sea
Nuclear winter spreading disease
The day of final conflict
All pay the price
The third world war
Rapes peace, takes life
Back to the start, talk of the part
When the Earth was cold as ice
Total dismay as the sun passed away
And the days where black as night

Eradication of Earth's
Population loves Polaris`
