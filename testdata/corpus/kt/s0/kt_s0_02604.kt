package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome21470 {
    data class Success(val payload: String) : Outcome21470()
    data object Failure : Outcome21470()
}

fun describe21470(outcome: Outcome21470): String = when (outcome) {
    is Outcome21470.Success -> outcome.payload
    Outcome21470.Failure -> "failed"
}
fun <T> List<T>.secondOrNull21471(): T? = if (size > 1) this[1] else null

fun render21471(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
