package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome19220 {
    data class Success(val payload: String) : Outcome19220()
    data object Failure : Outcome19220()
}

fun describe19220(outcome: Outcome19220): String = when (outcome) {
    is Outcome19220.Success -> outcome.payload
    Outcome19220.Failure -> "failed"
}
sealed class Outcome19221 {
    data class Success(val payload: String) : Outcome19221()
    data object Failure : Outcome19221()
}

fun describe19221(outcome: Outcome19221): String = when (outcome) {
    is Outcome19221.Success -> outcome.payload
    Outcome19221.Failure -> "failed"
}
@JvmName("step19222")
fun step19222(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 76
}
fun <T, R> List<T>.mapChecked19223(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull19224(): T? = if (size > 1) this[1] else null

fun render19224(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
