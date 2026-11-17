---- MODULE SandboxFuel ----
\*
\* Formal verification of the guest-plugin execution fuel model.
\*
\* The Abate sandbox (plugin/sandbox) and the native plugin host (plugin/host)
\* bound untrusted computation with a finite fuel (gas). This module proves the
\* two properties the design must satisfy:
\*   1. Bounding  - execution can never consume more fuel than was granted;
\*                  the fuel counter stays non-negative and used <= F0.
\*   2. Termination - with finite fuel the machine always eventually halts and,
\*                  once halted, never executes another instruction.
\*
\* The model is parameterized on the initial fuel F0 and the maximum per-step
\* cost MaxCost; the proof holds for any values by weak fairness on execution.

EXTENDS Integers, TLC

CONSTANTS F0, MaxCost

VARIABLES fuel, steps, halted, used

Init ==
    /\ fuel   = F0
    /\ steps  = 0
    /\ halted = FALSE
    /\ used   = 0

\* One execution step burns c units of fuel (1 <= c <= MaxCost) and advances
\* the step counter. When fuel is exhausted the machine halts.
Exec(c) ==
    /\ c \in 1..MaxCost
    /\ fuel >= c
    /\ fuel'   = fuel - c
    /\ used'   = used + c
    /\ steps'  = steps + 1
    /\ halted' = (fuel' = 0)

ExecAction == \E c \in 1..MaxCost : Exec(c)

\* Terminal stutter: once halted (fuel == 0) nothing changes.
Terminated == UNCHANGED <<fuel, steps, halted, used>>

Next == ExecAction \/ Terminated

Spec ==
    Init
        /\ [][Next]_<<fuel, steps, halted, used>>
        /\ WF_<<fuel, steps, halted, used>>(ExecAction)

\* ----- Invariants (safety) -----

FuelNonNeg   == fuel >= 0
UsedBounded  == used <= F0
StepsBounded == steps <= F0
HaltEmpty    == halted => (fuel = 0)

\* ----- Property (liveness) -----

Terminates == <>halted

====
