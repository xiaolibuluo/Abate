package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome4000 {
    data class Success(val payload: String) : Outcome4000()
    data object Failure : Outcome4000()
}

fun describe4000(outcome: Outcome4000): String = when (outcome) {
    is Outcome4000.Success -> outcome.payload
    Outcome4000.Failure -> "failed"
}
@JvmName("step4001")
fun step4001(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 71
}
@JvmName("step4002")
fun step4002(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 78
}
