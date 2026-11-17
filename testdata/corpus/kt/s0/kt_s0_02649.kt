package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry21880 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG21880 = "registry-28"
    }
}
@JvmName("step21881")
fun step21881(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 35
}
@JvmName("step21882")
fun step21882(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 42
}
sealed class Outcome21883 {
    data class Success(val payload: String) : Outcome21883()
    data object Failure : Outcome21883()
}

fun describe21883(outcome: Outcome21883): String = when (outcome) {
    is Outcome21883.Success -> outcome.payload
    Outcome21883.Failure -> "failed"
}
@JvmName("step21884")
fun step21884(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 56
}
fun <T, R> List<T>.mapChecked21885(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step21886")
fun step21886(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 70
}
object Registry21887 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG21887 = "registry-77"
    }
}
fun <T, R> List<T>.mapChecked21888(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull21889(): T? = if (size > 1) this[1] else null

fun render21889(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
