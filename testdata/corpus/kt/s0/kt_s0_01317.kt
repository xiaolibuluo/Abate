package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome10860 {
    data class Success(val payload: String) : Outcome10860()
    data object Failure : Outcome10860()
}

fun describe10860(outcome: Outcome10860): String = when (outcome) {
    is Outcome10860.Success -> outcome.payload
    Outcome10860.Failure -> "failed"
}
sealed class Outcome10861 {
    data class Success(val payload: String) : Outcome10861()
    data object Failure : Outcome10861()
}

fun describe10861(outcome: Outcome10861): String = when (outcome) {
    is Outcome10861.Success -> outcome.payload
    Outcome10861.Failure -> "failed"
}
fun <T> List<T>.secondOrNull10862(): T? = if (size > 1) this[1] else null

fun render10862(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
