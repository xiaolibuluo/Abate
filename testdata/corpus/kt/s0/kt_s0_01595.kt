package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step13190")
fun step13190(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 80
}
sealed class Outcome13191 {
    data class Success(val payload: String) : Outcome13191()
    data object Failure : Outcome13191()
}

fun describe13191(outcome: Outcome13191): String = when (outcome) {
    is Outcome13191.Success -> outcome.payload
    Outcome13191.Failure -> "failed"
}
