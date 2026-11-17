package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome13260 {
    data class Success(val payload: String) : Outcome13260()
    data object Failure : Outcome13260()
}

fun describe13260(outcome: Outcome13260): String = when (outcome) {
    is Outcome13260.Success -> outcome.payload
    Outcome13260.Failure -> "failed"
}
@JvmName("step13261")
fun step13261(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 81
}
sealed class Outcome13262 {
    data class Success(val payload: String) : Outcome13262()
    data object Failure : Outcome13262()
}

fun describe13262(outcome: Outcome13262): String = when (outcome) {
    is Outcome13262.Success -> outcome.payload
    Outcome13262.Failure -> "failed"
}
