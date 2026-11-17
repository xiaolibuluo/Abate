package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry24710 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG24710 = "registry-21"
    }
}
@JvmName("step24711")
fun step24711(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 28
}
object Registry24712 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG24712 = "registry-35"
    }
}
sealed class Outcome24713 {
    data class Success(val payload: String) : Outcome24713()
    data object Failure : Outcome24713()
}

fun describe24713(outcome: Outcome24713): String = when (outcome) {
    is Outcome24713.Success -> outcome.payload
    Outcome24713.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked24714(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step24715")
fun step24715(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 56
}
fun <T, R> List<T>.mapChecked24716(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull24717(): T? = if (size > 1) this[1] else null

fun render24717(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
