(defpackage :lanternfish
  (:use :cl)
  (:export :run-lanternfish)
)
(in-package :lanternfish)

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

(defun worker (inputData generations)
  (setq buckets (make-list 9 :initial-element 0))
  (dolist (cur inputData)
    (setf (nth cur buckets) (+ (nth cur buckets) 1))
  )
  (loop for generation from 0 to (- generations 1)
    do (setq newBuckets (make-list 9 :initial-element 0))
    (loop for i from 0 to 8
      do (setq count (nth i buckets))
      (if (= i 0) (progn
          (setf (nth 6 newBuckets) (+ (nth 6 newBuckets) count))
          (setf (nth 8 newBuckets) (+ (nth 8 newBuckets) count))
        ) (progn
          (setf (nth (- i 1) newBuckets) (+ (nth (- i 1) newBuckets) count))
        )
      )
    )
    (setq buckets (copy-list newBuckets))
  )
  (apply #'+ buckets)
)

(defun challenge1 (inputData)
  (worker inputData 80)
)

(defun challenge2 (inputData)
  (worker inputData 256)
)

(defun run-lanternfish ()
  (format t "Day 6 - Lantern Fish~%")
  (setq fileData (get-file "lanternfish/input.txt"))
  (format t "Challenge 1: ~d~%" (challenge1 fileData))
  (format t "Challenge 2: ~d~%" (challenge2 fileData))
)
