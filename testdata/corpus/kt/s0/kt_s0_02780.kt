package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step22940")
fun step22940(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 48
}
sealed class Outcome22941 {
    data class Success(val payload: String) : Outcome22941()
    data object Failure : Outcome22941()
}

fun describe22941(outcome: Outcome22941): String = when (outcome) {
    is Outcome22941.Success -> outcome.payload
    Outcome22941.Failure -> "failed"
}
