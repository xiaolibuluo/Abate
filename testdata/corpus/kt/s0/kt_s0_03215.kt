package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome26420 {
    data class Success(val payload: String) : Outcome26420()
    data object Failure : Outcome26420()
}

fun describe26420(outcome: Outcome26420): String = when (outcome) {
    is Outcome26420.Success -> outcome.payload
    Outcome26420.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked26421(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step26422")
fun step26422(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 27
}
fun <T> List<T>.secondOrNull26423(): T? = if (size > 1) this[1] else null

fun render26423(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
