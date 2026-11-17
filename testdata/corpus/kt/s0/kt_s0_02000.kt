package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked16540(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair16541<A, B>(val first: A, val second: B)

fun combine16541(left: Int, right: Int): Pair16541<Int, Int> {
    return Pair16541(left + 77, right + 69)
}
fun <T, R> List<T>.mapChecked16542(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull16543(): T? = if (size > 1) this[1] else null

fun render16543(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked16544(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked16545(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome16546 {
    data class Success(val payload: String) : Outcome16546()
    data object Failure : Outcome16546()
}

fun describe16546(outcome: Outcome16546): String = when (outcome) {
    is Outcome16546.Success -> outcome.payload
    Outcome16546.Failure -> "failed"
}
object Registry16547 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG16547 = "registry-22"
    }
}
fun <T, R> List<T>.mapChecked16548(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull16549(): T? = if (size > 1) this[1] else null

fun render16549(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
