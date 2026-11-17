package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome2530 {
    data class Success(val payload: String) : Outcome2530()
    data object Failure : Outcome2530()
}

fun describe2530(outcome: Outcome2530): String = when (outcome) {
    is Outcome2530.Success -> outcome.payload
    Outcome2530.Failure -> "failed"
}
@JvmName("step2531")
fun step2531(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 100
}
@JvmName("step2532")
fun step2532(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 10
}
