package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome24380 {
    data class Success(val payload: String) : Outcome24380()
    data object Failure : Outcome24380()
}

fun describe24380(outcome: Outcome24380): String = when (outcome) {
    is Outcome24380.Success -> outcome.payload
    Outcome24380.Failure -> "failed"
}
@JvmName("step24381")
fun step24381(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 84
}
