package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step21320")
fun step21320(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 76
}
sealed class Outcome21321 {
    data class Success(val payload: String) : Outcome21321()
    data object Failure : Outcome21321()
}

fun describe21321(outcome: Outcome21321): String = when (outcome) {
    is Outcome21321.Success -> outcome.payload
    Outcome21321.Failure -> "failed"
}
fun <T> List<T>.secondOrNull21322(): T? = if (size > 1) this[1] else null

fun render21322(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry21323 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG21323 = "registry-97"
    }
}
fun <T> List<T>.secondOrNull21324(): T? = if (size > 1) this[1] else null

fun render21324(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step21325")
fun step21325(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 14
}
fun <T, R> List<T>.mapChecked21326(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked21327(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
