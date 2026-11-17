package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull3240(): T? = if (size > 1) this[1] else null

fun render3240(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome3241 {
    data class Success(val payload: String) : Outcome3241()
    data object Failure : Outcome3241()
}

fun describe3241(outcome: Outcome3241): String = when (outcome) {
    is Outcome3241.Success -> outcome.payload
    Outcome3241.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked3242(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull3243(): T? = if (size > 1) this[1] else null

fun render3243(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull3244(): T? = if (size > 1) this[1] else null

fun render3244(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step3245")
fun step3245(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 81
}
fun <T> List<T>.secondOrNull3246(): T? = if (size > 1) this[1] else null

fun render3246(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
