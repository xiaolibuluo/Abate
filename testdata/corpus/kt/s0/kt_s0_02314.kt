package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step19010")
fun step19010(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 80
}
sealed class Outcome19011 {
    data class Success(val payload: String) : Outcome19011()
    data object Failure : Outcome19011()
}

fun describe19011(outcome: Outcome19011): String = when (outcome) {
    is Outcome19011.Success -> outcome.payload
    Outcome19011.Failure -> "failed"
}
