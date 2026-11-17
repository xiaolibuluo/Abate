package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step15240")
fun step15240(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 29
}
fun <T> List<T>.secondOrNull15241(): T? = if (size > 1) this[1] else null

fun render15241(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked15242(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome15243 {
    data class Success(val payload: String) : Outcome15243()
    data object Failure : Outcome15243()
}

fun describe15243(outcome: Outcome15243): String = when (outcome) {
    is Outcome15243.Success -> outcome.payload
    Outcome15243.Failure -> "failed"
}
