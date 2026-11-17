package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome20020 {
    data class Success(val payload: String) : Outcome20020()
    data object Failure : Outcome20020()
}

fun describe20020(outcome: Outcome20020): String = when (outcome) {
    is Outcome20020.Success -> outcome.payload
    Outcome20020.Failure -> "failed"
}
@JvmName("step20021")
fun step20021(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 42
}
