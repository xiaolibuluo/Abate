package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked26040(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step26041")
fun step26041(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 11
}
fun <T> List<T>.secondOrNull26042(): T? = if (size > 1) this[1] else null

fun render26042(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked26043(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome26044 {
    data class Success(val payload: String) : Outcome26044()
    data object Failure : Outcome26044()
}

fun describe26044(outcome: Outcome26044): String = when (outcome) {
    is Outcome26044.Success -> outcome.payload
    Outcome26044.Failure -> "failed"
}
@JvmName("step26045")
fun step26045(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 39
}
