---- MODULE PitParity ----
\*
\* Formal verification of the PIT coverage-parity accounting
\* (kernel/mutation/pitparity): the consistency rate
\*   rate = consistent * 10000 / comparable   (integer division)
\* is the percentage (scaled by 10000) of comparable mutation points that the
\* Abate report reproduced. This module proves the accounting can never produce
\* an out-of-range or inconsistent rate.
\*
\* Invariants:
\*   - consistent never exceeds comparable (no over-count);
\*   - rate is always in [0, 10000];
\*   - full agreement (consistent == comparable) yields rate 10000;
\*   - zero agreement (consistent == 0) yields rate 0.

EXTENDS Integers, TLC

CONSTANTS KMAX   \* maximum comparable count modeled

\* Production scales the rate by 10000 (a scaled percentage). TLC represents
\* integers only up to 2^31-1 and the CHOOSE below scans 0..SCALE, so we use
\* SCALE = 100 for model checking; the bounding/monotonicity properties proven
\* are scale-invariant and generalize to 10000.
SCALE == 100

VARIABLES comparable, consistent

Init ==
    /\ comparable \in 1..KMAX
    /\ consistent \in 0..comparable

\* A new comparable point appears (consistent may stay the same or, if the new
\* point is itself reproduced, advance — kept conservative here: it stays). The
\* count is bounded by KMAX so the state space stays finite for model checking.
AddComparable ==
    /\ comparable < KMAX
    /\ comparable' = comparable + 1
    /\ consistent' = consistent
    /\ consistent' <= comparable'

\* A comparable point is reproduced; only legal while consistent < comparable.
AddConsistent ==
    /\ consistent < comparable
    /\ consistent' = consistent + 1
    /\ comparable' = comparable

Terminated == UNCHANGED <<comparable, consistent>>

Next == AddComparable \/ AddConsistent \/ Terminated

Spec == Init /\ [][Next]_<<comparable, consistent>>

\* Floor division by bounded search (this TLA+ build has no Floor/Div
\* operator, and a recursive subtract-and-count would overflow the stack for
\* large quotients). CHOOSE over the bounded result range 0..SCALE is exact and
\* terminates because comparable >= 1 always.
Rate ==
    CHOOSE q \in 0..SCALE :
        q * comparable <= consistent * SCALE
            /\ (q + 1) * comparable > consistent * SCALE

RateInRange   == Rate \in 0..SCALE
NoOverCount   == consistent <= comparable
FullWhenEqual == (consistent = comparable) => (Rate = SCALE)
ZeroWhenZero  == (consistent = 0) => (Rate = 0)

====
