package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked12810(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step12811")
fun step12811(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 78
}
fun <T> List<T>.secondOrNull12812(): T? = if (size > 1) this[1] else null

fun render12812(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome12813 {
    data class Success(val payload: String) : Outcome12813()
    data object Failure : Outcome12813()
}

fun describe12813(outcome: Outcome12813): String = when (outcome) {
    is Outcome12813.Success -> outcome.payload
    Outcome12813.Failure -> "failed"
}
