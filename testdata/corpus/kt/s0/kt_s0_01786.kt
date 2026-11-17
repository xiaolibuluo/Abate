package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step14650")
fun step14650(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 38
}
sealed class Outcome14651 {
    data class Success(val payload: String) : Outcome14651()
    data object Failure : Outcome14651()
}

fun describe14651(outcome: Outcome14651): String = when (outcome) {
    is Outcome14651.Success -> outcome.payload
    Outcome14651.Failure -> "failed"
}
