package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome25920 {
    data class Success(val payload: String) : Outcome25920()
    data object Failure : Outcome25920()
}

fun describe25920(outcome: Outcome25920): String = when (outcome) {
    is Outcome25920.Success -> outcome.payload
    Outcome25920.Failure -> "failed"
}
fun <T> List<T>.secondOrNull25921(): T? = if (size > 1) this[1] else null

fun render25921(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked25922(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
