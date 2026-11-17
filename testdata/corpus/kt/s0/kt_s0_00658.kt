package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome5400 {
    data class Success(val payload: String) : Outcome5400()
    data object Failure : Outcome5400()
}

fun describe5400(outcome: Outcome5400): String = when (outcome) {
    is Outcome5400.Success -> outcome.payload
    Outcome5400.Failure -> "failed"
}
@JvmName("step5401")
fun step5401(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 48
}
