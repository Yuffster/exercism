(ns rna-transcription)

(defn to-rna [s]
  (if (> (count s) 1)   ; Map string by character.
      (->> s 
         (map str)
         (map to-rna)
         (apply str))
      (condp = s        ; Transcribe each character.
        "G" "C"
        "C" "G"
        "T" "A"
        "A" "U"
        (assert false (apply str ["Invalid code: " s])))))

