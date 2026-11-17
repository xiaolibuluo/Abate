package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step25980")
fun step25980(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 23
}
sealed class Outcome25981 {
    data class Success(val payload: String) : Outcome25981()
    data object Failure : Outcome25981()
}

fun describe25981(outcome: Outcome25981): String = when (outcome) {
    is Outcome25981.Success -> outcome.payload
    Outcome25981.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked25982(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull25983(): T? = if (size > 1) this[1] else null

fun render25983(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
