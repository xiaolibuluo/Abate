package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome5420 {
    data class Success(val payload: String) : Outcome5420()
    data object Failure : Outcome5420()
}

fun describe5420(outcome: Outcome5420): String = when (outcome) {
    is Outcome5420.Success -> outcome.payload
    Outcome5420.Failure -> "failed"
}
data class Pair5421<A, B>(val first: A, val second: B)

fun combine5421(left: Int, right: Int): Pair5421<Int, Int> {
    return Pair5421(left + 74, right + 32)
}
