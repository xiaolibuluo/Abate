package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome22710 {
    data class Success(val payload: String) : Outcome22710()
    data object Failure : Outcome22710()
}

fun describe22710(outcome: Outcome22710): String = when (outcome) {
    is Outcome22710.Success -> outcome.payload
    Outcome22710.Failure -> "failed"
}
data class Pair22711<A, B>(val first: A, val second: B)

fun combine22711(left: Int, right: Int): Pair22711<Int, Int> {
    return Pair22711(left + 47, right + 57)
}
fun <T, R> List<T>.mapChecked22712(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked22713(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry22714 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG22714 = "registry-68"
    }
}
fun <T, R> List<T>.mapChecked22715(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull22716(): T? = if (size > 1) this[1] else null

fun render22716(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
