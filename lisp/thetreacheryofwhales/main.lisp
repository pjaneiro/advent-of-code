(defpackage :thetreacheryofwhales
  (:use :cl)
  (:export :run-thetreacheryofwhales)
)
(in-package :thetreacheryofwhales)

(defun delimiterp (c) (or (char= c #\Space) (char= c #\,)))

(defun split (string &key (delimiterp #'delimiterp))
  (loop :for beg = (position-if-not delimiterp string)
    :then (position-if-not delimiterp string :start (1+ end))
    :for end = (and beg (position-if delimiterp string :start beg))
    :when beg :collect (subseq string beg end)
    :while end))

(defun get-file (filename)
  (with-open-file (stream filename)
    (map 'list #'parse-integer (split (read-line stream nil)))
  )
)

(defun challenge1 (inputData)
  (setq minimum most-positive-fixnum)
  (setq maximum 0)
  (setq totalFuel most-positive-fixnum)
  (setq weights (make-hash-table))
  (dolist (cur inputData)
    (if (< cur minimum) (setq minimum cur))
    (if (> cur maximum) (setq maximum cur))
    (setf (gethash cur weights) (+ (if (gethash cur weights) (gethash cur weights) 0) 1))
  )
  (loop for i from minimum to maximum
    do (setq curFuel 0)
    (loop for k being the hash-keys in weights using (hash-value v)
      do (setq curDistance (abs (- k i)))
      (setq curFuel (+ curFuel (* v curDistance)))
    )
    (if (< curFuel totalFuel) (setq totalFuel curFuel))
  )
  totalFuel
)

(defun challenge2 (inputData)
  (setq minimum most-positive-fixnum)
  (setq maximum 0)
  (setq totalFuel most-positive-fixnum)
  (setq weights (make-hash-table))
  (dolist (cur inputData)
    (if (< cur minimum) (setq minimum cur))
    (if (> cur maximum) (setq maximum cur))
    (setf (gethash cur weights) (+ (if (gethash cur weights) (gethash cur weights) 0) 1))
  )
  (loop for i from minimum to maximum
    do (setq curFuel 0)
    (loop for k being the hash-keys in weights using (hash-value v)
      do (setq curDistance (abs (- k i)))
      (setq crabFuel (/ (* (+ curDistance 1) curDistance) 2))
      (setq curFuel (+ curFuel (* v crabFuel)))
    )
    (if (< curFuel totalFuel) (setq totalFuel curFuel))
  )
  totalFuel
)

(defun run-thetreacheryofwhales ()
  (format t "Day 7 - The Treachery of Whales~C" #\linefeed)
  (setq fileData (get-file "thetreacheryofwhales/input.txt"))
  (format t "Challenge 1: ~d~C" (challenge1 fileData) #\linefeed)
  (format t "Challenge 2: ~d~C" (challenge2 fileData) #\linefeed)
)
