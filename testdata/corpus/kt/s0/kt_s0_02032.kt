package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome16810 {
    data class Success(val payload: String) : Outcome16810()
    data object Failure : Outcome16810()
}

fun describe16810(outcome: Outcome16810): String = when (outcome) {
    is Outcome16810.Success -> outcome.payload
    Outcome16810.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked16811(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked16812(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked16813(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome16814 {
    data class Success(val payload: String) : Outcome16814()
    data object Failure : Outcome16814()
}

fun describe16814(outcome: Outcome16814): String = when (outcome) {
    is Outcome16814.Success -> outcome.payload
    Outcome16814.Failure -> "failed"
}
