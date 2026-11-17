---- MODULE StreamDeterminism ----
\*
\* Formal verification of the Abate deterministic random stream
\* (sim/rand: Stream + splitMix64).
\*
\* The load-bearing guarantee is reproducibility: a stream derived from a
\* single 64-bit seed always yields the same sequence, and Restore(seed, d)
\* must reproduce exactly the state a freshly advanced stream would hold at
\* draw d. This is the property the SIM-REC replay certification relies on.
\*
\* The SplitMix64 *value* mix is a bijection over the full word and does not
\* affect the stored `state`; `state` is pure arithmetic
\*   state = origin + draws * Gamma   (mod 2^WIDTH)
\* which is exactly what makes Restore mathematically exact. We therefore
\* model the state machine at that arithmetic level (faithful to the Go code
\* in sim/rand/rand.go and sim/rand/splitmix.go) and prove the equivalence.
\*
\* The bit width is a constant overridden to a small value for model
\* checking. Because every relation is modular arithmetic uniform in the
\* width, a proof at width W generalizes to width 64 by the same reasoning.

EXTENDS Integers, TLC, Sequences

CONSTANTS
    WIDTH,          \* modeled word width (override small for checking)
    SEEDS,          \* seeds the environment may inject
    DMAX            \* maximum draw index modeled (bounds state space)

M == 2^WIDTH
\* SplitMix64 golden-ratio step constant. TLC represents integers only up to
\* 2^31-1, and `d * Gamma` must stay within that bound, so we model Gamma with a
\* small odd constant. The reproducibility / restore-equivalence property proven
\* here is parametric in Gamma and the modulus: the algebra
\*   state = (origin + draws * Gamma) % M
\* holds for any Gamma, so the result generalizes to the production 64-bit
\* constant 0x9E3779B97F4A7C15 and full word width.
Gamma == 9

VARIABLES S   \* function {"A","B"} -> [origin, state, draws]

Streams == [origin : 0..(M-1), state : 0..(M-1), draws : 0..DMAX]

Init ==
    S = [s \in {"A","B"} |-> [origin |-> 0, state |-> 0, draws |-> 0]]

fresh(seed) ==
    [origin |-> (seed % M), state |-> (seed % M), draws |-> 0]

advanced(st) ==
    [origin |-> st.origin,
     state  |-> ((st.state + Gamma) % M),
     draws  |-> st.draws + 1]

\* Restore may only target the stream's own origin; an alien seed is rejected
\* by the real code (ErrInvalidRestore). The restored state is exactly the
\* arithmetic position origin + d*Gamma, which is what a stream advanced d
\* times would hold.
restored(st, d) ==
    [origin |-> st.origin,
     state  |-> ((st.origin + d * Gamma) % M),
     draws  |-> d]

NewA(seed) == S' = [S EXCEPT !["A"] = fresh(seed)]
NewB(seed) == S' = [S EXCEPT !["B"] = fresh(seed)]

AdvanceA == /\ S["A"].draws < DMAX
            /\ S' = [S EXCEPT !["A"] = advanced(S["A"])]
AdvanceB == /\ S["B"].draws < DMAX
            /\ S' = [S EXCEPT !["B"] = advanced(S["B"])]

RestoreA(d) == /\ d \in 0..DMAX
               /\ S' = [S EXCEPT !["A"] = restored(S["A"], d)]
RestoreB(d) == /\ d \in 0..DMAX
               /\ S' = [S EXCEPT !["B"] = restored(S["B"], d)]

\* Fork(child) = New(parent.next()) in the Go code: the child origin is the
\* parent's current (post-increment) state. The two then evolve independently.
ForkAtoB == S' = [S EXCEPT !["B"] = fresh(S["A"].state)]
ForkBtoA == S' = [S EXCEPT !["A"] = fresh(S["B"].state)]

Next ==
    \/ \E seed \in SEEDS : NewA(seed) \/ NewB(seed)
    \/ AdvanceA \/ AdvanceB
    \/ \E d \in 0..DMAX : RestoreA(d) \/ RestoreB(d)
    \/ ForkAtoB \/ ForkBtoA

Spec == Init /\ [][Next]_S

\* ----- Invariants -----

TypeOK == S \in [{"A","B"} -> Streams]

\* Core mechanism: the stored state is wholly determined by (origin, draws).
StateConsistency ==
    \A s \in {"A","B"} :
        S[s].state = ((S[s].origin + S[s].draws * Gamma) % M)

\* The guarantee the rest of the system depends on: a stream restored to
\* (seed, d) is indistinguishable from one advanced to d draws. Two streams
\* sharing origin and draw count hold the same state.
Reproducibility ==
    (S["A"].origin = S["B"].origin /\ S["A"].draws = S["B"].draws)
        => S["A"].state = S["B"].state

====
