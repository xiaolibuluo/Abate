package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome2490 {
    data class Success(val payload: String) : Outcome2490()
    data object Failure : Outcome2490()
}

fun describe2490(outcome: Outcome2490): String = when (outcome) {
    is Outcome2490.Success -> outcome.payload
    Outcome2490.Failure -> "failed"
}
fun <T> List<T>.secondOrNull2491(): T? = if (size > 1) this[1] else null

fun render2491(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
