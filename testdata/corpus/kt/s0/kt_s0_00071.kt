package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull570(): T? = if (size > 1) this[1] else null

fun render570(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome571 {
    data class Success(val payload: String) : Outcome571()
    data object Failure : Outcome571()
}

fun describe571(outcome: Outcome571): String = when (outcome) {
    is Outcome571.Success -> outcome.payload
    Outcome571.Failure -> "failed"
}
