package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome9320 {
    data class Success(val payload: String) : Outcome9320()
    data object Failure : Outcome9320()
}

fun describe9320(outcome: Outcome9320): String = when (outcome) {
    is Outcome9320.Success -> outcome.payload
    Outcome9320.Failure -> "failed"
}
@JvmName("step9321")
fun step9321(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 100
}
