package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome3620 {
    data class Success(val payload: String) : Outcome3620()
    data object Failure : Outcome3620()
}

fun describe3620(outcome: Outcome3620): String = when (outcome) {
    is Outcome3620.Success -> outcome.payload
    Outcome3620.Failure -> "failed"
}
fun <T> List<T>.secondOrNull3621(): T? = if (size > 1) this[1] else null

fun render3621(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step3622")
fun step3622(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 69
}
fun <T, R> List<T>.mapChecked3623(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked3624(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
