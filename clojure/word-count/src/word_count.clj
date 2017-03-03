(ns word-count
  (:require [clojure.string :as string]
            [clojure.string :as str]))

(defn word-count [words]
  (->> words
       (re-seq #"\w+")
       (map str/lower-case)
       frequencies))
