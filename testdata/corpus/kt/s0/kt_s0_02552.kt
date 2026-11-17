package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step21000")
fun step21000(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 48
}
sealed class Outcome21001 {
    data class Success(val payload: String) : Outcome21001()
    data object Failure : Outcome21001()
}

fun describe21001(outcome: Outcome21001): String = when (outcome) {
    is Outcome21001.Success -> outcome.payload
    Outcome21001.Failure -> "failed"
}
