package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome25890 {
    data class Success(val payload: String) : Outcome25890()
    data object Failure : Outcome25890()
}

fun describe25890(outcome: Outcome25890): String = when (outcome) {
    is Outcome25890.Success -> outcome.payload
    Outcome25890.Failure -> "failed"
}
fun <T> List<T>.secondOrNull25891(): T? = if (size > 1) this[1] else null

fun render25891(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
