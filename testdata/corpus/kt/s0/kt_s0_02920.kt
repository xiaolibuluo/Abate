package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step24090")
fun step24090(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 88
}
@JvmName("step24091")
fun step24091(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 95
}
@JvmName("step24092")
fun step24092(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 5
}
sealed class Outcome24093 {
    data class Success(val payload: String) : Outcome24093()
    data object Failure : Outcome24093()
}

fun describe24093(outcome: Outcome24093): String = when (outcome) {
    is Outcome24093.Success -> outcome.payload
    Outcome24093.Failure -> "failed"
}
sealed class Outcome24094 {
    data class Success(val payload: String) : Outcome24094()
    data object Failure : Outcome24094()
}

fun describe24094(outcome: Outcome24094): String = when (outcome) {
    is Outcome24094.Success -> outcome.payload
    Outcome24094.Failure -> "failed"
}
