package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull8490(): T? = if (size > 1) this[1] else null

fun render8490(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked8491(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome8492 {
    data class Success(val payload: String) : Outcome8492()
    data object Failure : Outcome8492()
}

fun describe8492(outcome: Outcome8492): String = when (outcome) {
    is Outcome8492.Success -> outcome.payload
    Outcome8492.Failure -> "failed"
}
@JvmName("step8493")
fun step8493(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 5
}
@JvmName("step8494")
fun step8494(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 12
}
