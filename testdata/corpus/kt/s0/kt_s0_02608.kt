package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome21500 {
    data class Success(val payload: String) : Outcome21500()
    data object Failure : Outcome21500()
}

fun describe21500(outcome: Outcome21500): String = when (outcome) {
    is Outcome21500.Success -> outcome.payload
    Outcome21500.Failure -> "failed"
}
@JvmName("step21501")
fun step21501(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 26
}
@JvmName("step21502")
fun step21502(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 33
}
