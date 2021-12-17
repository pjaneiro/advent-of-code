(defpackage :giantsquid
  (:use :cl)
  (:export :run-giantsquid)
)
(in-package :giantsquid)

(defstruct input
  order
  boards
)

(defstruct house
  number
  marked
)

(defun delimiterp (c) (or (char= c #\Space) (char= c #\,)))

(defun split (string &key (delimiterp #'delimiterp))
  (loop :for beg = (position-if-not delimiterp string)
    :then (position-if-not delimiterp string :start (1+ end))
    :for end = (and beg (position-if delimiterp string :start beg))
    :when beg :collect (subseq string beg end)
    :while end))

(defun parse-house (input)
  (make-house :number (parse-integer input) :marked nil)
)

(defun winning-board (board)
  (loop for index from 0 to 4
    do (if (and (eql t (house-marked (nth index (nth 0 board)))) (eql t (house-marked (nth index (nth 1 board)))) (eql t (house-marked (nth index (nth 2 board)))) (eql t (house-marked (nth index (nth 3 board)))) (eql t (house-marked (nth index (nth 4 board))))) (return-from winning-board t))
    (if (and (eql t (house-marked (nth 0 (nth index board)))) (eql t (house-marked (nth 1 (nth index board)))) (eql t (house-marked (nth 2 (nth index board)))) (eql t (house-marked (nth 3 (nth index board)))) (eql t (house-marked (nth 4 (nth index board))))) (return-from winning-board t))
  )
  nil
)

(defun mark-house-on-board (board number)
  (loop for row from 0 to 4
    do (loop for col from 0 to 4
      do (if (= number (house-number (nth col (nth row board)))) (progn
        (setf (house-marked (nth col (nth row board))) t)
        (return-from mark-house-on-board t)
      ))
    )
  )
  nil
)

(defun get-file (filename)
  (setq result (make-input))
  (with-open-file (stream filename)
    (setq firstLine (read-line stream nil))
    (setf (input-order result) (map 'list #'parse-integer (split firstLine)))
    (setq currentBoard (make-list 1))
    (setf (nth 0 currentBoard) (make-list 5))
    (loop for line = (read-line stream nil)
      while line
      do (if (string= line "") (progn
          (setq counter 0)
          (setq currentBoard (make-list 1))
          (setf (nth 0 currentBoard) (make-list 5))
        ) (progn
          (setf (nth counter (nth 0 currentBoard)) (map 'list #'parse-house (split line)))
          (if (= counter 4) (progn
            (setf (input-boards result) (append (input-boards result) currentBoard))
          ))
          (setq counter (+ counter 1))
        )
      )
    )
  )
  result
)

(defun challenge1 (inputData)
  (setq order (copy-list (input-order inputData)))
  (setq boards (copy-list (input-boards inputData)))
  (setq score 0)
  (loop for index from 0 to (- (length order) 1)
    do (setq curNumber (nth index order))
    (dolist (curBoard boards)
      (mark-house-on-board curBoard curNumber)
      (if (winning-board curBoard) (progn
        (loop for i from 0 to 4
          do (loop for j from 0 to 4
            do (if (eql nil (house-marked (nth j (nth i curBoard)))) (setq score (+ score (house-number (nth j (nth i curBoard))))))
          )
        )
        (return-from challenge1 (* score curNumber))
      ))
    )
  )
  0
)

(defun challenge2 (inputData)
  (setq order (copy-list (input-order inputData)))
  (setq boards (copy-list (input-boards inputData)))
  (setq score 0)
  (setq closed 0)
  (setq statuses (make-list (length boards)))
  (loop for index from 0 to (- (length order) 1)
    do (setq curNumber (nth index order))
    (loop for k from 0 to (- (length boards) 1)
      do (setq curBoard (nth k boards))
      (mark-house-on-board curBoard curNumber)
      (if (and (not (nth k statuses)) (winning-board curBoard)) (progn
        (setf (nth k statuses) t)
        (setq closed (+ closed 1))
      ))
      (if (= closed (length boards)) (progn
        (loop for i from 0 to 4
          do (loop for j from 0 to 4
            do (if (eql nil (house-marked (nth j (nth i curBoard)))) (setq score (+ score (house-number (nth j (nth i curBoard))))))
          )
        )
        (return-from challenge2 (* score curNumber))
      ))
    )
  )
  0
)

(defun run-giantsquid ()
  (format t "Day 4 - Giant Squid~%")
  (setq fileData (get-file "giantsquid/input.txt"))
  (format t "Challenge 1: ~d~%" (challenge1 fileData))
  (format t "Challenge 2: ~d~%" (challenge2 fileData))
)
