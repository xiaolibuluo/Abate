---------------- MODULE ces ----------------
(************************************************************)
(* Abate CES v0.1 状态机模型（附录 A，CES-3.4 强制发布）    *)
(* 覆盖第 3 章状态定义、合法迁移表与不变量。                *)
(************************************************************)
EXTENDS Naturals

CONSTANTS
    Pending, Preflight, Injecting, Observing, Recovering, Verdict

VARIABLES
    state,     (* 当前生命周期状态 *)
    injected   (* 是否残留注入副作用 *)

vars == << state, injected >>

Init == /\ state = Pending
        /\ injected = FALSE

(* Pending -> Preflight：受理进入准入检查 *)
Admit == /\ state = Pending /\ state' = Preflight /\ injected' = injected

(* Pending -> Verdict：准入前取消，判定 CANCELLED *)
Cancel == /\ state = Pending /\ state' = Verdict /\ injected' = injected

(* Preflight -> Injecting：准入通过（CES-3.1.4） *)
Pass == /\ state = Preflight /\ state' = Injecting /\ injected' = injected

(* Preflight -> Verdict：准入失败，判定 REJECTED（CES-3.1.4） *)
Reject == /\ state = Preflight /\ state' = Verdict /\ injected' = injected

(* 红线触发：适用状态内强制迁移 Recovering（CES-3.1.1） *)
RedLine == /\ state \in {Preflight, Injecting, Observing}
           /\ state' = Recovering /\ injected' = injected

(* Injecting -> Observing：注入窗口结束 *)
WindowEnd == /\ state = Injecting /\ state' = Observing /\ injected' = injected

(* Observing -> Recovering：观测结束 *)
ObserveEnd == /\ state = Observing /\ state' = Recovering /\ injected' = injected

(* Recovering -> Verdict：恢复完成，无残留（CES-3.1.2） *)
Recover == /\ state = Recovering /\ state' = Verdict /\ injected' = FALSE

(* Verdict 终态自环：终态非死锁（CES-3.4） *)
Done == /\ state = Verdict /\ UNCHANGED vars

Next == Admit \/ Cancel \/ Pass \/ Reject \/ RedLine
        \/ WindowEnd \/ ObserveEnd \/ Recover \/ Done

Spec == Init /\ [][Next]_vars

(* 不变量 1：状态取值合法 *)
TypeOK == state \in {Pending, Preflight, Injecting, Observing, Recovering, Verdict}

(* 不变量 2：终态无残留（CES-3.1.2） *)
ResidueFree == state = Verdict => injected = FALSE
================