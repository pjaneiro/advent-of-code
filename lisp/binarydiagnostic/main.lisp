(defpackage :binarydiagnostic
  (:use :cl)
  (:export :run-binarydiagnostic)
)
(in-package :binarydiagnostic)

(defun get-file (filename)
  (with-open-file (stream filename)
    (loop for line = (read-line stream nil)
      while line
      collect line
    )
  )
)

(defun countTheOnes (data len)
  (setf ones (make-array len :initial-element 0))
  (dolist (entry data)
    (loop for index from 0 to (- len 1)
      do (if (char= (char entry index) #\1) (setf (aref ones index) (+ (aref ones index) 1)))
    )
  )
  ; (pprint ones)
  ones
)

(defun challenge1 (inputData)
  (setq gamma 0)
  (setq epsilon 0)
  (setq elemLength (length (nth 0 inputData)))
  (setq count (length inputData))
  (setq ones (countTheOnes inputData elemLength))
  (loop for index from 0 to (- elemLength 1)
    do (if (> (aref ones index) (/ count 2)) (setq gamma (+ gamma (ash 1 (- (- elemLength 1) index)))) (setq epsilon (+ epsilon (ash 1 (- (- elemLength 1) index)))))
  )
  (* gamma epsilon)
)

(defun challenge2 (inputData)
  0
)

(defun run-binarydiagnostic ()
  (format t "Day 3 - Binary Diagnostic~C" #\linefeed)
  (setq fileData (get-file "binarydiagnostic/input.txt"))
  (format t "Challenge 1: ~d~C" (challenge1 fileData) #\linefeed)
  (format t "Challenge 2: ~d~C" (challenge2 fileData) #\linefeed)
)
