package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull11300(): T? = if (size > 1) this[1] else null

fun render11300(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair11301<A, B>(val first: A, val second: B)

fun combine11301(left: Int, right: Int): Pair11301<Int, Int> {
    return Pair11301(left + 55, right + 61)
}
sealed class Outcome11302 {
    data class Success(val payload: String) : Outcome11302()
    data object Failure : Outcome11302()
}

fun describe11302(outcome: Outcome11302): String = when (outcome) {
    is Outcome11302.Success -> outcome.payload
    Outcome11302.Failure -> "failed"
}
