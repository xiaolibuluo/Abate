package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step12830")
fun step12830(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 97
}
fun <T> List<T>.secondOrNull12831(): T? = if (size > 1) this[1] else null

fun render12831(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome12832 {
    data class Success(val payload: String) : Outcome12832()
    data object Failure : Outcome12832()
}

fun describe12832(outcome: Outcome12832): String = when (outcome) {
    is Outcome12832.Success -> outcome.payload
    Outcome12832.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked12833(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
