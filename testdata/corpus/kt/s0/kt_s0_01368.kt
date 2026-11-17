package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step11270")
fun step11270(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 9
}
sealed class Outcome11271 {
    data class Success(val payload: String) : Outcome11271()
    data object Failure : Outcome11271()
}

fun describe11271(outcome: Outcome11271): String = when (outcome) {
    is Outcome11271.Success -> outcome.payload
    Outcome11271.Failure -> "failed"
}
