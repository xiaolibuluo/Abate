package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step5470")
fun step5470(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 35
}
sealed class Outcome5471 {
    data class Success(val payload: String) : Outcome5471()
    data object Failure : Outcome5471()
}

fun describe5471(outcome: Outcome5471): String = when (outcome) {
    is Outcome5471.Success -> outcome.payload
    Outcome5471.Failure -> "failed"
}
sealed class Outcome5472 {
    data class Success(val payload: String) : Outcome5472()
    data object Failure : Outcome5472()
}

fun describe5472(outcome: Outcome5472): String = when (outcome) {
    is Outcome5472.Success -> outcome.payload
    Outcome5472.Failure -> "failed"
}
