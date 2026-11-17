package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked21920(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome21921 {
    data class Success(val payload: String) : Outcome21921()
    data object Failure : Outcome21921()
}

fun describe21921(outcome: Outcome21921): String = when (outcome) {
    is Outcome21921.Success -> outcome.payload
    Outcome21921.Failure -> "failed"
}
