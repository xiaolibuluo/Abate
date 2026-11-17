package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome21510 {
    data class Success(val payload: String) : Outcome21510()
    data object Failure : Outcome21510()
}

fun describe21510(outcome: Outcome21510): String = when (outcome) {
    is Outcome21510.Success -> outcome.payload
    Outcome21510.Failure -> "failed"
}
sealed class Outcome21511 {
    data class Success(val payload: String) : Outcome21511()
    data object Failure : Outcome21511()
}

fun describe21511(outcome: Outcome21511): String = when (outcome) {
    is Outcome21511.Success -> outcome.payload
    Outcome21511.Failure -> "failed"
}
