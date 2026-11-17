package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome25450 {
    data class Success(val payload: String) : Outcome25450()
    data object Failure : Outcome25450()
}

fun describe25450(outcome: Outcome25450): String = when (outcome) {
    is Outcome25450.Success -> outcome.payload
    Outcome25450.Failure -> "failed"
}
@JvmName("step25451")
fun step25451(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 20
}
sealed class Outcome25452 {
    data class Success(val payload: String) : Outcome25452()
    data object Failure : Outcome25452()
}

fun describe25452(outcome: Outcome25452): String = when (outcome) {
    is Outcome25452.Success -> outcome.payload
    Outcome25452.Failure -> "failed"
}
